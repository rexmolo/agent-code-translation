#!/usr/bin/env python3
"""Filter CodeNet Python/Go pairs to problems with verified sample I/O.

This keeps only pairs whose ``problem_id`` has:
1. ``input.txt`` and ``output.txt`` under the derived CodeNet input/output tree
2. no entry in ``unverified_accepted_solutions.txt``

Usage:
    uv run python src/scripts/filter_codenet_pairs_by_verified_io.py

By default the script rewrites the existing pair file in place.
"""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from rich.console import Console

from src.config import REPO_ROOT

console = Console()

DEFAULT_PAIRS_FILE = REPO_ROOT / "data" / "processed" / "parallel_corpus" / "codeNet" / "python_go_pairs.jsonl"
DEFAULT_CODENET_ROOT = Path("/Volumes/MyZhiTai/DEV/www/thesis/Project_CodeNet")


def _load_unverified_problem_ids(path: Path) -> set[str]:
    if not path.exists():
        return set()

    ids: set[str] = set()
    for line in path.read_text(encoding="utf-8", errors="replace").splitlines():
        text = line.strip()
        if text.startswith("p"):
            ids.add(text.split()[0])
    return ids


def _has_verified_sample_io(problem_id: str, *, io_root: Path, unverified_ids: set[str]) -> bool:
    if problem_id in unverified_ids:
        return False
    problem_io_dir = io_root / problem_id
    return (problem_io_dir / "input.txt").is_file() and (problem_io_dir / "output.txt").is_file()


def filter_pairs(input_file: Path, output_file: Path, codenet_root: Path) -> tuple[int, int]:
    io_root = codenet_root / "derived" / "input_output" / "data"
    unverified_file = codenet_root / "derived" / "input_output" / "unverified_accepted_solutions.txt"
    unverified_ids = _load_unverified_problem_ids(unverified_file)

    kept = 0
    total = 0
    output_file.parent.mkdir(parents=True, exist_ok=True)

    with input_file.open(encoding="utf-8") as src, output_file.open("w", encoding="utf-8") as dst:
        for line in src:
            if not line.strip():
                continue
            total += 1
            record = json.loads(line)
            problem_id = record.get("problem_id", "")
            if _has_verified_sample_io(problem_id, io_root=io_root, unverified_ids=unverified_ids):
                dst.write(json.dumps(record, ensure_ascii=False) + "\n")
                kept += 1

    return total, kept


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--input-file",
        type=Path,
        default=DEFAULT_PAIRS_FILE,
        help="Path to the existing CodeNet Python/Go pair JSONL.",
    )
    parser.add_argument(
        "--output-file",
        type=Path,
        default=DEFAULT_PAIRS_FILE,
        help="Where to write the filtered JSONL. Defaults to in-place rewrite.",
    )
    parser.add_argument(
        "--codenet-root",
        type=Path,
        default=DEFAULT_CODENET_ROOT,
        help="Path to the local Project_CodeNet root.",
    )
    args = parser.parse_args()

    if not args.input_file.exists():
        raise FileNotFoundError(f"Input file not found: {args.input_file}")

    if args.input_file.resolve() == args.output_file.resolve():
        temp_output = args.output_file.with_suffix(args.output_file.suffix + ".tmp")
    else:
        temp_output = args.output_file

    total, kept = filter_pairs(args.input_file, temp_output, args.codenet_root)

    if temp_output != args.output_file:
        temp_output.replace(args.output_file)

    console.print(f"[green]Filtered pairs written to:[/green] {args.output_file}")
    console.print(f"Total input pairs: [cyan]{total}[/cyan]")
    console.print(f"Kept verified-I/O pairs: [cyan]{kept}[/cyan]")


if __name__ == "__main__":
    main()
