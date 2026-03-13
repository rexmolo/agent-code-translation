"""Hybrid retrieval: BM25 (sparse) + dense retrieval + Reciprocal Rank Fusion.

Supports two dense backends:
  - ChromaDB (default): local vector DB with configurable embedding model
  - Gemini: Vertex AI Vector Search with Gemini embeddings

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
    GRAMMAR_MAPPINGS_FILE,
)
from src.rag.embeddings import get_embedding_function, load_rag_config


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

    def __init__(self, collection_name: str, documents: list[dict], text_key: str, rrf_k: int = 60, mode: str = "hybrid"):
        from src.rag.store import get_chroma_client, get_or_create_collection

        self._ef = get_embedding_function()
        client = get_chroma_client()
        self._collection = get_or_create_collection(client, collection_name, self._ef)
        self._documents = documents
        self._text_key = text_key
        self._rrf_k = rrf_k
        self._mode = mode

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

        dense_ids = []
        bm25_ids = []

        # Dense retrieval from ChromaDB
        if self._mode in ("hybrid", "dense"):
            dense_results = self._collection.query(
                query_texts=[query], n_results=min(fetch_k, count)
            )
            dense_ids = dense_results["ids"][0] if dense_results["ids"] else []

            if self._mode == "dense":
                return [self._id_to_doc[doc_id] for doc_id in dense_ids[:n_results] if doc_id in self._id_to_doc]

        # BM25 retrieval
        if self._mode in ("hybrid", "sparse"):
            tokenized_query = _tokenize_code(query)
            bm25_scores = self._bm25.get_scores(tokenized_query)
            top_indices = sorted(
                range(len(bm25_scores)), key=lambda i: bm25_scores[i], reverse=True
            )[:fetch_k]
            bm25_ids = [self._doc_ids[i] for i in top_indices]

            if self._mode == "sparse":
                return [self._id_to_doc[doc_id] for doc_id in bm25_ids[:n_results] if doc_id in self._id_to_doc]

        # Reciprocal Rank Fusion
        merged_ids = _reciprocal_rank_fusion([dense_ids, bm25_ids], k=self._rrf_k)

        results = []
        for doc_id in merged_ids[:n_results]:
            if doc_id in self._id_to_doc:
                results.append(self._id_to_doc[doc_id])
        return results


class VertexAIRetriever:
    """Combines Vertex AI Vector Search dense retrieval with BM25 sparse retrieval."""

    def __init__(self, collection_name: str, documents: list[dict], text_key: str, rrf_k: int = 60, mode: str = "hybrid"):
        from src.rag.embeddings import GeminiEmbeddingFunction
        from src.rag.vertex_store import (
            ensure_deployed,
            get_or_create_endpoint,
            get_or_create_index,
        )

        cfg = load_rag_config()
        model = cfg["embedding"]["gemini"]["model"]
        self._ef = GeminiEmbeddingFunction(model_name=model)
        self._collection_name = collection_name
        self._rrf_k = rrf_k
        self._mode = mode

        # Vertex AI resources (lazy, cached at module level)
        self._endpoint = _get_vertex_endpoint()
        self._deployed_index_id = _get_vertex_deployed_id()

        # Build ID -> document index
        self._id_to_doc: dict[str, dict] = {}
        corpus_texts = []
        for doc in documents:
            doc_id = doc["_id"]
            self._id_to_doc[doc_id] = doc
            corpus_texts.append(doc[text_key])

        # Build BM25 index (same as HybridRetriever)
        tokenized = [_tokenize_code(t) for t in corpus_texts]
        self._bm25 = BM25Okapi(tokenized)
        self._doc_ids = [doc["_id"] for doc in documents]

    def retrieve(self, query: str, n_results: int = 5) -> list[dict]:
        from src.rag.vertex_store import query_neighbors

        fetch_k = n_results * 3
        dense_ids = []
        bm25_ids = []

        # Dense retrieval from Vertex AI
        if self._mode in ("hybrid", "dense"):
            query_embedding = self._ef.embed_query(query)
            dense_ids = query_neighbors(
                self._endpoint,
                self._deployed_index_id,
                query_embedding,
                n_results=fetch_k,
                collection_name=self._collection_name,
            )

            if self._mode == "dense":
                return [self._id_to_doc[doc_id] for doc_id in dense_ids[:n_results] if doc_id in self._id_to_doc]

        # BM25 retrieval
        if self._mode in ("hybrid", "sparse"):
            tokenized_query = _tokenize_code(query)
            bm25_scores = self._bm25.get_scores(tokenized_query)
            top_indices = sorted(
                range(len(bm25_scores)), key=lambda i: bm25_scores[i], reverse=True
            )[:fetch_k]
            bm25_ids = [self._doc_ids[i] for i in top_indices]

            if self._mode == "sparse":
                return [self._id_to_doc[doc_id] for doc_id in bm25_ids[:n_results] if doc_id in self._id_to_doc]

        # Reciprocal Rank Fusion
        merged_ids = _reciprocal_rank_fusion([dense_ids, bm25_ids], k=self._rrf_k)

        results = []
        for doc_id in merged_ids[:n_results]:
            if doc_id in self._id_to_doc:
                results.append(self._id_to_doc[doc_id])
        return results


# ---------------------------------------------------------------------------
# Vertex AI resource singletons (lazy-initialized)
# ---------------------------------------------------------------------------
_vertex_endpoint = None
_vertex_deployed_id: str | None = None


def _get_vertex_endpoint():
    global _vertex_endpoint, _vertex_deployed_id
    if _vertex_endpoint is None:
        from src.rag.vertex_store import (
            ensure_deployed,
            get_or_create_endpoint,
            get_or_create_index,
        )
        index = get_or_create_index()
        _vertex_endpoint = get_or_create_endpoint()
        _vertex_deployed_id = ensure_deployed(index, _vertex_endpoint)
    return _vertex_endpoint


def _get_vertex_deployed_id() -> str:
    if _vertex_deployed_id is None:
        _get_vertex_endpoint()  # triggers initialization
    return _vertex_deployed_id


# ---------------------------------------------------------------------------
# Singleton retrievers (lazy-initialized), keyed by backend
# ---------------------------------------------------------------------------
_retrievers: dict[tuple[str, str], HybridRetriever | VertexAIRetriever] = {}
_rag_cfg: dict | None = None

# Runtime overrides for knowledge base toggles (set by configure_kb_for_experiment)
_kb_overrides: dict[str, bool] | None = None

# Experiment name -> knowledge base toggles
_EXPERIMENT_KB_PRESETS: dict[str, dict[str, bool]] = {
    "rag":             {"code_snippets": True,  "api_mappings": True,  "documentation": True},
    "rag-no-snippets": {"code_snippets": False, "api_mappings": True,  "documentation": True},
    "rag-no-mappings": {"code_snippets": True,  "api_mappings": False, "documentation": True},
    "rag-no-docs":     {"code_snippets": True,  "api_mappings": True,  "documentation": False},
}


def configure_kb_for_experiment(experiment: str) -> dict[str, bool] | None:
    """Apply knowledge base toggles for a given experiment name.

    Returns the active KB toggles, or None if the experiment doesn't use RAG
    (e.g. baseline).
    """
    global _kb_overrides

    if experiment == "baseline":
        _kb_overrides = None
        return None

    if experiment in _EXPERIMENT_KB_PRESETS:
        _kb_overrides = _EXPERIMENT_KB_PRESETS[experiment]
        return dict(_kb_overrides)

    # Unknown experiment name: fall back to YAML knowledge_bases section
    _kb_overrides = None
    return _cfg().get("knowledge_bases")


def get_active_kb_toggles(experiment: str) -> dict[str, bool] | None:
    """Return the effective KB toggles for display purposes.

    Returns None for experiments that don't use RAG (baseline).
    """
    if experiment == "baseline":
        return None
    if experiment in _EXPERIMENT_KB_PRESETS:
        return _EXPERIMENT_KB_PRESETS[experiment]
    return _cfg().get("knowledge_bases", {
        "code_snippets": True, "api_mappings": True, "documentation": True,
    })


def _cfg() -> dict:
    global _rag_cfg
    if _rag_cfg is None:
        _rag_cfg = load_rag_config()
    return _rag_cfg


def _get_kb_toggles() -> dict[str, bool]:
    """Return effective KB toggles (runtime overrides take precedence over YAML)."""
    if _kb_overrides is not None:
        return _kb_overrides
    return _cfg().get("knowledge_bases", {
        "code_snippets": True, "api_mappings": True, "documentation": True,
    })


def _make_retriever(
    backend: str,
    collection_name: str,
    documents: list[dict],
    text_key: str,
    mode: str = "hybrid",
) -> HybridRetriever | VertexAIRetriever:
    """Create a retriever for the given backend, with caching."""
    cache_key = (backend, collection_name)
    if cache_key not in _retrievers:
        rrf_k = _cfg()["retrieval"]["rrf_k"]
        if backend == "gemini":
            _retrievers[cache_key] = VertexAIRetriever(
                collection_name, documents, text_key, rrf_k=rrf_k, mode=mode,
            )
        else:
            _retrievers[cache_key] = HybridRetriever(
                collection_name, documents, text_key, rrf_k=rrf_k, mode=mode,
            )
    return _retrievers[cache_key]


def _get_grammar_retriever(backend: str = "chromadb") -> HybridRetriever | VertexAIRetriever:
    cache_key = (backend, "grammar_mappings")
    if cache_key not in _retrievers:
        records = _load_jsonl(GRAMMAR_MAPPINGS_FILE)
        docs = []
        for i, r in enumerate(records):
            docs.append({
                "_id": f"grammar_{r['category']}_{i}",
                "text": f"{r['category']}: {r['python_pattern']} -> {r['go_pattern']}. {r['description']}",
                "category": r["category"],
                "python_pattern": r["python_pattern"],
                "go_pattern": r["go_pattern"],
                "description": r["description"],
            })
        _retrievers[cache_key] = _make_retriever(backend, "grammar_mappings", docs, "text", mode="hybrid")
    return _retrievers[cache_key]


def _get_api_retriever(backend: str = "chromadb") -> HybridRetriever | VertexAIRetriever:
    cache_key = (backend, "api_mappings")
    if cache_key not in _retrievers:
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
        _retrievers[cache_key] = _make_retriever(backend, "api_mappings", docs, "text", mode="hybrid")
    return _retrievers[cache_key]


def _get_godoc_retriever(backend: str = "chromadb") -> HybridRetriever | VertexAIRetriever:
    cache_key = (backend, "go_docs")
    if cache_key not in _retrievers:
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
        _retrievers[cache_key] = _make_retriever(backend, "go_docs", docs, "text", mode="hybrid")
    return _retrievers[cache_key]


class RAGResult:
    """Container for RAG retrieval results and the formatted context."""

    __slots__ = ("api_mappings", "documentation", "grammar_mappings", "context")

    def __init__(self) -> None:
        self.api_mappings: list[dict] = []
        self.documentation: list[dict] = []
        self.grammar_mappings: list[dict] = []
        self.context: str = ""


def build_translation_context(
    python_code: str,
    embedding_backend: str = "chromadb",
) -> RAGResult:
    """Structured RAG pipeline for Python-to-Go translation.

    Pipeline:
      1. Extract Python APIs using tree-sitter
      2. Query api_mappings using extracted API names (precise)
      3. Query go_docs using matched Go APIs + error patterns if needed
      4. Query grammar_mappings for abstract Python -> Go structural patterns
      5. Format all context for the LLM prompt

    Args:
        python_code: The source Python code to translate.
        embedding_backend: "chromadb" for ChromaDB or "gemini" for Vertex AI.

    Returns a RAGResult with raw retrieved items and the formatted context string.
    """
    from src.rag.api_extractor import extract_api_info

    cfg = _cfg()["retrieval"]
    kb_toggles = _get_kb_toggles()
    use_grammar_mappings = kb_toggles.get("code_snippets", True) # Keep using the existing config toggle
    use_api_mappings = kb_toggles.get("api_mappings", True)
    use_documentation = kb_toggles.get("documentation", True)
    sections = []
    result = RAGResult()

    # Step 1: Extract APIs from Python code
    api_info = extract_api_info(python_code)

    # Step 2: Query api_mappings using extracted API names (not raw code)
    mappings = []
    if use_api_mappings:
        api_ret = _get_api_retriever(embedding_backend)
        api_query = api_info["query_apis"]
        if api_info["query_imports"]:
            api_query += " " + api_info["query_imports"]
        mappings = api_ret.retrieve(api_query, n_results=cfg["api_mappings_k"]) if api_query.strip() else []
        result.api_mappings = mappings

        if mappings:
            mapping_lines = [
                f"- `{m['python_api']}` -> `{m['go_api']}`: {m['description']}"
                for m in mappings
            ]
            sections.append("## Relevant API Mappings\n" + "\n".join(mapping_lines))

    # Step 3: Query go_docs using Go API names from mappings
    docs = []
    if use_documentation:
        godoc_ret = _get_godoc_retriever(embedding_backend)
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
    result.documentation = docs

    # Step 4: Query grammar mappings for abstract syntactic rules
    grammar_matches = []
    if use_grammar_mappings:
        grammar_ret = _get_grammar_retriever(embedding_backend)
        grammar_matches = grammar_ret.retrieve(python_code, n_results=cfg.get("grammar_k", 3))
        if grammar_matches:
            examples = []
            for i, match in enumerate(grammar_matches, 1):
                examples.append(
                    f"### Grammar Rule: {match['category']}\n"
                    f"{match['description']}\n\n"
                    f"**Python:**\n```python\n{match['python_pattern']}\n```\n"
                    f"**Go:**\n```go\n{match['go_pattern']}\n```"
                )
            sections.append("## Go Grammar Implementation Patterns\n" + "\n\n".join(examples))
    result.grammar_mappings = grammar_matches

    result.context = "\n\n".join(sections) if sections else ""
    return result
