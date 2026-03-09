"""Pipeline orchestration: translate and evaluate.

High-level dispatch functions that coordinate the translation and
evaluation workflows across datasets (local, HumanEval-X) and models.
"""

import os
import threading
from concurrent.futures import ThreadPoolExecutor, as_completed
from pathlib import Path

from rich.console import Console
from rich.progress import Progress

from src.config import (
    TRANSLATION_SOURCE_DIR,
    TRANSLATION_TARGET_DIR,
    LOCAL_TARGET_DIR,
    HUMANEVAL_X_TARGET_DIR,
    load_eval_config,
)
from src.core import agents as _agents
from src.core import reporting as _reporting
from src.core.evaluation import discover_python_files, mirror_path, evaluate_file
from src.core.error_db import save_error
from src.core.logger import (
    log_translation_start, log_prompt, log_response,
    log_translation_done, log_translation_error,
    log_eval_start, log_eval_success, log_eval_error,
)
from src.core.schemas import TranslationResult, EvaluationRecord
from src.providers.registry import get_enabled_models, get_model_env_var, get_model_id, get_model_vertex_env_vars


def _get_rag_context(python_code: str) -> str:
    """Retrieve RAG context, returning empty string if unavailable."""
    try:
        from src.rag.retriever import build_translation_context
        return build_translation_context(python_code)
    except Exception:
        return ""


def _parse_target_path(target_dir: Path) -> tuple[str, str, str]:
    """Extract (provider, variant, experiment) from a target directory path."""
    parts = target_dir.parts
    return parts[-3], parts[-2], parts[-1]


def _classify_error(record: EvaluationRecord) -> str | None:
    """Return error type string or None if the record is a success."""
    if not record.compiles:
        return "compile_error"
    if record.compiles and record.notes and "imed out" in record.notes:
        return "timeout"
    if record.compiles and not record.runs_successfully:
        return "runtime_error"
    if record.compiles and record.runs_successfully and not record.pass_at_1:
        return "test_failure"
    return None


def _handle_eval_record(
    record: EvaluationRecord,
    file_name: str,
    provider: str,
    variant: str,
    experiment: str,
    model_id: str,
) -> None:
    """Log and persist evaluation result if it's a failure."""
    err_type = _classify_error(record)
    if err_type is None:
        log_eval_success(file_name)
        return
    log_eval_error(file_name, err_type, record.notes or "")
    save_error(
        file_name=file_name,
        err_type=err_type,
        err_log_context=record.notes,
        provider=provider,
        model=model_id,
        variant=variant,
        experiment_type=experiment,
    )


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
    experiment: str = "baseline",
) -> None:
    """Dispatch translation to the appropriate pipeline based on dataset."""
    if dataset == "local":
        _translate_local(source_dir, LOCAL_TARGET_DIR, skip_preflight, sample=sample, experiment=experiment)
    elif dataset == "humaneval-x":
        _translate_humaneval_x(HUMANEVAL_X_TARGET_DIR, skip_preflight, sample=sample, experiment=experiment)
    else:
        Console().print(f"[red]Unknown dataset: {dataset}[/red]")


def _translate_local(
    source_dir: Path,
    target_dir: Path,
    skip_preflight: bool = False,
    sample: int | None = None,
    experiment: str = "baseline",
) -> None:
    """Translate local Python files to Go.

    Output: target/local/<provider>/<variant>/<experiment>/<mirrored_source>.go
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
        model_target_dir = target_dir / provider_key / variant_key / experiment
        label = f"{provider_key}/{variant_key}"
        console.print(f"\n[bold blue]── Model: {label} ──[/bold blue]")
        console.print(f"   Experiment: {experiment}")
        console.print(f"   Output: {model_target_dir}\n")

        translator = _agents.create_translation_agent(model)
        records: list[dict] = []
        interrupted = False

        with Progress() as progress:
            task = progress.add_task(f"Translating ({label})...", total=len(py_files))

            try:
                for py_file in py_files:
                    log_translation_start(py_file.name, provider_key, variant_key)
                    python_code = py_file.read_text(encoding="utf-8")
                    target_file = mirror_path(py_file, source_dir, model_target_dir, ".go")
                    target_file.parent.mkdir(parents=True, exist_ok=True)

                    rag_context = _get_rag_context(python_code)
                    prompt = (
                        f"{rag_context}\n\n" if rag_context else ""
                    ) + (
                        f"Translate the following Python code to Go:\n\n"
                        f"```python\n{python_code}\n```"
                    )
                    log_prompt(prompt)

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
                            log_response(result, str(target_file))
                            log_translation_done(py_file.name, str(target_file))
                            progress.console.print(
                                f"  [green]OK[/green] {py_file.name} -> {target_file.name}"
                            )
                        else:
                            records.append({
                                "source": str(py_file),
                                "target": str(target_file),
                                "status": "no structured output",
                            })
                            log_response(result)
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
                        log_translation_error(py_file.name, e)
                        progress.console.print(f"  [red]FAIL[/red] {py_file.name}: {e}")

                    progress.advance(task)
            except KeyboardInterrupt:
                interrupted = True

        if interrupted:
            console.print("\n[yellow]Interrupted. Partial results saved.[/yellow]")

        ok_count = sum(1 for r in records if r["status"] == "ok")
        console.print(
            f"\n[bold]{label}:[/bold] {ok_count}/{len(records)} files translated successfully."
        )


def _translate_humaneval_x(
    target_dir: Path,
    skip_preflight: bool = False,
    sample: int | None = None,
    experiment: str = "baseline",
) -> None:
    """Translate HumanEval-X Python problems to Go.

    Output: target/humaneval-x/<provider>/<variant>/<experiment>/Go_<N>.go
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
        model_target_dir = target_dir / provider_key / variant_key / experiment
        model_target_dir.mkdir(parents=True, exist_ok=True)
        label = f"{provider_key}/{variant_key}"
        console.print(f"\n[bold blue]── Model: {label} ──[/bold blue]")
        console.print(f"   Experiment: {experiment}")
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
                    log_translation_start(pair["task_id"], provider_key, variant_key)

                    rag_context = _get_rag_context(pair["py_solution"])
                    prompt = (
                        f"{rag_context}\n\n" if rag_context else ""
                    ) + (
                        f"Translate the following Python code to Go.\n"
                        f"Use this Go function signature:\n"
                        f"```go\n{pair['declaration']}\n```\n\n"
                        f"Python code:\n"
                        f"```python\n{pair['py_solution']}\n```"
                    )
                    log_prompt(prompt)

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
                            log_response(result, str(target_file))
                            log_translation_done(pair["task_id"], str(target_file))
                            progress.console.print(
                                f"  [green]OK[/green] {pair['task_id']} -> {target_file.name}"
                            )
                        else:
                            records.append({
                                "source": pair["task_id"],
                                "target": str(target_file),
                                "status": "no structured output",
                            })
                            log_response(result)
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
                        log_translation_error(pair["task_id"], e)
                        progress.console.print(
                            f"  [red]FAIL[/red] {pair['task_id']}: {e}"
                        )

                    progress.advance(task)
            except KeyboardInterrupt:
                interrupted = True

        if interrupted:
            console.print("\n[yellow]Interrupted. Partial results saved.[/yellow]")

        ok_count = sum(1 for r in records if r["status"] == "ok")
        console.print(
            f"\n[bold]{label}:[/bold] {ok_count}/{len(records)} problems translated successfully."
        )


def evaluate(
    source_dir: Path = TRANSLATION_SOURCE_DIR,
    eval_target_dir: Path | None = None,
    dataset: str = "local",
    batch_size: int | None = None,
) -> None:
    """Dispatch evaluation to the appropriate pipeline based on dataset."""
    if dataset == "local":
        _evaluate_local(source_dir, eval_target_dir)
    elif dataset == "humaneval-x":
        _evaluate_humaneval_x(eval_target_dir, batch_size=batch_size)
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

    provider, variant, experiment = _parse_target_path(eval_target_dir)
    model_id_str = get_model_id(provider, variant)

    records: list[EvaluationRecord] = []

    for i, py_file in enumerate(py_files, 1):
        go_file = mirror_path(py_file, source_dir, eval_target_dir, ".go")
        log_eval_start(py_file.name)
        console.print(f"[bold]({i}/{len(py_files)}) {py_file.name}[/bold]")
        record = evaluate_file(py_file, go_file, console=console)
        _handle_eval_record(record, py_file.name, provider, variant, experiment, model_id_str)
        records.append(record)

    _reporting.display_per_file_table(records)
    summary = _reporting.compute_summary(records)
    _reporting.display_summary_table(summary)


def _evaluate_humaneval_x(
    eval_target_dir: Path | None = None,
    batch_size: int | None = None,
) -> None:
    """Evaluate HumanEval-X translations using Docker.

    Runs evaluations in parallel using ThreadPoolExecutor. Batch size
    is read from config/eval_config.yaml but can be overridden via
    the batch_size parameter or --batch-size CLI flag.
    """
    from src.data.humaneval_x import load_humaneval_x

    from src.core.docker_eval import (
        check_docker_available,
        ensure_go_image,
        ensure_go_mod_cache,
        evaluate_single_task,
        DEFAULT_GO_IMAGE,
    )

    console = Console()

    if eval_target_dir is None:
        console.print("[red]No target directory specified.[/red]")
        return

    if not eval_target_dir.exists():
        console.print(f"[yellow]Directory does not exist: {eval_target_dir}[/yellow]")
        return

    # Load eval config
    eval_config = load_eval_config()
    if batch_size is None:
        batch_size = eval_config["parallel"]["batch_size"]
    timeout = eval_config["docker"]["timeout"]

    console.print(f"\n[bold blue]── Evaluating (HumanEval-X): {eval_target_dir} ──[/bold blue]\n")

    provider, variant, experiment = _parse_target_path(eval_target_dir)
    model_id_str = get_model_id(provider, variant)

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
    go_files = sorted(eval_target_dir.glob("Go_*.go"), key=lambda f: int(f.stem.replace("Go_", "")))
    if not go_files:
        console.print("[yellow]No Go_*.go files found in target directory.[/yellow]")
        return

    console.print(f"Found [bold]{len(go_files)}[/bold] translated Go files to evaluate.")
    console.print(f"[dim]Parallel batch size: {batch_size}[/dim]\n")

    # Build work items (skip files with no matching task)
    work_items: list[tuple[Path, dict]] = []
    skipped = 0
    for go_file in go_files:
        task_num = go_file.stem.replace("Go_", "")
        pair = task_lookup.get(task_num)
        if pair is None:
            console.print(f"  [yellow]SKIP[/yellow] {go_file.name}: no matching HumanEval-X task")
            skipped += 1
            continue
        work_items.append((go_file, pair))

    records: list[EvaluationRecord] = []
    print_lock = threading.Lock()

    def _eval_one(go_file: Path, pair: dict) -> tuple[Path, dict, EvaluationRecord]:
        """Evaluate a single file — called from the thread pool."""
        generated_code = go_file.read_text(encoding="utf-8")
        test_code = pair["test"]
        log_eval_start(go_file.name)
        record = evaluate_single_task(
            task_id=pair["task_id"],
            generated_code=generated_code,
            test_code=test_code,
            timeout=timeout,
        )
        _handle_eval_record(record, go_file.name, provider, variant, experiment, model_id_str)
        return go_file, pair, record

    with Progress() as progress:
        task = progress.add_task("Evaluating...", total=len(go_files))

        # Account for skipped files in progress
        if skipped:
            progress.advance(task, advance=skipped)

        try:
            with ThreadPoolExecutor(max_workers=batch_size) as executor:
                futures = {
                    executor.submit(_eval_one, go_file, pair): (go_file, pair)
                    for go_file, pair in work_items
                }

                for future in as_completed(futures):
                    try:
                        go_file, pair, record = future.result()

                        status = "[green]PASS[/green]" if record.pass_at_1 else (
                            "[yellow]COMPILE[/yellow]" if record.compiles else "[red]FAIL[/red]"
                        )
                        with print_lock:
                            progress.console.print(f"  {status} {go_file.name} ({pair['task_id']})")
                            if record.notes:
                                progress.console.print(f"         [dim]{record.notes[:100]}[/dim]")

                        records.append(record)
                    except Exception as exc:
                        go_file, pair = futures[future]
                        with print_lock:
                            progress.console.print(
                                f"  [red]ERROR[/red] {go_file.name}: {exc}"
                            )
                        records.append(EvaluationRecord(
                            source_file=pair["task_id"],
                            target_file=go_file.name,
                            dataset="humaneval-x",
                            notes=f"Worker error: {exc}",
                        ))

                    progress.advance(task)
        except KeyboardInterrupt:
            console.print("\n[yellow]Interrupted. Cancelling pending tasks...[/yellow]")
            executor.shutdown(wait=False, cancel_futures=True)

    _reporting.display_per_file_table(records)
    summary = _reporting.compute_summary(records)
    _reporting.display_summary_table(summary)
