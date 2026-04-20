"""Helpers for the CodeNet discovery run layout."""

from __future__ import annotations

import json
from dataclasses import dataclass
from pathlib import Path

from src.config import TRANSLATION_TARGET_DIR

CODENET_TARGET_DIR = TRANSLATION_TARGET_DIR / "codenet"


@dataclass(frozen=True)
class CodeNetTaskPaths:
    """Filesystem paths for one CodeNet translation/evaluation bundle."""

    run_root: Path
    problem_id: str

    @property
    def task_dir(self) -> Path:
        return self.run_root / "tasks" / self.problem_id

    @property
    def metadata_json(self) -> Path:
        return self.task_dir / "metadata.json"

    @property
    def source_py(self) -> Path:
        return self.task_dir / "source.py"

    @property
    def reference_go(self) -> Path:
        return self.task_dir / "reference.go"

    @property
    def prompt_json(self) -> Path:
        return self.task_dir / "prompt.json"

    @property
    def llm_raw_json(self) -> Path:
        return self.task_dir / "llm_raw.json"

    @property
    def translation_go(self) -> Path:
        return self.task_dir / "translation.go"

    @property
    def translation_result_json(self) -> Path:
        return self.task_dir / "translation_result.json"

    @property
    def evaluation_dir(self) -> Path:
        return self.task_dir / "evaluation"

    @property
    def evaluation_solution_go(self) -> Path:
        return self.evaluation_dir / "solution.go"

    @property
    def evaluation_input_txt(self) -> Path:
        return self.evaluation_dir / "input.txt"

    @property
    def evaluation_expected_output_txt(self) -> Path:
        return self.evaluation_dir / "expected_output.txt"

    @property
    def evaluation_stdout_txt(self) -> Path:
        return self.evaluation_dir / "stdout.txt"

    @property
    def evaluation_stderr_txt(self) -> Path:
        return self.evaluation_dir / "stderr.txt"

    @property
    def evaluation_result_json(self) -> Path:
        return self.evaluation_dir / "result.json"


@dataclass(frozen=True)
class CodeNetRunPaths:
    """Filesystem paths for a CodeNet discovery run."""

    run_root: Path

    @property
    def manifest_json(self) -> Path:
        return self.run_root / "manifest.json"

    @property
    def tasks_dir(self) -> Path:
        return self.run_root / "tasks"

    @property
    def translation_summary_json(self) -> Path:
        return self.run_root / "translation_summary.json"

    @property
    def evaluation_results_dir(self) -> Path:
        return self.run_root / "evaluation" / "results"

    @property
    def control_dir(self) -> Path:
        return self.run_root / "control"

    @property
    def control_state_json(self) -> Path:
        return self.control_dir / "state.json"

    @property
    def control_ready_jsonl(self) -> Path:
        return self.control_dir / "ready.jsonl"

    @property
    def control_done_flag(self) -> Path:
        return self.control_dir / "done.flag"

    @property
    def evaluation_per_task_jsonl(self) -> Path:
        return self.evaluation_results_dir / "per_task.jsonl"

    @property
    def evaluation_summary_json(self) -> Path:
        return self.evaluation_results_dir / "summary.json"

    def task(self, problem_id: str) -> CodeNetTaskPaths:
        return CodeNetTaskPaths(self.run_root, problem_id)

    def ensure_translation_dirs(self) -> None:
        self.tasks_dir.mkdir(parents=True, exist_ok=True)
        self.control_dir.mkdir(parents=True, exist_ok=True)

    def ensure_evaluation_dirs(self) -> None:
        self.evaluation_results_dir.mkdir(parents=True, exist_ok=True)

    def iter_task_dirs(self) -> list[CodeNetTaskPaths]:
        if not self.tasks_dir.exists():
            return []
        paths: list[CodeNetTaskPaths] = []
        for task_dir in sorted(self.tasks_dir.iterdir()):
            if task_dir.is_dir():
                paths.append(CodeNetTaskPaths(self.run_root, task_dir.name))
        return paths


def codenet_run_root(
    provider: str,
    variant: str,
    experiment: str,
    run_name: str,
    *,
    root: Path = CODENET_TARGET_DIR,
) -> Path:
    """Return the canonical CodeNet discovery run root."""
    return root / provider / variant / experiment / run_name


def append_jsonl_record(path: Path, payload: object) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("a", encoding="utf-8") as fh:
        fh.write(json.dumps(payload, sort_keys=True) + "\n")
