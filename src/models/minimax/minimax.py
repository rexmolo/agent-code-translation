"""MiniMax model class using Anthropic-compatible API."""

from dataclasses import dataclass, field
from os import getenv
from typing import Any, Dict, List, Optional, Type, Union

import httpx
from pydantic import BaseModel

from agno.models.anthropic import Claude

try:
    from anthropic import Anthropic as AnthropicClient
except ImportError as e:
    raise ImportError("`anthropic` not installed. Please install it with `pip install anthropic`") from e


@dataclass
class MiniMax(Claude):
    """
    MiniMax model accessed via the Anthropic-compatible API.

    Uses the Anthropic SDK pointed at MiniMax's endpoint:
    https://api.minimaxi.com/anthropic

    Supported models: MiniMax-M2.5, MiniMax-M2.1, MiniMax-M2
    """

    id: str = "MiniMax-M2.5"
    name: str = "MiniMax"
    provider: str = "MiniMax"
    base_url: str = "https://api.minimaxi.com/anthropic"

    def __post_init__(self):
        if self.client_params is None:
            self.client_params = {}
        self.client_params.setdefault("base_url", self.base_url)
        super().__post_init__()

    def _get_client_params(self) -> Dict[str, Any]:
        if self.api_key is None:
            self.api_key = getenv("MINIMAX_API_KEY")
        return super()._get_client_params()

    def get_client(self) -> AnthropicClient:
        """Return an Anthropic client with a dedicated httpx client for MiniMax.

        Avoids the shared global httpx singleton to prevent stale keepalive
        connections between the team coordinator and member agent calls.
        """
        if self.client and not self.client.is_closed():
            return self.client

        _client_params = self._get_client_params()
        if not isinstance(self.http_client, httpx.Client):
            self.http_client = httpx.Client(
                timeout=httpx.Timeout(300.0, connect=10.0),
                follow_redirects=True,
            )
        _client_params["http_client"] = self.http_client
        self.client = AnthropicClient(**_client_params)
        return self.client

    def _supports_structured_outputs(self) -> bool:
        """MiniMax does not support Anthropic's native structured outputs."""
        return False

    def _using_structured_outputs(
        self,
        response_format: Optional[Union[Dict, Type[BaseModel]]] = None,
        tools: Optional[List[Dict[str, Any]]] = None,
    ) -> bool:
        """MiniMax does not support structured outputs; skip without warning."""
        return False

    def _has_beta_features(
        self,
        response_format: Optional[Union[Dict, Type[BaseModel]]] = None,
        tools: Optional[List[Dict[str, Any]]] = None,
    ) -> bool:
        """MiniMax does not support Anthropic beta API features."""
        return False
