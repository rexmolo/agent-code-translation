from pathlib import Path

import pytest

from src.scripts.analyze_dimension_variance import (
    ParsedRow,
    aggregate_baseline_runs,
    build_dataframe,
    load_jsonl,
    parse_target,
    render_report,
)


def _summary(compilation: float, passed: float) -> str:
    return f"Compilation@1 {compilation:.1f}% | Pass@1 {passed:.1f}%"


def _write_jsonl(path: Path, rows: list[dict]) -> None:
    path.write_text("\n".join(__import__("json").dumps(row) for row in rows), encoding="utf-8")


def _report_df(rows: list[ParsedRow]):
    return aggregate_baseline_runs(build_dataframe(rows))


def test_parse_target_supports_baseline_and_rag_shapes():
    assert parse_target("minimax/M2.5/baseline") == ("minimax", "M2.5", None, None, "baseline", None)
    assert parse_target("minimax/M2.5/baseline/run-3") == ("minimax", "M2.5", None, None, "baseline", 3)
    assert parse_target("minimax/M2.5/vec-gemini/rag-full") == ("minimax", "M2.5", "vec-gemini", None, "rag-full", 1)
    assert parse_target("minimax/M2.5/vec-chroma-768/run-2/rag-full") == (
        "minimax",
        "M2.5",
        "vec-chroma-768",
        768,
        "rag-full",
        2,
    )


def test_parse_target_rejects_malformed_shapes():
    with pytest.raises(ValueError):
        parse_target("minimax")
    with pytest.raises(ValueError):
        parse_target("minimax/M2.5")
    with pytest.raises(ValueError):
        parse_target("minimax/M2.5/baseline/run-x")


def test_load_jsonl_skips_malformed_successful_rows(tmp_path: Path):
    path = tmp_path / "results.jsonl"
    _write_jsonl(
        path,
        [
            {"success": True, "summary": _summary(10.0, 20.0)},
            {"success": True, "target": "minimax/M2.5/baseline/run-1"},
            {"success": True, "target": "minimax/M2.5/baseline/run-2", "summary": "bad"},
            {
                "success": True,
                "target": "minimax/M2.5/baseline/run-3",
                "summary": _summary(30.0, 40.0),
            },
            {"success": False, "target": "minimax/M2.5/baseline/run-4", "summary": _summary(90.0, 90.0)},
        ],
    )

    rows = load_jsonl(path)

    assert len(rows) == 1
    assert rows[0].provider == "minimax"
    assert rows[0].run_id == 3
    assert rows[0].compilation_at_1 == 0.30
    assert rows[0].pass_at_1 == 0.40


def test_aggregate_baseline_runs_by_provider_model():
    df = build_dataframe(
        [
            ParsedRow("minimax", "M2.5", None, None, "baseline", 1, 0.4, 0.2),
            ParsedRow("minimax", "M2.5", None, None, "baseline", 2, 0.8, 0.6),
            ParsedRow("minimax", "M2.5", "vec-chroma-768", 768, "rag-full", 1, 0.9, 0.7),
        ]
    )

    aggregated = aggregate_baseline_runs(df)
    baseline = aggregated[(aggregated["provider"] == "minimax") & (aggregated["experiment"] == "baseline")].iloc[0]

    assert len(aggregated) == 2
    assert baseline["compilation_at_1"] == pytest.approx(0.6)
    assert baseline["pass_at_1"] == pytest.approx(0.4)
    assert baseline["baseline_runs"] == 2


def test_render_report_omits_minimax_line_section_without_minimax_baseline(tmp_path: Path):
    rows = [
        ParsedRow("minimax", "M2.5", "vec-chroma-768", 768, "rag-full", 1, 0.8, 0.7),
        ParsedRow("minimax", "M2.5", "vec-chroma-768", 768, "rag-full", 2, 0.7, 0.6),
    ]
    df = _report_df(rows)
    report_path = tmp_path / "report.md"

    render_report(df, df[(df["provider"] == "minimax") & (df["backend"] == "vec-chroma-768")], ["pass_at_1"], tmp_path, report_path, [])

    report = report_path.read_text(encoding="utf-8")
    assert "Grouped bars — mean ± 1 SD" in report
    assert "Line across dimensions with ±1 SD bands" not in report


def test_render_report_omits_openai_section_without_openai_baseline(tmp_path: Path):
    rows = [
        ParsedRow("minimax", "M2.5", None, None, "baseline", 1, 0.9, 0.8),
        ParsedRow("minimax", "M2.5", "vec-chroma-768", 768, "rag-full", 1, 0.8, 0.7),
        ParsedRow("minimax", "M2.5", "vec-chroma-768", 768, "rag-full", 2, 0.7, 0.6),
        ParsedRow("openai", "GPT-5.4", "vec-chroma-3072", 3072, "rag-full", 1, 0.95, 0.91),
    ]
    df = _report_df(rows)
    report_path = tmp_path / "report.md"

    render_report(df, df[(df["provider"] == "minimax") & (df["backend"] == "vec-chroma-768")], ["pass_at_1"], tmp_path, report_path, [])

    report = report_path.read_text(encoding="utf-8")
    assert "### C2 — OpenAI GPT-5.4 RAG variants (chroma-3072)" not in report


def test_render_report_omits_cross_model_baseline_section_when_no_baselines(tmp_path: Path):
    rows = [
        ParsedRow("minimax", "M2.5", "vec-chroma-768", 768, "rag-full", 1, 0.8, 0.7),
        ParsedRow("minimax", "M2.5", "vec-chroma-768", 768, "rag-full", 2, 0.7, 0.6),
        ParsedRow("openai", "GPT-5.4", "vec-chroma-3072", 3072, "rag-full", 1, 0.95, 0.91),
    ]
    df = _report_df(rows)
    report_path = tmp_path / "report.md"

    render_report(df, df[(df["provider"] == "minimax") & (df["backend"] == "vec-chroma-768")], ["pass_at_1"], tmp_path, report_path, [])

    report = report_path.read_text(encoding="utf-8")
    assert "### C1 — Baseline comparison across providers" not in report
