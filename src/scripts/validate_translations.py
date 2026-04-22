#!/usr/bin/env python3
"""Validate translated Go files for a HumanEval-X run.

Flags files that are missing, truncated, prose-only, or otherwise not parseable
as Go source. Returns a list of problem numbers to re-translate.

Usage:
    uv run python src/scripts/validate_translations.py \\
        --target-dir data/translation/target/humaneval-x/minimax/M2.5/vec-chroma-3072/run-101/rag-pattern-only \\
        --expected 164

Exit code 0 if everything is valid; 1 if any file is invalid (list printed on
stdout as comma-separated problem numbers for use with `--problems`).
"""

from __future__ import annotations

import json
import re
import sys
from pathlib import Path

import click


_PACKAGE_RE = re.compile(r"(?m)^\s*package\s+[A-Za-z_][A-Za-z0-9_]*\s*$")
_FUNC_RE = re.compile(r"(?m)^\s*func\s+")
_ERROR_HINTS = (
    "\"error\":", "i'm sorry", "i am sorry", "i apologize",
    "rate limit", "429", "unauthorized", "unavailable",
)


def _classify(path: Path) -> tuple[bool, str]:
    """Return (is_valid, reason). A file is valid if it looks like Go source."""
    if not path.exists():
        return False, "missing"
    try:
        text = path.read_text(encoding="utf-8")
    except Exception as exc:
        return False, f"unreadable: {exc}"

    stripped = text.strip()
    if not stripped:
        return False, "empty"
    if len(stripped) < 30:
        return False, "too_short"

    if not _PACKAGE_RE.search(text):
        return False, "no_package_decl"
    if not _FUNC_RE.search(text):
        return False, "no_func_decl"

    lowered = stripped.lower()
    for hint in _ERROR_HINTS:
        if hint in lowered:
            return False, f"error_hint:{hint}"

    return True, "ok"


def _task_num(task_dir: Path) -> int | None:
    name = task_dir.name
    if not name.startswith("Go_"):
        return None
    suffix = name.split("_", 1)[1]
    if not suffix.isdigit():
        return None
    return int(suffix)


@click.command()
@click.option("--target-dir", required=True, type=click.Path(path_type=Path, exists=True))
@click.option("--expected", type=int, default=164, show_default=True,
              help="Expected number of tasks (0..expected-1).")
@click.option("--report-json", type=click.Path(path_type=Path), default=None,
              help="Optional path to write a JSON report.")
def main(target_dir: Path, expected: int, report_json: Path | None) -> None:
    tasks_dir = target_dir / "tasks"
    if not tasks_dir.is_dir():
        click.echo(f"No tasks/ directory under {target_dir}", err=True)
        sys.exit(2)

    problems: list[tuple[int, str]] = []
    seen: set[int] = set()

    for task_dir in sorted(tasks_dir.iterdir()):
        if not task_dir.is_dir():
            continue
        n = _task_num(task_dir)
        if n is None:
            continue
        seen.add(n)
        go_file = task_dir / "translation.go"
        ok, reason = _classify(go_file)
        if not ok:
            problems.append((n, reason))

    missing = sorted(set(range(expected)) - seen)
    for n in missing:
        problems.append((n, "missing_task_dir"))

    problems.sort()

    click.echo(f"Checked {target_dir}")
    click.echo(f"  Expected tasks: {expected}")
    click.echo(f"  Seen tasks:     {len(seen)}")
    click.echo(f"  Invalid:        {len(problems)}")
    for n, reason in problems:
        click.echo(f"    Go_{n}: {reason}")

    if report_json is not None:
        report_json.parent.mkdir(parents=True, exist_ok=True)
        report_json.write_text(
            json.dumps(
                {
                    "target_dir": str(target_dir),
                    "expected": expected,
                    "seen": sorted(seen),
                    "invalid": [{"task": n, "reason": r} for n, r in problems],
                },
                indent=2,
            ),
            encoding="utf-8",
        )

    if problems:
        ids = ",".join(str(n) for n, _ in problems)
        click.echo(f"\nRetry list: {ids}")
        sys.exit(1)
    sys.exit(0)


if __name__ == "__main__":
    main()
