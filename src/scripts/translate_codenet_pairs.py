#!/usr/bin/env python3
"""Translate verified CodeNet Python programs to Go.

The input JSONL is expected to contain verified-I/O CodeNet pairs, such as the
filtered ``data/processed/parallel_corpus/codeNet/python_go_pairs.jsonl`` file.

Artifacts are written under:
    data/translation/target/codenet/<provider>/<variant>/<experiment>/<run_name>/tasks/<problem_id>/
"""

from __future__ import annotations

import argparse
import json
import os
import time
from pathlib import Path
from typing import Iterable

from rich.console import Console
from rich.progress import Progress

from src.core.codenet_artifacts import CodeNetRunPaths, append_jsonl_record, codenet_run_root
from src.core.humaneval_artifacts import append_jsonl, write_json, write_text

DEFAULT_PAIRS_FILE = Path("data/processed/parallel_corpus/codeNet/python_go_pairs.jsonl")

_PROGRAM_OUTPUT_CONTRACT = """\
Program output contract:
- Emit a complete standalone Go program.
- Use `package main`.
- Include all required imports.
- Preserve the Python program's stdin/stdout behavior and output formatting exactly.
- Do not include prose before or after the code.
"""


def _default_run_name() -> str:
    return time.strftime("%Y%m%d-%H%M%S")


def _load_pairs(path: Path) -> list[dict]:
    rows: list[dict] = []
    with path.open(encoding="utf-8") as fh:
        for line in fh:
            text = line.strip()
            if text:
                rows.append(json.loads(text))
    return rows


def _select_pairs(rows: list[dict], *, problem_ids: set[str], sample: int | None) -> list[dict]:
    selected = [row for row in rows if not problem_ids or row["problem_id"] in problem_ids]
    if sample is not None:
        selected = selected[:sample]
    return selected


def _serialize_raw_response(response: object) -> dict:
    for method_name in ("model_dump", "dict", "to_dict"):
        method = getattr(response, method_name, None)
        if callable(method):
            try:
                payload = method()
            except Exception:
                continue
            return {
                "available": True,
                "format": "json",
                "payload": payload,
                "note": f"Serialized via response.{method_name}()",
            }

    for attr_name in ("raw", "raw_response", "response"):
        payload = getattr(response, attr_name, None)
        if payload is None:
            continue
        if isinstance(payload, (str, bytes)):
            if isinstance(payload, bytes):
                payload = payload.decode("utf-8", errors="replace")
            return {
                "available": True,
                "format": "text",
                "payload": payload,
                "note": f"Captured from response.{attr_name}",
            }
        return {
            "available": True,
            "format": "json",
            "payload": json.loads(json.dumps(payload, default=str)),
            "note": f"Captured from response.{attr_name}",
        }

    return {
        "available": False,
        "format": "unavailable",
        "payload": None,
        "note": "Raw provider payload is not exposed by the SDK response object.",
    }


def _build_prompt(python_code: str) -> str:
    return (
        "Translate the Python code below to Go.\n\n"
        f"Python code:\n```python\n{python_code}\n```\n\n"
        f"{_PROGRAM_OUTPUT_CONTRACT}"
    )


def _prompt_payload(translator: object, user_prompt: str, *, provider: str, variant: str, problem_id: str, experiment: str) -> dict:
    return {
        "system_prompt": getattr(translator, "instructions", ""),
        "user_prompt": user_prompt,
        "provider": provider,
        "variant": variant,
        "problem_id": problem_id,
        "experiment": experiment,
    }


def _load_single_model(provider: str, variant: str):
    from src.providers.registry import enable_model, get_enabled_models, reset

    reset()
    enable_model(provider, variant)
    enabled = get_enabled_models()
    if len(enabled) != 1:
        raise RuntimeError(f"Expected one enabled model, got {len(enabled)}")
    return enabled[0]


def _translation_record(problem_id: str, task_dir: Path, status: str) -> dict:
    return {
        "problem_id": problem_id,
        "task_dir": str(task_dir),
        "status": status,
    }


def _iter_problem_ids(rows: Iterable[dict]) -> set[str]:
    return {row["problem_id"] for row in rows}


def _write_control_state(run_paths: CodeNetRunPaths, *, status: str, task_count: int, translated_ok: int, translated_failed: int) -> None:
    write_json(
        run_paths.control_state_json,
        {
            "status": status,
            "pid": os.getpid(),
            "task_count": task_count,
            "translated_ok": translated_ok,
            "translated_failed": translated_failed,
            "updated_at": time.strftime("%Y-%m-%dT%H:%M:%S"),
        },
    )


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--provider", required=True, help="Model provider key, e.g. minimax.")
    parser.add_argument("--variant", required=True, help="Model variant key, e.g. M2.5.")
    parser.add_argument("--experiment", default="baseline", help="Experiment label for this run.")
    parser.add_argument("--run-name", default=_default_run_name(), help="Run directory name.")
    parser.add_argument("--pairs-file", type=Path, default=DEFAULT_PAIRS_FILE, help="Verified CodeNet pair JSONL.")
    parser.add_argument("--sample", type=int, default=None, help="Translate only the first N selected pairs.")
    parser.add_argument("--problem-id", action="append", default=[], help="One or more problem IDs to translate.")
    parser.add_argument("--skip-existing", action="store_true", help="Skip tasks that already have translation.go.")
    args = parser.parse_args()

    console = Console()
    pairs_file = args.pairs_file.resolve()
    rows = _load_pairs(pairs_file)
    selected = _select_pairs(rows, problem_ids=set(args.problem_id), sample=args.sample)
    if not selected:
        console.print("[yellow]No CodeNet pairs selected.[/yellow]")
        return

    from src.core import agents as _agents
    from src.core.schemas import TranslationResult
    from src.providers.registry import get_model_id

    provider_key, variant_key, model = _load_single_model(args.provider, args.variant)
    translator = _agents.create_translation_agent(model)

    run_root = codenet_run_root(provider_key, variant_key, args.experiment, args.run_name)
    run_paths = CodeNetRunPaths(run_root)
    run_paths.ensure_translation_dirs()
    run_paths.control_ready_jsonl.write_text("", encoding="utf-8")
    if run_paths.control_done_flag.exists():
        run_paths.control_done_flag.unlink()

    write_json(
        run_paths.manifest_json,
        {
            "dataset": "codenet",
            "provider": provider_key,
            "variant": variant_key,
            "model_id": get_model_id(provider_key, variant_key),
            "experiment": args.experiment,
            "run_name": args.run_name,
            "pairs_file": str(pairs_file),
            "selected_problem_ids": sorted(_iter_problem_ids(selected)),
            "task_count": len(selected),
        },
    )
    _write_control_state(
        run_paths,
        status="running",
        task_count=len(selected),
        translated_ok=0,
        translated_failed=0,
    )

    console.print(f"[bold blue]Run Root:[/bold blue] {run_root}")
    console.print(f"[bold blue]Model:[/bold blue] {provider_key}/{variant_key}")
    console.print(f"[bold blue]Tasks:[/bold blue] {len(selected)}")

    records: list[dict] = []

    with Progress(console=console) as progress:
        task = progress.add_task("Translating CodeNet pairs", total=len(selected))
        for row in selected:
            problem_id = row["problem_id"]
            task_paths = run_paths.task(problem_id)
            if args.skip_existing and task_paths.translation_go.exists():
                records.append(_translation_record(problem_id, task_paths.task_dir, "skipped_existing"))
                progress.advance(task)
                continue

            write_json(
                task_paths.metadata_json,
                {
                    "problem_id": problem_id,
                    "provider": provider_key,
                    "variant": variant_key,
                    "experiment": args.experiment,
                },
            )
            write_text(task_paths.source_py, row["python_code"])
            write_text(task_paths.reference_go, row.get("go_code", ""))

            prompt = _build_prompt(row["python_code"])
            write_json(
                task_paths.prompt_json,
                _prompt_payload(
                    translator,
                    prompt,
                    provider=provider_key,
                    variant=variant_key,
                    problem_id=problem_id,
                    experiment=args.experiment,
                ),
            )

            status = "error"
            response = None
            error_text = ""
            try:
                response = translator.run(prompt, stream=False)
                result = getattr(response, "content", None)
                write_json(task_paths.llm_raw_json, _serialize_raw_response(response))
                if isinstance(result, TranslationResult):
                    write_text(task_paths.translation_go, result.go_code)
                    status = "ok"
                    append_jsonl_record(
                        run_paths.control_ready_jsonl,
                        {
                            "problem_id": problem_id,
                            "status": "ready",
                            "task_dir": str(task_paths.task_dir),
                            "translation_path": str(task_paths.translation_go),
                            "published_at": time.strftime("%Y-%m-%dT%H:%M:%S"),
                        },
                    )
                    console.print(f"[green]READY[/green] {problem_id}")
                else:
                    error_text = "Provider response did not yield TranslationResult."
                    status = "no_structured_output"
            except Exception as exc:
                error_text = str(exc)
                write_json(
                    task_paths.llm_raw_json,
                    {
                        "available": False,
                        "format": "unavailable",
                        "payload": None,
                        "note": f"Translation request failed before a raw response could be captured: {exc}",
                    },
                )

            write_json(
                task_paths.translation_result_json,
                {
                    "problem_id": problem_id,
                    "status": status,
                    "error": error_text,
                    "has_translation_file": task_paths.translation_go.exists(),
                },
            )
            records.append(_translation_record(problem_id, task_paths.task_dir, status))
            ok_count = sum(1 for row in records if row["status"] == "ok")
            fail_count = len(records) - ok_count
            _write_control_state(
                run_paths,
                status="running",
                task_count=len(selected),
                translated_ok=ok_count,
                translated_failed=fail_count,
            )
            progress.advance(task)

    ok_count = sum(1 for row in records if row["status"] == "ok")
    write_json(
        run_paths.translation_summary_json,
        {
            "task_count": len(records),
            "translated_ok": ok_count,
            "translated_failed": len(records) - ok_count,
        },
    )
    append_jsonl(run_root / "translation_records.jsonl", records)
    run_paths.control_done_flag.write_text(
        time.strftime("%Y-%m-%dT%H:%M:%S"),
        encoding="utf-8",
    )
    _write_control_state(
        run_paths,
        status="done",
        task_count=len(records),
        translated_ok=ok_count,
        translated_failed=len(records) - ok_count,
    )

    console.print(f"[green]Completed:[/green] {ok_count}/{len(records)} translations written")
    console.print(f"[green]Artifacts:[/green] {run_root}")


if __name__ == "__main__":
    main()
