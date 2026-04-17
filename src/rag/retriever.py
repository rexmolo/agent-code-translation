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
    GO_API_SEQUENCES_FILE,
    GRAMMAR_MAPPINGS_FILE,
    PARALLEL_CORPUS_FILE,
)
from src.rag.embeddings import get_embedding_function, load_rag_config
from src.rag.rendering import sanitize_parallel_go_reference
from src.rag.schema import RAG_SOURCES, rag_result_has_usable_items
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


def _rrf_details(ranked_lists: list[list[str]], k: int = 60) -> list[dict]:
    scores: dict[str, float] = {}
    ranks: dict[str, list[int | None]] = {}
    for list_index, ranked in enumerate(ranked_lists):
        for rank, doc_id in enumerate(ranked, start=1):
            scores[doc_id] = scores.get(doc_id, 0.0) + 1.0 / (k + rank)
            if doc_id not in ranks:
                ranks[doc_id] = [None] * len(ranked_lists)
            ranks[doc_id][list_index] = rank

    merged_ids = sorted(scores, key=scores.get, reverse=True)
    details = []
    for merged_rank, doc_id in enumerate(merged_ids, start=1):
        detail = {
            "doc_id": doc_id,
            "merged_rank": merged_rank,
            "rrf_score": scores[doc_id],
        }
        if len(ranked_lists) > 0:
            detail["dense_rank"] = ranks[doc_id][0]
        if len(ranked_lists) > 1:
            detail["sparse_rank"] = ranks[doc_id][1]
        details.append(detail)
    return details


def _clean_number(value):
    if isinstance(value, bool) or value is None:
        return value
    if isinstance(value, int):
        return value
    if isinstance(value, float):
        return round(value, 8)
    return value


def _copy_source_rules(config: dict, source_name: str) -> dict:
    source_rules = config.get("source_rules", {})
    rules = source_rules.get(source_name, {})
    return deepcopy(rules) if isinstance(rules, dict) else {}


def _normalize_source_trace(source_name: str, trace: dict | None) -> dict:
    trace = deepcopy(trace or {})
    trace.setdefault("source", source_name)
    trace.setdefault("queried", False)
    trace.setdefault("accepted", False)
    trace.setdefault("query_text", "")
    trace.setdefault("returned_count", 0)
    trace.setdefault("accepted_count", 0)
    trace.setdefault("items", [])
    trace.setdefault("acceptance", {})
    trace["acceptance"].setdefault("enabled", False)
    trace["acceptance"].setdefault("rules", {})
    trace["acceptance"].setdefault("fallback_to_no_retrieval", False)
    return trace


def _empty_source_traces(config: dict | None = None) -> dict[str, dict]:
    retrieval_cfg = config or _cfg()["retrieval"]
    return {
        source: _normalize_source_trace(
            source,
            {
                "acceptance": {
                    "enabled": bool(retrieval_cfg.get("enable_confidence_gate", False)),
                    "rules": _copy_source_rules(retrieval_cfg, source),
                    "fallback_to_no_retrieval": bool(
                        retrieval_cfg.get("hard_fallback_to_no_retrieval", False)
                    ),
                },
            },
        )
        for source in RAG_SOURCES
    }


def _route_kb_toggles(
    base_kb_toggles: dict[str, bool],
    *,
    api_info: dict,
    grammar_patterns: list[dict],
    router_config: dict,
) -> tuple[dict[str, bool], dict]:
    normalized = _normalize_kb_toggles(base_kb_toggles)
    api_call_count = len(api_info.get("calls", []))
    import_count = len(api_info.get("imports", []))
    grammar_pattern_count = len(grammar_patterns)
    has_error_handling = bool(api_info.get("has_error_handling", False))

    routed = dict(normalized)
    reasons: dict[str, list[str]] = {source: [] for source in RAG_SOURCES}

    for source_name, base_enabled in (
        ("grammar_mappings", normalized.get("grammar", False)),
        ("parallel_corpus", normalized.get("parallel_corpus", False)),
        ("api_mappings", normalized.get("api_mappings", False)),
        ("documentation", normalized.get("documentation", False)),
        ("api_sequences", normalized.get("api_sequences", False)),
    ):
        if not base_enabled:
            reasons[source_name].append("disabled_by_base_preset")

    routed["grammar"] = normalized.get("grammar", False) and (
        grammar_pattern_count >= router_config.get("min_grammar_patterns_for_grammar", 1)
    )
    reasons["grammar_mappings"].append(
        f"grammar_pattern_count={grammar_pattern_count}"
    )

    routed["api_mappings"] = normalized.get("api_mappings", False) and (
        api_call_count >= router_config.get("min_api_calls_for_api_mappings", 1)
        or import_count >= router_config.get("min_imports_for_api_mappings", 1)
    )
    reasons["api_mappings"].append(
        f"api_call_count={api_call_count}, import_count={import_count}"
    )

    routed["api_sequences"] = normalized.get("api_sequences", False) and (
        api_call_count >= router_config.get("min_api_calls_for_api_sequences", 2)
    )
    reasons["api_sequences"].append(
        f"api_call_count={api_call_count}"
    )

    parallel_allowed = (
        api_call_count <= router_config.get("max_api_calls_for_parallel", 1)
        and import_count <= router_config.get("max_imports_for_parallel", 1)
        and grammar_pattern_count <= router_config.get("max_grammar_patterns_for_parallel", 1)
    )
    if router_config.get("disable_parallel_on_error_handling", True) and has_error_handling:
        parallel_allowed = False
        reasons["parallel_corpus"].append("disabled_by_error_handling")
    routed["parallel_corpus"] = normalized.get("parallel_corpus", False) and parallel_allowed
    reasons["parallel_corpus"].append(
        f"api_call_count={api_call_count}, import_count={import_count}, grammar_pattern_count={grammar_pattern_count}"
    )

    routed["documentation"] = normalized.get("documentation", False) and routed["api_mappings"]
    reasons["documentation"].append(
        f"documentation_follows_api_mappings={routed['api_mappings']}"
    )

    metadata = {
        "mode": "rule_based",
        "decision": "extractor_signals",
        "signals": {
            "api_call_count": api_call_count,
            "import_count": import_count,
            "grammar_pattern_count": grammar_pattern_count,
            "has_error_handling": has_error_handling,
        },
        "base_kb_toggles": normalized,
        "selected_kb_toggles": routed,
        "notes": [],
        "reasons": reasons,
    }
    return routed, metadata
def _accept_retrieved_items(
    source_name: str,
    items: list[dict],
    *,
    config: dict,
) -> tuple[list[dict], dict]:
    rules = _copy_source_rules(config, source_name)
    gate_enabled = bool(config.get("enable_confidence_gate", False))
    hard_fallback = bool(config.get("hard_fallback_to_no_retrieval", False))
    max_items = rules.get("max_accepted_items")
    max_merged_rank = rules.get("max_merged_rank")
    max_dense_distance = rules.get("max_dense_distance")
    min_sparse_score = rules.get("min_sparse_score")

    accepted_items: list[dict] = []
    trace_items: list[dict] = []

    for item in items:
        candidate = deepcopy(item)
        retrieval = deepcopy(candidate.get("retrieval", {}))
        reasons: list[str] = []
        accepted = True

        if gate_enabled:
            if isinstance(max_merged_rank, int):
                merged_rank = retrieval.get("merged_rank")
                if not isinstance(merged_rank, int) or merged_rank > max_merged_rank:
                    accepted = False
                    reasons.append("merged_rank_out_of_range")

            if isinstance(max_dense_distance, (int, float)):
                dense_distance = retrieval.get("dense_distance")
                if dense_distance is None or dense_distance > max_dense_distance:
                    accepted = False
                    reasons.append("dense_distance_too_high")

            if isinstance(min_sparse_score, (int, float)):
                sparse_score = retrieval.get("sparse_score")
                if sparse_score is None or sparse_score < min_sparse_score:
                    accepted = False
                    reasons.append("sparse_score_too_low")

        if accepted and isinstance(max_items, int) and len(accepted_items) >= max_items:
            accepted = False
            reasons.append("over_source_item_limit")

        retrieval["accepted"] = accepted
        retrieval["rejection_reasons"] = reasons
        candidate["retrieval"] = {key: _clean_number(value) for key, value in retrieval.items()}
        trace_items.append(candidate)
        if accepted:
            accepted_items.append(deepcopy(candidate))

    trace = _normalize_source_trace(
        source_name,
        {
            "queried": True,
            "accepted": bool(accepted_items),
            "returned_count": len(items),
            "accepted_count": len(accepted_items),
            "items": trace_items,
            "acceptance": {
                "enabled": gate_enabled,
                "rules": rules,
                "fallback_to_no_retrieval": hard_fallback,
            },
        },
    )
    return accepted_items, trace


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

    def retrieve(self, query: str, n_results: int = 5, include_metadata: bool = False) -> list[dict]:
        fetch_k = n_results * 3
        count = self._collection.count()
        if count == 0:
            return []

        dense_ids = []
        bm25_ids = []
        dense_distance_by_id: dict[str, float] = {}

        # Dense retrieval from ChromaDB — pre-compute embedding to avoid
        # ChromaDB's internal numpy conversion issues with custom EF
        if self._mode in ("hybrid", "dense"):
            query_embedding = self._ef([query])[0]
            dense_results = self._collection.query(
                query_embeddings=[query_embedding],
                n_results=min(fetch_k, count),
                include=["distances"],
            )
            dense_ids = dense_results["ids"][0] if dense_results["ids"] else []
            dense_distances = dense_results.get("distances", [[]])
            dense_distances = dense_distances[0] if dense_distances else []
            dense_distance_by_id = {
                doc_id: dense_distances[index]
                for index, doc_id in enumerate(dense_ids)
                if index < len(dense_distances)
            }

            if self._mode == "dense":
                results = []
                for rank, doc_id in enumerate(dense_ids[:n_results], start=1):
                    if doc_id not in self._id_to_doc:
                        continue
                    doc = deepcopy(self._id_to_doc[doc_id])
                    if include_metadata:
                        doc["retrieval"] = {
                            "backend": "chromadb",
                            "mode": self._mode,
                            "doc_id": doc_id,
                            "merged_rank": rank,
                            "rrf_score": None,
                            "dense_rank": rank,
                            "dense_distance": _clean_number(dense_distance_by_id.get(doc_id)),
                            "sparse_rank": None,
                            "sparse_score": None,
                        }
                    results.append(doc)
                return results

        # BM25 retrieval
        sparse_score_by_id: dict[str, float] = {}
        if self._mode in ("hybrid", "sparse"):
            tokenized_query = _tokenize_code(query)
            bm25_scores = self._bm25.get_scores(tokenized_query)
            top_indices = sorted(
                range(len(bm25_scores)), key=lambda i: bm25_scores[i], reverse=True
            )[:fetch_k]
            bm25_ids = [self._doc_ids[i] for i in top_indices]
            sparse_score_by_id = {
                self._doc_ids[index]: bm25_scores[index]
                for index in top_indices
            }

            if self._mode == "sparse":
                results = []
                for rank, doc_id in enumerate(bm25_ids[:n_results], start=1):
                    if doc_id not in self._id_to_doc:
                        continue
                    doc = deepcopy(self._id_to_doc[doc_id])
                    if include_metadata:
                        doc["retrieval"] = {
                            "backend": "chromadb",
                            "mode": self._mode,
                            "doc_id": doc_id,
                            "merged_rank": rank,
                            "rrf_score": None,
                            "dense_rank": None,
                            "dense_distance": None,
                            "sparse_rank": rank,
                            "sparse_score": _clean_number(sparse_score_by_id.get(doc_id)),
                        }
                    results.append(doc)
                return results

        # Reciprocal Rank Fusion
        merged = _rrf_details([dense_ids, bm25_ids], k=self._rrf_k)

        results = []
        for detail in merged[:n_results]:
            doc_id = detail["doc_id"]
            if doc_id in self._id_to_doc:
                doc = deepcopy(self._id_to_doc[doc_id])
                if include_metadata:
                    doc["retrieval"] = {
                        "backend": "chromadb",
                        "mode": self._mode,
                        "doc_id": doc_id,
                        "merged_rank": detail.get("merged_rank"),
                        "rrf_score": _clean_number(detail.get("rrf_score")),
                        "dense_rank": detail.get("dense_rank"),
                        "dense_distance": _clean_number(dense_distance_by_id.get(doc_id)),
                        "sparse_rank": detail.get("sparse_rank"),
                        "sparse_score": _clean_number(sparse_score_by_id.get(doc_id)),
                    }
                results.append(doc)
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

    def retrieve(self, query: str, n_results: int = 5, include_metadata: bool = False) -> list[dict]:
        from src.rag.vertex_store import query_neighbors

        fetch_k = n_results * 3
        dense_ids = []
        bm25_ids = []
        dense_distance_by_id: dict[str, float] = {}

        # Dense retrieval from Vertex AI
        if self._mode in ("hybrid", "dense"):
            query_embedding = self._ef.embed_query(query)
            dense_neighbors = query_neighbors(
                self._endpoint,
                self._deployed_index_id,
                query_embedding,
                n_results=fetch_k,
                collection_name=self._collection_name,
            )
            dense_ids = [neighbor["id"] for neighbor in dense_neighbors]
            dense_distance_by_id = {
                neighbor["id"]: neighbor.get("distance")
                for neighbor in dense_neighbors
            }

            if self._mode == "dense":
                results = []
                for rank, doc_id in enumerate(dense_ids[:n_results], start=1):
                    if doc_id not in self._id_to_doc:
                        continue
                    doc = deepcopy(self._id_to_doc[doc_id])
                    if include_metadata:
                        doc["retrieval"] = {
                            "backend": "gemini",
                            "mode": self._mode,
                            "doc_id": doc_id,
                            "merged_rank": rank,
                            "rrf_score": None,
                            "dense_rank": rank,
                            "dense_distance": _clean_number(dense_distance_by_id.get(doc_id)),
                            "sparse_rank": None,
                            "sparse_score": None,
                        }
                    results.append(doc)
                return results

        # BM25 retrieval
        sparse_score_by_id: dict[str, float] = {}
        if self._mode in ("hybrid", "sparse"):
            tokenized_query = _tokenize_code(query)
            bm25_scores = self._bm25.get_scores(tokenized_query)
            top_indices = sorted(
                range(len(bm25_scores)), key=lambda i: bm25_scores[i], reverse=True
            )[:fetch_k]
            bm25_ids = [self._doc_ids[i] for i in top_indices]
            sparse_score_by_id = {
                self._doc_ids[index]: bm25_scores[index]
                for index in top_indices
            }

            if self._mode == "sparse":
                results = []
                for rank, doc_id in enumerate(bm25_ids[:n_results], start=1):
                    if doc_id not in self._id_to_doc:
                        continue
                    doc = deepcopy(self._id_to_doc[doc_id])
                    if include_metadata:
                        doc["retrieval"] = {
                            "backend": "gemini",
                            "mode": self._mode,
                            "doc_id": doc_id,
                            "merged_rank": rank,
                            "rrf_score": None,
                            "dense_rank": None,
                            "dense_distance": None,
                            "sparse_rank": rank,
                            "sparse_score": _clean_number(sparse_score_by_id.get(doc_id)),
                        }
                    results.append(doc)
                return results

        # Reciprocal Rank Fusion
        merged = _rrf_details([dense_ids, bm25_ids], k=self._rrf_k)

        results = []
        for detail in merged[:n_results]:
            doc_id = detail["doc_id"]
            if doc_id in self._id_to_doc:
                doc = deepcopy(self._id_to_doc[doc_id])
                if include_metadata:
                    doc["retrieval"] = {
                        "backend": "gemini",
                        "mode": self._mode,
                        "doc_id": doc_id,
                        "merged_rank": detail.get("merged_rank"),
                        "rrf_score": _clean_number(detail.get("rrf_score")),
                        "dense_rank": detail.get("dense_rank"),
                        "dense_distance": _clean_number(dense_distance_by_id.get(doc_id)),
                        "sparse_rank": detail.get("sparse_rank"),
                        "sparse_score": _clean_number(sparse_score_by_id.get(doc_id)),
                    }
                results.append(doc)
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
_router_mode: str | None = None

# Experiment name -> knowledge base toggles
_EXPERIMENT_KB_PRESETS: dict[str, dict[str, bool]] = {
    "rag-pattern-only":     {"grammar": True,  "parallel_corpus": False, "api_mappings": False, "documentation": False, "api_sequences": False},
    "rag-pattern-samples":  {"grammar": True,  "parallel_corpus": True,  "api_mappings": False, "documentation": False, "api_sequences": False},
    "rag-pattern-api-docs": {"grammar": True,  "parallel_corpus": False, "api_mappings": True,  "documentation": True,  "api_sequences": False},
    "rag-full":             {"grammar": True,  "parallel_corpus": True,  "api_mappings": True,  "documentation": True,  "api_sequences": False},
    "rag-routed":           {"grammar": True,  "parallel_corpus": True,  "api_mappings": True,  "documentation": True,  "api_sequences": True},
}


from src.core.humaneval_artifacts import base_experiment_name as _base_experiment_name


def configure_kb_for_experiment(experiment: str) -> dict[str, bool] | None:
    """Apply knowledge base toggles for a given experiment name.

    Returns the active KB toggles, or None if the experiment doesn't use RAG
    (e.g. baseline).
    """
    global _kb_overrides, _router_mode
    experiment = _base_experiment_name(experiment)

    if experiment == "baseline":
        _kb_overrides = None
        _router_mode = None
        return None

    if experiment in _EXPERIMENT_KB_PRESETS:
        _kb_overrides = _EXPERIMENT_KB_PRESETS[experiment]
        _router_mode = "rule_based" if experiment == "rag-routed" else None
        return dict(_kb_overrides)

    # Unknown experiment name: fall back to YAML knowledge_bases section
    _kb_overrides = None
    _router_mode = None
    return _cfg().get("knowledge_bases")


def get_active_kb_toggles(experiment: str) -> dict[str, bool] | None:
    """Return the effective KB toggles for display purposes.

    Returns None for experiments that don't use RAG (baseline).
    """
    experiment = _base_experiment_name(experiment)
    if experiment == "baseline":
        return None
    if experiment in _EXPERIMENT_KB_PRESETS:
        return _EXPERIMENT_KB_PRESETS[experiment]
    return _cfg().get("knowledge_bases", {
        "code_snippets": True, "api_mappings": True, "documentation": True, "api_sequences": False,
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
        "code_snippets": True, "api_mappings": True, "documentation": True, "api_sequences": False,
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


def _get_api_sequence_retriever(backend: str = "chromadb") -> HybridRetriever | VertexAIRetriever:
    name = "api_sequences" if backend == "gemini" else collection_name_with_dim("api_sequences")
    cache_key = (backend, name)
    if cache_key not in _retrievers:
        if not GO_API_SEQUENCES_FILE.is_file():
            raise FileNotFoundError(
                f"API-sequence corpus not found at {GO_API_SEQUENCES_FILE}. "
                "Build it first with src/scripts/build_api_sequence_corpus.py."
            )
        records = _load_jsonl(GO_API_SEQUENCES_FILE)
        docs = []
        for index, record in enumerate(records):
            docs.append({
                "_id": record.get("_id", f"api_sequence_{index}"),
                "text": record["sequence_text"],
                "language": record.get("language", "go"),
                "source_corpus": record.get("source_corpus", "project_codenet_go"),
                "file_path": record.get("file_path", ""),
                "function_name": record.get("function_name", ""),
                "sequence_text": record["sequence_text"],
                "apis": record.get("apis", []),
                "imports": record.get("imports", []),
                "function_code": record.get("function_code", ""),
            })
        _retrievers[cache_key] = _make_retriever(
            backend, "api_sequences", docs, "text", mode="hybrid"
        )
    return _retrievers[cache_key]


class RAGResult:
    """Container for RAG retrieval results and the formatted context."""

    __slots__ = (
        "api_mappings",
        "documentation",
        "grammar_mappings",
        "parallel_corpus",
        "api_sequences",
        "context",
        "source_traces",
        "router_metadata",
        "prompt_metadata",
    )

    def __init__(self) -> None:
        retrieval_cfg = _cfg()["retrieval"]
        self.api_mappings: list[dict] = []
        self.documentation: list[dict] = []
        self.grammar_mappings: list[dict] = []
        self.parallel_corpus: list[dict] = []
        self.api_sequences: list[dict] = []
        self.context: str = ""
        self.source_traces: dict[str, dict] = _empty_source_traces()
        self.router_metadata: dict = {
            "mode": "preset",
            "decision": "static_kb_toggles",
            "signals": {},
            "notes": [],
        }
        self.prompt_metadata: dict = {
            "format": retrieval_cfg.get("prompt_format", "verbose"),
            "retrieval_contract": "on" if retrieval_cfg.get("retrieval_contract", True) else "off",
            "includes_retrieval": False,
        }

    def has_usable_items(self) -> bool:
        return rag_result_has_usable_items(self)

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
        "api_sequences": deepcopy(rag_result.api_sequences),
    }
    source_traces = _empty_source_traces(config)
    for source_name, trace in getattr(rag_result, "source_traces", {}).items():
        source_traces[source_name] = _normalize_source_trace(source_name, trace)

    prompt_metadata = deepcopy(getattr(rag_result, "prompt_metadata", {}))
    if "prompt_format" in config:
        prompt_metadata["format"] = config["prompt_format"]
    else:
        prompt_metadata.setdefault("format", "verbose")
    if "retrieval_contract" in config:
        prompt_metadata["retrieval_contract"] = "on" if config["retrieval_contract"] else "off"
    else:
        prompt_metadata.setdefault("retrieval_contract", "on")
    prompt_metadata.setdefault("includes_retrieval", rag_result_has_usable_items(rag_result))

    router_metadata = deepcopy(getattr(rag_result, "router_metadata", {}))
    router_metadata.setdefault("mode", "preset")
    router_metadata.setdefault("decision", "static_kb_toggles")
    router_metadata.setdefault("signals", {})
    router_metadata.setdefault("notes", [])

    queried_sources = [
        source_name for source_name, trace in source_traces.items()
        if trace.get("queried")
    ]
    accepted_sources = [
        source_name for source_name, trace in source_traces.items()
        if trace.get("accepted")
    ]
    return {
        "embedding_backend": embedding_backend,
        "kb_toggles": _normalize_kb_toggles(
            kb_toggles if kb_toggles is not None else _get_kb_toggles()
        ),
        "retrieval_config": {
            "parallel_corpus_k": config.get("parallel_corpus_k"),
            "api_mappings_k": config.get("api_mappings_k"),
            "go_docs_k": config.get("go_docs_k"),
            "api_sequences_k": config.get("api_sequences_k"),
            "rrf_k": config.get("rrf_k"),
            "enable_confidence_gate": config.get("enable_confidence_gate"),
            "hard_fallback_to_no_retrieval": config.get("hard_fallback_to_no_retrieval"),
            "prompt_format": config.get("prompt_format"),
            "retrieval_contract": config.get("retrieval_contract"),
            "source_rules": deepcopy(config.get("source_rules", {})),
        },
        "prompt": prompt_metadata,
        "router": router_metadata,
        "summary": {
            "has_usable_items": rag_result_has_usable_items(rag_result),
            "queried_sources": queried_sources,
            "accepted_sources": accepted_sources,
            "total_candidates": sum(trace.get("returned_count", 0) for trace in source_traces.values()),
            "total_accepted": sum(trace.get("accepted_count", 0) for trace in source_traces.values()),
        },
        "retrieval_counts": {
            key: len(value)
            for key, value in items.items()
        },
        "items": items,
        "sources": source_traces,
        "rendered_context": rag_result.context,
    }


def build_empty_retrieval_artifact(
    *,
    embedding_backend: str,
    kb_toggles: dict[str, bool] | None = None,
    retrieval_config: dict | None = None,
) -> dict:
    config = retrieval_config or _cfg()["retrieval"]
    rag_result = RAGResult()
    rag_result.source_traces = _empty_source_traces(config)
    rag_result.prompt_metadata = {
        "format": config.get("prompt_format", "verbose"),
        "retrieval_contract": "on" if config.get("retrieval_contract", True) else "off",
        "includes_retrieval": False,
    }
    return build_retrieval_artifact(
        rag_result,
        embedding_backend=embedding_backend,
        kb_toggles=kb_toggles,
        retrieval_config=config,
    )


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
    stripped = strip_python_comments(python_code)
    sections = []
    cfg = _cfg()["retrieval"]
    base_kb = _get_kb_toggles()
    result = RAGResult()
    result.source_traces = _empty_source_traces(cfg)
    accepted_mappings: list[dict] = []
    accepted_api_sequences: list[dict] = []
    api_info: dict | None = None
    grammar_patterns: list[dict] = []

    if base_kb.get("grammar", False) or _router_mode == "rule_based":
        from src.rag.grammar_extractor import extract_grammar_patterns

        grammar_patterns = extract_grammar_patterns(python_code)

    if (
        base_kb.get("api_mappings", False)
        or base_kb.get("api_sequences", False)
        or _router_mode == "rule_based"
    ):
        from src.rag.api_extractor import extract_api_info

        api_info = extract_api_info(python_code)

    if _router_mode == "rule_based":
        kb, router_metadata = _route_kb_toggles(
            base_kb,
            api_info=api_info or {
                "calls": [],
                "imports": [],
                "has_error_handling": False,
            },
            grammar_patterns=grammar_patterns,
            router_config=cfg.get("router", {}),
        )
        result.router_metadata = router_metadata
    else:
        kb = base_kb
        result.router_metadata = {
            "mode": "preset",
            "decision": "static_kb_toggles",
            "signals": {},
            "base_kb_toggles": _normalize_kb_toggles(base_kb),
            "selected_kb_toggles": _normalize_kb_toggles(kb),
            "notes": [],
        }

    use_grammar = kb.get("grammar", False)
    use_parallel = kb.get("parallel_corpus", False)
    use_api = kb.get("api_mappings", False)
    use_docs = kb.get("documentation", False)
    use_api_sequences = kb.get("api_sequences", False)

    # Step A: grammar_mappings — tree-sitter detects which constructs are present,
    # then query once per detected construct for a precise per-category match.
    if use_grammar:
        grammar_ret = _get_grammar_retriever(embedding_backend)
        grammar_matches = []
        seen_categories: set[str] = set()
        for pat in grammar_patterns:
            hits = grammar_ret.retrieve(pat["fragment"], n_results=1, include_metadata=True)
            for r in hits:
                if r["category"] not in seen_categories:
                    seen_categories.add(r["category"])
                    grammar_matches.append(r)
        accepted_matches, trace = _accept_retrieved_items(
            "grammar_mappings",
            grammar_matches,
            config=cfg,
        )
        trace["query_text"] = stripped
        result.grammar_mappings = accepted_matches
        result.source_traces["grammar_mappings"] = trace
        if accepted_matches:
            examples = []
            for match in accepted_matches:
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
        corpus_matches = corpus_ret.retrieve(
            stripped,
            n_results=cfg["parallel_corpus_k"],
            include_metadata=True,
        )
        accepted_matches, trace = _accept_retrieved_items(
            "parallel_corpus",
            corpus_matches,
            config=cfg,
        )
        trace["query_text"] = stripped
        result.parallel_corpus = accepted_matches
        result.source_traces["parallel_corpus"] = trace
        if accepted_matches:
            pairs = []
            for m in accepted_matches:
                go_code = sanitize_parallel_go_reference(m["go_code"])
                pairs.append(
                    f"**Python:**\n```python\n{m['python_code']}\n```\n"
                    f"**Go:**\n```go\n{go_code}\n```"
                )
            sections.append("## Python-Go Code Pairs\n\n" + "\n\n".join(pairs))

    # Step C: api_mappings — tree-sitter extracts API calls and imports,
    # then query with those specific names for a precise match.
    mappings: list[dict] = []
    if use_api:
        api_query = api_info["query_apis"]
        if api_info["query_imports"]:
            api_query += " " + api_info["query_imports"]
        if api_query.strip():
            api_ret = _get_api_retriever(embedding_backend)
            mappings = api_ret.retrieve(
                api_query,
                n_results=cfg["api_mappings_k"],
                include_metadata=True,
            )
        accepted_mappings, trace = _accept_retrieved_items(
            "api_mappings",
            mappings,
            config=cfg,
        )
        trace["query_text"] = api_query.strip()
        result.api_mappings = accepted_mappings
        result.source_traces["api_mappings"] = trace
        if accepted_mappings:
            lines = [
                f"- `{m['python_api']}` -> `{m['go_api']}`: {m['description']}"
                for m in accepted_mappings
            ]
            sections.append("## Relevant API Mappings\n" + "\n".join(lines))

    if use_api_sequences and api_info and api_info["query_apis"].strip():
        api_sequence_ret = _get_api_sequence_retriever(embedding_backend)
        api_sequences = api_sequence_ret.retrieve(
            api_info["query_apis"],
            n_results=cfg["api_sequences_k"],
            include_metadata=True,
        )
        accepted_api_sequences, trace = _accept_retrieved_items(
            "api_sequences",
            api_sequences,
            config=cfg,
        )
        trace["query_text"] = api_info["query_apis"]
        result.api_sequences = accepted_api_sequences
        result.source_traces["api_sequences"] = trace
        if accepted_api_sequences:
            sequence_lines = [
                f"- `{record['sequence_text']}`"
                for record in accepted_api_sequences
            ]
            sections.append("## Relevant Go API Usage Sequences\n" + "\n".join(sequence_lines))

    # Step D: go_docs — bridge query uses Go API names from step C
    if use_docs and accepted_mappings:
        go_api_query = " ".join(m["go_api"] for m in accepted_mappings)
        godoc_ret = _get_godoc_retriever(embedding_backend)
        docs = godoc_ret.retrieve(
            go_api_query,
            n_results=cfg["go_docs_k"],
            include_metadata=True,
        )
        accepted_docs, trace = _accept_retrieved_items(
            "documentation",
            docs,
            config=cfg,
        )
        trace["query_text"] = go_api_query
        result.documentation = accepted_docs
        result.source_traces["documentation"] = trace
        if accepted_docs:
            doc_lines = []
            for d in accepted_docs:
                line = f"- **{d['api']}**: {d['description']}"
                if d.get("example"):
                    line += f"\n  ```go\n  {d['example']}\n  ```"
                doc_lines.append(line)
            sections.append("## Go Documentation & Patterns\n" + "\n".join(doc_lines))

    result.context = "\n\n".join(sections) if sections else ""
    result.prompt_metadata["includes_retrieval"] = rag_result_has_usable_items(result)
    return result
