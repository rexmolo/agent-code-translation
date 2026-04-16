from pathlib import Path

from src.core.humaneval_artifacts import (
    base_experiment_name,
    is_baseline_experiment,
    parse_humaneval_run_root,
    repair_enabled_for_experiment,
)
from src.core.pipeline import _finalize_repair_record
from src.core.schemas import EvaluationRecord


def test_repair_experiment_helpers():
    assert base_experiment_name("rag-full-repair") == "rag-full"
    assert repair_enabled_for_experiment("rag-full-repair") is True
    assert repair_enabled_for_experiment("rag-full") is False
    assert is_baseline_experiment("baseline-repair") is True
    assert is_baseline_experiment("baseline") is True
    assert is_baseline_experiment("rag-full-repair") is False


def test_parse_baseline_repair_run_root():
    path = Path("/data/humaneval-x/minimax/M2.5/baseline-repair/run-7")
    provider, variant, experiment, backend, run_id = parse_humaneval_run_root(path)
    assert provider == "minimax"
    assert variant == "M2.5"
    assert experiment == "baseline-repair"
    assert backend is None
    assert run_id == 7


def test_finalize_repair_record_success():
    first_pass = EvaluationRecord(
        source_file="Go/0",
        target_file="Go_0.go",
        dataset="humaneval-x",
        compiles=False,
        runs_successfully=False,
        pass_at_1=False,
        notes="undefined: foo",
    )
    repaired = EvaluationRecord(
        source_file="Go/0",
        target_file="Go_0.go",
        dataset="humaneval-x",
        compiles=True,
        runs_successfully=True,
        pass_at_1=True,
        notes="",
    )

    record = _finalize_repair_record(first_pass, repaired)

    assert record.compiles is True
    assert record.pass_at_1 is True
    assert record.repair_attempted is True
    assert record.repair_succeeded is True
    assert record.first_pass_compiles is False
    assert record.first_pass_pass_at_1 is False
    assert record.first_pass_notes == "undefined: foo"
    assert "Compilation repaired" in record.repair_notes


def test_finalize_repair_record_failure():
    first_pass = EvaluationRecord(
        source_file="Go/1",
        target_file="Go_1.go",
        dataset="humaneval-x",
        compiles=False,
        runs_successfully=False,
        pass_at_1=False,
        notes="syntax error",
    )

    record = _finalize_repair_record(first_pass, None, "Repair request returned no structured output")

    assert record.compiles is False
    assert record.repair_attempted is True
    assert record.repair_succeeded is False
    assert record.first_pass_compiles is False
    assert record.first_pass_notes == "syntax error"
    assert record.repair_notes == "Repair request returned no structured output"
