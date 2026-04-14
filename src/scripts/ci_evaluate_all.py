"""CI evaluation script: evaluate all HumanEval-X experiment folders.

Discovers experiment folders under data/translation/humaneval-x/,
runs Docker-based evaluation on each, and produces:
  - Per-experiment JSON results
  - Markdown comparison table
  - Line plot (PNG)
  - GitHub Actions Job Summary (if $GITHUB_STEP_SUMMARY is set)

Usage:
    uv run python src/scripts/ci_evaluate_all.py
"""

from __future__ import annotations

import argparse
import json
import os
import sys
from pathlib import Path
from concurrent.futures import ThreadPoolExecutor, as_completed

# Ensure the project root is on sys.path
PROJECT_ROOT = Path(__file__).resolve().parent.parent.parent
sys.path.insert(0, str(PROJECT_ROOT))

from src.core.docker_eval import (
    check_docker_available,
    ensure_go_image,
    ensure_go_mod_cache,
    evaluate_single_task,
    prepare_evaluation_sources,
    DEFAULT_GO_IMAGE,
)
from src.core.humaneval_artifacts import HumanEvalRunPaths, is_humaneval_run_root, parse_humaneval_run_root, write_json, write_text
from src.core.reporting import compute_summary
from src.core.schemas import EvaluationRecord
from src.config import HUMANEVAL_X_DIR, load_eval_config

from rich.console import Console

console = Console()

# ---------------------------------------------------------------------------
# Discovery
# ---------------------------------------------------------------------------

def discover_experiment_dirs(root: Path) -> list[tuple[str, str, str, Path]]:
    """Find all experiment run roots: (provider, model, strategy[/backend], path)."""
    experiments: list[tuple[str, str, str, Path]] = []
    if not root.is_dir():
        return experiments

    for run_root in sorted(path for path in root.rglob("*") if is_humaneval_run_root(path)):
        relative_parts = run_root.relative_to(root).parts
        if any(part.startswith(".") for part in relative_parts):
            continue
        provider, model, experiment, backend, run_id = parse_humaneval_run_root(run_root)
        strategy_name = "baseline"
        if experiment != "baseline":
            strategy_name = f"{backend}/run-{run_id}/{experiment}"
        elif run_id is not None:
            strategy_name = f"baseline/run-{run_id}"
        experiments.append((provider, model, strategy_name, run_root))

    return experiments


# ---------------------------------------------------------------------------
# Evaluation
# ---------------------------------------------------------------------------

def evaluate_experiment(
    target_dir: Path,
    pairs: list[dict],
    batch_size: int = 10,
    timeout: int = 60,
) -> list[EvaluationRecord]:
    """Evaluate all Go files in a single experiment directory."""
    task_lookup: dict[str, dict] = {}
    for pair in pairs:
        task_num = pair["task_id"].split("/")[1]
        task_lookup[task_num] = pair

    run_paths = HumanEvalRunPaths(target_dir)
    task_paths = run_paths.iter_task_dirs()
    if not task_paths:
        return []

    work_items: list[tuple[object, dict]] = []
    for task_path in task_paths:
        task_num = task_path.task_name.replace("Go_", "")
        pair = task_lookup.get(task_num)
        if pair is not None:
            work_items.append((task_path, pair))

    records: list[EvaluationRecord] = []

    def _eval_one(task_path, pair: dict) -> EvaluationRecord:
        generated_code = task_path.translation_go.read_text(encoding="utf-8")
        solution_source, test_source = prepare_evaluation_sources(generated_code, pair["test"])
        write_text(task_path.evaluation_solution_go, solution_source)
        write_text(task_path.evaluation_test_go, test_source)
        record = evaluate_single_task(
            task_id=pair["task_id"],
            generated_code=generated_code,
            test_code=pair["test"],
            timeout=timeout,
        )
        write_json(task_path.evaluation_result_json, record.model_dump())
        return record

    with ThreadPoolExecutor(max_workers=batch_size) as executor:
        futures = {
            executor.submit(_eval_one, task_path, p): (task_path, p)
            for task_path, p in work_items
        }
        for future in as_completed(futures):
            task_path, p = futures[future]
            try:
                record = future.result()
                status = "PASS" if record.pass_at_1 else (
                    "COMPILE" if record.compiles else "FAIL"
                )
                console.print(f"  [{status}] {task_path.task_name}")
                records.append(record)
            except Exception as exc:
                console.print(f"  [ERROR] {task_path.task_name}: {exc}")
                records.append(EvaluationRecord(
                    source_file=p["task_id"],
                    target_file=task_path.task_name,
                    dataset="humaneval-x",
                    notes=f"Worker error: {exc}",
                ))

    return records


# ---------------------------------------------------------------------------
# Output generation
# ---------------------------------------------------------------------------

def build_comparison_table(
    all_results: list[dict],
) -> str:
    """Build a Markdown comparison table from all experiment summaries."""
    lines = [
        "# HumanEval-X Evaluation Comparison",
        "",
        "| Provider | Model | Strategy | Total | Compilation@1 | Pass@1 |",
        "|----------|-------|----------|------:|--------------:|-------:|",
    ]
    for r in all_results:
        s = r["summary"]
        lines.append(
            f"| {r['provider']} | {r['model']} | {r['strategy']} "
            f"| {s['total_files']} "
            f"| {s['compilation_at_1']:.1%} "
            f"| {s['pass_at_1']:.1%} |"
        )
    lines.append("")
    return "\n".join(lines)


def generate_plot(all_results: list[dict], output_path: Path) -> None:
    """Generate a comparison bar chart of Compilation@1 and Pass@1."""
    try:
        import matplotlib
        matplotlib.use("Agg")
        import matplotlib.pyplot as plt
    except ImportError:
        console.print("[yellow]matplotlib not installed, skipping plot.[/yellow]")
        return

    labels = [f"{r['provider']}/{r['model']}/{r['strategy']}" for r in all_results]
    compilation = [r["summary"]["compilation_at_1"] * 100 for r in all_results]
    pass_at_1 = [r["summary"]["pass_at_1"] * 100 for r in all_results]

    x = range(len(labels))
    bar_width = 0.35

    fig, ax = plt.subplots(figsize=(max(10, len(labels) * 2), 6))

    bars1 = ax.bar(
        [i - bar_width / 2 for i in x], compilation, bar_width,
        label="Compilation@1 (%)", color="#4A90D9", alpha=0.85,
    )
    bars2 = ax.bar(
        [i + bar_width / 2 for i in x], pass_at_1, bar_width,
        label="Pass@1 (%)", color="#2ECC71", alpha=0.85,
    )

    # Add value labels on bars
    for bar in bars1:
        height = bar.get_height()
        ax.annotate(
            f"{height:.1f}%",
            xy=(bar.get_x() + bar.get_width() / 2, height),
            xytext=(0, 3), textcoords="offset points",
            ha="center", va="bottom", fontsize=8,
        )
    for bar in bars2:
        height = bar.get_height()
        ax.annotate(
            f"{height:.1f}%",
            xy=(bar.get_x() + bar.get_width() / 2, height),
            xytext=(0, 3), textcoords="offset points",
            ha="center", va="bottom", fontsize=8,
        )

    ax.set_xlabel("Experiment")
    ax.set_ylabel("Rate (%)")
    ax.set_title("HumanEval-X Evaluation Results")
    ax.set_xticks(list(x))
    ax.set_xticklabels(labels, rotation=30, ha="right", fontsize=9)
    ax.set_ylim(0, 110)
    ax.legend()
    ax.grid(axis="y", alpha=0.3)

    plt.tight_layout()
    plt.savefig(output_path, dpi=150)
    plt.close()
    console.print(f"[green]Plot saved:[/green] {output_path}")


# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

def main() -> None:
    parser = argparse.ArgumentParser(description="Evaluate all HumanEval-X experiments")
    parser.add_argument("--batch-size", type=int, default=None, help="Override parallel batch size")
    parser.add_argument("--root", type=str, default=None, help="Root directory to search for experiments (default: HUMANEVAL_X_DIR)")
    parser.add_argument("--output-dir", type=str, default=None, help="Directory to save results (default: results/)")
    args = parser.parse_args()

    # Load HumanEval-X dataset
    console.print("[dim]Loading HumanEval-X dataset...[/dim]")
    from src.data.humaneval_x import load_humaneval_x
    pairs = load_humaneval_x()
    console.print(f"Loaded [bold]{len(pairs)}[/bold] HumanEval-X problems.\n")

    # Load eval config
    eval_config = load_eval_config()
    batch_size = args.batch_size if args.batch_size is not None else eval_config["parallel"]["batch_size"]
    timeout = eval_config["docker"]["timeout"]

    # Pre-flight: Docker
    if not check_docker_available(console):
        console.print("[red]Docker not available. Exiting.[/red]")
        sys.exit(1)
    if not ensure_go_image(console, DEFAULT_GO_IMAGE):
        console.print("[red]Failed to pull Go image. Exiting.[/red]")
        sys.exit(1)

    console.print("[dim]Ensuring Go module cache (testify)...[/dim]")
    if not ensure_go_mod_cache(DEFAULT_GO_IMAGE):
        console.print("[red]Failed to download Go modules. Exiting.[/red]")
        sys.exit(1)
    console.print("[green]OK[/green]   Go module cache ready\n")

    # Discover experiments
    search_root = Path(args.root) if args.root else HUMANEVAL_X_DIR
    experiments = discover_experiment_dirs(search_root)
    if not experiments:
        console.print("[yellow]No experiment folders found under "
                       f"{search_root}[/yellow]")
        sys.exit(0)

    console.print(f"Found [bold]{len(experiments)}[/bold] experiment(s):\n")
    for provider, model, strategy, path in experiments:
        go_count = len([task for task in HumanEvalRunPaths(path).iter_task_dirs() if task.translation_go.exists()])
        console.print(f"  • {provider}/{model}/{strategy} ({go_count} files)")
    console.print()

    # Prepare output directory
    results_dir = Path(args.output_dir) if args.output_dir else PROJECT_ROOT / "results"
    results_dir.mkdir(parents=True, exist_ok=True)

    # Evaluate each experiment
    all_results: list[dict] = []

    for provider, model, strategy, target_dir in experiments:
        label = f"{provider}/{model}/{strategy}"
        console.print(f"\n[bold blue]── Evaluating: {label} ──[/bold blue]")

        records = evaluate_experiment(
            target_dir, pairs,
            batch_size=batch_size, timeout=timeout,
        )
        summary = compute_summary(records, dataset="humaneval-x")

        result_entry = {
            "provider": provider,
            "model": model,
            "strategy": strategy,
            "summary": summary,
            "records": [r.model_dump() for r in records],
        }
        all_results.append(result_entry)

        # Save per-experiment JSON
        json_name = f"{provider}__{model}__{strategy.replace('/', '__')}.json"
        json_path = results_dir / json_name
        with open(json_path, "w", encoding="utf-8") as f:
            json.dump(result_entry, f, indent=2, default=str)
        console.print(f"  → Saved: {json_path}")

        # Print summary for this experiment
        console.print(
            f"  Compilation@1: {summary['compilation_at_1']:.1%}  "
            f"Pass@1: {summary['pass_at_1']:.1%}"
        )

    # Build comparison table
    console.print("\n[bold]── Comparison ──[/bold]\n")
    table_md = build_comparison_table(all_results)
    console.print(table_md)

    comparison_md_path = results_dir / "comparison.md"
    comparison_md_path.write_text(table_md, encoding="utf-8")
    console.print(f"\n[green]Table saved:[/green] {comparison_md_path}")

    # Generate plot
    plot_path = results_dir / "comparison.png"
    generate_plot(all_results, plot_path)

    # Write GitHub Actions Job Summary
    summary_path = os.environ.get("GITHUB_STEP_SUMMARY")
    if summary_path:
        with open(summary_path, "a", encoding="utf-8") as f:
            f.write(table_md)
            f.write("\n\n")
            f.write("![Comparison Plot](comparison.png)\n")
        console.print("[green]Job summary written.[/green]")

    console.print("\n[bold green]Done![/bold green]")


if __name__ == "__main__":
    main()
