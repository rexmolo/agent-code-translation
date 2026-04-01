"""Tests for OpenAI (GPT-5.4) provider registration and model class."""

from unittest.mock import patch

import pytest
from agno.models.openai import OpenAIChat

from src.providers.openai import GPT
from src.providers.registry import (
    _REGISTRY,
    enable_model,
    get_enabled_models,
    get_model_env_var,
    get_model_id,
    list_providers,
    list_variants,
    reset,
)

_PATCH_RESOLVE = "src.config.resolve_api_key"
_PATCH_LOAD = "src.config.load_providers_config"
_FAKE_CFG = {"openai": {"api_key_env": "OPENAI_API_KEY"}}


# ── GPT model class ──────────────────────────────────────────────────────────


class TestGPTClass:
    def test_is_subclass_of_openai_chat(self):
        assert issubclass(GPT, OpenAIChat)

    def test_default_fields(self):
        with patch(_PATCH_RESOLVE, return_value="sk-test"), \
             patch(_PATCH_LOAD, return_value=_FAKE_CFG):
            model = GPT(id="gpt-5.4")
        assert model.id == "gpt-5.4"
        assert model.name == "GPT"
        assert model.provider == "OpenAI"

    def test_resolves_api_key_from_config(self):
        with patch(_PATCH_RESOLVE, return_value="sk-from-yaml"), \
             patch(_PATCH_LOAD, return_value=_FAKE_CFG):
            model = GPT(id="gpt-5.4")
        assert model.api_key == "sk-from-yaml"

    def test_explicit_api_key_skips_config(self):
        model = GPT(id="gpt-5.4", api_key="sk-explicit")
        assert model.api_key == "sk-explicit"


# ── Registry ──────────────────────────────────────────────────────────────────


class TestOpenAIRegistry:
    def setup_method(self):
        reset()

    def test_openai_in_providers(self):
        providers = list_providers()
        keys = [p["key"] for p in providers]
        assert "openai" in keys

    def test_openai_env_var(self):
        assert get_model_env_var("openai") == "OPENAI_API_KEY"

    def test_gpt54_variant_listed(self):
        variants = list_variants("openai")
        keys = [v["key"] for v in variants]
        assert "gpt-5.4" in keys

    def test_gpt54_model_id(self):
        assert get_model_id("openai", "gpt-5.4") == "gpt-5.4"

    def test_enable_and_get_model(self):
        enable_model("openai", "gpt-5.4")
        with patch(_PATCH_RESOLVE, return_value="sk-test"), \
             patch(_PATCH_LOAD, return_value=_FAKE_CFG):
            models = get_enabled_models()
        assert len(models) == 1
        provider_key, variant_key, model = models[0]
        assert provider_key == "openai"
        assert variant_key == "gpt-5.4"
        assert isinstance(model, GPT)
        assert isinstance(model, OpenAIChat)

    def test_enable_invalid_variant_raises(self):
        with pytest.raises(KeyError, match="Unknown variant"):
            enable_model("openai", "nonexistent")

    def test_enable_invalid_provider_raises(self):
        with pytest.raises(KeyError, match="Unknown provider"):
            enable_model("nonexistent", "gpt-5.4")

    def test_factory_produces_correct_model_id(self):
        enable_model("openai", "gpt-5.4")
        with patch(_PATCH_RESOLVE, return_value="sk-test"), \
             patch(_PATCH_LOAD, return_value=_FAKE_CFG):
            models = get_enabled_models()
        _, _, model = models[0]
        assert model.id == "gpt-5.4"
