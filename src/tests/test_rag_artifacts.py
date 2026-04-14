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

from src.rag.retriever import RAGResult, build_retrieval_artifact


def test_build_retrieval_artifact_serializes_counts_and_items():
    rag_result = RAGResult()
    rag_result.grammar_mappings = [{"category": "loop", "go_pattern": "for"}]
    rag_result.parallel_corpus = [{"problem_id": "p0", "python_code": "x", "go_code": "y"}]
    rag_result.api_mappings = [{"python_api": "len", "go_api": "len"}]
    rag_result.documentation = [{"api": "strings.TrimSpace", "description": "trim"}]
    rag_result.context = "retrieved context"

    artifact = build_retrieval_artifact(
        rag_result,
        embedding_backend="chromadb",
        kb_toggles={"grammar": True, "code_snippets": True, "api_mappings": False},
        retrieval_config={"parallel_corpus_k": 1, "api_mappings_k": 2, "go_docs_k": 1, "rrf_k": 60},
    )

    assert artifact["embedding_backend"] == "chromadb"
    assert artifact["kb_toggles"]["parallel_corpus"] is True
    assert "code_snippets" not in artifact["kb_toggles"]
    assert artifact["retrieval_counts"] == {
        "grammar_mappings": 1,
        "parallel_corpus": 1,
        "api_mappings": 1,
        "documentation": 1,
    }
    assert artifact["items"]["parallel_corpus"][0]["problem_id"] == "p0"
    assert artifact["rendered_context"] == "retrieved context"


def test_rag_result_to_artifact_uses_same_payload_shape():
    rag_result = RAGResult()
    rag_result.parallel_corpus = [{"problem_id": "p1"}]

    artifact = rag_result.to_artifact(
        embedding_backend="gemini",
        kb_toggles={"parallel_corpus": True},
        retrieval_config={"parallel_corpus_k": 1, "api_mappings_k": 2, "go_docs_k": 1, "rrf_k": 60},
    )

    assert artifact["embedding_backend"] == "gemini"
    assert artifact["retrieval_counts"]["parallel_corpus"] == 1
    assert artifact["items"]["parallel_corpus"][0]["problem_id"] == "p1"
