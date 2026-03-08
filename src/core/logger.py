"""Verbose pipeline logger with Rich formatting.

Toggle via VERBOSE_LOG env var or programmatically with set_verbose().
All log_* functions are no-ops when verbose is disabled.
"""

import os

from rich.console import Console

_verbose = bool(os.environ.get("VERBOSE_LOG", ""))
_console = Console()


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
    _console.print(
        f"\n[bold cyan]▶ TRANSLATE[/bold cyan] {file_name}  "
        f"[dim]({provider}/{variant})[/dim]"
    )


def log_prompt(prompt: str, truncate: int = 500) -> None:
    if not _verbose:
        return
    display = prompt[:truncate] + ("..." if len(prompt) > truncate else "")
    _console.print(f"  [dim]prompt:[/dim] {display}")


def log_response(result: object, target_path: str | None = None) -> None:
    if not _verbose:
        return
    type_name = type(result).__name__
    _console.print(f"  [dim]response type:[/dim] {type_name}")
    if target_path:
        _console.print(f"  [dim]written to:[/dim] {target_path}")


def log_translation_done(file_name: str, target_path: str) -> None:
    if not _verbose:
        return
    _console.print(f"  [green]✓[/green] {file_name} → {target_path}")


def log_translation_error(file_name: str, error: Exception) -> None:
    if not _verbose:
        return
    _console.print(f"  [red]✗ ERROR[/red] {file_name}: {error}")


# ---------------------------------------------------------------------------
# Evaluation logging
# ---------------------------------------------------------------------------

def log_eval_start(file_name: str) -> None:
    if not _verbose:
        return
    _console.print(f"\n[bold cyan]▶ EVALUATE[/bold cyan] {file_name}")


def log_eval_success(file_name: str) -> None:
    if not _verbose:
        return
    _console.print(f"  [green]✓ PASS[/green] {file_name}")


def log_eval_error(file_name: str, err_type: str, stderr: str) -> None:
    if not _verbose:
        return
    _console.print(f"  [red]✗ {err_type}[/red] {file_name}")
    if stderr:
        truncated = stderr[:300] + ("..." if len(stderr) > 300 else "")
        _console.print(f"  [dim]{truncated}[/dim]")
