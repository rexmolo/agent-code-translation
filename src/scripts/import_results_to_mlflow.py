#!/usr/bin/env python3
"""Import existing result JSONs from docs/results/ into MLflow.

Usage:
    uv run python src/scripts/import_results_to_mlflow.py
"""

import json
import re
from pathlib import Path

import mlflow

RESULTS_DIR = Path(__file__).resolve().parent.parent.parent / "docs" / "results"


def _parse_strategy(strategy: str) -> tuple[str | None, str]:
    """Parse strategy field into (backend, experiment).

    Examples:
        "baseline"                  -> (None, "baseline")
        "vec-chroma/rag-full"       -> ("vec-chroma", "rag-full")
        "vec-gemini/rag-pattern-only" -> ("vec-gemini", "rag-pattern-only")
    """
    if "/" in strategy:
        backend, experiment = strategy.split("/", 1)
        return backend, experiment
    return None, strategy


def import_result(filepath: Path) -> None:
    data = json.loads(filepath.read_text(encoding="utf-8"))

    provider = data["provider"]
    variant = data["model"]
    strategy = data["strategy"]
    summary = data["summary"]

    backend, experiment = _parse_strategy(strategy)

    # Extract dimensions from backend (e.g. "vec-chroma-768" -> 768)
    dims = None
    if backend and (m := re.search(r"-(\d+)$", backend)):
        dims = int(m.group(1))

    # Build run name
    run_name = f"{provider}/{variant}"
    if backend:
        run_name += f"/{backend}"
    run_name += f"/{experiment}"

    with mlflow.start_run(run_name=run_name):
        # Parameters
        params = {
            "provider": provider,
            "variant": variant,
            "experiment": experiment,
            "dataset": summary.get("dataset", "humaneval-x"),
        }
        if backend:
            params["embedding_backend"] = backend
        if dims:
            params["embedding_dimensions"] = dims

        mlflow.log_params(params)

        # Metrics (as percentages)
        mlflow.log_metrics({
            "total_files": summary["total_files"],
            "compilation_at_1": round(summary["compilation_at_1"] * 100, 1),
            "pass_at_1": round(summary["pass_at_1"] * 100, 1),
        })

        # Log the original JSON as artifact
        mlflow.log_artifact(str(filepath))

    print(f"  Logged: {run_name}")


def main():
    mlflow.set_experiment("thesis-code-translation")

    json_files = sorted(RESULTS_DIR.glob("*.json"))
    print(f"Found {len(json_files)} result files in {RESULTS_DIR}\n")

    for filepath in json_files:
        import_result(filepath)

    print(f"\nDone! {len(json_files)} runs imported to MLflow.")


if __name__ == "__main__":
    main()
