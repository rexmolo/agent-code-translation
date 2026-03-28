"""Embedding function factory — reads provider from rag_config.yaml."""

from __future__ import annotations

import os
from pathlib import Path

import yaml
import chromadb.utils.embedding_functions as embedding_functions

_CONFIG_PATH = Path(__file__).resolve().parent.parent.parent / "config" / "rag_config.yaml"


def load_rag_config() -> dict:
    with open(_CONFIG_PATH, encoding="utf-8") as f:
        return yaml.safe_load(f)


def _google_cfg() -> dict:
    """Return the google section from providers.yaml (cached via load_providers_config)."""
    from src.config import load_providers_config
    return load_providers_config()["google"]


class GeminiEmbeddingFunction:
    """Wrap Google GenAI embedding API for use with Vertex AI Vector Search.

    Also implements the ChromaDB EmbeddingFunction protocol so it can be used
    interchangeably where needed.
    """

    def __init__(self, model_name: str = "gemini-embedding-001"):
        from google import genai

        gcfg = _google_cfg()
        vertex_cfg = gcfg["vertex_ai"]

        use_vertex = os.environ.get(vertex_cfg["enabled_env"], "").lower() == "true"

        if use_vertex:
            # Vertex AI mode: uses Application Default Credentials (gcloud auth)
            project = os.environ.get(vertex_cfg["project_env"], "")
            location = os.environ.get(
                vertex_cfg["location_env"], vertex_cfg["default_location"],
            )
            if not project:
                raise ValueError(
                    f"{vertex_cfg['project_env']} is not set. "
                    "Set it in .env to use Gemini embeddings via Vertex AI."
                )
            self._client = genai.Client(
                vertexai=True,
                project=project,
                location=location,
            )
        else:
            # Standard API key mode
            from src.config import resolve_api_key
            api_key = resolve_api_key(gcfg)
            if not api_key:
                raise ValueError(
                    "Google API key is not configured. "
                    "Set it in config/providers.yaml or enable Vertex AI mode "
                    f"({vertex_cfg['enabled_env']}=true) to use Gemini embeddings."
                )
            self._client = genai.Client(api_key=api_key)

        self._model = model_name

    def __call__(self, input: list[str]) -> list[list[float]]:
        """Embed a list of texts, returning a list of float vectors."""
        if not input:
            return []
        result = self._client.models.embed_content(
            model=self._model,
            contents=input,
        )
        return [e.values for e in result.embeddings]

    def embed_query(self, text: str | None = None, input=None) -> list[float] | list[list[float]]:
        """Embed query string(s).

        - Called with `text` (str): single query, returns list[float]. Used by VertexAIRetriever.
        - Called with `input` (list[str]): batch of queries, returns list[list[float]]. Used by ChromaDB.
        """
        if isinstance(input, list):
            return self(input)
        query = text if text is not None else input
        return self([query])[0]

    def name(self) -> str:
        """ChromaDB embedding function protocol — identifies this function by model name."""
        return self._model


def get_embedding_function(
    provider_override: str | None = None,
) -> embedding_functions.EmbeddingFunction | GeminiEmbeddingFunction:
    """Return the embedding function specified in rag_config.yaml.

    If *provider_override* is given, it takes precedence over the YAML config.
    """
    cfg = load_rag_config()
    provider = provider_override or cfg["embedding"]["provider"]

    if provider == "default":
        return embedding_functions.DefaultEmbeddingFunction()

    if provider == "openai":
        from src.config import load_providers_config, resolve_api_key
        api_key = resolve_api_key(load_providers_config()["openai"])
        if not api_key:
            raise ValueError(
                "OpenAI API key is not configured. "
                "Set it in config/providers.yaml or switch embedding.provider to 'default' in rag_config.yaml."
            )
        model = cfg["embedding"]["openai"]["model"]
        return embedding_functions.OpenAIEmbeddingFunction(
            api_key=api_key,
            model_name=model,
        )

    if provider == "gemini":
        model = cfg["embedding"]["gemini"]["model"]
        return GeminiEmbeddingFunction(model_name=model)

    raise ValueError(f"Unknown embedding provider: {provider}")
