"""ChromaDB client and collection management."""

import chromadb

from src.config import CHROMA_PERSIST_DIR


def get_chroma_client() -> chromadb.ClientAPI:
    CHROMA_PERSIST_DIR.mkdir(parents=True, exist_ok=True)
    return chromadb.PersistentClient(path=str(CHROMA_PERSIST_DIR))


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
