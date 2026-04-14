"""Tests for ci_evaluate_all.py (no Docker, no network).

    uv run pytest src/tests/test_ci_evaluate_all.py -v
"""

from pathlib import Path

import pytest

from src.core.humaneval_artifacts import HumanEvalRunPaths
from src.scripts.ci_evaluate_all import discover_experiment_dirs, build_comparison_table


# ---------------------------------------------------------------------------
# discover_experiment_dirs
# ---------------------------------------------------------------------------

class TestDiscoverExperimentDirs:
    @staticmethod
    def _make_task_bundle(run_root: Path, task_ids: tuple[int, ...] = (0, 1)) -> None:
        run_paths = HumanEvalRunPaths(run_root)
        run_paths.ensure_translation_dirs()
        for task_id in task_ids:
            task = run_paths.task(task_id)
            task.translation_go.parent.mkdir(parents=True, exist_ok=True)
            task.translation_go.write_text("package main\n", encoding="utf-8")

    def test_finds_baseline_run_roots(self, tmp_path):
        d = tmp_path / "openai" / "gpt-4" / "baseline" / "run-1"
        self._make_task_bundle(d)

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 1
        provider, model, strategy, path = results[0]
        assert provider == "openai"
        assert model == "gpt-4"
        assert strategy == "baseline/run-1"
        assert path == d

    def test_finds_rag_run_roots(self, tmp_path):
        d = tmp_path / "openai" / "gpt-4" / "vec-chroma-768" / "run-2" / "rag-full"
        self._make_task_bundle(d, task_ids=(0,))

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 1
        provider, model, strategy, path = results[0]
        assert strategy == "vec-chroma-768/run-2/rag-full"
        assert path == d

    def test_ignores_hidden_dirs(self, tmp_path):
        d = tmp_path / ".hidden" / "model" / "baseline" / "run-1"
        self._make_task_bundle(d, task_ids=(0,))

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 0

    def test_ignores_dirs_without_task_bundles(self, tmp_path):
        d = tmp_path / "openai" / "gpt-4" / "baseline" / "run-1"
        (d / "notes.txt").parent.mkdir(parents=True)
        (d / "notes.txt").write_text("no task bundles here", encoding="utf-8")

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 0

    def test_empty_root(self, tmp_path):
        results = discover_experiment_dirs(tmp_path)
        assert results == []

    def test_nonexistent_root(self, tmp_path):
        results = discover_experiment_dirs(tmp_path / "nonexistent")
        assert results == []

    def test_multiple_experiments(self, tmp_path):
        for run_root in [
            tmp_path / "openai" / "gpt-4" / "baseline" / "run-1",
            tmp_path / "openai" / "gpt-4" / "vec-chroma-768" / "run-1" / "rag-full",
        ]:
            self._make_task_bundle(run_root, task_ids=(0,))

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 2
        strategies = {r[2] for r in results}
        assert strategies == {"baseline/run-1", "vec-chroma-768/run-1/rag-full"}


# ---------------------------------------------------------------------------
# build_comparison_table
# ---------------------------------------------------------------------------

class TestBuildComparisonTable:
    def test_single_result(self):
        results = [{
            "provider": "openai",
            "model": "gpt-4",
            "strategy": "baseline",
            "summary": {
                "total_files": 10,
                "compilation_at_1": 0.8,
                "pass_at_1": 0.5,
            },
        }]
        md = build_comparison_table(results)
        assert "openai" in md
        assert "gpt-4" in md
        assert "baseline" in md
        assert "80.0%" in md
        assert "50.0%" in md

    def test_empty_results(self):
        md = build_comparison_table([])
        assert "HumanEval-X" in md
        # Table header exists but no data rows
        assert "openai" not in md

    def test_multiple_results(self):
        results = [
            {
                "provider": "openai",
                "model": "gpt-4",
                "strategy": "baseline",
                "summary": {"total_files": 10, "compilation_at_1": 1.0, "pass_at_1": 0.9},
            },
            {
                "provider": "minimax",
                "model": "m2.5",
                "strategy": "rag/chromadb",
                "summary": {"total_files": 5, "compilation_at_1": 0.6, "pass_at_1": 0.4},
            },
        ]
        md = build_comparison_table(results)
        assert "openai" in md
        assert "minimax" in md
        assert "rag/chromadb" in md
