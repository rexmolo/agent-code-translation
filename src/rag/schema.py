"""Shared RAG schema helpers used outside the retriever hot path."""

from __future__ import annotations

RAG_SOURCES = (
    "grammar_mappings",
    "parallel_corpus",
    "api_mappings",
    "documentation",
    "api_sequences",
    "translation_traps",
)


def rag_result_has_usable_items(rag_result: object | None) -> bool:
    """Return True when a RAG result has any retrieved items in known sources."""
    if rag_result is None:
        return False

    for field in RAG_SOURCES:
        value = getattr(rag_result, field, None)
        if isinstance(value, list) and value:
            return True
    return False
