"""Aggregate metrics computation and Rich table display."""

import importlib
from pathlib import Path

from rich.console import Console
from rich.table import Table

_models = importlib.import_module("src.lab.00_get_hands_on.models")
EvaluationRecord = _models.EvaluationRecord


def compute_summary(records: list[EvaluationRecord]) -> dict:
    """Compute aggregate metrics from individual evaluation records."""
    total = len(records)
    if total == 0:
        return {
            "total_files": 0,
            "compilation_at_1": 0.0,
            "runs_rate": 0.0,
            "pass_at_1": 0.0,
            "avg_ast_similarity": 0.0,
        }

    has_ast = any(r.ast_similarity > 0 for r in records)

    return {
        "total_files": total,
        "compilation_at_1": sum(r.compiles for r in records) / total,
        "runs_rate": sum(r.runs_successfully for r in records) / total,
        "pass_at_1": sum(r.pass_at_1 for r in records) / total,
        "avg_ast_similarity": (
            sum(r.ast_similarity for r in records) / total if has_ast else 0.0
        ),
    }


def display_summary_table(summary: dict) -> None:
    """Print a Rich table of aggregate evaluation metrics."""
    console = Console()
    table = Table(title="Translation Evaluation Metrics")
    table.add_column("Metric", style="cyan", no_wrap=True)
    table.add_column("Value", justify="right", style="green")

    table.add_row("Total Files", str(summary["total_files"]))
    table.add_row(
        "Compilation@1",
        f"{summary['compilation_at_1']:.1%}",
    )
    table.add_row(
        "Runs Successfully",
        f"{summary['runs_rate']:.1%}",
    )
    table.add_row(
        "Pass@1",
        f"{summary['pass_at_1']:.1%}",
    )
    if summary["avg_ast_similarity"] > 0:
        table.add_row(
            "Match_ast (avg)",
            f"{summary['avg_ast_similarity']:.3f}",
        )
    console.print(table)


def display_per_file_table(records: list[EvaluationRecord]) -> None:
    """Print per-file evaluation results as a Rich table."""
    console = Console()
    table = Table(title="Per-File Evaluation Results")
    table.add_column("Source", style="cyan")
    table.add_column("Compiles", justify="center")
    table.add_column("Runs", justify="center")
    table.add_column("Pass@1", justify="center")

    has_ast = any(r.ast_similarity > 0 for r in records)
    if has_ast:
        table.add_column("AST", justify="center")

    table.add_column("Tests", justify="center")
    table.add_column("Notes", style="dim", max_width=40)

    def check(v: bool) -> str:
        return "[green]Y[/green]" if v else "[red]N[/red]"

    for r in records:
        if r.tests_total > 0:
            tests_str = f"[green]{r.tests_passed}[/green]/{r.tests_total}"
        else:
            tests_str = "[dim]-[/dim]"

        row = [
            Path(r.source_file).name,
            check(r.compiles),
            check(r.runs_successfully),
            check(r.pass_at_1),
        ]
        if has_ast:
            row.append(f"{r.ast_similarity:.2f}" if r.ast_similarity > 0 else "[dim]-[/dim]")
        row.append(tests_str)
        row.append(r.notes[:40] if r.notes else "")
        table.add_row(*row)

    console.print(table)
