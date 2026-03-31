#!/usr/bin/env python3
"""Statistical analysis of multi-run experiment results.

Discovers mlflow_results.json files, groups by dimension + experiment,
runs ANOVA across dimensions and pairwise t-tests if significant.

Usage:
    uv run python src/scripts/analyze_statistics.py
    uv run python src/scripts/analyze_statistics.py --metric pass_at_1
    uv run python src/scripts/analyze_statistics.py --metric compilation_at_1
"""

import json
import re
import sys
from collections import defaultdict
from pathlib import Path

import click
from rich.console import Console
from rich.table import Table
from scipy import stats

RESULTS_ROOT = Path(__file__).resolve().parent.parent.parent / "data" / "translation" / "target" / "humaneval-x"

SIGNIFICANCE_LEVEL = 0.05


def discover_results(root: Path) -> list[dict]:
    """Find all mlflow_results.json files and extract metadata from paths."""
    results = []
    for json_path in sorted(root.rglob("mlflow_results.json")):
        try:
            data = json.loads(json_path.read_text(encoding="utf-8"))
        except (json.JSONDecodeError, OSError):
            continue

        # Extract run_id from path (look for run-N directory)
        run_id = None
        for part in json_path.parts:
            if part.startswith("run-"):
                run_id = int(part.split("-", 1)[1])
                break

        # Extract dimension from backend (e.g. "vec-chroma-768" -> 768)
        backend = data.get("backend")
        dims = None
        if backend and (m := re.search(r"-(\d+)$", backend)):
            dims = int(m.group(1))

        experiment = data.get("strategy", "unknown")
        provider = data.get("provider", "unknown")
        model = data.get("model", "unknown")

        results.append({
            "path": str(json_path),
            "provider": provider,
            "model": model,
            "experiment": experiment,
            "backend": backend,
            "dimensions": dims,
            "run_id": run_id,
            "compilation_at_1": data["summary"]["compilation_at_1"],
            "pass_at_1": data["summary"]["pass_at_1"],
            "total_files": data["summary"]["total_files"],
        })

    return results


def group_results(results: list[dict], metric: str) -> dict[str, dict[int | str, list[float]]]:
    """Group results by experiment -> dimension -> list of metric values.

    Returns: {experiment: {dimension: [values]}}
    Baseline results are grouped under experiment="baseline", dimension="none".
    """
    grouped: dict[str, dict[int | str, list[float]]] = defaultdict(lambda: defaultdict(list))

    for r in results:
        experiment = r["experiment"]
        dims = r["dimensions"]
        value = r[metric]

        if experiment == "baseline":
            grouped["baseline"]["none"].append(value)
        elif dims is not None:
            grouped[experiment][dims].append(value)

    return dict(grouped)


def run_analysis(grouped: dict[str, dict[int | str, list[float]]], metric: str) -> None:
    """Run ANOVA and pairwise t-tests, display results."""
    console = Console()
    metric_label = metric.replace("_", "@").replace("at@", "@")

    for experiment, dim_groups in sorted(grouped.items()):
        console.print(f"\n[bold cyan]{'='*60}[/bold cyan]")
        console.print(f"[bold cyan]Experiment: {experiment}[/bold cyan]")
        console.print(f"[bold cyan]{'='*60}[/bold cyan]")

        # Summary table
        table = Table(title=f"{experiment} — {metric_label}")
        table.add_column("Dimension", style="cyan", justify="right")
        table.add_column("Runs", justify="center")
        table.add_column("Mean", justify="right", style="green")
        table.add_column("StdDev", justify="right")
        table.add_column("Min", justify="right", style="dim")
        table.add_column("Max", justify="right", style="dim")

        dim_keys = sorted(dim_groups.keys(), key=lambda x: (isinstance(x, str), x))
        for dim in dim_keys:
            values = dim_groups[dim]
            n = len(values)
            mean = sum(values) / n
            std = (sum((v - mean) ** 2 for v in values) / max(n - 1, 1)) ** 0.5
            dim_label = str(dim) if dim != "none" else "N/A"
            table.add_row(
                dim_label,
                str(n),
                f"{mean:.1%}",
                f"{std:.1%}" if n > 1 else "—",
                f"{min(values):.1%}",
                f"{max(values):.1%}",
            )

        console.print(table)

        # Skip ANOVA for baseline (only one group) or single dimension
        if experiment == "baseline" or len(dim_groups) < 2:
            if experiment == "baseline":
                console.print("[dim]Baseline — no dimension comparison needed.[/dim]")
            else:
                console.print("[dim]Only one dimension group — no comparison possible.[/dim]")
            continue

        # Check if we have enough data for statistical tests
        groups = [dim_groups[d] for d in dim_keys]
        if any(len(g) < 2 for g in groups):
            console.print("[yellow]⚠ Need at least 2 runs per dimension for statistical tests.[/yellow]")
            continue

        # ANOVA
        f_stat, p_value = stats.f_oneway(*groups)
        if p_value < SIGNIFICANCE_LEVEL:
            sig_label = f"[bold green]SIGNIFICANT (p={p_value:.4f})[/bold green]"
        else:
            sig_label = f"[bold red]NOT significant (p={p_value:.4f})[/bold red]"

        console.print(f"\n  ANOVA: F={f_stat:.3f}, p={p_value:.4f} → {sig_label}")

        # Pairwise t-tests only if ANOVA is significant
        if p_value < SIGNIFICANCE_LEVEL:
            console.print("\n  [bold]Pairwise t-tests:[/bold]")
            for i in range(len(dim_keys)):
                for j in range(i + 1, len(dim_keys)):
                    d1, d2 = dim_keys[i], dim_keys[j]
                    t_stat, t_p = stats.ttest_ind(dim_groups[d1], dim_groups[d2])
                    if t_p < SIGNIFICANCE_LEVEL:
                        t_sig = f"[green]significant (p={t_p:.4f})[/green]"
                    else:
                        t_sig = f"[dim]not significant (p={t_p:.4f})[/dim]"
                    console.print(f"    {d1} vs {d2}: t={t_stat:.3f}, {t_sig}")


@click.command()
@click.option(
    "--metric", type=click.Choice(["pass_at_1", "compilation_at_1"]),
    default="pass_at_1", show_default=True,
    help="Metric to analyze.",
)
@click.option(
    "--root", type=click.Path(exists=True, path_type=Path),
    default=None,
    help="Override results root directory.",
)
def main(metric: str, root: Path | None):
    """Analyze statistical significance of multi-run experiments."""
    console = Console()
    results_root = root or RESULTS_ROOT

    console.print(f"[bold]Scanning: {results_root}[/bold]")
    results = discover_results(results_root)

    if not results:
        console.print("[red]No mlflow_results.json files found.[/red]")
        sys.exit(1)

    console.print(f"Found [bold]{len(results)}[/bold] result files.\n")

    grouped = group_results(results, metric)
    run_analysis(grouped, metric)

    console.print("\n[dim]Done.[/dim]")


if __name__ == "__main__":
    main()
