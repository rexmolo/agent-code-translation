"""Tests for ci_evaluate_all.py (no Docker, no network).

    uv run pytest src/tests/test_ci_evaluate_all.py -v
"""

from pathlib import Path

import pytest

from src.scripts.ci_evaluate_all import discover_experiment_dirs, build_comparison_table


# ---------------------------------------------------------------------------
# discover_experiment_dirs
# ---------------------------------------------------------------------------

class TestDiscoverExperimentDirs:
    def test_finds_depth3_experiments(self, tmp_path):
        """provider/model/strategy with Go files at depth 3."""
        d = tmp_path / "openai" / "gpt-4" / "baseline"
        d.mkdir(parents=True)
        (d / "Go_0.go").write_text("package main")
        (d / "Go_1.go").write_text("package main")

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 1
        provider, model, strategy, path = results[0]
        assert provider == "openai"
        assert model == "gpt-4"
        assert strategy == "baseline"
        assert path == d

    def test_finds_depth4_experiments(self, tmp_path):
        """provider/model/strategy/backend with Go files at depth 4."""
        d = tmp_path / "openai" / "gpt-4" / "rag" / "chromadb"
        d.mkdir(parents=True)
        (d / "Go_0.go").write_text("package main")

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 1
        provider, model, strategy, path = results[0]
        assert strategy == "rag/chromadb"
        assert path == d

    def test_ignores_hidden_dirs(self, tmp_path):
        d = tmp_path / ".hidden" / "model" / "strategy"
        d.mkdir(parents=True)
        (d / "Go_0.go").write_text("package main")

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 0

    def test_ignores_dirs_without_go_files(self, tmp_path):
        d = tmp_path / "openai" / "gpt-4" / "baseline"
        d.mkdir(parents=True)
        (d / "notes.txt").write_text("no go files here")

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 0

    def test_empty_root(self, tmp_path):
        results = discover_experiment_dirs(tmp_path)
        assert results == []

    def test_nonexistent_root(self, tmp_path):
        results = discover_experiment_dirs(tmp_path / "nonexistent")
        assert results == []

    def test_multiple_experiments(self, tmp_path):
        for name in ["baseline", "rag"]:
            d = tmp_path / "openai" / "gpt-4" / name
            d.mkdir(parents=True)
            (d / "Go_0.go").write_text("package main")

        results = discover_experiment_dirs(tmp_path)
        assert len(results) == 2
        strategies = {r[2] for r in results}
        assert strategies == {"baseline", "rag"}


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
