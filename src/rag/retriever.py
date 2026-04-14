"""Hybrid retrieval: BM25 (sparse) + dense retrieval + Reciprocal Rank Fusion.

Supports two dense backends:
  - ChromaDB (default): local vector DB with configurable embedding model
  - Gemini: Vertex AI Vector Search with Gemini embeddings

Query pipeline (per translation request):
  1. strip_python_comments(source)        → stripped_source
  2. [if A] grammar_mappings.query(stripped_source)   → go_pattern + description
  3. [if B] parallel_corpus.query(stripped_source)    → python_code + go_code pairs
  4. [if C] api_mappings.query(stripped_source)       → python_api + go_api + description
            collect go_api names → go_api_query
  5. [if C and D] go_docs.query(go_api_query)         → api + description + example
"""

from __future__ import annotations

import json
import re
from copy import deepcopy

from rank_bm25 import BM25Okapi

from src.config import (
    API_MAPPINGS_FILE,
    GO_DOCS_FILE,
    GRAMMAR_MAPPINGS_FILE,
    PARALLEL_CORPUS_FILE,
)
from src.rag.embeddings import get_embedding_function, load_rag_config
from src.rag.store import collection_name_with_dim


def _tokenize_code(text: str) -> list[str]:
    """Tokenize code for BM25: split on whitespace and common delimiters."""
    return [t for t in re.split(r"[\s.,()\[\]{};:=<>+\-*/\"']+", text.lower()) if t]


def strip_python_comments(source: str) -> str:
    """Remove # line comments and triple-quoted docstrings from Python source."""
    source = re.sub(r'""".*?"""', "", source, flags=re.DOTALL)
    source = re.sub(r"'''.*?'''", "", source, flags=re.DOTALL)
    source = re.sub(r"#[^\n]*", "", source)
    return source.strip()


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

        # Dense retrieval from ChromaDB — pre-compute embedding to avoid
        # ChromaDB's internal numpy conversion issues with custom EF
        if self._mode in ("hybrid", "dense"):
            query_embedding = self._ef([query])[0]
            dense_results = self._collection.query(
                query_embeddings=[query_embedding], n_results=min(fetch_k, count)
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
    "rag-pattern-only":     {"grammar": True,  "parallel_corpus": False, "api_mappings": False, "documentation": False},
    "rag-pattern-samples":  {"grammar": True,  "parallel_corpus": True,  "api_mappings": False, "documentation": False},
    "rag-pattern-api-docs": {"grammar": True,  "parallel_corpus": False, "api_mappings": True,  "documentation": True},
    "rag-full":             {"grammar": True,  "parallel_corpus": True,  "api_mappings": True,  "documentation": True},
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
    """Create a retriever for the given backend, with caching.

    For ChromaDB backend, the collection name is suffixed with the configured
    embedding dimensions (e.g. ``grammar_mappings_768``).
    """
    if backend != "gemini":
        collection_name = collection_name_with_dim(collection_name)
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
    name = "grammar_mappings" if backend == "gemini" else collection_name_with_dim("grammar_mappings")
    cache_key = (backend, name)
    if cache_key not in _retrievers:
        records = _load_jsonl(GRAMMAR_MAPPINGS_FILE)
        docs = []
        for i, r in enumerate(records):
            docs.append({
                "_id": f"grammar_{r['category']}_{i}",
                "python_pattern": r["python_pattern"],
                "category": r["category"],
                "go_pattern": r["go_pattern"],
                "description": r["description"],
            })
        _retrievers[cache_key] = _make_retriever(backend, "grammar_mappings", docs, "python_pattern", mode="dense")
    return _retrievers[cache_key]


def _get_parallel_corpus_retriever(backend: str = "chromadb") -> HybridRetriever | VertexAIRetriever:
    name = "parallel_corpus" if backend == "gemini" else collection_name_with_dim("parallel_corpus")
    cache_key = (backend, name)
    if cache_key not in _retrievers:
        records = _load_jsonl(PARALLEL_CORPUS_FILE)
        docs = []
        for i, r in enumerate(records):
            docs.append({
                "_id": f"parallel_{r['problem_id']}_{i}",
                "python_code": r["python_code"],
                "go_code": r["go_code"],
                "problem_description": r.get("problem_description", ""),
            })
        _retrievers[cache_key] = _make_retriever(backend, "parallel_corpus", docs, "python_code", mode="dense")
    return _retrievers[cache_key]


def _get_api_retriever(backend: str = "chromadb") -> HybridRetriever | VertexAIRetriever:
    name = "api_mappings" if backend == "gemini" else collection_name_with_dim("api_mappings")
    cache_key = (backend, name)
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
    name = "go_docs" if backend == "gemini" else collection_name_with_dim("go_docs")
    cache_key = (backend, name)
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

    __slots__ = ("api_mappings", "documentation", "grammar_mappings", "parallel_corpus", "context")

    def __init__(self) -> None:
        self.api_mappings: list[dict] = []
        self.documentation: list[dict] = []
        self.grammar_mappings: list[dict] = []
        self.parallel_corpus: list[dict] = []
        self.context: str = ""

    def to_artifact(
        self,
        *,
        embedding_backend: str,
        kb_toggles: dict[str, bool] | None = None,
        retrieval_config: dict | None = None,
    ) -> dict:
        """Return a JSON-serializable retrieval artifact for persistence."""
        return build_retrieval_artifact(
            self,
            embedding_backend=embedding_backend,
            kb_toggles=kb_toggles,
            retrieval_config=retrieval_config,
        )


def _normalize_kb_toggles(kb_toggles: dict[str, bool] | None) -> dict[str, bool]:
    normalized = dict(kb_toggles or {})
    if "code_snippets" in normalized and "parallel_corpus" not in normalized:
        normalized["parallel_corpus"] = normalized.pop("code_snippets")
    return normalized


def build_retrieval_artifact(
    rag_result: RAGResult,
    *,
    embedding_backend: str,
    kb_toggles: dict[str, bool] | None = None,
    retrieval_config: dict | None = None,
) -> dict:
    """Serialize a RAGResult into a stable artifact payload.

    This is intentionally independent from pipeline path handling so callers can
    persist the returned dict wherever they want.
    """
    config = dict(retrieval_config or _cfg()["retrieval"])
    items = {
        "grammar_mappings": deepcopy(rag_result.grammar_mappings),
        "parallel_corpus": deepcopy(rag_result.parallel_corpus),
        "api_mappings": deepcopy(rag_result.api_mappings),
        "documentation": deepcopy(rag_result.documentation),
    }
    return {
        "embedding_backend": embedding_backend,
        "kb_toggles": _normalize_kb_toggles(
            kb_toggles if kb_toggles is not None else _get_kb_toggles()
        ),
        "retrieval_config": {
            "parallel_corpus_k": config.get("parallel_corpus_k"),
            "api_mappings_k": config.get("api_mappings_k"),
            "go_docs_k": config.get("go_docs_k"),
            "rrf_k": config.get("rrf_k"),
        },
        "retrieval_counts": {
            key: len(value)
            for key, value in items.items()
        },
        "items": items,
        "rendered_context": rag_result.context,
    }


def build_translation_context(
    python_code: str,
    embedding_backend: str = "chromadb",
) -> RAGResult:
    """Query pipeline for Python-to-Go translation context.

    Steps:
      1. Strip comments from Python source
      2. [A] Query grammar_mappings with stripped source
      3. [B] Query parallel_corpus with stripped source
      4. [C] Query api_mappings with stripped source; collect Go API names
      5. [C+D] Query go_docs with Go API names from step 4

    Args:
        python_code: The source Python code to translate.
        embedding_backend: "chromadb" for ChromaDB or "gemini" for Vertex AI.

    Returns a RAGResult with raw retrieved items and the formatted context string.
    """
    cfg = _cfg()["retrieval"]
    kb = _get_kb_toggles()
    use_grammar  = kb.get("grammar", False)
    use_parallel = kb.get("parallel_corpus", False)
    use_api      = kb.get("api_mappings", False)
    use_docs     = kb.get("documentation", False)

    stripped = strip_python_comments(python_code)
    sections = []
    result = RAGResult()

    # Step A: grammar_mappings — tree-sitter detects which constructs are present,
    # then query once per detected construct for a precise per-category match.
    if use_grammar:
        from src.rag.grammar_extractor import extract_grammar_patterns

        grammar_ret = _get_grammar_retriever(embedding_backend)
        grammar_matches = []
        seen_categories: set[str] = set()
        for pat in extract_grammar_patterns(python_code):
            hits = grammar_ret.retrieve(pat["fragment"], n_results=1)
            for r in hits:
                if r["category"] not in seen_categories:
                    seen_categories.add(r["category"])
                    grammar_matches.append(r)
        result.grammar_mappings = grammar_matches
        if grammar_matches:
            examples = []
            for match in grammar_matches:
                examples.append(
                    f"### {match['category']}\n"
                    f"{match['description']}\n\n"
                    f"**Python:**\n```python\n{match['python_pattern']}\n```\n"
                    f"**Go:**\n```go\n{match['go_pattern']}\n```"
                )
            sections.append("## Go Grammar Implementation Patterns\n\n" + "\n\n".join(examples))

    # Step B: parallel_corpus
    if use_parallel:
        corpus_ret = _get_parallel_corpus_retriever(embedding_backend)
        corpus_matches = corpus_ret.retrieve(stripped, n_results=cfg["parallel_corpus_k"])
        result.parallel_corpus = corpus_matches
        if corpus_matches:
            pairs = []
            for m in corpus_matches:
                pairs.append(
                    f"**Python:**\n```python\n{m['python_code']}\n```\n"
                    f"**Go:**\n```go\n{m['go_code']}\n```"
                )
            sections.append("## Python-Go Code Pairs\n\n" + "\n\n".join(pairs))

    # Step C: api_mappings — tree-sitter extracts API calls and imports,
    # then query with those specific names for a precise match.
    mappings: list[dict] = []
    if use_api:
        from src.rag.api_extractor import extract_api_info

        api_info = extract_api_info(python_code)
        api_query = api_info["query_apis"]
        if api_info["query_imports"]:
            api_query += " " + api_info["query_imports"]
        if api_query.strip():
            api_ret = _get_api_retriever(embedding_backend)
            mappings = api_ret.retrieve(api_query, n_results=cfg["api_mappings_k"])
        result.api_mappings = mappings
        if mappings:
            lines = [
                f"- `{m['python_api']}` -> `{m['go_api']}`: {m['description']}"
                for m in mappings
            ]
            sections.append("## Relevant API Mappings\n" + "\n".join(lines))

    # Step D: go_docs — bridge query uses Go API names from step C
    if use_docs and mappings:
        go_api_query = " ".join(m["go_api"] for m in mappings)
        godoc_ret = _get_godoc_retriever(embedding_backend)
        docs = godoc_ret.retrieve(go_api_query, n_results=cfg["go_docs_k"])
        result.documentation = docs
        if docs:
            doc_lines = []
            for d in docs:
                line = f"- **{d['api']}**: {d['description']}"
                if d.get("example"):
                    line += f"\n  ```go\n  {d['example']}\n  ```"
                doc_lines.append(line)
            sections.append("## Go Documentation & Patterns\n" + "\n".join(doc_lines))

    result.context = "\n\n".join(sections) if sections else ""
    return result
