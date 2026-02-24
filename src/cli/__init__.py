"""CLI entry point for experiment workflows.

Usage:
    uv run python -m src.cli translate [-e EXPERIMENT] [--skip-preflight]
    uv run python -m src.cli evaluate [-e EXPERIMENT]
"""

import importlib
from pathlib import Path

import click
from dotenv import load_dotenv

from src.config import TRANSLATION_SOURCE_DIR, TRANSLATION_TARGET_DIR


def _load_experiment(experiment: str):
    """Dynamically import an experiment's run module."""
    module_path = f"src.lab.{experiment}.run"
    try:
        return importlib.import_module(module_path)
    except ModuleNotFoundError as e:
        raise click.ClickException(
            f"Experiment '{experiment}' not found (tried {module_path}): {e}"
        )


@click.group()
def cli():
    """Thesis experiment CLI — translate Python to Go and evaluate results."""
    load_dotenv()


@cli.command()
@click.option(
    "-e", "--experiment", default="00_get_hands_on", show_default=True,
    help="Experiment ID (folder name under src/lab/).",
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
def translate(experiment, source_dir, target_dir, skip_preflight):
    """Run translation for an experiment."""
    run_mod = _load_experiment(experiment)
    kwargs = {"skip_preflight": skip_preflight}
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
    "--source-dir", type=click.Path(exists=True, path_type=Path), default=None,
    help="Override source directory.",
)
@click.option(
    "--target-dir", type=click.Path(exists=True, path_type=Path), default=None,
    help="Override target directory.",
)
def evaluate(experiment, source_dir, target_dir):
    """Evaluate existing translated files for an experiment."""
    run_mod = _load_experiment(experiment)
    kwargs = {}
    if source_dir is not None:
        kwargs["source_dir"] = source_dir
    if target_dir is not None:
        kwargs["target_dir"] = target_dir
    run_mod.evaluate(**kwargs)
