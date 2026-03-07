"""Hybrid retrieval: BM25 (sparse) + ChromaDB dense + Reciprocal Rank Fusion.

Pipeline (matching thesis design):
  1. Input: Python code
  2. Tool 1: Extract Python APIs using tree-sitter
  3. Tool 2: Query api_mappings for Go equivalents
  4. Tool 3: Query go_docs for usage of those Go APIs + error patterns
  5. Output: Combined context for the LLM prompt
"""

from __future__ import annotations

import json
import re

from rank_bm25 import BM25Okapi

from src.config import (
    API_MAPPINGS_FILE,
    GO_DOCS_FILE,
    PARALLEL_CORPUS_FILE,
)
from src.rag.embeddings import get_embedding_function, load_rag_config
from src.rag.store import get_chroma_client, get_or_create_collection


def _tokenize_code(text: str) -> list[str]:
    """Tokenize code for BM25: split on whitespace and common delimiters."""
    return [t for t in re.split(r"[\s.,()\[\]{};:=<>+\-*/\"']+", text.lower()) if t]


def _load_jsonl(path) -> list[dict]:
    records = []
    with open(path, encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if line:
                records.append(json.loads(line))
    return records


def _reciprocal_rank_fusion(
    ranked_lists: list[list[str]], k: int = 60
) -> list[str]:
    """Merge multiple ranked ID lists using RRF."""
    scores: dict[str, float] = {}
    for ranked in ranked_lists:
        for rank, doc_id in enumerate(ranked):
            scores[doc_id] = scores.get(doc_id, 0.0) + 1.0 / (k + rank + 1)
    return sorted(scores, key=scores.get, reverse=True)


class HybridRetriever:
    """Combines ChromaDB dense retrieval with BM25 sparse retrieval."""

    def __init__(self, collection_name: str, documents: list[dict], text_key: str, rrf_k: int = 60):
        self._ef = get_embedding_function()
        client = get_chroma_client()
        self._collection = get_or_create_collection(client, collection_name, self._ef)
        self._documents = documents
        self._text_key = text_key
        self._rrf_k = rrf_k

        # Build ID -> document index
        self._id_to_doc: dict[str, dict] = {}
        corpus_texts = []
        for doc in documents:
            doc_id = doc["_id"]
            self._id_to_doc[doc_id] = doc
            corpus_texts.append(doc[text_key])

        # Build BM25 index
        tokenized = [_tokenize_code(t) for t in corpus_texts]
        self._bm25 = BM25Okapi(tokenized)
        self._doc_ids = [doc["_id"] for doc in documents]

    def retrieve(self, query: str, n_results: int = 5) -> list[dict]:
        fetch_k = n_results * 3
        count = self._collection.count()
        if count == 0:
            return []

        # Dense retrieval from ChromaDB
        dense_results = self._collection.query(
            query_texts=[query], n_results=min(fetch_k, count)
        )
        dense_ids = dense_results["ids"][0] if dense_results["ids"] else []

        # BM25 retrieval
        tokenized_query = _tokenize_code(query)
        bm25_scores = self._bm25.get_scores(tokenized_query)
        top_indices = sorted(
            range(len(bm25_scores)), key=lambda i: bm25_scores[i], reverse=True
        )[:fetch_k]
        bm25_ids = [self._doc_ids[i] for i in top_indices]

        # Reciprocal Rank Fusion
        merged_ids = _reciprocal_rank_fusion([dense_ids, bm25_ids], k=self._rrf_k)

        results = []
        for doc_id in merged_ids[:n_results]:
            if doc_id in self._id_to_doc:
                results.append(self._id_to_doc[doc_id])
        return results


# ---------------------------------------------------------------------------
# Singleton retrievers (lazy-initialized)
# ---------------------------------------------------------------------------
_corpus_retriever: HybridRetriever | None = None
_api_retriever: HybridRetriever | None = None
_godoc_retriever: HybridRetriever | None = None
_rag_cfg: dict | None = None


def _cfg() -> dict:
    global _rag_cfg
    if _rag_cfg is None:
        _rag_cfg = load_rag_config()
    return _rag_cfg


def _get_corpus_retriever() -> HybridRetriever:
    global _corpus_retriever
    if _corpus_retriever is None:
        records = _load_jsonl(PARALLEL_CORPUS_FILE)
        docs = []
        for r in records:
            docs.append({
                "_id": f"corpus_{r['problem_id']}",
                "python_code": r["python_code"],
                "go_code": r["go_code"],
                "problem_description": r.get("problem_description", ""),
            })
        _corpus_retriever = HybridRetriever(
            "parallel_corpus", docs, "python_code",
            rrf_k=_cfg()["retrieval"]["rrf_k"],
        )
    return _corpus_retriever


def _get_api_retriever() -> HybridRetriever:
    global _api_retriever
    if _api_retriever is None:
        records = _load_jsonl(API_MAPPINGS_FILE)
        docs = []
        for i, r in enumerate(records):
            text = f"{r['category']}: {r['python_api']} -> {r['go_api']}. {r['description']}"
            docs.append({
                "_id": f"api_{r['category']}_{i}",
                "text": text,
                "category": r["category"],
                "python_api": r["python_api"],
                "go_api": r["go_api"],
                "description": r["description"],
            })
        _api_retriever = HybridRetriever(
            "api_mappings", docs, "text",
            rrf_k=_cfg()["retrieval"]["rrf_k"],
        )
    return _api_retriever


def _get_godoc_retriever() -> HybridRetriever:
    global _godoc_retriever
    if _godoc_retriever is None:
        records = _load_jsonl(GO_DOCS_FILE)
        docs = []
        for i, r in enumerate(records):
            text = f"{r['package']}: {r['api']}. {r['description']} Example: {r.get('example', '')}"
            docs.append({
                "_id": f"godoc_{r['package']}_{i}",
                "text": text,
                "package": r["package"],
                "api": r["api"],
                "description": r["description"],
                "example": r.get("example", ""),
            })
        _godoc_retriever = HybridRetriever(
            "go_docs", docs, "text",
            rrf_k=_cfg()["retrieval"]["rrf_k"],
        )
    return _godoc_retriever


def build_translation_context(python_code: str) -> str:
    """Structured RAG pipeline for Python-to-Go translation.

    Pipeline:
      1. Extract Python APIs using tree-sitter
      2. Query api_mappings using extracted API names (precise)
      3. Query go_docs using matched Go APIs + error patterns if needed
      4. Query parallel_corpus for similar full-code examples
      5. Format all context for the LLM prompt
    """
    from src.rag.api_extractor import extract_api_info

    cfg = _cfg()["retrieval"]
    sections = []

    # Step 1: Extract APIs from Python code
    api_info = extract_api_info(python_code)

    # Step 2: Query api_mappings using extracted API names (not raw code)
    api_ret = _get_api_retriever()
    api_query = api_info["query_apis"]
    if api_info["query_imports"]:
        api_query += " " + api_info["query_imports"]
    mappings = api_ret.retrieve(api_query, n_results=cfg["api_mappings_k"]) if api_query.strip() else []

    if mappings:
        mapping_lines = [
            f"- `{m['python_api']}` -> `{m['go_api']}`: {m['description']}"
            for m in mappings
        ]
        sections.append("## Relevant API Mappings\n" + "\n".join(mapping_lines))

    # Step 3: Query go_docs using Go API names from mappings
    godoc_ret = _get_godoc_retriever()
    go_api_query = " ".join(m["go_api"] for m in mappings) if mappings else ""

    # If error handling detected, also search for error patterns
    if api_info["has_error_handling"]:
        go_api_query += " error handling pattern try except"

    if go_api_query.strip():
        docs = godoc_ret.retrieve(go_api_query, n_results=cfg["go_docs_k"])
        if docs:
            doc_lines = []
            for d in docs:
                line = f"- **{d['api']}**: {d['description']}"
                if d.get("example"):
                    line += f"\n  ```go\n  {d['example']}\n  ```"
                doc_lines.append(line)
            sections.append("## Go Documentation & Patterns\n" + "\n".join(doc_lines))

    # Step 4: Query parallel corpus for similar full-code examples
    corpus_ret = _get_corpus_retriever()
    similar = corpus_ret.retrieve(python_code, n_results=cfg["parallel_corpus_k"])
    if similar:
        examples = []
        for i, s in enumerate(similar, 1):
            examples.append(
                f"### Example {i}\n"
                f"**Python:**\n```python\n{s['python_code']}\n```\n"
                f"**Go:**\n```go\n{s['go_code']}\n```"
            )
        sections.append("## Reference Translation Examples\n" + "\n\n".join(examples))

    if not sections:
        return ""
    return "\n\n".join(sections)
