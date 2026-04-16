#!/usr/bin/env python3
"""Compare task-bundle runs and surface baseline-pass / RAG-fail regressions.

Expected run layout:
  <run-root>/
    tasks/
      Go_<id>/
        prompt.json
        retrieval.json
        llm_raw.json
        translation.go
        evaluation/
          solution.go
          test.go
          result.json
"""

from __future__ import annotations

import argparse
import json
import re
import shutil
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Any

from rich.console import Console
from rich.table import Table

from src.rag.schema import RAG_SOURCES

console = Console()


@dataclass(frozen=True)
class TaskBundle:
    task_id: str
    task_dir: Path
    pass_at_1: bool
    result_path: Path
    prompt_path: Path
    retrieval_path: Path
    raw_path: Path
    translation_path: Path
    solution_path: Path
    test_path: Path


def _task_sort_key(task_id: str) -> tuple[int, str]:
    match = re.search(r"(\d+)$", task_id)
    return (int(match.group(1)), task_id) if match else (sys.maxsize, task_id)


def _safe_slug(path: Path) -> str:
    return re.sub(r"[^A-Za-z0-9._-]+", "_", str(path).strip("/")) or "run"


def _read_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


def _extract_pass_at_1(payload: dict[str, Any]) -> bool:
    value = payload.get("pass_at_1")
    if isinstance(value, bool):
        return value
    if isinstance(value, (int, float)):
        return bool(value)

    summary = payload.get("summary")
    if isinstance(summary, dict):
        nested = summary.get("pass_at_1")
        if isinstance(nested, bool):
            return nested
        if isinstance(nested, (int, float)):
            return bool(nested)

    raise ValueError("result.json does not contain a usable pass_at_1 field")


def load_task_bundles(run_root: Path) -> dict[str, TaskBundle]:
    tasks_dir = run_root / "tasks"
    if not tasks_dir.is_dir():
        raise FileNotFoundError(f"tasks directory not found: {tasks_dir}")

    bundles: dict[str, TaskBundle] = {}
    for task_dir in sorted(tasks_dir.glob("Go_*"), key=lambda p: _task_sort_key(p.name)):
        result_path = task_dir / "evaluation" / "result.json"
        if not result_path.is_file():
            continue

        payload = _read_json(result_path)
        bundles[task_dir.name] = TaskBundle(
            task_id=task_dir.name,
            task_dir=task_dir,
            pass_at_1=_extract_pass_at_1(payload),
            result_path=result_path,
            prompt_path=task_dir / "prompt.json",
            retrieval_path=task_dir / "retrieval.json",
            raw_path=task_dir / "llm_raw.json",
            translation_path=task_dir / "translation.go",
            solution_path=task_dir / "evaluation" / "solution.go",
            test_path=task_dir / "evaluation" / "test.go",
        )
    return bundles


def _source_counts(retrieval_path: Path) -> dict[str, int]:
    if not retrieval_path.is_file():
        return {}

    payload = _read_json(retrieval_path)
    if isinstance(payload.get("retrieval_counts"), dict):
        counts = {
            key: value
            for key, value in payload["retrieval_counts"].items()
            if isinstance(value, int)
        }
        if counts:
            return counts

    items = payload.get("items")
    if isinstance(items, dict):
        counts = {
            key: len(value)
            for key, value in items.items()
            if isinstance(value, list)
        }
        if counts:
            return counts

    legacy_counts = {}
    for key in RAG_SOURCES:
        value = payload.get(key)
        if isinstance(value, list):
            legacy_counts[key] = len(value)
    return legacy_counts


def analyze_regressions(
    baseline_run: Path,
    rag_runs: list[Path],
    *,
    task_filters: set[str] | None = None,
    limit: int | None = None,
) -> dict[str, Any]:
    baseline = load_task_bundles(baseline_run)
    requested = set(task_filters or ())

    run_summaries = []
    for rag_run in rag_runs:
        rag = load_task_bundles(rag_run)
        candidate_ids = sorted(set(baseline) & set(rag), key=_task_sort_key)
        if requested:
            candidate_ids = [task_id for task_id in candidate_ids if task_id in requested]

        regressions = []
        for task_id in candidate_ids:
            baseline_task = baseline[task_id]
            rag_task = rag[task_id]
            if baseline_task.pass_at_1 and not rag_task.pass_at_1:
                regressions.append(
                    {
                        "task_id": task_id,
                        "source_counts": _source_counts(rag_task.retrieval_path),
                        "artifacts": {
                            "prompt": str(rag_task.prompt_path),
                            "retrieval": str(rag_task.retrieval_path),
                            "llm_raw": str(rag_task.raw_path),
                            "translation": str(rag_task.translation_path),
                            "evaluation_solution": str(rag_task.solution_path),
                            "evaluation_test": str(rag_task.test_path),
                            "evaluation_result": str(rag_task.result_path),
                        },
                    }
                )

        if limit is not None:
            regressions = regressions[:limit]

        run_summaries.append(
            {
                "rag_run": str(rag_run),
                "baseline_run": str(baseline_run),
                "comparable_tasks": len(candidate_ids),
                "regressions": regressions,
            }
        )

    return {
        "baseline_run": str(baseline_run),
        "task_filter": sorted(requested),
        "limit": limit,
        "runs": run_summaries,
    }


def _copy_if_exists(source: Path, destination: Path) -> bool:
    if not source.is_file():
        return False
    destination.parent.mkdir(parents=True, exist_ok=True)
    shutil.copy2(source, destination)
    return True


def _format_source_counts(source_counts: dict[str, int]) -> str:
    if not source_counts:
        return "none"
    return ", ".join(
        f"{key}={value}"
        for key, value in sorted(source_counts.items())
    )


def write_reports(summary: dict[str, Any], output_dir: Path) -> None:
    output_dir.mkdir(parents=True, exist_ok=True)
    (output_dir / "summary.json").write_text(json.dumps(summary, indent=2), encoding="utf-8")

    lines = [
        "# RAG Regression Summary",
        "",
        f"- Baseline run: `{summary['baseline_run']}`",
        f"- Task filter: `{', '.join(summary['task_filter']) or 'all'}`",
        "",
    ]

    for run in summary["runs"]:
        lines.append(f"## `{run['rag_run']}`")
        lines.append("")
        lines.append(f"- Comparable tasks: {run['comparable_tasks']}")
        lines.append(f"- Regressions: {len(run['regressions'])}")
        lines.append("")

        run_dir = output_dir / _safe_slug(Path(run["rag_run"]))
        for case in run["regressions"]:
            task_id = case["task_id"]
            lines.append(
                f"- `{task_id}`: source counts = {_format_source_counts(case['source_counts'])}"
            )

            case_dir = run_dir / task_id
            artifacts = {
                "prompt.json": Path(case["artifacts"]["prompt"]),
                "retrieval.json": Path(case["artifacts"]["retrieval"]),
                "llm_raw.json": Path(case["artifacts"]["llm_raw"]),
                "translation.go": Path(case["artifacts"]["translation"]),
                "solution.go": Path(case["artifacts"]["evaluation_solution"]),
                "test.go": Path(case["artifacts"]["evaluation_test"]),
                "result.json": Path(case["artifacts"]["evaluation_result"]),
            }
            copied = []
            for filename, source in artifacts.items():
                if _copy_if_exists(source, case_dir / filename):
                    copied.append(filename)
            (case_dir / "case_summary.json").write_text(
                json.dumps(
                    {
                        "task_id": task_id,
                        "source_counts": case["source_counts"],
                        "copied_artifacts": copied,
                    },
                    indent=2,
                ),
                encoding="utf-8",
            )

        lines.append("")

    (output_dir / "summary.md").write_text("\n".join(lines), encoding="utf-8")


def print_summary(summary: dict[str, Any]) -> None:
    table = Table(title="Baseline-Pass / RAG-Fail Regressions")
    table.add_column("RAG Run", style="cyan")
    table.add_column("Comparable", justify="right")
    table.add_column("Regressions", justify="right")
    table.add_column("Parallel Corpus", justify="right")

    for run in summary["runs"]:
        parallel_cases = sum(1 for case in run["regressions"] if case["parallel_corpus_hits"] > 0)
        table.add_row(
            run["rag_run"],
            str(run["comparable_tasks"]),
            str(len(run["regressions"])),
            str(parallel_cases),
        )

    console.print(table)

    for run in summary["runs"]:
        if not run["regressions"]:
            console.print(f"[green]No regressions found for[/green] {run['rag_run']}")
            continue

        console.print(f"\n[bold]{run['rag_run']}[/bold]")
        for case in run["regressions"]:
            suffix = ""
            if case["parallel_corpus_hits"] > 0:
                suffix = f" [yellow](parallel corpus hits: {case['parallel_corpus_hits']})[/yellow]"
            console.print(f"  - {case['task_id']}{suffix}")


def parse_args(argv: list[str] | None = None) -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--baseline-run", required=True, type=Path)
    parser.add_argument("--rag-run", required=True, type=Path, action="append")
    parser.add_argument("--task", dest="tasks", action="append", default=[])
    parser.add_argument("--limit", type=int, default=None)
    parser.add_argument("--output-dir", type=Path, default=None)
    return parser.parse_args(argv)


def main(argv: list[str] | None = None) -> int:
    args = parse_args(argv)
    summary = analyze_regressions(
        args.baseline_run,
        args.rag_run,
        task_filters=set(args.tasks),
        limit=args.limit,
    )
    print_summary(summary)
    if args.output_dir is not None:
        write_reports(summary, args.output_dir)
        console.print(f"\n[green]Wrote diagnostics to[/green] {args.output_dir}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
