"""Tests for metric computation and display.

    uv run pytest src/lab/00_get_hands_on/tests/test_metrics.py -v
"""

import importlib

import pytest

_models = importlib.import_module("src.lab.00_get_hands_on.models")
_metrics = importlib.import_module("src.lab.00_get_hands_on.metrics")

EvaluationRecord = _models.EvaluationRecord


class TestComputeSummary:
    def test_empty_records(self):
        summary = _metrics.compute_summary([])
        assert summary["total_files"] == 0
        assert summary["compilation_success_rate"] == 0.0
        assert summary["test_pass_rate"] == 0.0

    def test_all_pass(self):
        records = [
            EvaluationRecord(
                source_file="a.py",
                target_file="a.go",
                compiles=True,
                runs_successfully=True,
                io_equivalent=True,
                computational_accuracy=True,
                tests_total=3,
                tests_passed=3,
                test_pass_rate=1.0,
            ),
        ]
        summary = _metrics.compute_summary(records)
        assert summary["total_files"] == 1
        assert summary["compilation_success_rate"] == 1.0
        assert summary["successful_translation_rate"] == 1.0
        assert summary["computational_accuracy"] == 1.0
        assert summary["io_equivalence_rate"] == 1.0
        assert summary["test_pass_rate"] == 1.0

    def test_all_fail(self):
        records = [
            EvaluationRecord(source_file="a.py", target_file="a.go"),
        ]
        summary = _metrics.compute_summary(records)
        assert summary["compilation_success_rate"] == 0.0
        assert summary["successful_translation_rate"] == 0.0
        assert summary["test_pass_rate"] == 0.0

    def test_mixed_results(self):
        records = [
            EvaluationRecord(
                source_file="a.py",
                target_file="a.go",
                compiles=True,
                runs_successfully=True,
                io_equivalent=True,
                computational_accuracy=True,
                tests_total=2,
                tests_passed=2,
                test_pass_rate=1.0,
            ),
            EvaluationRecord(
                source_file="b.py",
                target_file="b.go",
                compiles=True,
                runs_successfully=False,
            ),
            EvaluationRecord(
                source_file="c.py",
                target_file="c.go",
                compiles=False,
            ),
        ]
        summary = _metrics.compute_summary(records)
        assert summary["total_files"] == 3
        assert abs(summary["compilation_success_rate"] - 2 / 3) < 0.01
        assert abs(summary["successful_translation_rate"] - 1 / 3) < 0.01
        assert abs(summary["computational_accuracy"] - 1 / 3) < 0.01

    def test_test_pass_rate_only_counts_files_with_tests(self):
        records = [
            EvaluationRecord(
                source_file="a.py",
                target_file="a.go",
                tests_total=4,
                tests_passed=2,
                test_pass_rate=0.5,
            ),
            EvaluationRecord(
                source_file="b.py",
                target_file="b.go",
                tests_total=0,
                tests_passed=0,
                test_pass_rate=0.0,
            ),
        ]
        summary = _metrics.compute_summary(records)
        # Only a.py has tests, so test_pass_rate = 0.5 (not averaged with b.py's 0.0)
        assert summary["test_pass_rate"] == 0.5


class TestDisplayTables:
    """Smoke tests — just verify they don't crash."""

    def test_display_summary_table_runs(self, capsys):
        summary = {
            "total_files": 1,
            "compilation_success_rate": 1.0,
            "successful_translation_rate": 1.0,
            "computational_accuracy": 1.0,
            "io_equivalence_rate": 1.0,
            "test_pass_rate": 1.0,
        }
        _metrics.display_summary_table(summary)

    def test_display_per_file_table_runs(self, capsys):
        records = [
            EvaluationRecord(source_file="x.py", target_file="x.go", compiles=True),
        ]
        _metrics.display_per_file_table(records)
