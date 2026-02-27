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

from src.config import (
    TRANSLATION_SOURCE_DIR,
    TRANSLATION_TARGET_DIR,
    LOCAL_TARGET_DIR,
    HUMANEVAL_X_TARGET_DIR,
)
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
                record.notes = comp.stderr.strip()[:200]
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
                record.notes = go_run.stderr.strip()[:200]
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
    dataset: str = "local",
    sample: int | None = None,
) -> None:
    """Dispatch translation to the appropriate pipeline based on dataset."""
    if dataset == "local":
        _translate_local(source_dir, LOCAL_TARGET_DIR, skip_preflight, sample=sample)
    elif dataset == "humaneval-x":
        _translate_humaneval_x(HUMANEVAL_X_TARGET_DIR, skip_preflight, sample=sample)
    else:
        Console().print(f"[red]Unknown dataset: {dataset}[/red]")


def _translate_local(
    source_dir: Path,
    target_dir: Path,
    skip_preflight: bool = False,
    sample: int | None = None,
) -> None:
    """Translate local Python files to Go.

    Output: target/local/<provider>/<variant>/<mirrored_source>.go
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

    if sample is not None:
        py_files = py_files[:sample]

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
        interrupted = False

        with Progress() as progress:
            task = progress.add_task(f"Translating ({label})...", total=len(py_files))

            try:
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
            except KeyboardInterrupt:
                interrupted = True

        if interrupted:
            console.print("\n[yellow]⚡ Interrupted. Partial results saved.[/yellow]")

        ok_count = sum(1 for r in records if r["status"] == "ok")
        console.print(
            f"\n[bold]{label}:[/bold] {ok_count}/{len(records)} files translated successfully."
        )


def _translate_humaneval_x(
    target_dir: Path,
    skip_preflight: bool = False,
    sample: int | None = None,
) -> None:
    """Translate HumanEval-X Python problems to Go.

    Output: target/humaneval-x/<provider>/<variant>/Go_<N>.go
    """
    from src.data.humaneval_x import load_humaneval_x

    console = Console()

    enabled = get_enabled_models()
    if not enabled:
        console.print("[red]No models enabled. Enable at least one model.[/red]")
        return

    if not skip_preflight:
        if not preflight_check(console, enabled):
            return

    console.print("[dim]Loading HumanEval-X dataset...[/dim]")
    pairs = load_humaneval_x()
    console.print(f"Loaded [bold]{len(pairs)}[/bold] HumanEval-X problems.")

    if sample is not None:
        pairs = pairs[:sample]
        console.print(f"[dim]Sample mode: translating [bold]{len(pairs)}[/bold] problem(s).\n[/dim]")
    else:
        console.print()

    for provider_key, variant_key, model in enabled:
        model_target_dir = target_dir / provider_key / variant_key
        model_target_dir.mkdir(parents=True, exist_ok=True)
        label = f"{provider_key}/{variant_key}"
        console.print(f"\n[bold blue]── Model: {label} ──[/bold blue]")
        console.print(f"   Output: {model_target_dir}\n")

        translator = _agents.create_translation_agent(model)
        records: list[dict] = []
        interrupted = False

        with Progress() as progress:
            task = progress.add_task(f"Translating ({label})...", total=len(pairs))

            try:
                for pair in pairs:
                    task_num = pair["task_id"].split("/")[1]
                    target_file = model_target_dir / f"Go_{task_num}.go"

                    prompt = (
                        f"Translate the following Python function to Go.\n"
                        f"Use this Go function signature:\n"
                        f"```go\n{pair['declaration']}\n```\n\n"
                        f"Python code:\n"
                        f"```python\n{pair['py_solution']}\n```"
                    )

                    try:
                        response = translator.run(prompt, stream=False)
                        result = response.content

                        if isinstance(result, TranslationResult):
                            target_file.write_text(result.go_code, encoding="utf-8")
                            records.append({
                                "source": pair["task_id"],
                                "target": str(target_file),
                                "status": "ok",
                            })
                            progress.console.print(
                                f"  [green]OK[/green] {pair['task_id']} -> {target_file.name}"
                            )
                        else:
                            records.append({
                                "source": pair["task_id"],
                                "target": str(target_file),
                                "status": "no structured output",
                            })
                            progress.console.print(
                                f"  [yellow]WARN[/yellow] {pair['task_id']}: "
                                f"unexpected response type ({type(result).__name__})"
                            )
                    except Exception as e:
                        records.append({
                            "source": pair["task_id"],
                            "target": str(target_file),
                            "status": f"error: {e}",
                        })
                        progress.console.print(
                            f"  [red]FAIL[/red] {pair['task_id']}: {e}"
                        )

                    progress.advance(task)
            except KeyboardInterrupt:
                interrupted = True

        if interrupted:
            console.print("\n[yellow]⚡ Interrupted. Partial results saved.[/yellow]")

        ok_count = sum(1 for r in records if r["status"] == "ok")
        console.print(
            f"\n[bold]{label}:[/bold] {ok_count}/{len(records)} problems translated successfully."
        )


def evaluate(
    source_dir: Path = TRANSLATION_SOURCE_DIR,
    eval_target_dir: Path | None = None,
    dataset: str = "local",
) -> None:
    """Dispatch evaluation to the appropriate pipeline based on dataset."""
    if dataset == "local":
        _evaluate_local(source_dir, eval_target_dir)
    elif dataset == "humaneval-x":
        _evaluate_humaneval_x(eval_target_dir)
    else:
        Console().print(f"[red]Unknown dataset: {dataset}[/red]")


def _evaluate_local(
    source_dir: Path,
    eval_target_dir: Path | None = None,
) -> None:
    """Evaluate local translated Go files against their Python sources."""
    console = Console()

    if eval_target_dir is None:
        console.print("[red]No target directory specified.[/red]")
        return

    if not eval_target_dir.exists():
        console.print(f"[yellow]Directory does not exist: {eval_target_dir}[/yellow]")
        return

    py_files = discover_python_files(source_dir)
    if not py_files:
        console.print("[yellow]No Python files found.[/yellow]")
        return

    console.print(f"\n[bold blue]── Evaluating (local): {eval_target_dir} ──[/bold blue]\n")

    records: list[EvaluationRecord] = []

    for i, py_file in enumerate(py_files, 1):
        go_file = mirror_path(py_file, source_dir, eval_target_dir, ".go")
        console.print(f"[bold]({i}/{len(py_files)}) {py_file.name}[/bold]")
        record = evaluate_file(py_file, go_file, console=console)
        records.append(record)

    _metrics.display_per_file_table(records)
    summary = _metrics.compute_summary(records)
    _metrics.display_summary_table(summary)


def _evaluate_humaneval_x(
    eval_target_dir: Path | None = None,
) -> None:
    """Evaluate HumanEval-X translations using Docker.

    For each Go_N.go file in eval_target_dir:
    1. Merge translated code with HumanEval-X test harness
    2. Run combined source in a Docker container (golang:1.23-alpine)
    3. Report Compilation@1 and Pass@1
    """
    from src.data.humaneval_x import load_humaneval_x

    _docker_eval = importlib.import_module("src.lab.00_get_hands_on.docker_eval")
    check_docker_available = _docker_eval.check_docker_available
    ensure_go_image = _docker_eval.ensure_go_image
    ensure_go_mod_cache = _docker_eval.ensure_go_mod_cache
    evaluate_single_task = _docker_eval.evaluate_single_task
    DEFAULT_GO_IMAGE = _docker_eval.DEFAULT_GO_IMAGE

    console = Console()

    if eval_target_dir is None:
        console.print("[red]No target directory specified.[/red]")
        return

    if not eval_target_dir.exists():
        console.print(f"[yellow]Directory does not exist: {eval_target_dir}[/yellow]")
        return

    console.print(f"\n[bold blue]── Evaluating (HumanEval-X): {eval_target_dir} ──[/bold blue]\n")

    # Pre-flight: Docker
    if not check_docker_available(console):
        return
    if not ensure_go_image(console, DEFAULT_GO_IMAGE):
        return

    # Pre-download Go modules (testify) into a Docker volume
    console.print("[dim]Ensuring Go module cache (testify)...[/dim]")
    if not ensure_go_mod_cache(DEFAULT_GO_IMAGE):
        console.print("[red]FAIL[/red] Could not download Go modules")
        return
    console.print("[green]OK[/green]   Go module cache ready")

    # Load dataset and build lookup by task number
    console.print("[dim]Loading HumanEval-X dataset...[/dim]")
    pairs = load_humaneval_x()
    task_lookup: dict[str, dict] = {}
    for pair in pairs:
        task_num = pair["task_id"].split("/")[1]
        task_lookup[task_num] = pair
    console.print(f"Loaded [bold]{len(pairs)}[/bold] HumanEval-X problems.\n")

    # Discover Go_*.go files
    go_files = sorted(eval_target_dir.glob("Go_*.go"))
    if not go_files:
        console.print("[yellow]No Go_*.go files found in target directory.[/yellow]")
        return

    console.print(f"Found [bold]{len(go_files)}[/bold] translated Go files to evaluate.\n")

    records: list[EvaluationRecord] = []

    with Progress() as progress:
        task = progress.add_task("Evaluating...", total=len(go_files))

        for go_file in go_files:
            # Extract task number from filename: Go_42.go -> "42"
            match = go_file.stem.replace("Go_", "")
            task_num = match

            pair = task_lookup.get(task_num)
            if pair is None:
                progress.console.print(
                    f"  [yellow]SKIP[/yellow] {go_file.name}: no matching HumanEval-X task"
                )
                progress.advance(task)
                continue

            generated_code = go_file.read_text(encoding="utf-8")
            test_code = pair["test"]

            record = evaluate_single_task(
                task_id=pair["task_id"],
                generated_code=generated_code,
                test_code=test_code,
            )

            status = "[green]PASS[/green]" if record.pass_at_1 else (
                "[yellow]COMPILE[/yellow]" if record.compiles else "[red]FAIL[/red]"
            )
            progress.console.print(f"  {status} {go_file.name} ({pair['task_id']})")
            if record.notes:
                progress.console.print(f"         [dim]{record.notes[:100]}[/dim]")

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
