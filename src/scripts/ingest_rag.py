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
    GO_API_SEQUENCES_FILE,
    GRAMMAR_MAPPINGS_FILE,
    PARALLEL_CORPUS_FILE,
)
from src.rag.embeddings import get_embedding_function, load_rag_config
from src.rag.store import collection_name_with_dim, get_chroma_client, get_or_create_collection

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


def ingest_grammar_mappings(client, ef, dimensions: int | None = None):
    coll_name = collection_name_with_dim("grammar_mappings", dimensions)
    console.print(f"\n[bold]Ingesting grammar mappings[/bold] from {GRAMMAR_MAPPINGS_FILE}")
    records = _load_jsonl(GRAMMAR_MAPPINGS_FILE)
    console.print(f"  Loaded {len(records)} records")

    # Delete stale collection to remove entries from removed categories
    try:
        client.delete_collection(coll_name)
        console.print(f"  Cleared existing {coll_name} collection.")
    except Exception:
        pass  # Collection doesn't exist yet
    collection = get_or_create_collection(client, coll_name, ef)

    ids = []
    documents = []
    metadatas = []
    for i, r in enumerate(records):
        doc_id = f"grammar_{r['category']}_{i}"
        ids.append(doc_id)
        documents.append(f"Category: {r['category']}\n{r['python_pattern']}")
        metadatas.append({
            "category": r["category"],
            "python_pattern": r["python_pattern"],
            "go_pattern": r["go_pattern"],
            "description": r["description"],
        })

    _upsert_batched(collection, ids, documents, metadatas)
    console.print(f"  [green]Done![/green] Collection '{collection.name}' has {collection.count()} entries.")


def ingest_api_mappings(client, ef, dimensions: int | None = None):
    coll_name = collection_name_with_dim("api_mappings", dimensions)
    console.print(f"\n[bold]Ingesting API mappings[/bold] from {API_MAPPINGS_FILE}")
    records = _load_jsonl(API_MAPPINGS_FILE)
    console.print(f"  Loaded {len(records)} records")

    try:
        client.delete_collection(coll_name)
        console.print(f"  Cleared existing {coll_name} collection.")
    except Exception:
        pass
    collection = get_or_create_collection(client, coll_name, ef)

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


def ingest_parallel_corpus(client, ef, dimensions: int | None = None):
    coll_name = collection_name_with_dim("parallel_corpus", dimensions)
    console.print(f"\n[bold]Ingesting parallel corpus[/bold] from {PARALLEL_CORPUS_FILE}")
    records = _load_jsonl(PARALLEL_CORPUS_FILE)
    console.print(f"  Loaded {len(records)} records")

    try:
        client.delete_collection(coll_name)
        console.print(f"  Cleared existing {coll_name} collection.")
    except Exception:
        pass
    collection = get_or_create_collection(client, coll_name, ef)

    ids = []
    documents = []
    metadatas = []
    for i, r in enumerate(records):
        ids.append(f"parallel_{r['problem_id']}_{i}")
        documents.append(r["python_code"])
        metadatas.append({
            "problem_id": r["problem_id"],
            "go_code": r["go_code"],
            "problem_description": r.get("problem_description", ""),
        })

    _upsert_batched(collection, ids, documents, metadatas)
    console.print(f"  [green]Done![/green] Collection '{collection.name}' has {collection.count()} entries.")


def ingest_go_docs(client, ef, dimensions: int | None = None):
    coll_name = collection_name_with_dim("go_docs", dimensions)
    console.print(f"\n[bold]Ingesting Go docs[/bold] from {GO_DOCS_FILE}")
    records = _load_jsonl(GO_DOCS_FILE)
    console.print(f"  Loaded {len(records)} records")

    try:
        client.delete_collection(coll_name)
        console.print(f"  Cleared existing {coll_name} collection.")
    except Exception:
        pass
    collection = get_or_create_collection(client, coll_name, ef)

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


def ingest_api_sequences(client, ef, dimensions: int | None = None):
    coll_name = collection_name_with_dim("api_sequences", dimensions)
    console.print(f"\n[bold]Ingesting API sequences[/bold] from {GO_API_SEQUENCES_FILE}")
    records = _load_jsonl(GO_API_SEQUENCES_FILE)
    console.print(f"  Loaded {len(records)} records")

    try:
        client.delete_collection(coll_name)
        console.print(f"  Cleared existing {coll_name} collection.")
    except Exception:
        pass
    collection = get_or_create_collection(client, coll_name, ef)

    ids = []
    documents = []
    metadatas = []
    for i, r in enumerate(records):
        doc_id = r.get("_id") or f"api_seq_go_{i + 1:06d}"
        ids.append(doc_id)
        documents.append(r["sequence_text"])
        metadatas.append({
            "language": r.get("language", "go"),
            "source_corpus": r.get("source_corpus", ""),
            "file_path": r.get("file_path", ""),
            "function_name": r.get("function_name", ""),
            "apis": json.dumps(r.get("apis", []), ensure_ascii=True),
            "imports": json.dumps(r.get("imports", []), ensure_ascii=True),
        })

    _upsert_batched(collection, ids, documents, metadatas)
    console.print(f"  [green]Done![/green] Collection '{collection.name}' has {collection.count()} entries.")


@click.command()
@click.option(
    "--collection",
    type=click.Choice(["grammar_mappings", "parallel_corpus", "api_mappings", "go_docs", "api_sequences", "all"]),
    default="all",
    help="Which collection(s) to ingest.",
)
@click.option(
    "--dimensions",
    type=int,
    default=None,
    help="Override embedding dimensions (Gemini only). Supported: 768, 1050, 1536, 3072. "
         "Defaults to the value in rag_config.yaml.",
)
def main(collection: str, dimensions: int | None):
    """Ingest RAG data sources into ChromaDB.

    Embedding provider is controlled by config/rag_config.yaml.
    Set provider to 'gemini' to use Gemini embeddings stored in ChromaDB.
    Set provider to 'default' for local all-MiniLM-L6-v2 (384 dims, no API key).

    Use --dimensions to create separate collections per dimension for ablation
    experiments (e.g. grammar_mappings_768, grammar_mappings_1536).
    """
    cfg = load_rag_config()
    provider = cfg["embedding"]["provider"]
    effective_dims = dimensions or cfg["embedding"].get("gemini", {}).get("dimensions", 3072)
    console.print(f"Embedding provider: [bold]{provider}[/bold]")
    console.print(f"Embedding dimensions: [bold]{effective_dims}[/bold]")

    ef = get_embedding_function(dimensions_override=dimensions)
    client = get_chroma_client()

    if collection in ("grammar_mappings", "all"):
        ingest_grammar_mappings(client, ef, dimensions)
    if collection in ("parallel_corpus", "all"):
        ingest_parallel_corpus(client, ef, dimensions)
    if collection in ("api_mappings", "all"):
        ingest_api_mappings(client, ef, dimensions)
    if collection in ("go_docs", "all"):
        ingest_go_docs(client, ef, dimensions)
    if collection in ("api_sequences", "all"):
        ingest_api_sequences(client, ef, dimensions)

    console.print("\n[bold green]All done![/bold green]")


if __name__ == "__main__":
    main()
