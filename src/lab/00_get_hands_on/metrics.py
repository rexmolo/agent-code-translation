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
            "compilation_success_rate": 0.0,
            "successful_translation_rate": 0.0,
            "computational_accuracy": 0.0,
            "io_equivalence_rate": 0.0,
            "test_pass_rate": 0.0,
        }

    # Test pass rate: average of per-file test pass rates (only files with tests)
    files_with_tests = [r for r in records if r.tests_total > 0]
    avg_test_pass_rate = (
        sum(r.test_pass_rate for r in files_with_tests) / len(files_with_tests)
        if files_with_tests
        else 0.0
    )

    return {
        "total_files": total,
        "compilation_success_rate": sum(r.compiles for r in records) / total,
        "successful_translation_rate": sum(
            r.compiles and r.runs_successfully for r in records
        )
        / total,
        "computational_accuracy": sum(r.computational_accuracy for r in records)
        / total,
        "io_equivalence_rate": sum(r.io_equivalent for r in records) / total,
        "test_pass_rate": avg_test_pass_rate,
    }


def display_summary_table(summary: dict) -> None:
    """Print a Rich table of aggregate evaluation metrics."""
    console = Console()
    table = Table(title="Translation Evaluation Metrics")
    table.add_column("Metric", style="cyan", no_wrap=True)
    table.add_column("Value", justify="right", style="green")

    table.add_row("Total Files", str(summary["total_files"]))
    table.add_row(
        "Compilation Success Rate",
        f"{summary['compilation_success_rate']:.1%}",
    )
    table.add_row(
        "Successful Translation Rate",
        f"{summary['successful_translation_rate']:.1%}",
    )
    table.add_row(
        "Computational Accuracy (CA)",
        f"{summary['computational_accuracy']:.1%}",
    )
    table.add_row(
        "I/O Equivalence Rate",
        f"{summary['io_equivalence_rate']:.1%}",
    )
    table.add_row(
        "Test Pass Rate",
        f"{summary['test_pass_rate']:.1%}",
    )
    console.print(table)


def display_per_file_table(records: list[EvaluationRecord]) -> None:
    """Print per-file evaluation results as a Rich table."""
    console = Console()
    table = Table(title="Per-File Evaluation Results")
    table.add_column("Source File", style="cyan")
    table.add_column("Compiles", justify="center")
    table.add_column("Runs", justify="center")
    table.add_column("CA", justify="center")
    table.add_column("I/O Eq", justify="center")
    table.add_column("Tests", justify="center")
    table.add_column("Notes", style="dim", max_width=40)

    def check(v: bool) -> str:
        return "[green]Y[/green]" if v else "[red]N[/red]"

    for r in records:
        if r.tests_total > 0:
            tests_str = f"[green]{r.tests_passed}[/green]/{r.tests_total}"
        else:
            tests_str = "[dim]-[/dim]"
        table.add_row(
            Path(r.source_file).name,
            check(r.compiles),
            check(r.runs_successfully),
            check(r.computational_accuracy),
            check(r.io_equivalent),
            tests_str,
            r.notes[:40] if r.notes else "",
        )
    console.print(table)
