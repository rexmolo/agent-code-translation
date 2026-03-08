"""File-level evaluation and utility functions.

Handles file discovery, path mirroring, and single-file Go evaluation
(compilation, execution, test running, output comparison).
"""

import subprocess
import tempfile
from pathlib import Path

from rich.console import Console

from src.core.schemas import EvaluationRecord


def discover_python_files(root: Path) -> list[Path]:
    """Recursively find all .py files under a directory, excluding test files."""
    return sorted(
        f
        for f in root.rglob("*.py")
        if not f.name.startswith("test_") and not f.name.endswith("_test.py")
    )


def find_test_file(source_file: Path) -> Path | None:
    """Find an existing Python test file for the given source file."""
    parent = source_file.parent
    stem = source_file.stem

    candidates = [
        parent / f"test_{source_file.name}",
        parent / f"{stem}_test.py",
        parent / "tests" / f"test_{source_file.name}",
        parent / "tests" / f"{stem}_test.py",
    ]
    for candidate in candidates:
        if candidate.exists():
            return candidate
    return None


def mirror_path(
    source_file: Path, source_root: Path, target_root: Path, new_ext: str
) -> Path:
    """Compute the mirrored target path with a new extension."""
    relative = source_file.relative_to(source_root)
    return target_root / relative.with_suffix(new_ext)


def evaluate_file(
    source_file: Path,
    target_file: Path,
    go_test_code: str | None = None,
    console: Console | None = None,
) -> EvaluationRecord:
    """Programmatically evaluate a translated Go file."""
    record = EvaluationRecord(
        source_file=str(source_file),
        target_file=str(target_file),
    )
    log = console or Console()
    filename = target_file.name

    if not target_file.exists():
        record.notes = "Target file not found"
        log.print(f"  [red]SKIP[/red] {filename}: target file not found")
        return record

    go_code = target_file.read_text(encoding="utf-8")
    python_code = source_file.read_text(encoding="utf-8")

    # 1. Compilation check
    log.print(f"  [dim]go build[/dim] {filename}...")
    with tempfile.TemporaryDirectory() as tmpdir:
        go_file = Path(tmpdir) / "main.go"
        go_file.write_text(go_code, encoding="utf-8")
        try:
            comp = subprocess.run(
                ["go", "build", "-o", str(Path(tmpdir) / "main"), str(go_file)],
                capture_output=True,
                text=True,
                timeout=30,
            )
            record.compiles = comp.returncode == 0
            if not record.compiles:
                record.notes = comp.stderr.strip()
                log.print(f"  [red]FAIL[/red] go build: {record.notes}")
                return record
            else:
                log.print(f"  [green]OK[/green]   go build: compiled successfully")
        except (subprocess.TimeoutExpired, FileNotFoundError) as e:
            record.notes = str(e)
            log.print(f"  [red]FAIL[/red] go build: {e}")
            return record

    # 2. Run Go code
    log.print(f"  [dim]go run[/dim]   {filename}...")
    go_stdout = ""
    with tempfile.TemporaryDirectory() as tmpdir:
        go_file = Path(tmpdir) / "main.go"
        go_file.write_text(go_code, encoding="utf-8")
        try:
            go_run = subprocess.run(
                ["go", "run", str(go_file)],
                input="",
                capture_output=True,
                text=True,
                timeout=30,
            )
            record.runs_successfully = go_run.returncode == 0
            go_stdout = go_run.stdout
            if not record.runs_successfully:
                record.notes = go_run.stderr.strip()
                log.print(f"  [red]FAIL[/red] go run: {record.notes}")
            else:
                log.print(f"  [green]OK[/green]   go run: exit 0")
                if go_stdout.strip():
                    log.print(f"  [dim]       stdout: {go_stdout.strip()[:100]}[/dim]")
        except subprocess.TimeoutExpired:
            record.notes = "Go execution timed out"
            log.print(f"  [red]FAIL[/red] go run: timed out")
            return record

    # 3. Run Go tests if available (Pass@1: all tests must pass)
    if go_test_code and record.compiles:
        log.print(f"  [dim]go test[/dim]  {filename}...")
        with tempfile.TemporaryDirectory() as tmpdir:
            subprocess.run(
                ["go", "mod", "init", "testmod"],
                capture_output=True,
                text=True,
                cwd=tmpdir,
            )
            go_file = Path(tmpdir) / "main.go"
            go_file.write_text(go_code, encoding="utf-8")
            test_file = Path(tmpdir) / "main_test.go"
            test_file.write_text(go_test_code, encoding="utf-8")
            try:
                test_run = subprocess.run(
                    ["go", "test", "-v", "./..."],
                    capture_output=True,
                    text=True,
                    timeout=30,
                    cwd=tmpdir,
                )
                output = test_run.stdout + test_run.stderr
                passed = output.count("--- PASS:")
                failed = output.count("--- FAIL:")
                total = passed + failed
                record.tests_total = total
                record.tests_passed = passed
                record.pass_at_1 = (failed == 0 and total > 0)
                if record.pass_at_1:
                    log.print(f"  [green]OK[/green]   go test: {passed}/{total} passed (Pass@1: Y)")
                else:
                    log.print(f"  [red]FAIL[/red] go test: {passed}/{total} passed (Pass@1: N)")
                    if output.strip():
                        log.print(f"  [dim]{output[:300]}[/dim]")
            except subprocess.TimeoutExpired:
                record.notes += " | Go tests timed out"
                log.print(f"  [red]FAIL[/red] go test: timed out")
    elif not go_test_code and record.runs_successfully:
        # Fallback: compare stdout if no test suite available
        log.print(f"  [dim]python3[/dim]  {source_file.name}...")
        with tempfile.TemporaryDirectory() as tmpdir:
            py_file = Path(tmpdir) / "main.py"
            py_file.write_text(python_code, encoding="utf-8")
            try:
                py_run = subprocess.run(
                    ["python3", str(py_file)],
                    input="",
                    capture_output=True,
                    text=True,
                    timeout=30,
                )
                if py_run.returncode == 0:
                    outputs_match = py_run.stdout.strip() == go_stdout.strip()
                    record.pass_at_1 = outputs_match
                    if outputs_match:
                        log.print(f"  [green]OK[/green]   output match: Go == Python")
                    else:
                        log.print(f"  [yellow]DIFF[/yellow] output mismatch:")
                        log.print(f"         Python: {py_run.stdout.strip()[:80]}")
                        log.print(f"         Go:     {go_stdout.strip()[:80]}")
                else:
                    record.notes = "Python source failed to run"
                    log.print(f"  [red]FAIL[/red] python3: {py_run.stderr.strip()[:100]}")
            except subprocess.TimeoutExpired:
                record.notes = "Python execution timed out"
                log.print(f"  [red]FAIL[/red] python3: timed out")

    log.print()
    return record
