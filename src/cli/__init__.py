"""CLI entry point for experiment workflows.

Usage:
    uv run python -m src.cli                          # interactive mode
    uv run python -m src.cli translate [-e EXPERIMENT] [--skip-preflight]
    uv run python -m src.cli evaluate [-e EXPERIMENT]
"""

import importlib
import inspect
from pathlib import Path

import click
import questionary
from dotenv import load_dotenv
from questionary import Choice, Style
from rich.console import Console
from rich.panel import Panel

from src.config import LAB_DIR, TRANSLATION_SOURCE_DIR, TRANSLATION_TARGET_DIR
from src.models.registry import enable_model, list_providers, list_variants

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


def _load_experiment(experiment: str):
    """Dynamically import an experiment's run module."""
    module_path = f"src.lab.{experiment}.run"
    try:
        return importlib.import_module(module_path)
    except ModuleNotFoundError as e:
        raise click.ClickException(
            f"Experiment '{experiment}' not found (tried {module_path}): {e}"
        )


def _discover_experiments() -> list[str]:
    """Return sorted list of experiment folder names under src/lab/."""
    if not LAB_DIR.is_dir():
        return []
    return sorted(
        d.name for d in LAB_DIR.iterdir()
        if d.is_dir() and not d.name.startswith("__")
    )


def _discover_actions(run_mod) -> list[str]:
    """Discover available actions from an experiment's run module.

    If the module defines ACTIONS, use that. Otherwise introspect public functions.
    """
    if hasattr(run_mod, "ACTIONS"):
        return list(run_mod.ACTIONS)

    return [
        name
        for name, obj in inspect.getmembers(run_mod, inspect.isfunction)
        if not name.startswith("_")
        and name not in _SKIP_FUNCTIONS
        and obj.__module__ == run_mod.__name__
    ]


def _ask_or_abort(result):
    """If the user pressed Ctrl+C (result is None), exit cleanly."""
    if result is None:
        console.print("\n\nAborted.")
        raise SystemExit(0)
    return result


def _pick_eval_target(dataset: str) -> Path:
    """Discover target folders (provider/variant) and let user pick one."""
    target_root = TRANSLATION_TARGET_DIR / dataset
    folders: list[tuple[str, Path]] = []

    if target_root.is_dir():
        for provider_dir in sorted(target_root.iterdir()):
            if not provider_dir.is_dir() or provider_dir.name.startswith("."):
                continue
            for variant_dir in sorted(provider_dir.iterdir()):
                if not variant_dir.is_dir() or variant_dir.name.startswith("."):
                    continue
                label = f"{provider_dir.name}/{variant_dir.name}"
                folders.append((label, variant_dir))

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
    """Interactive mode: arrow-key selection for experiment, action, models."""
    console.print(Panel(
        "[bold]Thesis Experiment CLI[/bold]\nPython → Go Translation & Evaluation",
        border_style="blue",
    ))

    # Step 1: Pick experiment
    experiments = _discover_experiments()
    if not experiments:
        console.print("[red]No experiments found under src/lab/[/red]")
        raise SystemExit(1)

    experiment = _ask_or_abort(questionary.select(
        "Select experiment:",
        choices=experiments,
        style=_style,
    ).ask())

    # Step 2: Pick dataset
    selected_dataset = _ask_or_abort(questionary.select(
        "Select dataset:",
        choices=_DATASETS,
        style=_style,
    ).ask())

    # Step 3: Discover and pick action from that experiment
    run_mod = _load_experiment(experiment)
    actions = _discover_actions(run_mod)
    if not actions:
        console.print(f"[red]No actions found in {experiment}/run.py[/red]")
        raise SystemExit(1)

    action = _ask_or_abort(questionary.select(
        "Select action:",
        choices=actions,
        style=_style,
    ).ask())

    action_fn = getattr(run_mod, action)
    sig = inspect.signature(action_fn)
    kwargs: dict = {}

    if "dataset" in sig.parameters:
        kwargs["dataset"] = selected_dataset

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

    # Confirm
    console.print(
        f"\n→ [cyan]{experiment}[/cyan] / [magenta]{selected_dataset}[/magenta] "
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
    "-e", "--experiment", default="00_get_hands_on", show_default=True,
    help="Experiment ID (folder name under src/lab/).",
)
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
def translate(experiment, dataset, source_dir, target_dir, skip_preflight):
    """Run translation for an experiment."""
    run_mod = _load_experiment(experiment)
    kwargs = {"skip_preflight": skip_preflight, "dataset": dataset}
    if source_dir is not None:
        kwargs["source_dir"] = source_dir
    if target_dir is not None:
        kwargs["target_dir"] = target_dir
    run_mod.translate(**kwargs)


@cli.command()
@click.option(
    "-e", "--experiment", default="00_get_hands_on", show_default=True,
    help="Experiment ID (folder name under src/lab/).",
)
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
def evaluate(experiment, dataset, source_dir, target_dir):
    """Evaluate existing translated files for an experiment."""
    run_mod = _load_experiment(experiment)
    kwargs = {"dataset": dataset}
    if source_dir is not None:
        kwargs["source_dir"] = source_dir
    if target_dir is not None:
        kwargs["eval_target_dir"] = target_dir
    run_mod.evaluate(**kwargs)
