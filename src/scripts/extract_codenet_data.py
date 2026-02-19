#!/usr/bin/env python3
"""
Extract parallel Python and Go code pairs from the IBM CodeNet dataset.

For each problem that has both Python and Go accepted submissions,
selects the shortest (by code_size) accepted submission for each language
and writes the pair to a JSON lines file.

Usage:
    uv run src/scripts/extract_codenet_data.py
"""

import json
import os
import sys
from pathlib import Path

import pandas as pd
from rich.console import Console
from rich.progress import BarColumn, MofNCompleteColumn, Progress, TextColumn, TimeRemainingColumn

console = Console()

# ---------------------------------------------------------------------------
# Paths
# ---------------------------------------------------------------------------
REPO_ROOT = Path(__file__).resolve().parents[2]
CODENET_ROOT = REPO_ROOT / "data" / "RAG" / "unprocessed" / "Project_CodeNet"
PROBLEM_LIST_CSV = CODENET_ROOT / "metadata" / "problem_list.csv"
METADATA_DIR = CODENET_ROOT / "metadata"
DATA_DIR = CODENET_ROOT / "data"
OUTPUT_DIR = REPO_ROOT / "data" / "processed" / "parallel_corpus" / "codeNet"
OUTPUT_FILE = OUTPUT_DIR / "python_go_pairs.jsonl"


def read_accepted_submissions(problem_id: str, language: str) -> pd.DataFrame:
    """Return a DataFrame of accepted submissions for a given problem and language."""
    meta_csv = METADATA_DIR / f"{problem_id}.csv"
    if not meta_csv.exists():
        return pd.DataFrame()

    try:
        df = pd.read_csv(meta_csv, dtype=str)
    except Exception:
        return pd.DataFrame()

    mask = (df["language"] == language) & (df["status"] == "Accepted")
    return df[mask].copy()


def shortest_accepted_code(
    problem_id: str, language: str, lang_dir: Path, ext: str
) -> str | None:
    """
    Find the shortest accepted submission for a problem/language pair.

    Returns the source code as a string, or None if nothing is found.
    """
    accepted = read_accepted_submissions(problem_id, language)
    if accepted.empty:
        return None

    # Sort by code_size ascending so we try smallest first
    accepted["code_size"] = pd.to_numeric(accepted["code_size"], errors="coerce")
    accepted = accepted.dropna(subset=["code_size"]).sort_values("code_size")

    for _, row in accepted.iterrows():
        sid = row["submission_id"]
        filepath = lang_dir / f"{sid}.{ext}"
        if not filepath.exists():
            continue
        try:
            return filepath.read_text(encoding="utf-8", errors="replace")
        except Exception:
            continue

    return None


def read_description(problem_id: str) -> str:
    """Read the HTML problem description if it exists."""
    desc_path = DATA_DIR / problem_id / "description.html"
    if desc_path.exists():
        try:
            return desc_path.read_text(encoding="utf-8", errors="replace")
        except Exception:
            pass
    return ""


def main() -> None:
    console.print(f"[bold cyan]CodeNet Parallel Pair Extractor[/bold cyan]")
    console.print(f"Reading problem list from: [green]{PROBLEM_LIST_CSV}[/green]")

    if not PROBLEM_LIST_CSV.exists():
        console.print(f"[red]ERROR:[/red] problem_list.csv not found at {PROBLEM_LIST_CSV}")
        sys.exit(1)

    problem_list = pd.read_csv(PROBLEM_LIST_CSV, dtype=str)
    problem_ids = problem_list["id"].dropna().tolist()
    console.print(f"Total problems: [yellow]{len(problem_ids)}[/yellow]")

    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)

    pairs_found = 0
    skipped_no_dirs = 0
    skipped_no_accepted = 0

    progress = Progress(
        TextColumn("[progress.description]{task.description}"),
        BarColumn(),
        MofNCompleteColumn(),
        TimeRemainingColumn(),
        console=console,
    )

    with OUTPUT_FILE.open("w", encoding="utf-8") as out_f, progress:
        task = progress.add_task("Processing problems", total=len(problem_ids))

        for problem_id in problem_ids:
            python_dir = DATA_DIR / problem_id / "Python"
            go_dir = DATA_DIR / problem_id / "Go"

            if not python_dir.is_dir() or not go_dir.is_dir():
                skipped_no_dirs += 1
                progress.advance(task)
                continue

            python_code = shortest_accepted_code(problem_id, "Python", python_dir, "py")
            go_code = shortest_accepted_code(problem_id, "Go", go_dir, "go")

            if python_code is None or go_code is None:
                skipped_no_accepted += 1
                progress.advance(task)
                continue

            description = read_description(problem_id)

            record = {
                "problem_id": problem_id,
                "python_code": python_code,
                "go_code": go_code,
                "problem_description": description,
            }
            out_f.write(json.dumps(record, ensure_ascii=False) + "\n")
            pairs_found += 1
            progress.advance(task)

    console.print(f"\n[bold green]Done![/bold green]")
    console.print(f"  Pairs extracted   : [green]{pairs_found}[/green]")
    console.print(f"  Skipped (no dirs) : [yellow]{skipped_no_dirs}[/yellow]")
    console.print(f"  Skipped (no acc.) : [yellow]{skipped_no_accepted}[/yellow]")
    console.print(f"  Output file       : [cyan]{OUTPUT_FILE}[/cyan]")


if __name__ == "__main__":
    main()
