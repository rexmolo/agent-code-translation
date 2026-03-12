#!/usr/bin/env python3
"""Ingest RAG data into ChromaDB collections.

Usage:
    uv run python src/scripts/ingest_rag.py [--collection parallel_corpus|api_mappings|go_docs|all]
"""

import json

import click
from rich.console import Console
from rich.progress import Progress

from src.config import (
    API_MAPPINGS_FILE,
    GO_DOCS_FILE,
    PARALLEL_CORPUS_FILE,
)
from src.rag.embeddings import get_embedding_function, load_rag_config
from src.rag.store import get_chroma_client, get_or_create_collection

console = Console()
BATCH_SIZE = 100


def _load_jsonl(path) -> list[dict]:
    records = []
    with open(path, encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if line:
                records.append(json.loads(line))
    return records


def _upsert_batched(collection, ids, documents, metadatas):
    """Upsert into ChromaDB in batches."""
    total = len(ids)
    with Progress() as progress:
        task = progress.add_task(f"Upserting {collection.name}...", total=total)
        for start in range(0, total, BATCH_SIZE):
            end = min(start + BATCH_SIZE, total)
            collection.upsert(
                ids=ids[start:end],
                documents=documents[start:end],
                metadatas=metadatas[start:end],
            )
            progress.update(task, advance=end - start)


def ingest_parallel_corpus(client, ef):
    console.print(f"\n[bold]Ingesting parallel corpus[/bold] from {PARALLEL_CORPUS_FILE}")
    records = _load_jsonl(PARALLEL_CORPUS_FILE)
    console.print(f"  Loaded {len(records)} records")

    collection = get_or_create_collection(client, "parallel_corpus", ef)

    ids = []
    documents = []
    metadatas = []
    for r in records:
        doc_id = f"corpus_{r['problem_id']}"
        go_code = r["go_code"]
        # Truncate go_code in metadata if too large (ChromaDB limit)
        if len(go_code) > 8000:
            go_code = go_code[:8000] + "\n// ... truncated"
        ids.append(doc_id)
        documents.append(r["python_code"])
        metadatas.append({
            "problem_id": r["problem_id"],
            "go_code": go_code,
            "problem_description": r.get("problem_description", ""),
        })

    _upsert_batched(collection, ids, documents, metadatas)
    console.print(f"  [green]Done![/green] Collection '{collection.name}' has {collection.count()} entries.")


def ingest_api_mappings(client, ef):
    console.print(f"\n[bold]Ingesting API mappings[/bold] from {API_MAPPINGS_FILE}")
    records = _load_jsonl(API_MAPPINGS_FILE)
    console.print(f"  Loaded {len(records)} records")

    collection = get_or_create_collection(client, "api_mappings", ef)

    ids = []
    documents = []
    metadatas = []
    for i, r in enumerate(records):
        doc_id = f"api_{r['category']}_{i}"
        text = f"{r['category']}: {r['python_api']} -> {r['go_api']}. {r['description']}"
        ids.append(doc_id)
        documents.append(text)
        metadatas.append({
            "category": r["category"],
            "python_api": r["python_api"],
            "go_api": r["go_api"],
            "description": r["description"],
        })

    _upsert_batched(collection, ids, documents, metadatas)
    console.print(f"  [green]Done![/green] Collection '{collection.name}' has {collection.count()} entries.")


def ingest_go_docs(client, ef):
    console.print(f"\n[bold]Ingesting Go docs[/bold] from {GO_DOCS_FILE}")
    records = _load_jsonl(GO_DOCS_FILE)
    console.print(f"  Loaded {len(records)} records")

    collection = get_or_create_collection(client, "go_docs", ef)

    ids = []
    documents = []
    metadatas = []
    for i, r in enumerate(records):
        doc_id = f"godoc_{r['package']}_{i}"
        text = f"{r['package']}: {r['api']}. {r['description']} Example: {r.get('example', '')}"
        ids.append(doc_id)
        documents.append(text)
        metadatas.append({
            "package": r["package"],
            "api": r["api"],
            "description": r["description"],
            "example": r.get("example", ""),
        })

    _upsert_batched(collection, ids, documents, metadatas)
    console.print(f"  [green]Done![/green] Collection '{collection.name}' has {collection.count()} entries.")


@click.command()
@click.option(
    "--collection",
    type=click.Choice(["parallel_corpus", "api_mappings", "go_docs", "all"]),
    default="all",
    help="Which collection(s) to ingest.",
)
def main(collection: str):
    """Ingest RAG data sources into ChromaDB."""
    cfg = load_rag_config()
    provider = cfg["embedding"]["provider"]
    console.print(f"Embedding provider: [bold]{provider}[/bold]")

    ef = get_embedding_function()
    client = get_chroma_client()

    if collection in ("parallel_corpus", "all"):
        ingest_parallel_corpus(client, ef)
    if collection in ("api_mappings", "all"):
        ingest_api_mappings(client, ef)
    if collection in ("go_docs", "all"):
        ingest_go_docs(client, ef)

    console.print("\n[bold green]All done![/bold green]")


if __name__ == "__main__":
    main()
