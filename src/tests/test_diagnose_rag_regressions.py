import json
from pathlib import Path

from src.scripts.diagnose_rag_regressions import (
    analyze_regressions,
    load_task_bundles,
    main,
    write_reports,
)


def _write_json(path: Path, payload: dict) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2), encoding="utf-8")


def _write_task_bundle(
    run_root: Path,
    task_id: str,
    *,
    pass_at_1: bool,
    parallel_corpus_hits: int = 0,
) -> None:
    task_dir = run_root / "tasks" / task_id
    task_dir.mkdir(parents=True, exist_ok=True)
    _write_json(task_dir / "prompt.json", {"task_id": task_id, "user_prompt": "prompt"})
    _write_json(
        task_dir / "retrieval.json",
        {
            "retrieval_counts": {"parallel_corpus": parallel_corpus_hits},
            "items": {"parallel_corpus": [{} for _ in range(parallel_corpus_hits)]},
        },
    )
    _write_json(task_dir / "llm_raw.json", {"available": False, "format": None, "payload": None, "note": "n/a"})
    (task_dir / "translation.go").write_text("func Solve() int { return 1 }\n", encoding="utf-8")
    eval_dir = task_dir / "evaluation"
    eval_dir.mkdir(parents=True, exist_ok=True)
    (eval_dir / "solution.go").write_text("package main\nfunc Solve() int { return 1 }\n", encoding="utf-8")
    (eval_dir / "test.go").write_text("package main\n", encoding="utf-8")
    _write_json(eval_dir / "result.json", {"pass_at_1": pass_at_1})


def test_analyze_regressions_identifies_baseline_pass_rag_fail(tmp_path):
    baseline = tmp_path / "baseline" / "run-1"
    rag = tmp_path / "vec-chroma-768" / "run-1" / "rag-full"
    _write_task_bundle(baseline, "Go_0", pass_at_1=True)
    _write_task_bundle(baseline, "Go_1", pass_at_1=True)
    _write_task_bundle(rag, "Go_0", pass_at_1=False, parallel_corpus_hits=1)
    _write_task_bundle(rag, "Go_1", pass_at_1=True, parallel_corpus_hits=0)

    summary = analyze_regressions(baseline, [rag])

    assert summary["baseline_run"] == str(baseline)
    assert len(summary["runs"]) == 1
    run_summary = summary["runs"][0]
    assert run_summary["comparable_tasks"] == 2
    assert len(run_summary["regressions"]) == 1
    assert run_summary["regressions"][0]["task_id"] == "Go_0"
    assert run_summary["regressions"][0]["parallel_corpus_hits"] == 1


def test_write_reports_copies_case_artifacts(tmp_path):
    baseline = tmp_path / "baseline" / "run-1"
    rag = tmp_path / "vec-chroma-768" / "run-1" / "rag-full"
    _write_task_bundle(baseline, "Go_3", pass_at_1=True)
    _write_task_bundle(rag, "Go_3", pass_at_1=False, parallel_corpus_hits=2)

    summary = analyze_regressions(baseline, [rag])
    output_dir = tmp_path / "diagnostics"
    write_reports(summary, output_dir)

    assert (output_dir / "summary.json").is_file()
    assert (output_dir / "summary.md").is_file()

    copied_case = next(output_dir.glob("**/Go_3/case_summary.json"))
    case_dir = copied_case.parent
    assert (case_dir / "prompt.json").is_file()
    assert (case_dir / "retrieval.json").is_file()
    assert (case_dir / "translation.go").is_file()
    assert (case_dir / "solution.go").is_file()
    assert json.loads(copied_case.read_text(encoding="utf-8"))["parallel_corpus_hits"] == 2


def test_main_handles_empty_regressions_and_filters(tmp_path):
    baseline = tmp_path / "baseline" / "run-1"
    rag = tmp_path / "vec-chroma-768" / "run-1" / "rag-full"
    _write_task_bundle(baseline, "Go_8", pass_at_1=True)
    _write_task_bundle(rag, "Go_8", pass_at_1=True)

    assert load_task_bundles(baseline)["Go_8"].pass_at_1 is True

    exit_code = main(
        [
            "--baseline-run",
            str(baseline),
            "--rag-run",
            str(rag),
            "--task",
            "Go_8",
            "--output-dir",
            str(tmp_path / "out"),
        ]
    )

    assert exit_code == 0
    summary = json.loads((tmp_path / "out" / "summary.json").read_text(encoding="utf-8"))
    assert summary["runs"][0]["regressions"] == []
