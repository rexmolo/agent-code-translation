"""Central model registry for multi-model translation.

Supports multiple providers, each with multiple model variants.
See MODELS.md in this directory for the convention.

Usage:
    from src.providers.registry import enable_model, get_enabled_models

    enable_model("gemini", "2.5_flash_lite")

    for provider, variant, model in get_enabled_models():
        ...
"""

from __future__ import annotations

from typing import Any, Callable

# Two-level registry: provider_key -> {env_var, label, variants: {variant_key -> ...}}
_REGISTRY: dict[str, dict[str, Any]] = {}
_enabled: set[tuple[str, str]] = set()


def _register_defaults() -> None:
    """Populate the registry with known providers and variants (lazy imports)."""
    if _REGISTRY:
        return

    def _minimax_factory(model_id: str) -> Callable:
        def factory():
            from src.providers.minimax import MiniMax
            return MiniMax(id=model_id)
        return factory

    def _gemini_factory(model_id: str, location: str | None = None) -> Callable:
        def factory():
            from agno.models.google import Gemini
            return Gemini(id=model_id, location=location)
        return factory

    _REGISTRY["gemini"] = {
        "env_var": "GOOGLE_API_KEY",
        "vertex_env_vars": [
            "GOOGLE_GENAI_USE_VERTEXAI",
            "GOOGLE_CLOUD_PROJECT",
            "GOOGLE_CLOUD_LOCATION",
        ],
        "label": "Google Gemini",
        "variants": {
            "2.5_flash_lite": {
                "model_id": "gemini-2.5-flash-lite",
                "label": "Gemini 2.5 Flash Lite",
                "factory": _gemini_factory("gemini-2.5-flash-lite"),
            },
            "2.5_flash": {
                "model_id": "gemini-2.5-flash",
                "label": "Gemini 2.5 Flash",
                "factory": _gemini_factory("gemini-2.5-flash"),
            },
            "2.5_pro": {
                "model_id": "gemini-2.5-pro",
                "label": "Gemini 2.5 Pro",
                "factory": _gemini_factory("gemini-2.5-pro"),
            },
            "3_flash_preview": {
                "model_id": "gemini-3-flash-preview",
                "label": "Gemini 3 Flash Preview",
                "factory": _gemini_factory("gemini-3-flash-preview", location="global"),
            },
            "3_pro_preview": {
                "model_id": "gemini-3-pro-preview",
                "label": "Gemini 3 Pro Preview",
                "factory": _gemini_factory("gemini-3-pro-preview", location="global"),
            },
            "3.1_pro_preview": {
                "model_id": "gemini-3.1-pro-preview",
                "label": "Gemini 3.1 Pro Preview",
                "factory": _gemini_factory("gemini-3.1-pro-preview", location="global"),
            },
        },
    }

    _REGISTRY["minimax"] = {
        "env_var": "MINIMAX_API_KEY",
        "label": "MiniMax",
        "variants": {
            "M2.5": {
                "model_id": "MiniMax-M2.5",
                "label": "MiniMax M2.5",
                "factory": _minimax_factory("MiniMax-M2.5"),
            },
            "M2.1": {
                "model_id": "MiniMax-M2.1",
                "label": "MiniMax M2.1",
                "factory": _minimax_factory("MiniMax-M2.1"),
            },
            "M2": {
                "model_id": "MiniMax-M2",
                "label": "MiniMax M2",
                "factory": _minimax_factory("MiniMax-M2"),
            },
        },
    }


def list_providers() -> list[dict[str, Any]]:
    """Return info about all registered providers."""
    _register_defaults()
    return [
        {"key": k, "label": v["label"], "env_var": v["env_var"]}
        for k, v in _REGISTRY.items()
    ]


def list_variants(provider_key: str) -> list[dict[str, Any]]:
    """Return variant info for a specific provider."""
    _register_defaults()
    if provider_key not in _REGISTRY:
        raise KeyError(f"Unknown provider: {provider_key!r}. Available: {list(_REGISTRY)}")
    return [
        {"key": k, "model_id": v["model_id"], "label": v["label"]}
        for k, v in _REGISTRY[provider_key]["variants"].items()
    ]


def enable_model(provider_key: str, variant_key: str) -> None:
    """Enable a specific provider+variant combination."""
    _register_defaults()
    if provider_key not in _REGISTRY:
        raise KeyError(f"Unknown provider: {provider_key!r}. Available: {list(_REGISTRY)}")
    variants = _REGISTRY[provider_key]["variants"]
    if variant_key not in variants:
        raise KeyError(f"Unknown variant: {variant_key!r}. Available: {list(variants)}")
    _enabled.add((provider_key, variant_key))


def disable_model(provider_key: str, variant_key: str) -> None:
    """Disable a specific provider+variant combination."""
    _enabled.discard((provider_key, variant_key))


def get_enabled_models() -> list[tuple[str, str, Any]]:
    """Return (provider_key, variant_key, model_instance) triples for enabled models."""
    _register_defaults()
    results = []
    for provider_key, variant_key in _enabled:
        variant = _REGISTRY[provider_key]["variants"][variant_key]
        model = variant["factory"]()
        results.append((provider_key, variant_key, model))
    return results


def get_model_id(provider_key: str, variant_key: str) -> str:
    """Return the model_id string for a provider+variant combination."""
    _register_defaults()
    return _REGISTRY[provider_key]["variants"][variant_key]["model_id"]


def get_model_env_var(provider_key: str) -> str:
    """Return the environment variable name required for a provider."""
    _register_defaults()
    return _REGISTRY[provider_key]["env_var"]


def get_model_vertex_env_vars(provider_key: str) -> list[str] | None:
    """Return Vertex AI environment variable names for a provider, or None if unsupported."""
    _register_defaults()
    return _REGISTRY[provider_key].get("vertex_env_vars")


def reset() -> None:
    """Reset enabled set (useful for testing)."""
    _enabled.clear()
