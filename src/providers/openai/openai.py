"""OpenAI GPT model class using Agno's built-in OpenAIChat."""

from dataclasses import dataclass

from agno.models.openai import OpenAIChat


@dataclass
class GPT(OpenAIChat):
    """GPT model that resolves its API key from providers.yaml."""

    name: str = "GPT"
    provider: str = "OpenAI"

    def __post_init__(self):
        if self.api_key is None:
            from src.config import load_providers_config, resolve_api_key
            self.api_key = resolve_api_key(load_providers_config()["openai"])
        super().__post_init__()
