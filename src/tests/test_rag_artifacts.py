import sys
import types

sys.modules.setdefault("rank_bm25", types.SimpleNamespace(BM25Okapi=object))
sys.modules.setdefault(
    "src.rag.embeddings",
    types.SimpleNamespace(
        get_embedding_function=lambda: None,
        load_rag_config=lambda: {
            "retrieval": {
                "parallel_corpus_k": 1,
                "api_mappings_k": 2,
                "go_docs_k": 1,
                "rrf_k": 60,
            }
        },
    ),
)
sys.modules.setdefault(
    "src.rag.store",
    types.SimpleNamespace(collection_name_with_dim=lambda name: name),
)

from src.rag.retriever import (
    RAGResult,
    _route_kb_toggles,
    build_empty_retrieval_artifact,
    build_retrieval_artifact,
    rag_result_has_usable_items,
)


def test_build_retrieval_artifact_serializes_counts_and_items():
    rag_result = RAGResult()
    rag_result.grammar_mappings = [{"category": "loop", "go_pattern": "for"}]
    rag_result.parallel_corpus = [{"problem_id": "p0", "python_code": "x", "go_code": "y"}]
    rag_result.api_mappings = [{"python_api": "len", "go_api": "len"}]
    rag_result.documentation = [{"api": "strings.TrimSpace", "description": "trim"}]
    rag_result.api_sequences = [{"sequence_text": "strings.TrimSpace -> strings.Split"}]
    rag_result.context = "retrieved context"
    rag_result.source_traces["api_mappings"] = {
        "source": "api_mappings",
        "queried": True,
        "accepted": True,
        "query_text": "len",
        "returned_count": 1,
        "accepted_count": 1,
        "items": [{"python_api": "len", "go_api": "len", "retrieval": {"merged_rank": 1, "accepted": True}}],
        "acceptance": {"enabled": True, "rules": {"max_merged_rank": 3}, "fallback_to_no_retrieval": True},
    }

    artifact = build_retrieval_artifact(
        rag_result,
        embedding_backend="chromadb",
        kb_toggles={"grammar": True, "code_snippets": True, "api_mappings": False},
        retrieval_config={
            "parallel_corpus_k": 1,
            "api_mappings_k": 2,
            "go_docs_k": 1,
            "api_sequences_k": 2,
            "rrf_k": 60,
            "enable_confidence_gate": True,
            "hard_fallback_to_no_retrieval": True,
            "prompt_format": "verbose",
            "retrieval_contract": True,
            "source_rules": {"api_mappings": {"max_merged_rank": 3}},
        },
    )

    assert artifact["embedding_backend"] == "chromadb"
    assert artifact["kb_toggles"]["parallel_corpus"] is True
    assert "code_snippets" not in artifact["kb_toggles"]
    assert artifact["prompt"]["format"] == "verbose"
    assert artifact["router"]["mode"] == "preset"
    assert artifact["summary"]["has_usable_items"] is True
    assert "api_mappings" in artifact["summary"]["accepted_sources"]
    assert artifact["retrieval_counts"] == {
        "grammar_mappings": 1,
        "parallel_corpus": 1,
        "api_mappings": 1,
        "documentation": 1,
        "api_sequences": 1,
    }
    assert artifact["items"]["parallel_corpus"][0]["problem_id"] == "p0"
    assert artifact["sources"]["api_mappings"]["queried"] is True
    assert artifact["sources"]["api_mappings"]["items"][0]["retrieval"]["merged_rank"] == 1
    assert artifact["rendered_context"] == "retrieved context"


def test_rag_result_to_artifact_uses_same_payload_shape():
    rag_result = RAGResult()
    rag_result.parallel_corpus = [{"problem_id": "p1"}]

    artifact = rag_result.to_artifact(
        embedding_backend="gemini",
        kb_toggles={"parallel_corpus": True},
        retrieval_config={
            "parallel_corpus_k": 1,
            "api_mappings_k": 2,
            "go_docs_k": 1,
            "api_sequences_k": 2,
            "rrf_k": 60,
            "enable_confidence_gate": True,
            "hard_fallback_to_no_retrieval": True,
            "prompt_format": "verbose",
            "retrieval_contract": True,
            "source_rules": {},
        },
    )

    assert artifact["embedding_backend"] == "gemini"
    assert artifact["retrieval_counts"]["parallel_corpus"] == 1
    assert artifact["items"]["parallel_corpus"][0]["problem_id"] == "p1"
    assert artifact["summary"]["has_usable_items"] is True


def test_rag_result_has_usable_items_checks_all_sources():
    rag_result = RAGResult()
    assert rag_result_has_usable_items(rag_result) is False
    assert rag_result.has_usable_items() is False

    rag_result.documentation = [{"api": "strings.TrimSpace"}]
    assert rag_result_has_usable_items(rag_result) is True
    assert rag_result.has_usable_items() is True


def test_build_empty_retrieval_artifact_uses_frozen_schema():
    artifact = build_empty_retrieval_artifact(
        embedding_backend="chromadb",
        kb_toggles={"parallel_corpus": False},
        retrieval_config={
            "parallel_corpus_k": 1,
            "api_mappings_k": 2,
            "go_docs_k": 1,
            "api_sequences_k": 2,
            "rrf_k": 60,
            "enable_confidence_gate": True,
            "hard_fallback_to_no_retrieval": True,
            "prompt_format": "compact",
            "retrieval_contract": False,
            "source_rules": {"parallel_corpus": {"max_merged_rank": 1}},
        },
    )

    assert artifact["summary"]["has_usable_items"] is False
    assert artifact["prompt"]["format"] == "compact"
    assert artifact["prompt"]["retrieval_contract"] == "off"
    assert artifact["prompt"]["includes_retrieval"] is False
    assert set(artifact["sources"]) == {
        "grammar_mappings",
        "parallel_corpus",
        "api_mappings",
        "documentation",
        "api_sequences",
    }
    assert artifact["sources"]["parallel_corpus"]["acceptance"]["enabled"] is True


def test_route_kb_toggles_prefers_api_sources_for_api_heavy_tasks():
    routed, metadata = _route_kb_toggles(
        {
            "grammar": True,
            "parallel_corpus": True,
            "api_mappings": True,
            "documentation": True,
            "api_sequences": True,
        },
        api_info={
            "calls": ["json.loads", "items.sort", "os.path.join"],
            "imports": ["json", "os"],
            "has_error_handling": True,
        },
        grammar_patterns=[{"category": "List Comprehensions", "fragment": "x"}],
        router_config={
            "min_api_calls_for_api_mappings": 1,
            "min_imports_for_api_mappings": 1,
            "min_api_calls_for_api_sequences": 2,
            "min_grammar_patterns_for_grammar": 1,
            "max_api_calls_for_parallel": 1,
            "max_imports_for_parallel": 1,
            "max_grammar_patterns_for_parallel": 1,
            "disable_parallel_on_error_handling": True,
        },
    )

    assert routed["grammar"] is True
    assert routed["api_mappings"] is True
    assert routed["documentation"] is True
    assert routed["api_sequences"] is True
    assert routed["parallel_corpus"] is False
    assert metadata["mode"] == "rule_based"
    assert metadata["signals"]["api_call_count"] == 3


def test_route_kb_toggles_allows_parallel_for_simple_tasks():
    routed, metadata = _route_kb_toggles(
        {
            "grammar": True,
            "parallel_corpus": True,
            "api_mappings": True,
            "documentation": True,
            "api_sequences": True,
        },
        api_info={
            "calls": ["len"],
            "imports": [],
            "has_error_handling": False,
        },
        grammar_patterns=[],
        router_config={
            "min_api_calls_for_api_mappings": 1,
            "min_imports_for_api_mappings": 1,
            "min_api_calls_for_api_sequences": 2,
            "min_grammar_patterns_for_grammar": 1,
            "max_api_calls_for_parallel": 1,
            "max_imports_for_parallel": 1,
            "max_grammar_patterns_for_parallel": 1,
            "disable_parallel_on_error_handling": True,
        },
    )

    assert routed["parallel_corpus"] is True
    assert routed["api_sequences"] is False
    assert routed["grammar"] is False
    assert metadata["selected_kb_toggles"]["parallel_corpus"] is True


def test_route_kb_toggles_records_sources_disabled_by_base_preset():
    routed, metadata = _route_kb_toggles(
        {
            "grammar": True,
            "parallel_corpus": False,
            "api_mappings": True,
            "documentation": True,
            "api_sequences": False,
        },
        api_info={
            "calls": ["json.loads", "os.path.join", "items.sort"],
            "imports": ["json", "os"],
            "has_error_handling": False,
        },
        grammar_patterns=[{"category": "List Comprehensions", "fragment": "x"}],
        router_config={
            "min_api_calls_for_api_mappings": 1,
            "min_imports_for_api_mappings": 1,
            "min_api_calls_for_api_sequences": 2,
            "min_grammar_patterns_for_grammar": 1,
            "max_api_calls_for_parallel": 10,
            "max_imports_for_parallel": 10,
            "max_grammar_patterns_for_parallel": 10,
            "disable_parallel_on_error_handling": False,
        },
    )

    assert routed["parallel_corpus"] is False
    assert routed["api_sequences"] is False
    assert "disabled_by_base_preset" in metadata["reasons"]["parallel_corpus"]
    assert "disabled_by_base_preset" in metadata["reasons"]["api_sequences"]
