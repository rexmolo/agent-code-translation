"""LM Studio local model class using Agno's OpenAIChat.

LM Studio exposes an OpenAI-compatible server (default http://localhost:1234/v1).
The API key is ignored by LM Studio but the SDK requires a non-empty string.
"""

from dataclasses import dataclass

from agno.models.openai import OpenAIChat


@dataclass
class LMStudio(OpenAIChat):
    name: str = "LMStudio"
    provider: str = "LMStudio"
    # Hard cap to prevent runaway generation on local coder models
    max_tokens: int | None = 8192

    def __post_init__(self):
        from src.config import load_providers_config
        cfg = load_providers_config().get("lmstudio", {})
        if self.base_url is None:
            self.base_url = cfg.get("base_url", "http://localhost:1234/v1")
        if self.api_key is None:
            self.api_key = cfg.get("api_key") or "lm-studio"
        super().__post_init__()
