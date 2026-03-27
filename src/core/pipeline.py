"""Pipeline orchestration: translate and evaluate.

High-level dispatch functions that coordinate the translation and
evaluation workflows across datasets (local, HumanEval-X) and models.
"""

import multiprocessing
import os
import threading
import time
from concurrent.futures import ThreadPoolExecutor, as_completed
from pathlib import Path

from rich.console import Console
from rich.live import Live
from rich.progress import Progress
from rich.table import Table

from src.config import (
    TRANSLATION_SOURCE_DIR,
    TRANSLATION_TARGET_DIR,
    LOCAL_TARGET_DIR,
    HUMANEVAL_X_TARGET_DIR,
    load_eval_config,
)
from src.core import agents as _agents
from src.core import reporting as _reporting
from src.core.prompt_builder import PromptBuilder
from src.core.evaluation import discover_python_files, mirror_path, evaluate_file
from src.core.error_db import save_error
from src.core.logger import (
    log_translation_start, log_prompt, log_response,
    log_translation_done, log_translation_error,
    log_eval_start, log_eval_success, log_eval_error,
    log_rag_retrieval,
)
from src.core.schemas import TranslationResult, EvaluationRecord
from src.providers.registry import get_enabled_models, get_model_env_var, get_model_id, get_model_vertex_env_vars, resolve_provider_api_key


def _get_rag_result(
    python_code: str,
    experiment: str = "baseline",
    embedding_backend: str = "chromadb",
):
    """Retrieve RAG result, returning None if unavailable or not needed."""
    if experiment == "baseline":
        return None
    try:
        from src.rag.retriever import build_translation_context
        rag_result = build_translation_context(python_code, embedding_backend=embedding_backend)
        log_rag_retrieval(rag_result)
        return rag_result
    except Exception:
        return None


def _setup_and_display_kb(
    experiment: str,
    console: Console,
    embedding_backend: str = "chromadb",
) -> None:
    """Configure knowledge base toggles for the experiment and display status."""
    from src.rag.retriever import configure_kb_for_experiment, get_active_kb_toggles

    configure_kb_for_experiment(experiment)
    toggles = get_active_kb_toggles(experiment)

    if toggles is None:
        console.print("   RAG: [dim]disabled (baseline)[/dim]")
        return

    labels = {
        "grammar":         "Grammar Patterns",
        "parallel_corpus": "Parallel Corpus",
        "api_mappings":    "API Mappings",
        "documentation":   "Go Docs",
    }
    parts = []
    for key, label in labels.items():
        enabled = toggles.get(key, False)
        tag = "[green]ON[/green]" if enabled else "[red]OFF[/red]"
        parts.append(f"{label}: {tag}")
    console.print(f"   RAG: {' | '.join(parts)}")
    backend_label = "Vertex AI + Gemini" if embedding_backend == "gemini" else "ChromaDB"
    console.print(f"   Embedding: [cyan]{backend_label}[/cyan]")


def _parse_target_path(target_dir: Path) -> tuple[str, str, str]:
    """Extract provider, variant, and experiment from a target directory path.

    Path structure: .../humaneval-x/<provider>/<variant>/<experiment>
    """
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
        if vertex_vars and os.getenv(vertex_vars[0], "").lower() == "true":
            missing = [v for v in vertex_vars if not os.getenv(v)]
            if missing:
                console.print(f"[red]FAIL[/red] Vertex AI vars missing: {', '.join(missing)}")
                ok = False
            else:
                console.print(f"[green]OK[/green]   Vertex AI credentials set for {provider_key}")
        else:
            key = resolve_provider_api_key(provider_key)
            if not key or len(key) < 10:
                console.print(f"[red]FAIL[/red] API key for {provider_key} is not configured in providers.yaml or env")
                ok = False
            else:
                console.print(f"[green]OK[/green]   API key set for {provider_key}")

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
    embedding_backend: str = "chromadb",
) -> None:
    """Dispatch translation to the appropriate pipeline based on dataset."""
    if dataset == "local":
        _translate_local(source_dir, LOCAL_TARGET_DIR, skip_preflight, sample=sample, experiment=experiment, embedding_backend=embedding_backend)
    elif dataset == "humaneval-x":
        _translate_humaneval_x(HUMANEVAL_X_TARGET_DIR, skip_preflight, sample=sample, experiment=experiment, embedding_backend=embedding_backend)
    else:
        Console().print(f"[red]Unknown dataset: {dataset}[/red]")


def _translate_local(
    source_dir: Path,
    target_dir: Path,
    skip_preflight: bool = False,
    sample: int | None = None,
    experiment: str = "baseline",
    embedding_backend: str = "chromadb",
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

    translation_cfg = load_eval_config()["translation"]
    batch_size = translation_cfg["batch_size"]
    stagger = translation_cfg.get("thread_stagger_seconds", 1)

    console.print(
        f"Found [bold]{len(py_files)}[/bold] Python file(s) to translate.\n"
    )

    for provider_key, variant_key, model in enabled:
        model_target_dir = target_dir / provider_key / variant_key / experiment

        label = f"{provider_key}/{variant_key}"
        console.print(f"\n[bold blue]── Model: {label} ──[/bold blue]")
        console.print(f"   Experiment: {experiment}")
        _setup_and_display_kb(experiment, console, embedding_backend)
        console.print(f"   Output: {model_target_dir}")
        console.print(f"   [dim]Parallel batch size: {batch_size}[/dim]\n")

        print_lock = threading.Lock()
        records: list[dict] = []

        from src.rag.retriever import get_active_kb_toggles as _get_kb_toggles
        _kb_toggles = _get_kb_toggles(experiment)
        _prompt_builder = PromptBuilder()

        def _translate_one(py_file: Path, stagger_delay: float = 0) -> dict:
            time.sleep(stagger_delay)
            translator = _agents.create_translation_agent(model, kb_toggles=_kb_toggles)
            log_translation_start(py_file.name, provider_key, variant_key)
            python_code = py_file.read_text(encoding="utf-8")
            target_file = mirror_path(py_file, source_dir, model_target_dir, ".go")
            target_file.parent.mkdir(parents=True, exist_ok=True)

            rag_result = _get_rag_result(python_code, experiment, embedding_backend)
            prompt = _prompt_builder.build_local(python_code, rag_result=rag_result)
            log_prompt(prompt)

            try:
                response = translator.run(prompt, stream=False)
                result = response.content

                if isinstance(result, TranslationResult):
                    target_file.write_text(result.go_code, encoding="utf-8")
                    log_response(result, str(target_file))
                    log_translation_done(py_file.name, str(target_file))
                    return {"source": str(py_file), "target": str(target_file), "status": "ok"}
                else:
                    log_response(result)
                    return {"source": str(py_file), "target": str(target_file), "status": "no structured output"}
            except Exception as e:
                log_translation_error(py_file.name, e)
                return {"source": str(py_file), "target": str(target_file), "status": f"error: {e}"}

        with Progress() as progress:
            task = progress.add_task(f"Translating ({label})...", total=len(py_files))

            try:
                with ThreadPoolExecutor(max_workers=batch_size) as executor:
                    futures = {
                        executor.submit(_translate_one, py_file, (i % batch_size) * stagger): py_file
                        for i, py_file in enumerate(py_files)
                    }

                    for future in as_completed(futures):
                        py_file = futures[future]
                        try:
                            record = future.result()
                            status_tag = (
                                "[green]OK[/green]" if record["status"] == "ok"
                                else "[yellow]WARN[/yellow]" if record["status"] == "no structured output"
                                else "[red]FAIL[/red]"
                            )
                            with print_lock:
                                progress.console.print(f"  {status_tag} {py_file.name}")
                                if record["status"] not in ("ok",):
                                    progress.console.print(f"         [dim]{record['status']}[/dim]")
                            records.append(record)
                        except Exception as exc:
                            with print_lock:
                                progress.console.print(f"  [red]ERROR[/red] {py_file.name}: {exc}")
                            records.append({
                                "source": str(py_file), "target": "", "status": f"error: {exc}",
                            })

                        progress.advance(task)
            except KeyboardInterrupt:
                console.print("\n[yellow]Interrupted. Cancelling pending tasks...[/yellow]")
                executor.shutdown(wait=False, cancel_futures=True)

        ok_count = sum(1 for r in records if r["status"] == "ok")
        console.print(
            f"\n[bold]{label}:[/bold] {ok_count}/{len(records)} files translated successfully."
        )


def _translate_humaneval_x(
    target_dir: Path,
    skip_preflight: bool = False,
    sample: int | None = None,
    experiment: str = "baseline",
    embedding_backend: str = "chromadb",
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

    translation_cfg = load_eval_config()["translation"]
    batch_size = translation_cfg["batch_size"]
    stagger = translation_cfg.get("thread_stagger_seconds", 1)

    for provider_key, variant_key, model in enabled:
        model_target_dir = target_dir / provider_key / variant_key / experiment
        model_target_dir.mkdir(parents=True, exist_ok=True)
        label = f"{provider_key}/{variant_key}"
        console.print(f"\n[bold blue]── Model: {label} ──[/bold blue]")
        console.print(f"   Experiment: {experiment}")
        _setup_and_display_kb(experiment, console, embedding_backend)
        console.print(f"   Output: {model_target_dir}")
        console.print(f"   [dim]Parallel batch size: {batch_size}[/dim]\n")

        print_lock = threading.Lock()
        records: list[dict] = []

        from src.rag.retriever import get_active_kb_toggles as _get_kb_toggles
        _kb_toggles = _get_kb_toggles(experiment)
        _prompt_builder = PromptBuilder()

        def _translate_one(pair: dict, stagger_delay: float = 0) -> dict:
            time.sleep(stagger_delay)
            translator = _agents.create_translation_agent(model, kb_toggles=_kb_toggles)
            task_num = pair["task_id"].split("/")[1]
            target_file = model_target_dir / f"Go_{task_num}.go"
            log_translation_start(pair["task_id"], provider_key, variant_key)

            rag_result = _get_rag_result(pair["py_solution"], experiment, embedding_backend)
            prompt = _prompt_builder.build_humaneval_x(
                pair["py_solution"],
                go_signature=pair["declaration"],
                rag_result=rag_result,
            )
            log_prompt(prompt)

            try:
                response = translator.run(prompt, stream=False)
                result = response.content

                if isinstance(result, TranslationResult):
                    target_file.write_text(result.go_code, encoding="utf-8")
                    log_response(result, str(target_file))
                    log_translation_done(pair["task_id"], str(target_file))
                    return {"source": pair["task_id"], "target": str(target_file), "status": "ok"}
                else:
                    log_response(result)
                    return {"source": pair["task_id"], "target": str(target_file), "status": "no structured output"}
            except Exception as e:
                log_translation_error(pair["task_id"], e)
                return {"source": pair["task_id"], "target": str(target_file), "status": f"error: {e}"}

        with Progress() as progress:
            task = progress.add_task(f"Translating ({label})...", total=len(pairs))

            try:
                with ThreadPoolExecutor(max_workers=batch_size) as executor:
                    futures = {
                        executor.submit(_translate_one, pair, (i % batch_size) * stagger): pair
                        for i, pair in enumerate(pairs)
                    }

                    for future in as_completed(futures):
                        pair = futures[future]
                        try:
                            record = future.result()
                            status_tag = (
                                "[green]OK[/green]" if record["status"] == "ok"
                                else "[yellow]WARN[/yellow]" if record["status"] == "no structured output"
                                else "[red]FAIL[/red]"
                            )
                            with print_lock:
                                progress.console.print(
                                    f"  {status_tag} {pair['task_id']} -> Go_{pair['task_id'].split('/')[1]}.go"
                                )
                                if record["status"] not in ("ok",):
                                    progress.console.print(f"         [dim]{record['status']}[/dim]")
                            records.append(record)
                        except Exception as exc:
                            with print_lock:
                                progress.console.print(f"  [red]ERROR[/red] {pair['task_id']}: {exc}")
                            records.append({
                                "source": pair["task_id"], "target": "", "status": f"error: {exc}",
                            })

                        progress.advance(task)
            except KeyboardInterrupt:
                console.print("\n[yellow]Interrupted. Cancelling pending tasks...[/yellow]")
                executor.shutdown(wait=False, cancel_futures=True)

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

    provider, variant, experiment = _parse_target_path(eval_target_dir)

    console.print(f"\n[bold blue]── Evaluating (local): {eval_target_dir} ──[/bold blue]")
    console.print(f"   Experiment: {experiment}")
    _setup_and_display_kb(experiment, console)
    console.print()
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

    provider, variant, experiment = _parse_target_path(eval_target_dir)

    console.print(f"\n[bold blue]── Evaluating (HumanEval-X): {eval_target_dir} ──[/bold blue]")
    console.print(f"   Experiment: {experiment}")
    _setup_and_display_kb(experiment, console)
    console.print()
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

    _reporting.display_per_file_table(records, dataset="humaneval-x")
    summary = _reporting.compute_summary(records, dataset="humaneval-x")
    _reporting.display_summary_table(summary)


# ---------------------------------------------------------------------------
# Parallel experiment runner
# ---------------------------------------------------------------------------

_ALL_EXPERIMENTS = [
    "baseline",
    "rag-pattern-only",
    "rag-pattern-samples",
    "rag-pattern-api-docs",
    "rag-full",
]


def _run_experiment_subprocess(args: tuple) -> None:
    """Top-level picklable subprocess entry point for run_all_humaneval_x.

    Must be top-level (not nested) so multiprocessing can pickle it.
    """
    provider_key, variant_key, experiment, sample, embedding_backend = args
    from src.providers.registry import enable_model
    enable_model(provider_key, variant_key)
    _translate_humaneval_x(
        HUMANEVAL_X_TARGET_DIR,
        skip_preflight=True,
        sample=sample,
        experiment=experiment,
        embedding_backend=embedding_backend,
    )


def run_all_humaneval_x(
    provider_key: str,
    variant_key: str,
    mode: str = "smoke",
    embedding_backend: str = "gemini",
) -> None:
    """Run all 5 experiments in parallel — one process per experiment.

    Args:
        provider_key: Provider identifier (e.g. "minimax").
        variant_key: Variant identifier (e.g. "M2.5").
        mode: "smoke" translates 10 files; "full" translates all 164.
        embedding_backend: "gemini" (Vertex AI) or "chromadb".
    """
    sample = 10 if mode == "smoke" else None
    console = Console()

    # Preflight once in main process
    from src.providers.registry import enable_model, get_enabled_models
    enable_model(provider_key, variant_key)
    enabled = get_enabled_models()
    console.print(
        f"\n[bold]run-all: {provider_key}/{variant_key}  mode={mode}  "
        f"experiments={len(_ALL_EXPERIMENTS)}[/bold]\n"
    )
    if not preflight_check(console, enabled):
        return

    # Spawn one process per experiment
    processes: dict[str, multiprocessing.Process] = {}
    for experiment in _ALL_EXPERIMENTS:
        args = (provider_key, variant_key, experiment, sample, embedding_backend)
        p = multiprocessing.Process(target=_run_experiment_subprocess, args=(args,))
        p.start()
        processes[experiment] = p

    def _make_status_table() -> Table:
        table = Table(title=f"run-all  {provider_key}/{variant_key}  [{mode}]")
        table.add_column("Experiment", style="cyan", min_width=24)
        table.add_column("Status")
        for exp, p in processes.items():
            if p.is_alive():
                status = "[yellow]⟳ Running[/yellow]"
            elif p.exitcode == 0:
                status = "[green]✓ Done[/green]"
            else:
                status = f"[red]✗ Failed (exit {p.exitcode})[/red]"
            table.add_row(exp, status)
        return table

    # Live status loop
    with Live(_make_status_table(), refresh_per_second=2, console=console) as live:
        while any(p.is_alive() for p in processes.values()):
            time.sleep(0.5)
            live.update(_make_status_table())

    console.print()
    console.print(_make_status_table())

    failed = [exp for exp, p in processes.items() if p.exitcode != 0]
    if failed:
        console.print(f"\n[red]Failed: {', '.join(failed)}[/red]")
    else:
        console.print(f"\n[green]All {len(_ALL_EXPERIMENTS)} experiments completed.[/green]")
