"""Entry point: translate Python files to Go.

Usage:
    uv run python src/lab/00_get_hands_on/run.py [--source-dir PATH]
    uv run python -m src.cli translate
    uv run python -m src.cli evaluate

Evaluation can be run separately after translation completes.
"""

import argparse
import importlib
import os
import subprocess
import tempfile
from pathlib import Path

from dotenv import load_dotenv
from rich.console import Console
from rich.progress import Progress

from src.config import TRANSLATION_SOURCE_DIR, TRANSLATION_TARGET_DIR
from src.models.registry import get_enabled_models, get_model_env_var, get_model_vertex_env_vars

_models = importlib.import_module("src.lab.00_get_hands_on.models")
_metrics = importlib.import_module("src.lab.00_get_hands_on.metrics")
_agents = importlib.import_module("src.lab.00_get_hands_on.agents")

TranslationResult = _models.TranslationResult
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


def preflight_check(console: Console, enabled_models: list[tuple[str, str, object]]) -> bool:
    """Verify environment before running the pipeline. Returns True if all checks pass."""
    import shutil

    ok = True

    # 1. Check API keys for each enabled provider (deduplicate per provider)
    checked_providers: set[str] = set()
    for provider_key, variant_key, _model in enabled_models:
        if provider_key in checked_providers:
            continue
        checked_providers.add(provider_key)

        vertex_vars = get_model_vertex_env_vars(provider_key)
        if vertex_vars and os.getenv("GOOGLE_GENAI_USE_VERTEXAI", "").lower() == "true":
            missing = [v for v in vertex_vars if not os.getenv(v)]
            if missing:
                console.print(f"[red]FAIL[/red] Vertex AI vars missing: {', '.join(missing)}")
                ok = False
            else:
                console.print(f"[green]OK[/green]   Vertex AI credentials set for {provider_key}")
        else:
            env_var = get_model_env_var(provider_key)
            key = os.getenv(env_var)
            if not key or len(key) < 10:
                console.print(f"[red]FAIL[/red] {env_var} is not set or invalid in .env")
                ok = False
            else:
                console.print(f"[green]OK[/green]   {env_var} is set")

    # 2. Go compiler
    if shutil.which("go"):
        console.print("[green]OK[/green]   Go compiler found")
    else:
        console.print("[red]FAIL[/red] Go compiler not found on PATH")
        ok = False

    # 3. API connection test for each enabled model variant
    if ok:
        for provider_key, variant_key, model in enabled_models:
            label = f"{provider_key}/{variant_key}"
            console.print(f"[dim]     Testing {label} API connection...[/dim]")
            try:
                agent = _agents.create_translation_agent(model)
                response = agent.run("Translate to Go: print('hello')", stream=False)
                if response and response.content:
                    console.print(f"[green]OK[/green]   {label} API responded")
                else:
                    console.print(f"[red]FAIL[/red] {label} API returned empty response")
                    ok = False
            except Exception as e:
                console.print(f"[red]FAIL[/red] {label} API connection failed: {e}")
                ok = False

    if not ok:
        console.print("\n[bold red]Preflight failed. Fix the issues above before running.[/bold red]")
    else:
        console.print()
    return ok


def translate(
    source_dir: Path = TRANSLATION_SOURCE_DIR,
    target_dir: Path = TRANSLATION_TARGET_DIR,
    skip_preflight: bool = False,
) -> None:
    """Translate all Python files in source_dir to Go files in target_dir.

    Loops over enabled models. Each model's output is saved under
    target_dir/<provider>/<variant>/ (e.g. target/gemini/2.5_flash_lite/).
    """
    console = Console()

    enabled = get_enabled_models()
    if not enabled:
        console.print("[red]No models enabled. Enable at least one model.[/red]")
        return

    if not skip_preflight:
        if not preflight_check(console, enabled):
            return

    py_files = discover_python_files(source_dir)
    if not py_files:
        console.print("[yellow]No Python files found.[/yellow]")
        return

    console.print(
        f"Found [bold]{len(py_files)}[/bold] Python file(s) to translate.\n"
    )

    for provider_key, variant_key, model in enabled:
        model_target_dir = target_dir / provider_key / variant_key
        label = f"{provider_key}/{variant_key}"
        console.print(f"\n[bold blue]── Model: {label} ──[/bold blue]")
        console.print(f"   Output: {model_target_dir}\n")

        translator = _agents.create_translation_agent(model)
        records: list[dict] = []

        with Progress() as progress:
            task = progress.add_task(f"Translating ({label})...", total=len(py_files))

            for py_file in py_files:
                python_code = py_file.read_text(encoding="utf-8")
                target_file = mirror_path(py_file, source_dir, model_target_dir, ".go")
                target_file.parent.mkdir(parents=True, exist_ok=True)

                prompt = (
                    f"Translate the following Python code to Go:\n\n"
                    f"```python\n{python_code}\n```"
                )

                try:
                    response = translator.run(prompt, stream=False)
                    result = response.content

                    if isinstance(result, TranslationResult):
                        target_file.write_text(result.go_code, encoding="utf-8")
                        records.append({
                            "source": str(py_file),
                            "target": str(target_file),
                            "status": "ok",
                        })
                        progress.console.print(
                            f"  [green]OK[/green] {py_file.name} -> {target_file.name}"
                        )
                    else:
                        records.append({
                            "source": str(py_file),
                            "target": str(target_file),
                            "status": "no structured output",
                        })
                        progress.console.print(
                            f"  [yellow]WARN[/yellow] {py_file.name}: "
                            f"unexpected response type ({type(result).__name__})"
                        )
                except Exception as e:
                    records.append({
                        "source": str(py_file),
                        "target": str(target_file),
                        "status": f"error: {e}",
                    })
                    progress.console.print(f"  [red]FAIL[/red] {py_file.name}: {e}")

                progress.advance(task)

        ok_count = sum(1 for r in records if r["status"] == "ok")
        console.print(
            f"\n[bold]{label}:[/bold] {ok_count}/{len(records)} files translated successfully."
        )


def evaluate(
    source_dir: Path = TRANSLATION_SOURCE_DIR,
    target_dir: Path = TRANSLATION_TARGET_DIR,
) -> None:
    """Evaluate translated Go files against their Python sources.

    Loops over enabled models. Each model's output is expected under
    target_dir/<provider>/<variant>/.
    """
    console = Console()

    enabled = get_enabled_models()
    if not enabled:
        console.print("[red]No models enabled. Enable at least one model.[/red]")
        return

    py_files = discover_python_files(source_dir)
    if not py_files:
        console.print("[yellow]No Python files found.[/yellow]")
        return

    for provider_key, variant_key, _model in enabled:
        model_target_dir = target_dir / provider_key / variant_key
        label = f"{provider_key}/{variant_key}"
        console.print(f"\n[bold blue]── Model: {label} ──[/bold blue]")
        console.print(f"   Evaluating: {model_target_dir}\n")

        if not model_target_dir.exists():
            console.print(f"[yellow]No translations found for {label} (directory does not exist).[/yellow]")
            continue

        records: list[EvaluationRecord] = []

        with Progress() as progress:
            task = progress.add_task(f"Evaluating ({label})...", total=len(py_files))

            for py_file in py_files:
                go_file = mirror_path(py_file, source_dir, model_target_dir, ".go")
                record = evaluate_file(py_file, go_file)
                records.append(record)
                progress.advance(task)

        _metrics.display_per_file_table(records)
        summary = _metrics.compute_summary(records)
        _metrics.display_summary_table(summary)


def main():
    load_dotenv()

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

    translate(
        source_dir=args.source_dir,
        target_dir=args.target_dir,
        skip_preflight=args.skip_preflight,
    )


if __name__ == "__main__":
    main()
