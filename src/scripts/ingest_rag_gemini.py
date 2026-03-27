#!/usr/bin/env python3
"""Ingest RAG data into Vertex AI Vector Search using Gemini embeddings.

Usage:
    uv run python src/scripts/ingest_rag_gemini.py [--collection parallel_corpus|api_mappings|go_docs|all]

Prerequisites:
    - Google credentials configured (see config/providers.yaml for env var names)
    - gcloud CLI authenticated (Application Default Credentials)

Note: First run creates the Vertex AI index and endpoint, which takes ~20-30 min.
      Subsequent runs reuse existing resources and only upsert data.
"""

import json

import click
from rich.console import Console
from rich.progress import Progress

from src.config import (
    API_MAPPINGS_FILE,
    GO_DOCS_FILE,
    GRAMMAR_MAPPINGS_FILE,
    PARALLEL_CORPUS_FILE,
)
from src.rag.embeddings import GeminiEmbeddingFunction, load_rag_config
from src.rag.vertex_store import (
    ensure_deployed,
    get_or_create_endpoint,
    get_or_create_index,
    upsert_datapoints,
)

console = Console()
EMBED_BATCH_SIZE = 100


def _load_jsonl(path) -> list[dict]:
    records = []
    with open(path, encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if line:
                records.append(json.loads(line))
    return records


def _embed_texts(ef: GeminiEmbeddingFunction, texts: list[str]) -> list[list[float]]:
    """Embed texts in batches to respect API limits."""
    all_embeddings: list[list[float]] = []
    with Progress() as progress:
        task = progress.add_task("Embedding...", total=len(texts))
        for start in range(0, len(texts), EMBED_BATCH_SIZE):
            end = min(start + EMBED_BATCH_SIZE, len(texts))
            batch = texts[start:end]
            embeddings = ef(batch)
            all_embeddings.extend(embeddings)
            progress.update(task, advance=end - start)
    return all_embeddings


def ingest_grammar_mappings(index, ef: GeminiEmbeddingFunction):
    console.print(f"\n[bold]Ingesting grammar mappings[/bold] from {GRAMMAR_MAPPINGS_FILE}")
    records = _load_jsonl(GRAMMAR_MAPPINGS_FILE)
    console.print(f"  Loaded {len(records)} records")

    ids = []
    documents = []
    for i, r in enumerate(records):
        ids.append(f"grammar_{r['category']}_{i}")
        documents.append(r["python_pattern"])

    # Note: Vertex AI upsert overwrites existing IDs but does not remove stale ones.
    # If the index previously had more grammar entries, consider recreating the index
    # to remove stale datapoints from removed categories.

    console.print("  Generating Gemini embeddings...")
    embeddings = _embed_texts(ef, documents)

    console.print("  Upserting to Vertex AI...")
    upsert_datapoints(index, ids, embeddings, "grammar_mappings")
    console.print(f"  [green]Done![/green] Upserted {len(ids)} grammar mapping entries.")


def ingest_api_mappings(index, ef: GeminiEmbeddingFunction):
    console.print(f"\n[bold]Ingesting API mappings[/bold] from {API_MAPPINGS_FILE}")
    records = _load_jsonl(API_MAPPINGS_FILE)
    console.print(f"  Loaded {len(records)} records")

    ids = []
    documents = []
    for i, r in enumerate(records):
        ids.append(f"api_{r['category']}_{i}")
        documents.append(r["python_api"])

    console.print("  Generating Gemini embeddings...")
    embeddings = _embed_texts(ef, documents)

    console.print("  Upserting to Vertex AI...")
    upsert_datapoints(index, ids, embeddings, "api_mappings")
    console.print(f"  [green]Done![/green] Upserted {len(ids)} API mapping entries.")


def ingest_parallel_corpus(index, ef: GeminiEmbeddingFunction):
    console.print(f"\n[bold]Ingesting parallel corpus[/bold] from {PARALLEL_CORPUS_FILE}")
    records = _load_jsonl(PARALLEL_CORPUS_FILE)
    console.print(f"  Loaded {len(records)} records")

    ids = []
    documents = []
    for i, r in enumerate(records):
        ids.append(f"parallel_{r['problem_id']}_{i}")
        documents.append(r["python_code"])

    console.print("  Generating Gemini embeddings...")
    embeddings = _embed_texts(ef, documents)

    console.print("  Upserting to Vertex AI...")
    upsert_datapoints(index, ids, embeddings, "parallel_corpus")
    console.print(f"  [green]Done![/green] Upserted {len(ids)} parallel corpus entries.")


def ingest_go_docs(index, ef: GeminiEmbeddingFunction):
    console.print(f"\n[bold]Ingesting Go docs[/bold] from {GO_DOCS_FILE}")
    records = _load_jsonl(GO_DOCS_FILE)
    console.print(f"  Loaded {len(records)} records")

    ids = []
    documents = []
    for i, r in enumerate(records):
        ids.append(f"godoc_{r['package']}_{i}")
        text = f"{r['package']}: {r['api']}. {r['description']} Example: {r.get('example', '')}"
        documents.append(text)

    console.print("  Generating Gemini embeddings...")
    embeddings = _embed_texts(ef, documents)

    console.print("  Upserting to Vertex AI...")
    upsert_datapoints(index, ids, embeddings, "go_docs")
    console.print(f"  [green]Done![/green] Upserted {len(ids)} Go docs entries.")


@click.command()
@click.option(
    "--collection",
    type=click.Choice(["grammar_mappings", "api_mappings", "go_docs", "parallel_corpus", "all"]),
    default="all",
    help="Which collection(s) to ingest.",
)
def main(collection: str):
    """Ingest RAG data sources into Vertex AI Vector Search with Gemini embeddings."""
    cfg = load_rag_config()
    model = cfg["embedding"]["gemini"]["model"]
    console.print(f"Embedding model: [bold]{model}[/bold]")

    # Initialise Gemini embedding function
    ef = GeminiEmbeddingFunction(model_name=model)

    # Set up Vertex AI resources
    console.print("\n[bold]Setting up Vertex AI Vector Search...[/bold]")
    index = get_or_create_index()
    endpoint = get_or_create_endpoint()
    deployed_id = ensure_deployed(index, endpoint)

    # Ingest collections
    if collection in ("grammar_mappings", "all"):
        ingest_grammar_mappings(index, ef)
    if collection in ("api_mappings", "all"):
        ingest_api_mappings(index, ef)
    if collection in ("go_docs", "all"):
        ingest_go_docs(index, ef)
    if collection in ("parallel_corpus", "all"):
        ingest_parallel_corpus(index, ef)

    console.print("\n[bold green]All done![/bold green]")
    console.print(f"  Index: {index.resource_name}")
    console.print(f"  Endpoint: {endpoint.resource_name}")
    console.print(f"  Deployed ID: {deployed_id}")


if __name__ == "__main__":
    main()
