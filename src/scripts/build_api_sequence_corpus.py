#!/usr/bin/env python3
"""Build a Go API-sequence retrieval corpus from Project CodeNet Go files.

Usage:
    uv run python src/scripts/build_api_sequence_corpus.py
    uv run python src/scripts/build_api_sequence_corpus.py --codenet-root /path/to/Project_CodeNet/data
"""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from rich.console import Console
from rich.progress import Progress

from src.config import GO_API_SEQUENCES_FILE
from src.rag.api_extractor import extract_go_api_sequences_from_file

console = Console()

DEFAULT_CODENET_GO_ROOT = Path("/Volumes/MyZhiTai/DEV/www/thesis/Project_CodeNet/data")


def build_api_sequence_records(go_root: Path, *, limit_files: int | None = None) -> list[dict]:
    records: list[dict] = []
    seen_sequences: set[str] = set()

    go_files = sorted(go_root.glob("*/Go/*.go"))
    if limit_files is not None:
        go_files = go_files[:limit_files]

    with Progress(console=console) as progress:
        task = progress.add_task("Extracting Go API sequences", total=len(go_files))
        for file_index, go_file in enumerate(go_files, start=1):
            try:
                extracted = extract_go_api_sequences_from_file(go_file)
            except Exception as exc:
                console.print(f"[yellow]Skip[/yellow] {go_file}: {exc}")
                progress.advance(task)
                continue

            for function_index, record in enumerate(extracted, start=1):
                sequence_text = record["sequence_text"]
                if sequence_text in seen_sequences:
                    continue
                seen_sequences.add(sequence_text)
                records.append(
                    {
                        "_id": f"api_seq_go_{len(records) + 1:06d}",
                        "language": "go",
                        "source_corpus": "project_codenet_go",
                        "file_path": str(go_file.relative_to(go_root)),
                        "function_name": record["function_name"],
                        "sequence_text": sequence_text,
                        "apis": record["apis"],
                        "imports": record["imports"],
                        "function_code": record["function_code"],
                    }
                )
            progress.advance(task)

    return records


def write_jsonl(records: list[dict], output_path: Path) -> None:
    output_path.parent.mkdir(parents=True, exist_ok=True)
    with output_path.open("w", encoding="utf-8") as handle:
        for record in records:
            handle.write(json.dumps(record, ensure_ascii=False) + "\n")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--codenet-root",
        type=Path,
        default=DEFAULT_CODENET_GO_ROOT,
        help="Path to Project_CodeNet/data root containing <problem>/Go/*.go files.",
    )
    parser.add_argument(
        "--output",
        type=Path,
        default=GO_API_SEQUENCES_FILE,
        help="Output JSONL path.",
    )
    parser.add_argument(
        "--limit-files",
        type=int,
        default=None,
        help="Optional limit for smoke-testing the builder.",
    )
    return parser.parse_args()


def main() -> None:
    args = parse_args()
    go_root = args.codenet_root
    if not go_root.is_dir():
        raise FileNotFoundError(f"CodeNet Go root not found: {go_root}")

    console.print(f"[bold cyan]Building Go API sequence corpus[/bold cyan]")
    console.print(f"Source: [green]{go_root}[/green]")
    console.print(f"Output: [green]{args.output}[/green]")

    records = build_api_sequence_records(go_root, limit_files=args.limit_files)
    write_jsonl(records, args.output)

    console.print(f"[bold green]Done[/bold green] wrote {len(records)} records")


if __name__ == "__main__":
    main()
