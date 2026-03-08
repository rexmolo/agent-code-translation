"""Embedding function factory — reads provider from rag_config.yaml."""

import os
from pathlib import Path

import yaml
import chromadb.utils.embedding_functions as embedding_functions

_CONFIG_PATH = Path(__file__).resolve().parent.parent.parent / "config" / "rag_config.yaml"


def load_rag_config() -> dict:
    with open(_CONFIG_PATH, encoding="utf-8") as f:
        return yaml.safe_load(f)


def get_embedding_function() -> embedding_functions.EmbeddingFunction:
    """Return the embedding function specified in rag_config.yaml."""
    cfg = load_rag_config()
    provider = cfg["embedding"]["provider"]

    if provider == "default":
        return embedding_functions.DefaultEmbeddingFunction()

    if provider == "openai":
        api_key = os.environ.get("OPENAI_API_KEY", "")
        if not api_key:
            raise ValueError(
                "OPENAI_API_KEY is not set. "
                "Set it in .env or switch embedding.provider to 'default' in rag_config.yaml."
            )
        model = cfg["embedding"]["openai"]["model"]
        return embedding_functions.OpenAIEmbeddingFunction(
            api_key=api_key,
            model_name=model,
        )

    raise ValueError(f"Unknown embedding provider: {provider}")
