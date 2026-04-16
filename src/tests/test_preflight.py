"""Preflight checks: API connection, Go compiler, environment.

Run these first. If any fail, the pipeline cannot work.

    uv run pytest src/tests/test_preflight.py -v
"""

import os
import subprocess
from io import StringIO
from pathlib import Path

import pytest
import yaml
from rich.console import Console

import src.config as config
import src.providers.registry as registry
from src.core.agents import _PlainTextTranslationAgent, create_translation_agent
from src.core.pipeline import preflight_check
from src.providers.minimax import MiniMax


class TestEnvironment:
    """Verify that required environment variables and tools are available."""

    def test_minimax_api_key_is_set(self):
        from src.config import load_providers_config, resolve_api_key
        key = resolve_api_key(load_providers_config()["minimax"])
        assert key is not None, "MiniMax API key is not configured in providers.yaml or env"
        assert len(key) > 10, "MiniMax API key looks too short to be valid"

    def test_go_compiler_available(self):
        result = subprocess.run(
            ["go", "version"], capture_output=True, text=True, timeout=10
        )
        assert result.returncode == 0, f"Go compiler not found: {result.stderr}"

    def test_python3_available(self):
        result = subprocess.run(
            ["python3", "--version"], capture_output=True, text=True, timeout=10
        )
        assert result.returncode == 0, f"python3 not found: {result.stderr}"


class TestAPIConnection:
    """Verify that the MiniMax API accepts our key and returns a response."""

    def test_minimax_api_responds(self):
        """Send a minimal prompt and verify we get a non-empty response."""
        model = MiniMax(id="MiniMax-M2.5")
        agent = create_translation_agent(model)
        response = agent.run("Translate to Go: print('hello')", stream=False)
        assert response is not None, "Agent returned None"
        assert response.content is not None, "Response content is None"
        content = response.content
        # TranslationResult has go_code field
        if hasattr(content, "go_code"):
            assert len(content.go_code.strip()) > 0, "go_code is empty"
        else:
            assert len(str(content).strip()) > 0, "Response is empty"


def test_minimax_uses_plaintext_translation_wrapper():
    """MiniMax advertises no structured-output support, so use fenced-code fallback."""
    model = MiniMax(id="MiniMax-M2.5")
    agent = create_translation_agent(model)
    assert agent.__class__.__name__ == "_PlainTextTranslationAgent"


def test_plaintext_translation_wrapper_does_not_promote_error_content():
    class DummyResponse:
        def __init__(self):
            self.content = "Error code: 529"
            self.status = "ERROR"

    class DummyInner:
        def run(self, *_args, **_kwargs):
            return DummyResponse()

    agent = _PlainTextTranslationAgent(DummyInner())
    response = agent.run("prompt", stream=False)
    assert response.content == "Error code: 529"


def test_plaintext_translation_wrapper_accepts_completed_like_status():
    class DummyResponse:
        def __init__(self):
            self.content = "```go\npackage main\n```"
            self.status = "RunStatus.COMPLETED"

    class DummyInner:
        def run(self, *_args, **_kwargs):
            return DummyResponse()

    agent = _PlainTextTranslationAgent(DummyInner())
    response = agent.run("prompt", stream=False)
    assert hasattr(response.content, "go_code")
    assert response.content.go_code.strip() == "package main"


def test_lmstudio_preflight_does_not_require_api_key(monkeypatch):
    """LM Studio should pass the credential gate with the shipped example config."""

    example_cfg = yaml.safe_load(Path("config/providers.yaml.example").read_text(encoding="utf-8"))
    monkeypatch.setattr(config, "_providers_cfg", example_cfg)
    registry._REGISTRY.clear()

    monkeypatch.setattr("shutil.which", lambda _cmd: "/usr/bin/go")

    class DummyResponse:
        content = "```go\npackage main\n```"

    class DummyAgent:
        def run(self, *_args, **_kwargs):
            return DummyResponse()

    monkeypatch.setattr("src.core.pipeline._agents.create_translation_agent", lambda _model: DummyAgent())

    buffer = StringIO()
    console = Console(file=buffer, force_terminal=False, color_system=None)
    ok = preflight_check(console, [("lmstudio", "qwen3_coder_next", object())])

    assert ok is True
    assert "does not require an API key" in buffer.getvalue()
