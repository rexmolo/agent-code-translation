"""Tests for metric computation and display.

    uv run pytest src/tests/test_reporting.py -v
"""

import pytest

from src.core.schemas import EvaluationRecord
from src.core import reporting as _reporting


class TestComputeSummary:
    def test_empty_records(self):
        summary = _reporting.compute_summary([])
        assert summary["total_files"] == 0
        assert summary["compilation_at_1"] == 0.0
        assert summary["pass_at_1"] == 0.0

    def test_all_pass(self):
        records = [
            EvaluationRecord(
                source_file="a.py",
                target_file="a.go",
                compiles=True,
                runs_successfully=True,
                pass_at_1=True,
                tests_total=3,
                tests_passed=3,
            ),
        ]
        summary = _reporting.compute_summary(records)
        assert summary["total_files"] == 1
        assert summary["compilation_at_1"] == 1.0
        assert summary["runs_rate"] == 1.0
        assert summary["pass_at_1"] == 1.0

    def test_all_fail(self):
        records = [
            EvaluationRecord(source_file="a.py", target_file="a.go"),
        ]
        summary = _reporting.compute_summary(records)
        assert summary["compilation_at_1"] == 0.0
        assert summary["runs_rate"] == 0.0
        assert summary["pass_at_1"] == 0.0

    def test_mixed_results(self):
        records = [
            EvaluationRecord(
                source_file="a.py",
                target_file="a.go",
                compiles=True,
                runs_successfully=True,
                pass_at_1=True,
                tests_total=2,
                tests_passed=2,
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
        summary = _reporting.compute_summary(records)
        assert summary["total_files"] == 3
        assert abs(summary["compilation_at_1"] - 2 / 3) < 0.01
        assert abs(summary["runs_rate"] - 1 / 3) < 0.01
        assert abs(summary["pass_at_1"] - 1 / 3) < 0.01


class TestDisplayTables:
    """Smoke tests -- just verify they don't crash."""

    def test_display_summary_table_runs(self, capsys):
        summary = {
            "total_files": 1,
            "compilation_at_1": 1.0,
            "runs_rate": 1.0,
            "pass_at_1": 1.0,
            "avg_ast_similarity": 0.0,
        }
        _reporting.display_summary_table(summary)

    def test_display_per_file_table_runs(self, capsys):
        records = [
            EvaluationRecord(source_file="x.py", target_file="x.go", compiles=True),
        ]
        _reporting.display_per_file_table(records)
