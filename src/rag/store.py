"""ChromaDB client and collection management."""

import chromadb

from src.rag.embeddings import load_rag_config


def get_chroma_client() -> chromadb.ClientAPI:
    cfg = load_rag_config()
    chroma_cfg = cfg["chromadb"]
    return chromadb.HttpClient(
        host=chroma_cfg["host"],
        port=chroma_cfg["port"],
    )


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
