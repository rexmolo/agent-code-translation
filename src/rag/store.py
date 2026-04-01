"""ChromaDB client and collection management."""

import chromadb

from src.rag.embeddings import get_active_dimensions, load_rag_config


def get_chroma_client() -> chromadb.ClientAPI:
    cfg = load_rag_config()
    chroma_cfg = cfg["chromadb"]
    return chromadb.HttpClient(
        host=chroma_cfg["host"],
        port=chroma_cfg["port"],
    )


def collection_name_with_dim(base_name: str, dimensions: int | None = None) -> str:
    """Append ``_{dim}`` suffix to a collection name.

    If *dimensions* is None, reads the default from rag_config.yaml (Gemini only).
    """
    if dimensions is None:
        dimensions = get_active_dimensions()
    return f"{base_name}_{dimensions}"


def get_or_create_collection(
    client: chromadb.ClientAPI,
    name: str,
    embedding_function=None,
) -> chromadb.Collection:
    return client.get_or_create_collection(
        name=name,
        embedding_function=embedding_function,
        metadata={"hnsw:space": "cosine"},
    )
