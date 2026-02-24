"""Entry point: translate Python files to Go and evaluate the results.

Usage:
    uv run python src/lab/00_get_hands_on/run.py [--source-dir PATH]

Test Pass Rate workflow:
    1. Look for existing Python test files (test_*.py or *_test.py) in the source dir.
    2. If no tests exist, the team's TestGenerator agent will generate them.
    3. Python tests are verified against the source before proceeding.
    4. The TestTranslator agent translates verified Python tests to Go tests.
    5. Go tests are run with `go test` and results feed into the Test Pass Rate metric.
"""

import argparse
import importlib
import re
import subprocess
import tempfile
from pathlib import Path

from dotenv import load_dotenv
from rich.console import Console
from rich.progress import Progress

from src.config import TRANSLATION_SOURCE_DIR, TRANSLATION_TARGET_DIR

_models = importlib.import_module("src.lab.00_get_hands_on.models")
_metrics = importlib.import_module("src.lab.00_get_hands_on.metrics")
_team = importlib.import_module("src.lab.00_get_hands_on.team")

TranslationResult = _models.TranslationResult
TestTranslationResult = _models.TestTranslationResult
EvaluationRecord = _models.EvaluationRecord


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
    source_file: Path, target_file: Path, go_test_code: str | None = None
) -> EvaluationRecord:
    """Programmatically evaluate a translated Go file."""
    record = EvaluationRecord(
        source_file=str(source_file),
        target_file=str(target_file),
    )

    if not target_file.exists():
        record.notes = "Target file not found"
        return record

    go_code = target_file.read_text(encoding="utf-8")
    python_code = source_file.read_text(encoding="utf-8")

    # 1. Compilation check
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
                record.notes = comp.stderr.strip()[:200]
                return record
        except (subprocess.TimeoutExpired, FileNotFoundError) as e:
            record.notes = str(e)
            return record

    # 2. Run Go code
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
                record.notes = go_run.stderr.strip()[:200]
        except subprocess.TimeoutExpired:
            record.notes = "Go execution timed out"
            return record

    # 3. Run Python code and compare outputs
    if record.runs_successfully:
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
                    record.io_equivalent = outputs_match
                    record.computational_accuracy = outputs_match
                else:
                    record.notes = "Python source failed to run"
            except subprocess.TimeoutExpired:
                record.notes = "Python execution timed out"

    # 4. Run Go tests if available
    if go_test_code and record.compiles:
        with tempfile.TemporaryDirectory() as tmpdir:
            # Init go module
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
                record.test_pass_rate = passed / total if total > 0 else 0.0
            except subprocess.TimeoutExpired:
                record.notes += " | Go tests timed out"

    return record


def extract_go_code(response) -> str | None:
    """Extract Go code from the team response."""
    if isinstance(response.content, TranslationResult):
        return response.content.go_code

    # Check member runs for TranslationResult
    if hasattr(response, "member_runs") and response.member_runs:
        for member_run in response.member_runs:
            if isinstance(member_run.content, TranslationResult):
                return member_run.content.go_code
            member_text = str(member_run.content) if member_run.content else ""
            if "```go" in member_text:
                start = member_text.index("```go") + 5
                end = member_text.index("```", start)
                return member_text[start:end].strip()

    text = str(response.content) if response.content else ""
    if "```go" in text:
        start = text.index("```go") + 5
        end = text.index("```", start)
        return text[start:end].strip()

    if "```" in text:
        start = text.index("```") + 3
        newline = text.index("\n", start)
        start = newline + 1
        end = text.index("```", start)
        return text[start:end].strip()

    return None


def extract_go_test_code(response) -> str | None:
    """Extract Go test code from the team response."""
    if hasattr(response, "member_runs") and response.member_runs:
        for member_run in response.member_runs:
            if isinstance(member_run.content, TestTranslationResult):
                return member_run.content.test_code

    # Try to find Go test code blocks in the response text
    text = str(response.content) if response.content else ""
    # Look for code that contains "func Test" (Go test signature)
    for block_match in re.finditer(r"```go\n(.*?)```", text, re.DOTALL):
        code = block_match.group(1).strip()
        if "func Test" in code:
            return code

    return None


def preflight_check(console: Console) -> bool:
    """Verify environment before running the pipeline. Returns True if all checks pass."""
    import os
    import shutil

    ok = True

    # 1. API key
    key = os.getenv("MINIMAX_API_KEY")
    if not key or len(key) < 10:
        console.print("[red]FAIL[/red] MINIMAX_API_KEY is not set or invalid in .env")
        ok = False
    else:
        console.print("[green]OK[/green]   MINIMAX_API_KEY is set")

    # 2. Go compiler
    if shutil.which("go"):
        console.print("[green]OK[/green]   Go compiler found")
    else:
        console.print("[red]FAIL[/red] Go compiler not found on PATH")
        ok = False

    # 3. API connection test
    if ok:
        console.print("[dim]     Testing API connection...[/dim]")
        try:
            _agents = importlib.import_module("src.lab.00_get_hands_on.agents")
            agent = _agents.create_translation_agent()
            response = agent.run("Translate to Go: print('hello')", stream=False)
            if response and response.content:
                console.print("[green]OK[/green]   MiniMax API responded")
            else:
                console.print("[red]FAIL[/red] MiniMax API returned empty response")
                ok = False
        except Exception as e:
            console.print(f"[red]FAIL[/red] MiniMax API connection failed: {e}")
            ok = False

    if not ok:
        console.print("\n[bold red]Preflight failed. Fix the issues above before running.[/bold red]")
    else:
        console.print()
    return ok


def main():
    load_dotenv()
    console = Console()

    parser = argparse.ArgumentParser(description="Translate Python to Go")
    parser.add_argument(
        "--source-dir",
        type=Path,
        default=TRANSLATION_SOURCE_DIR,
        help="Directory containing Python source files",
    )
    parser.add_argument(
        "--target-dir",
        type=Path,
        default=TRANSLATION_TARGET_DIR,
        help="Directory for translated Go files",
    )
    parser.add_argument(
        "--skip-preflight",
        action="store_true",
        help="Skip preflight checks (API connection test)",
    )
    args = parser.parse_args()

    # Preflight: verify environment and API before processing files
    if not args.skip_preflight:
        if not preflight_check(console):
            return

    py_files = discover_python_files(args.source_dir)
    if not py_files:
        console.print("[yellow]No Python files found.[/yellow]")
        return

    console.print(
        f"Found [bold]{len(py_files)}[/bold] Python file(s) to translate.\n"
    )

    # Create the team once (never in a loop)
    team = _team.create_translation_team()
    records: list[EvaluationRecord] = []

    with Progress() as progress:
        task = progress.add_task("Translating & evaluating...", total=len(py_files))

        for py_file in py_files:
            python_code = py_file.read_text(encoding="utf-8")
            target_file = mirror_path(
                py_file, args.source_dir, args.target_dir, ".go"
            )
            target_file.parent.mkdir(parents=True, exist_ok=True)

            # Check for existing Python tests
            test_file = find_test_file(py_file)
            test_context = ""
            if test_file:
                test_code = test_file.read_text(encoding="utf-8")
                test_context = (
                    f"\n\nExisting Python tests for this file:\n\n"
                    f"```python\n{test_code}\n```\n\n"
                    f"These tests have been verified. Translate them to Go tests as well."
                )
            else:
                test_context = (
                    "\n\nNo existing Python tests found. "
                    "Generate Python tests first, verify they pass, "
                    "then translate them to Go tests."
                )

            prompt = (
                f"Translate the following Python code to Go:\n\n"
                f"```python\n{python_code}\n```"
                f"{test_context}"
            )

            try:
                response = team.run(prompt, stream=False)

                go_code = extract_go_code(response)
                go_test_code = extract_go_test_code(response)

                if go_code:
                    target_file.write_text(go_code, encoding="utf-8")
                    # Save Go test file alongside translation
                    if go_test_code:
                        test_target = target_file.with_name(
                            target_file.stem + "_test.go"
                        )
                        test_target.write_text(go_test_code, encoding="utf-8")
                    record = evaluate_file(py_file, target_file, go_test_code)
                else:
                    record = EvaluationRecord(
                        source_file=str(py_file),
                        target_file=str(target_file),
                        notes="Could not extract Go code from response",
                    )
            except Exception as e:
                record = EvaluationRecord(
                    source_file=str(py_file),
                    target_file=str(target_file),
                    notes=f"Error: {e}",
                )

            records.append(record)
            progress.advance(task)

    # Display results
    console.print()
    _metrics.display_per_file_table(records)
    console.print()
    summary = _metrics.compute_summary(records)
    _metrics.display_summary_table(summary)


if __name__ == "__main__":
    main()
