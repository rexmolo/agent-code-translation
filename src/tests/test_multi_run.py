"""Tests for multi-run experiment support.

Covers:
- _parse_target_path() with run-N paths
- Path construction in _translate_humaneval_x (via target dir logic)
- ci_evaluate_all discover_experiment_dirs with run-N
- analyze_statistics discovery and grouping
- run_all_batches queue building and state management

    uv run pytest src/tests/test_multi_run.py -v
"""

import json
from pathlib import Path

import pytest

from src.core.humaneval_artifacts import HumanEvalRunPaths


# ---------------------------------------------------------------------------
# _parse_target_path
# ---------------------------------------------------------------------------

class TestParseTargetPath:
    """Test the updated _parse_target_path function."""

    def setup_method(self):
        from src.core.pipeline import _parse_target_path
        self.parse = _parse_target_path

    def test_legacy_baseline(self, tmp_path):
        """Legacy: .../provider/variant/baseline"""
        path = Path("/data/humaneval-x/minimax/M2.5/baseline")
        provider, variant, experiment, backend, run_id = self.parse(path)
        assert provider == "minimax"
        assert variant == "M2.5"
        assert experiment == "baseline"
        assert backend is None
        assert run_id is None

    def test_legacy_rag_with_backend(self):
        """Legacy: .../provider/variant/vec-chroma-768/rag-full"""
        path = Path("/data/humaneval-x/minimax/M2.5/vec-chroma-768/rag-full")
        provider, variant, experiment, backend, run_id = self.parse(path)
        assert provider == "minimax"
        assert variant == "M2.5"
        assert experiment == "rag-full"
        assert backend == "vec-chroma-768"
        assert run_id is None

    def test_new_rag_with_run(self):
        """New: .../provider/variant/vec-chroma-768/run-3/rag-full"""
        path = Path("/data/humaneval-x/minimax/M2.5/vec-chroma-768/run-3/rag-full")
        provider, variant, experiment, backend, run_id = self.parse(path)
        assert provider == "minimax"
        assert variant == "M2.5"
        assert experiment == "rag-full"
        assert backend == "vec-chroma-768"
        assert run_id == 3

    def test_new_rag_vec_gemini_with_run(self):
        """New: .../provider/variant/vec-gemini/run-2/rag-pattern-only"""
        path = Path("/data/humaneval-x/minimax/M2.5/vec-gemini/run-2/rag-pattern-only")
        provider, variant, experiment, backend, run_id = self.parse(path)
        assert experiment == "rag-pattern-only"
        assert backend == "vec-gemini"
        assert run_id == 2

    def test_legacy_vec_gemini(self):
        """Legacy: .../provider/variant/vec-gemini/rag-pattern-api-docs"""
        path = Path("/data/humaneval-x/minimax/M2.5/vec-gemini/rag-pattern-api-docs")
        provider, variant, experiment, backend, run_id = self.parse(path)
        assert experiment == "rag-pattern-api-docs"
        assert backend == "vec-gemini"
        assert run_id is None


# ---------------------------------------------------------------------------
# discover_experiment_dirs with run-N
# ---------------------------------------------------------------------------

class TestDiscoverRunN:
    """Test ci_evaluate_all.discover_experiment_dirs with run-N structure."""

    def setup_method(self):
        from src.scripts.ci_evaluate_all import discover_experiment_dirs
        self.discover = discover_experiment_dirs

    def _make_task_bundles(self, run_root: Path, n: int = 2):
        run_paths = HumanEvalRunPaths(run_root)
        run_paths.ensure_translation_dirs()
        for i in range(n):
            task = run_paths.task(i)
            task.translation_go.parent.mkdir(parents=True, exist_ok=True)
            task.translation_go.write_text("package main\n", encoding="utf-8")

    def test_baseline_run_n(self, tmp_path):
        """baseline/run-N with task bundles."""
        self._make_task_bundles(tmp_path / "minimax" / "M2.5" / "baseline" / "run-1")

        results = self.discover(tmp_path)
        assert len(results) == 1
        _, _, strategy, path = results[0]
        assert strategy == "baseline/run-1"

    def test_rag_run_n(self, tmp_path):
        """vec-chroma-768/run-1/rag-full/ with task bundles."""
        self._make_task_bundles(
            tmp_path / "minimax" / "M2.5" / "vec-chroma-768" / "run-1" / "rag-full"
        )

        results = self.discover(tmp_path)
        assert len(results) == 1
        _, _, strategy, path = results[0]
        assert strategy == "vec-chroma-768/run-1/rag-full"

    def test_multiple_runs(self, tmp_path):
        """Multiple runs for the same dimension."""
        base = tmp_path / "minimax" / "M2.5" / "vec-chroma-768"
        for run in range(1, 4):
            self._make_task_bundles(base / f"run-{run}" / "rag-full")

        results = self.discover(tmp_path)
        assert len(results) == 3
        strategies = {r[2] for r in results}
        assert strategies == {
            "vec-chroma-768/run-1/rag-full",
            "vec-chroma-768/run-2/rag-full",
            "vec-chroma-768/run-3/rag-full",
        }

    def test_baseline_and_run_n_rag(self, tmp_path):
        """Baseline/run-N plus run-N RAG structure."""
        self._make_task_bundles(tmp_path / "minimax" / "M2.5" / "baseline" / "run-1")
        self._make_task_bundles(
            tmp_path / "minimax" / "M2.5" / "vec-chroma-768" / "run-1" / "rag-full"
        )

        results = self.discover(tmp_path)
        assert len(results) == 2
        strategies = {r[2] for r in results}
        assert "baseline/run-1" in strategies
        assert "vec-chroma-768/run-1/rag-full" in strategies

    def test_multiple_experiments_per_run(self, tmp_path):
        """All 4 RAG experiments under one run."""
        base = tmp_path / "minimax" / "M2.5" / "vec-chroma-768" / "run-1"
        for exp in ["rag-full", "rag-pattern-only", "rag-pattern-samples", "rag-pattern-api-docs"]:
            self._make_task_bundles(base / exp)

        results = self.discover(tmp_path)
        assert len(results) == 4


# ---------------------------------------------------------------------------
# analyze_statistics
# ---------------------------------------------------------------------------

class TestAnalyzeStatistics:
    """Test statistics script discovery and grouping."""

    def setup_method(self):
        from src.scripts.analyze_statistics import discover_results, group_results
        self.discover = discover_results
        self.group = group_results

    def _write_result(self, path: Path, pass_at_1: float):
        path.parent.mkdir(parents=True, exist_ok=True)
        data = {
            "total_files": 164,
            "compilation_at_1": pass_at_1 + 0.05,
            "pass_at_1": pass_at_1,
            "runs_rate": pass_at_1,
            "avg_ast_similarity": 0.0,
        }
        path.write_text(json.dumps(data), encoding="utf-8")

    def test_discover_finds_results(self, tmp_path):
        self._write_result(
            tmp_path / "minimax" / "M2.5" / "vec-chroma-768" / "run-1" / "rag-full" / "evaluation" / "results" / "summary.json",
            0.42,
        )

        results = self.discover(tmp_path)
        assert len(results) == 1
        assert results[0]["experiment"] == "rag-full"
        assert results[0]["dimensions"] == 768
        assert results[0]["run_id"] == 1
        assert results[0]["pass_at_1"] == 0.42

    def test_discover_extracts_dimension(self, tmp_path):
        self._write_result(
            tmp_path / "minimax" / "M2.5" / "vec-chroma-3072" / "run-1" / "rag-full" / "evaluation" / "results" / "summary.json",
            0.38,
        )

        results = self.discover(tmp_path)
        assert results[0]["dimensions"] == 3072

    def test_discover_baseline_no_dimension(self, tmp_path):
        self._write_result(
            tmp_path / "minimax" / "M2.5" / "baseline" / "run-1" / "evaluation" / "results" / "summary.json",
            0.30,
        )

        results = self.discover(tmp_path)
        assert results[0]["experiment"] == "baseline"
        assert results[0]["dimensions"] is None

    def test_group_by_dimension(self, tmp_path):
        """Group results by experiment and dimension."""
        for dim in [768, 3072]:
            for run in range(1, 4):
                self._write_result(
                    tmp_path / "minimax" / "M2.5" / f"vec-chroma-{dim}" / f"run-{run}" / "rag-full" / "evaluation" / "results" / "summary.json",
                    0.40 + dim / 10000,
                )

        results = self.discover(tmp_path)
        grouped = self.group(results, "pass_at_1")

        assert "rag-full" in grouped
        assert 768 in grouped["rag-full"]
        assert 3072 in grouped["rag-full"]
        assert len(grouped["rag-full"][768]) == 3
        assert len(grouped["rag-full"][3072]) == 3

    def test_group_baseline(self, tmp_path):
        for run in range(1, 4):
            self._write_result(
                tmp_path / "minimax" / "M2.5" / "baseline" / f"run-{run}" / "evaluation" / "results" / "summary.json",
                0.30,
            )

        results = self.discover(tmp_path)
        grouped = self.group(results, "pass_at_1")

        assert "baseline" in grouped
        assert "none" in grouped["baseline"]
        assert len(grouped["baseline"]["none"]) == 3


# ---------------------------------------------------------------------------
# run_all_batches
# ---------------------------------------------------------------------------

class TestBatchRunner:
    """Test batch runner queue building and state management."""

    def setup_method(self):
        from src.scripts.run_all_batches import (
            build_experiment_queue, experiment_key, load_state, save_state,
        )
        self.build_queue = build_experiment_queue
        self.experiment_key = experiment_key
        self.load_state = load_state
        self.save_state = save_state

    def test_queue_size_without_baseline(self):
        queue = self.build_queue([768, 3072], runs=3, include_baseline=False)
        # 2 dims × 3 runs × 5 experiments = 30
        assert len(queue) == 30

    def test_queue_order_rag_experiments(self):
        queue = self.build_queue([768], runs=2, include_baseline=False)
        # First items should be RAG experiments for dim 768
        assert queue[0]["experiment"] == "rag-pattern-only"
        assert queue[0]["run_id"] == 1
        assert queue[0]["dimension"] == 768

    def test_queue_full_60_experiments(self):
        queue = self.build_queue([768, 1536, 3072], runs=5, include_baseline=False)
        # 3 dims × 5 runs × 5 experiments = 75
        assert len(queue) == 75

    def test_queue_includes_rag_routed(self):
        queue = self.build_queue([768], runs=1, include_baseline=False)
        experiments = [item["experiment"] for item in queue]
        assert experiments == [
            "rag-pattern-only",
            "rag-pattern-samples",
            "rag-pattern-api-docs",
            "rag-full",
            "rag-routed",
        ]

    def test_experiment_key_unique(self):
        queue = self.build_queue([768, 3072], runs=2, include_baseline=True)
        keys = [self.experiment_key(item) for item in queue]
        assert len(keys) == len(set(keys)), "Keys must be unique"

    def test_experiment_key_format(self):
        item = {"experiment": "rag-full", "dimension": 768, "run_id": 3}
        assert self.experiment_key(item) == "rag-full/dim-768/run-3"

    def test_experiment_key_baseline(self):
        item = {"experiment": "baseline", "dimension": None, "run_id": 1}
        assert self.experiment_key(item) == "baseline/dim-none/run-1"

    def test_state_save_and_load(self, tmp_path):
        state_file = tmp_path / "state.json"
        completed = {"baseline/dim-none/run-1", "rag-full/dim-768/run-1"}

        self.save_state(state_file, completed)
        loaded = self.load_state(state_file)

        assert loaded == completed

    def test_state_load_nonexistent(self, tmp_path):
        state_file = tmp_path / "nonexistent.json"
        loaded = self.load_state(state_file)
        assert loaded == set()

    def test_state_load_corrupted(self, tmp_path):
        state_file = tmp_path / "bad.json"
        state_file.write_text("not json", encoding="utf-8")
        loaded = self.load_state(state_file)
        assert loaded == set()
