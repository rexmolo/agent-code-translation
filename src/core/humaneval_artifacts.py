"""Helpers for the HumanEval-X run bundle layout."""

from __future__ import annotations

import json
from dataclasses import dataclass
from pathlib import Path
from typing import Iterable


def base_experiment_name(experiment: str) -> str:
    """Return the canonical experiment preset name."""
    return experiment


def is_baseline_experiment(experiment: str) -> bool:
    """Return True for the no-RAG baseline experiment."""
    return experiment == "baseline"


@dataclass(frozen=True)
class HumanEvalTaskPaths:
    """Filesystem paths for a single HumanEval-X task bundle."""

    run_root: Path
    task_name: str

    @property
    def task_dir(self) -> Path:
        return self.run_root / "tasks" / self.task_name

    @property
    def prompt_json(self) -> Path:
        return self.task_dir / "prompt.json"

    @property
    def retrieval_json(self) -> Path:
        return self.task_dir / "retrieval.json"

    @property
    def llm_raw_json(self) -> Path:
        return self.task_dir / "llm_raw.json"

    @property
    def translation_go(self) -> Path:
        return self.task_dir / "translation.go"

    @property
    def evaluation_dir(self) -> Path:
        return self.task_dir / "evaluation"

    @property
    def evaluation_solution_go(self) -> Path:
        return self.evaluation_dir / "solution.go"

    @property
    def evaluation_test_go(self) -> Path:
        return self.evaluation_dir / "test.go"

    @property
    def evaluation_result_json(self) -> Path:
        return self.evaluation_dir / "result.json"


@dataclass(frozen=True)
class HumanEvalRunPaths:
    """Filesystem paths for a HumanEval-X run bundle."""

    run_root: Path

    @property
    def manifest_json(self) -> Path:
        return self.run_root / "manifest.json"

    @property
    def tasks_dir(self) -> Path:
        return self.run_root / "tasks"

    @property
    def evaluation_results_dir(self) -> Path:
        return self.run_root / "evaluation" / "results"

    @property
    def per_task_jsonl(self) -> Path:
        return self.evaluation_results_dir / "per_task.jsonl"

    @property
    def summary_json(self) -> Path:
        return self.evaluation_results_dir / "summary.json"

    @property
    def diagnostics_dir(self) -> Path:
        return self.run_root / "diagnostics"

    def task(self, task_num: str | int) -> HumanEvalTaskPaths:
        return HumanEvalTaskPaths(self.run_root, f"Go_{task_num}")

    def ensure_translation_dirs(self) -> None:
        self.tasks_dir.mkdir(parents=True, exist_ok=True)

    def ensure_evaluation_dirs(self) -> None:
        self.evaluation_results_dir.mkdir(parents=True, exist_ok=True)

    def iter_task_dirs(self) -> list[HumanEvalTaskPaths]:
        if not self.tasks_dir.exists():
            return []
        task_paths = []
        for task_dir in sorted(self.tasks_dir.iterdir(), key=_task_sort_key):
            if task_dir.is_dir() and task_dir.name.startswith("Go_"):
                task_paths.append(HumanEvalTaskPaths(self.run_root, task_dir.name))
        return task_paths


def _task_sort_key(path: Path) -> tuple[int, str]:
    name = path.name
    if name.startswith("Go_"):
        suffix = name.split("_", 1)[1]
        if suffix.isdigit():
            return int(suffix), name
    return 10**9, name


def humaneval_run_root(
    root: Path,
    provider: str,
    variant: str,
    experiment: str,
    backend_label: str | None,
    run_id: int | None,
) -> Path:
    """Return the canonical HumanEval-X run root for a configuration."""

    if is_baseline_experiment(experiment):
        base = root / provider / variant / "baseline"
        return base / f"run-{run_id}" if run_id is not None else base

    if backend_label is None:
        raise ValueError("backend_label is required for RAG experiments")

    base = root / provider / variant / backend_label
    return base / f"run-{run_id}" / experiment if run_id is not None else base / experiment


def parse_humaneval_run_root(run_root: Path) -> tuple[str, str, str, str | None, int | None]:
    """Parse a HumanEval-X run root into provider, variant, experiment, backend, run_id."""

    parts = run_root.parts
    last = parts[-1]

    if last.startswith("run-"):
        run_id = int(last.split("-", 1)[1])
        parent = parts[-2]
        return parts[-4], parts[-3], parent, None, run_id

    parent = parts[-2]
    if parent.startswith("run-"):
        run_id = int(parent.split("-", 1)[1])
        backend = parts[-3]
        if backend.startswith("vec-"):
            return parts[-5], parts[-4], last, backend, run_id
        return parts[-4], parts[-3], last, None, run_id

    if parent.startswith("vec-"):
        return parts[-4], parts[-3], last, parent, None
    return parts[-3], parts[-2], last, None, None


def is_humaneval_run_root(path: Path) -> bool:
    """Return True if the path looks like a HumanEval-X run root."""

    if not path.is_dir():
        return False
    tasks_dir = path / "tasks"
    return tasks_dir.is_dir() and any(
        child.is_dir() and child.name.startswith("Go_")
        for child in tasks_dir.iterdir()
    )


def write_json(path: Path, payload: object) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2, sort_keys=True), encoding="utf-8")


def write_text(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")


def append_jsonl(path: Path, rows: Iterable[object]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", encoding="utf-8") as fh:
        for row in rows:
            fh.write(json.dumps(row, sort_keys=True) + "\n")
