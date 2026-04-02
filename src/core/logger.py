"""Verbose pipeline logger with Rich formatting.

Toggle via VERBOSE_LOG env var or programmatically with set_verbose().
All log_* functions are no-ops when verbose is disabled.
"""

import os
import threading

from rich.console import Console

_verbose = bool(os.environ.get("VERBOSE_LOG", ""))
_console = Console()
_log_lock = threading.Lock()


def set_verbose(enabled: bool) -> None:
    """Enable or disable verbose logging."""
    global _verbose
    _verbose = enabled


def is_verbose() -> bool:
    """Return current verbose state."""
    return _verbose


# ---------------------------------------------------------------------------
# Translation logging
# ---------------------------------------------------------------------------

def log_translation_start(file_name: str, provider: str, variant: str) -> None:
    if not _verbose:
        return
    with _log_lock:
        _console.print(
            f"\n[bold cyan]▶ TRANSLATE[/bold cyan] {file_name}  "
            f"[dim]({provider}/{variant})[/dim]"
        )


def log_prompt(prompt: str) -> None:
    if not _verbose:
        return
    with _log_lock:
        _console.rule("[bold cyan]Final Prompt[/bold cyan]")
        _console.print(prompt)
        _console.rule()


def log_rag_retrieval(rag_result: object) -> None:
    """Log raw items retrieved from each RAG knowledge base."""
    if not _verbose:
        return
    with _log_lock:
        _console.rule("[bold cyan]RAG Retrieval Details[/bold cyan]")

        # Grammar Mappings
        grammar = getattr(rag_result, "grammar_mappings", [])
        _console.print(f"  [bold]Grammar Mappings[/bold] ({len(grammar)} retrieved)")
        for i, s in enumerate(grammar, 1):
            _console.print(f"    [{i}] category: {s.get('category', 'N/A')}")
            _console.print(f"        Python: {s.get('python_pattern', '')[:120]}...")
            _console.print(f"        Go:     {s.get('go_pattern', '')[:120]}...")

        # Parallel Corpus
        corpus = getattr(rag_result, "parallel_corpus", [])
        _console.print(f"  [bold]Parallel Corpus[/bold] ({len(corpus)} retrieved)")
        for i, s in enumerate(corpus, 1):
            _console.print(f"    [{i}] problem: {s.get('_id', 'N/A')}")
            _console.print(f"        Python: {s.get('python_code', '')[:120]}...")
            _console.print(f"        Go:     {s.get('go_code', '')[:120]}...")

        # API Mappings
        mappings = getattr(rag_result, "api_mappings", [])
        _console.print(f"  [bold]API Mappings[/bold] ({len(mappings)} retrieved)")
        for i, m in enumerate(mappings, 1):
            _console.print(
                f"    [{i}] {m.get('python_api', '?')} -> {m.get('go_api', '?')}: "
                f"{m.get('description', '')}"
            )

        # Documentation
        docs = getattr(rag_result, "documentation", [])
        _console.print(f"  [bold]Documentation[/bold] ({len(docs)} retrieved)")
        for i, d in enumerate(docs, 1):
            _console.print(f"    [{i}] {d.get('api', '?')}: {d.get('description', '')}")
            if d.get("example"):
                _console.print(f"        example: {d['example'][:120]}...")

        _console.rule()


def log_response(result: object, target_path: str | None = None) -> None:
    if not _verbose:
        return
    type_name = type(result).__name__
    with _log_lock:
        _console.print(f"  [dim]response type:[/dim] {type_name}")
        if target_path:
            _console.print(f"  [dim]written to:[/dim] {target_path}")


def log_translation_done(file_name: str, target_path: str) -> None:
    if not _verbose:
        return
    with _log_lock:
        _console.print(f"  [green]✓[/green] {file_name} → {target_path}")


def log_translation_error(file_name: str, error: Exception) -> None:
    if not _verbose:
        return
    with _log_lock:
        _console.print(f"  [red]✗ ERROR[/red] {file_name}: {error}")


# ---------------------------------------------------------------------------
# Evaluation logging
# ---------------------------------------------------------------------------

def log_eval_start(file_name: str) -> None:
    if not _verbose:
        return
    with _log_lock:
        _console.print(f"\n[bold cyan]▶ EVALUATE[/bold cyan] {file_name}")


def log_eval_success(file_name: str) -> None:
    if not _verbose:
        return
    with _log_lock:
        _console.print(f"  [green]✓ PASS[/green] {file_name}")


def log_eval_error(file_name: str, err_type: str, stderr: str) -> None:
    if not _verbose:
        return
    with _log_lock:
        _console.print(f"  [red]✗ {err_type}[/red] {file_name}")
        if stderr:
            truncated = stderr[:300] + ("..." if len(stderr) > 300 else "")
            _console.print(f"  [dim]{truncated}[/dim]")
