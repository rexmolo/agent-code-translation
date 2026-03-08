"""CLI entry point for experiment workflows.

Usage:
    uv run python -m src.cli                          # interactive mode
    uv run python -m src.cli translate [--skip-preflight]
    uv run python -m src.cli evaluate
"""

import inspect
from pathlib import Path

import click
import questionary
from dotenv import load_dotenv
from questionary import Choice, Style
from rich.console import Console
from rich.panel import Panel

from src.config import TRANSLATION_SOURCE_DIR, TRANSLATION_TARGET_DIR
from src.core import pipeline as _pipeline
from src.core.logger import set_verbose
from src.providers.registry import enable_model, list_providers, list_variants

console = Console()

# Green highlight for the focused/selected item
_style = Style([
    ("highlighted", "fg:green bold"),
    ("selected", "fg:green"),
    ("pointer", "fg:green bold"),
    ("answer", "fg:green bold"),
])

# Dataset choices
_DATASETS = [
    Choice(title="Local source code", value="local"),
    Choice(title="HumanEval-X", value="humaneval-x"),
]

# Internal helpers to exclude when discovering actions from run.py
_SKIP_FUNCTIONS = {
    "main", "discover_python_files", "find_test_file",
    "mirror_path", "evaluate_file", "preflight_check",
    "load_humaneval_x",
}

_SAMPLE_OPTIONS = [1, 10, 20]  # fixed sample sizes; None means whole dataset


def _dataset_size(dataset: str) -> int | None:
    """Return the number of items in the dataset (best-effort, no heavy load)."""
    try:
        if dataset == "humaneval-x":
            from src.data.humaneval_x import load_humaneval_x
            return len(load_humaneval_x())
        elif dataset == "local":
            from src.core.evaluation import discover_python_files
            return len(discover_python_files(TRANSLATION_SOURCE_DIR))
    except Exception:
        pass
    return None


def _sample_choices(dataset: str) -> list[Choice]:
    """Build sample-size choices including a dynamic 'Whole' option.

    Uses 0 as sentinel for 'whole dataset' because questionary returns None
    when the user presses Ctrl+C, so we can't use None as a choice value.
    """
    size = _dataset_size(dataset)
    whole_label = f"Whole dataset ({size})" if size else "Whole dataset"
    choices = [Choice(title=f"{n} (smoke test)" if n == 1 else str(n), value=n)
               for n in _SAMPLE_OPTIONS]
    choices.append(Choice(title=whole_label, value=0))
    return choices


def _discover_actions() -> list[str]:
    """Discover available actions from the core run module.

    If the module defines ACTIONS, use that. Otherwise introspect public functions.
    """
    if hasattr(_pipeline, "ACTIONS"):
        return list(_pipeline.ACTIONS)

    return [
        name
        for name, obj in inspect.getmembers(_pipeline, inspect.isfunction)
        if not name.startswith("_")
        and name not in _SKIP_FUNCTIONS
        and obj.__module__ == _pipeline.__name__
    ]


def _ask_or_abort(result):
    """If the user pressed Ctrl+C (result is None), exit cleanly."""
    if result is None:
        console.print("\n\nAborted.")
        raise SystemExit(0)
    return result


def _pick_eval_target(dataset: str) -> Path:
    """Discover target folders (provider/variant/experiment) and let user pick one."""
    target_root = TRANSLATION_TARGET_DIR / dataset
    folders: list[tuple[str, Path]] = []

    if target_root.is_dir():
        for provider_dir in sorted(target_root.iterdir()):
            if not provider_dir.is_dir() or provider_dir.name.startswith("."):
                continue
            for variant_dir in sorted(provider_dir.iterdir()):
                if not variant_dir.is_dir() or variant_dir.name.startswith("."):
                    continue
                for experiment_dir in sorted(variant_dir.iterdir()):
                    if not experiment_dir.is_dir() or experiment_dir.name.startswith("."):
                        continue
                    label = f"{provider_dir.name}/{variant_dir.name}/{experiment_dir.name}"
                    folders.append((label, experiment_dir))

    if not folders:
        console.print(f"[red]No translated output found under target/{dataset}/[/red]")
        raise SystemExit(1)

    choices = [Choice(title=label, value=path) for label, path in folders]

    selected = _ask_or_abort(questionary.select(
        "Select target to evaluate:",
        choices=choices,
        style=_style,
    ).ask())

    return selected


def _interactive():
    """Interactive mode: arrow-key selection for action, dataset, models."""
    console.print(Panel(
        "[bold]Thesis Experiment CLI[/bold]\nPython → Go Translation & Evaluation",
        border_style="blue",
    ))

    # Step 1: Pick dataset
    selected_dataset = _ask_or_abort(questionary.select(
        "Select dataset:",
        choices=_DATASETS,
        style=_style,
    ).ask())

    # Step 2: Discover and pick action
    actions = _discover_actions()
    if not actions:
        console.print("[red]No actions found in core/pipeline.py[/red]")
        raise SystemExit(1)

    action = _ask_or_abort(questionary.select(
        "Select action:",
        choices=actions,
        style=_style,
    ).ask())

    action_fn = getattr(_pipeline, action)
    sig = inspect.signature(action_fn)
    kwargs: dict = {}

    if "dataset" in sig.parameters:
        kwargs["dataset"] = selected_dataset

    # Step 3: Sample size (only for translate-like actions that have a 'sample' param)
    if action != "evaluate":
        sample_choices = _sample_choices(selected_dataset)
        selected_sample = _ask_or_abort(questionary.select(
            "How many to translate?",
            choices=sample_choices,
            style=_style,
        ).ask())
        if "sample" in sig.parameters:
            # 0 is the sentinel for "whole dataset" → pass None (no slicing)
            kwargs["sample"] = selected_sample if selected_sample != 0 else None

    if action == "evaluate":
        # For evaluate: pick target folder directly instead of model
        target_dir = _pick_eval_target(selected_dataset)
        kwargs["eval_target_dir"] = target_dir
        display_label = str(target_dir)
    else:
        # For translate and others: pick provider + variant
        providers = list_providers()
        provider_choices = [
            Choice(title=f"{p['label']} ({p['key']})", value=p["key"])
            for p in providers
        ]

        selected_provider = _ask_or_abort(questionary.select(
            "Select provider:",
            choices=provider_choices,
            style=_style,
        ).ask())

        variants = list_variants(selected_provider)
        variant_choices = [
            Choice(title=f"{v['label']} ({v['model_id']})", value=v["key"])
            for v in variants
        ]

        selected_variant = _ask_or_abort(questionary.select(
            "Select model:",
            choices=variant_choices,
            style=_style,
        ).ask())

        enable_model(selected_provider, selected_variant)
        display_label = next(v["label"] for v in variants if v["key"] == selected_variant)

        # Step: Experiment name
        if "experiment" in sig.parameters:
            experiment_name = _ask_or_abort(questionary.text(
                "Experiment name:",
                default="baseline",
                style=_style,
            ).ask())
            kwargs["experiment"] = experiment_name

    # Verbose logging
    enable_verbose = _ask_or_abort(
        questionary.confirm("Enable verbose logging?", default=False, style=_style).ask()
    )
    if enable_verbose:
        set_verbose(True)

    # Confirm
    console.print(
        f"\n→ [magenta]{selected_dataset}[/magenta] "
        f"/ [green]{action}[/green] with [green]{display_label}[/green]\n"
    )

    proceed = _ask_or_abort(questionary.select(
        "Ready?",
        choices=["Confirm", "Cancel"],
        style=_style,
    ).ask())

    if proceed == "Cancel":
        console.print("Cancelled.")
        raise SystemExit(0)

    if "skip_preflight" in sig.parameters:
        kwargs["skip_preflight"] = _ask_or_abort(
            questionary.confirm("Skip preflight checks?", default=False, style=_style).ask()
        )

    console.print()
    action_fn(**kwargs)


# ---------------------------------------------------------------------------
# Click CLI (subcommand mode)
# ---------------------------------------------------------------------------

@click.group(invoke_without_command=True)
@click.pass_context
def cli(ctx):
    """Thesis experiment CLI — translate Python to Go and evaluate results."""
    load_dotenv()
    if ctx.invoked_subcommand is None:
        _interactive()


@cli.command()
@click.option(
    "-d", "--dataset", type=click.Choice(["local", "humaneval-x"]),
    default="local", show_default=True,
    help="Dataset to use for translation.",
)
@click.option(
    "--source-dir", type=click.Path(exists=True, path_type=Path), default=None,
    help="Override source directory.",
)
@click.option(
    "--target-dir", type=click.Path(path_type=Path), default=None,
    help="Override target directory.",
)
@click.option(
    "--skip-preflight", is_flag=True, default=False,
    help="Skip API/environment checks.",
)
@click.option(
    "-e", "--experiment", type=str, default="baseline", show_default=True,
    help="Experiment subfolder name (e.g. baseline, rag).",
)
@click.option(
    "-p", "--provider", type=str, default=None,
    help="Provider key (e.g. minimax, gemini).",
)
@click.option(
    "-v", "--variant", type=str, default=None,
    help="Model variant key (e.g. M2.5, 2.5_pro).",
)
@click.option(
    "-n", "--sample", type=int, default=None,
    help="Translate only the first N items.",
)
@click.option(
    "-V", "--verbose", is_flag=True, default=False,
    help="Enable verbose step-by-step logging.",
)
def translate(dataset, source_dir, target_dir, skip_preflight, experiment, provider, variant, sample, verbose):
    """Run translation pipeline."""
    if verbose:
        set_verbose(True)
    if provider and variant:
        enable_model(provider, variant)
    kwargs = {"skip_preflight": skip_preflight, "dataset": dataset, "experiment": experiment}
    if source_dir is not None:
        kwargs["source_dir"] = source_dir
    if target_dir is not None:
        kwargs["target_dir"] = target_dir
    if sample is not None:
        kwargs["sample"] = sample
    _pipeline.translate(**kwargs)


@cli.command()
@click.option(
    "-d", "--dataset", type=click.Choice(["local", "humaneval-x"]),
    default="local", show_default=True,
    help="Dataset being evaluated.",
)
@click.option(
    "--source-dir", type=click.Path(exists=True, path_type=Path), default=None,
    help="Override source directory.",
)
@click.option(
    "--target-dir", type=click.Path(exists=True, path_type=Path), default=None,
    help="Path to the specific translated output folder.",
)
@click.option(
    "-V", "--verbose", is_flag=True, default=False,
    help="Enable verbose step-by-step logging.",
)
def evaluate(dataset, source_dir, target_dir, verbose):
    """Evaluate existing translated files."""
    if verbose:
        set_verbose(True)
    kwargs = {"dataset": dataset}
    if source_dir is not None:
        kwargs["source_dir"] = source_dir
    if target_dir is not None:
        kwargs["eval_target_dir"] = target_dir
    _pipeline.evaluate(**kwargs)
