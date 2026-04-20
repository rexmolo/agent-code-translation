#!/usr/bin/env python3
"""Run Docker-isolated Go checks for a CodeNet translation run.

For each translated CodeNet task, this script:
1. normalizes the translated Go source
2. runs `go build` in Docker
3. runs `go vet` in Docker
4. runs the resulting program with CodeNet sample stdin in Docker
5. compares stdout against the verified sample output
"""

from __future__ import annotations

import argparse
import json
import subprocess
import tempfile
import time
from dataclasses import dataclass
from pathlib import Path

from rich.console import Console
from rich.progress import Progress

from src.core.codenet_artifacts import CodeNetRunPaths, append_jsonl_record
from src.core.docker_eval import (
    DEFAULT_GO_IMAGE,
    DEFAULT_TIMEOUT,
    build_solution_file,
    check_docker_available,
    ensure_go_image,
)
from src.core.humaneval_artifacts import append_jsonl, write_json, write_text

DEFAULT_CODENET_ROOT = Path("/Volumes/MyZhiTai/DEV/www/thesis/Project_CodeNet")

_GO_MOD = """\
module codenetcheck

go 1.26.0
"""


@dataclass
class DockerCheckResult:
    compiles: bool = False
    vet_passed: bool = False
    runs_successfully: bool = False
    timed_out: bool = False
    build_stderr: str = ""
    vet_stderr: str = ""
    run_stderr: str = ""
    stdout: str = ""
    exact_output_match: bool = False
    stripped_output_match: bool = False


def _docker_base(tmpdir: str, image: str) -> list[str]:
    return [
        "docker", "run", "--rm",
        "--network=none",
        "--memory=512m",
        "--pids-limit=256",
        "-e", "GOCACHE=/tmp/gocache",
        "-e", "GOMODCACHE=/tmp/gomodcache",
        "-v", f"{tmpdir}:/app",
        "-w", "/app",
        image,
    ]


def _run_command(command: list[str], *, timeout: int) -> subprocess.CompletedProcess[str]:
    return subprocess.run(
        command,
        capture_output=True,
        text=True,
        timeout=timeout,
    )


def _check_in_docker(
    go_source: str,
    stdin_input: str,
    expected_output: str,
    *,
    task_id: str,
    image: str,
    timeout: int,
) -> DockerCheckResult:
    result = DockerCheckResult()

    with tempfile.TemporaryDirectory(prefix=f"codenet_{task_id}_") as tmpdir:
        tmppath = Path(tmpdir)
        (tmppath / "go.mod").write_text(_GO_MOD, encoding="utf-8")
        (tmppath / "solution.go").write_text(go_source, encoding="utf-8")
        (tmppath / "input.txt").write_text(stdin_input, encoding="utf-8")

        docker_base = _docker_base(tmpdir, image)

        try:
            build = _run_command([*docker_base, "go", "build", "-o", "app", "."], timeout=timeout)
            result.build_stderr = build.stderr
            if build.returncode != 0:
                return result
            result.compiles = True

            vet = _run_command([*docker_base, "go", "vet", "."], timeout=timeout)
            result.vet_stderr = vet.stderr
            result.vet_passed = vet.returncode == 0
            if vet.returncode != 0:
                return result

            run = _run_command([*docker_base, "sh", "-lc", "./app < input.txt"], timeout=timeout)
            result.stdout = run.stdout
            result.run_stderr = run.stderr
            result.runs_successfully = run.returncode == 0
            if run.returncode != 0:
                return result

            result.exact_output_match = run.stdout == expected_output
            result.stripped_output_match = run.stdout.strip() == expected_output.strip()
            return result
        except subprocess.TimeoutExpired:
            result.timed_out = True
            return result


def _load_manifest(run_root: Path) -> dict:
    manifest_path = run_root / "manifest.json"
    if not manifest_path.exists():
        raise FileNotFoundError(f"Run manifest not found: {manifest_path}")
    return json.loads(manifest_path.read_text(encoding="utf-8"))


def _read_ready_records(path: Path, start_index: int) -> tuple[list[dict], int]:
    if not path.exists():
        return [], start_index
    rows: list[dict] = []
    with path.open(encoding="utf-8") as fh:
        for index, line in enumerate(fh):
            if index < start_index:
                continue
            text = line.strip()
            if text:
                rows.append(json.loads(text))
    return rows, start_index + len(rows)


def _translated_ok(task_path) -> bool:
    if not task_path.translation_go.exists():
        return False
    if not task_path.translation_result_json.exists():
        return False
    payload = json.loads(task_path.translation_result_json.read_text(encoding="utf-8"))
    return payload.get("status") == "ok"


def _queue_candidates(run_paths: CodeNetRunPaths, *, ready_index: int, selected_ids: set[str]) -> tuple[list[str], int]:
    records, next_index = _read_ready_records(run_paths.control_ready_jsonl, ready_index)
    pending: list[str] = []
    for record in records:
        problem_id = record.get("problem_id")
        if not problem_id:
            continue
        if selected_ids and problem_id not in selected_ids:
            continue
        pending.append(problem_id)
    return pending, next_index


def _scan_task_candidates(run_paths: CodeNetRunPaths, *, selected_ids: set[str], limit: int | None) -> list[str]:
    task_paths = run_paths.iter_task_dirs()
    if selected_ids:
        task_paths = [task for task in task_paths if task.problem_id in selected_ids]
    if limit is not None:
        task_paths = task_paths[:limit]

    pending: list[str] = []
    for task_path in task_paths:
        if task_path.evaluation_result_json.exists():
            continue
        if _translated_ok(task_path):
            pending.append(task_path.problem_id)
    return pending


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--run-root", type=Path, required=True, help="CodeNet translation run root.")
    parser.add_argument("--codenet-root", type=Path, default=DEFAULT_CODENET_ROOT, help="Local Project_CodeNet root.")
    parser.add_argument("--image", default=DEFAULT_GO_IMAGE, help="Docker image to use.")
    parser.add_argument("--timeout", type=int, default=DEFAULT_TIMEOUT, help="Per-step Docker timeout in seconds.")
    parser.add_argument("--sample", type=int, default=None, help="Check only the first N tasks.")
    parser.add_argument("--problem-id", action="append", default=[], help="One or more problem IDs to check.")
    parser.add_argument("--skip-existing", action="store_true", help="Skip tasks that already have evaluation/result.json.")
    parser.add_argument("--watch", action="store_true", help="Watch the translation control queue and evaluate tasks as they become ready.")
    parser.add_argument("--poll-interval", type=float, default=2.0, help="Polling interval in seconds when --watch is enabled.")
    args = parser.parse_args()

    console = Console()
    run_root = args.run_root.resolve()
    run_paths = CodeNetRunPaths(run_root)
    manifest = _load_manifest(run_root)

    if not check_docker_available(console):
        raise SystemExit(1)
    if not ensure_go_image(console, args.image):
        raise SystemExit(1)

    selected_ids = set(args.problem_id)
    task_paths = run_paths.iter_task_dirs()
    if selected_ids:
        task_paths = [task for task in task_paths if task.problem_id in selected_ids]
    if args.sample is not None:
        task_paths = task_paths[:args.sample]
    if not task_paths and not args.watch:
        console.print("[yellow]No CodeNet task bundles selected.[/yellow]")
        return

    run_paths.ensure_evaluation_dirs()
    io_root = args.codenet_root / "derived" / "input_output" / "data"

    console.print(f"[bold blue]Run Root:[/bold blue] {run_root}")
    console.print(f"[bold blue]Tasks:[/bold blue] {len(task_paths)}")
    console.print(f"[bold blue]Docker Image:[/bold blue] {args.image}")

    existing_records: list[dict] = []
    if run_paths.evaluation_per_task_jsonl.exists():
        with run_paths.evaluation_per_task_jsonl.open(encoding="utf-8") as fh:
            for line in fh:
                text = line.strip()
                if text:
                    existing_records.append(json.loads(text))
    records_by_problem = {row["problem_id"]: row for row in existing_records if "problem_id" in row}
    processed_ids = set(records_by_problem)

    pending_ids: list[str] = []
    ready_index = 0

    def enqueue(problem_id: str) -> None:
        if selected_ids and problem_id not in selected_ids:
            return
        if problem_id in processed_ids:
            return
        if problem_id in pending_ids:
            return
        pending_ids.append(problem_id)

    for problem_id in _scan_task_candidates(run_paths, selected_ids=selected_ids, limit=args.sample):
        enqueue(problem_id)

    progress_total = len(task_paths) if not args.watch else None
    with Progress(console=console) as progress:
        task = progress.add_task("Checking translations in Docker", total=progress_total)
        while True:
            queued_ids, ready_index = _queue_candidates(
                run_paths,
                ready_index=ready_index,
                selected_ids=selected_ids,
            )
            for problem_id in queued_ids:
                enqueue(problem_id)

            for problem_id in _scan_task_candidates(run_paths, selected_ids=selected_ids, limit=args.sample):
                enqueue(problem_id)

            if pending_ids:
                problem_id = pending_ids.pop(0)
                task_path = run_paths.task(problem_id)

                if args.skip_existing and task_path.evaluation_result_json.exists():
                    processed_ids.add(problem_id)
                    if progress_total is not None:
                        progress.advance(task)
                    continue

                translation_result = {}
                if task_path.translation_result_json.exists():
                    translation_result = json.loads(task_path.translation_result_json.read_text(encoding="utf-8"))

                if not task_path.translation_go.exists():
                    payload = {
                        "problem_id": task_path.problem_id,
                        "status": "missing_translation",
                        "compiles": False,
                        "vet_passed": False,
                        "runs_successfully": False,
                        "sample_io_match": False,
                        "note": "translation.go is missing",
                    }
                    write_json(task_path.evaluation_result_json, payload)
                    records_by_problem[task_path.problem_id] = payload
                    processed_ids.add(task_path.problem_id)
                    append_jsonl_record(run_paths.evaluation_per_task_jsonl, payload)
                    console.print(f"[red]MISS[/red] {task_path.problem_id} missing translation")
                    if progress_total is not None:
                        progress.advance(task)
                    continue

                input_path = io_root / task_path.problem_id / "input.txt"
                output_path = io_root / task_path.problem_id / "output.txt"
                if not input_path.exists() or not output_path.exists():
                    payload = {
                        "problem_id": task_path.problem_id,
                        "status": "missing_sample_io",
                        "compiles": False,
                        "vet_passed": False,
                        "runs_successfully": False,
                        "sample_io_match": False,
                        "note": "sample input/output files are missing",
                    }
                    write_json(task_path.evaluation_result_json, payload)
                    records_by_problem[task_path.problem_id] = payload
                    processed_ids.add(task_path.problem_id)
                    append_jsonl_record(run_paths.evaluation_per_task_jsonl, payload)
                    console.print(f"[red]MISS[/red] {task_path.problem_id} missing sample I/O")
                    if progress_total is not None:
                        progress.advance(task)
                    continue

                raw_go = task_path.translation_go.read_text(encoding="utf-8")
                normalized_go = build_solution_file(raw_go)
                stdin_input = input_path.read_text(encoding="utf-8", errors="replace")
                expected_output = output_path.read_text(encoding="utf-8", errors="replace")

                check = _check_in_docker(
                    normalized_go,
                    stdin_input,
                    expected_output,
                    task_id=task_path.problem_id,
                    image=args.image,
                    timeout=args.timeout,
                )

                write_text(task_path.evaluation_solution_go, normalized_go)
                write_text(task_path.evaluation_input_txt, stdin_input)
                write_text(task_path.evaluation_expected_output_txt, expected_output)
                write_text(task_path.evaluation_stdout_txt, check.stdout)
                write_text(
                    task_path.evaluation_stderr_txt,
                    "\n\n".join(
                        part for part in (
                            f"[build]\n{check.build_stderr}".strip(),
                            f"[vet]\n{check.vet_stderr}".strip(),
                            f"[run]\n{check.run_stderr}".strip(),
                        )
                        if part.strip() and part not in {"[build]", "[vet]", "[run]"}
                    ),
                )

                if check.timed_out:
                    status = "timeout"
                elif not check.compiles:
                    status = "build_failed"
                elif not check.vet_passed:
                    status = "vet_failed"
                elif not check.runs_successfully:
                    status = "run_failed"
                elif check.stripped_output_match:
                    status = "ok"
                else:
                    status = "wrong_output"

                payload = {
                    "problem_id": task_path.problem_id,
                    "status": status,
                    "provider": manifest.get("provider"),
                    "variant": manifest.get("variant"),
                    "experiment": manifest.get("experiment"),
                    "translation_status": translation_result.get("status"),
                    "compiles": check.compiles,
                    "vet_passed": check.vet_passed,
                    "runs_successfully": check.runs_successfully,
                    "timed_out": check.timed_out,
                    "sample_io_match": check.stripped_output_match,
                    "exact_output_match": check.exact_output_match,
                    "stdout_len": len(check.stdout),
                    "build_stderr": check.build_stderr,
                    "vet_stderr": check.vet_stderr,
                    "run_stderr": check.run_stderr,
                }
                write_json(task_path.evaluation_result_json, payload)
                records_by_problem[task_path.problem_id] = payload
                processed_ids.add(task_path.problem_id)
                append_jsonl_record(run_paths.evaluation_per_task_jsonl, payload)
                tag = "[green]OK[/green]" if status == "ok" else "[yellow]WARN[/yellow]" if status == "wrong_output" else "[red]FAIL[/red]"
                console.print(f"{tag} {task_path.problem_id} -> {status}")
                if progress_total is not None:
                    progress.advance(task)
                continue

            if not args.watch:
                break

            done = run_paths.control_done_flag.exists()
            if done and not pending_ids:
                break

            time.sleep(args.poll_interval)

    records = [records_by_problem[key] for key in sorted(records_by_problem)]
    summary = {
        "task_count": len(records),
        "build_failed": sum(1 for row in records if row["status"] == "build_failed"),
        "vet_failed": sum(1 for row in records if row["status"] == "vet_failed"),
        "run_failed": sum(1 for row in records if row["status"] == "run_failed"),
        "wrong_output": sum(1 for row in records if row["status"] == "wrong_output"),
        "ok": sum(1 for row in records if row["status"] == "ok"),
        "timeout": sum(1 for row in records if row["status"] == "timeout"),
        "missing_translation": sum(1 for row in records if row["status"] == "missing_translation"),
        "missing_sample_io": sum(1 for row in records if row["status"] == "missing_sample_io"),
    }
    write_json(run_paths.evaluation_summary_json, summary)

    console.print(f"[green]Evaluation complete.[/green] OK: {summary['ok']}/{summary['task_count']}")
    console.print(f"[green]Results:[/green] {run_paths.evaluation_results_dir}")


if __name__ == "__main__":
    main()
