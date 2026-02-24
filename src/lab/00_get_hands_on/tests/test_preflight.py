"""Preflight checks: API connection, Go compiler, environment.

Run these first. If any fail, the pipeline cannot work.

    uv run pytest src/lab/00_get_hands_on/tests/test_preflight.py -v
"""

import importlib
import os
import subprocess

import pytest
from dotenv import load_dotenv

load_dotenv()

_agents = importlib.import_module("src.lab.00_get_hands_on.agents")


class TestEnvironment:
    """Verify that required environment variables and tools are available."""

    def test_minimax_api_key_is_set(self):
        key = os.getenv("MINIMAX_API_KEY")
        assert key is not None, "MINIMAX_API_KEY is not set in .env"
        assert len(key) > 10, "MINIMAX_API_KEY looks too short to be valid"

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
        agent = _agents.create_translation_agent()
        response = agent.run("Translate to Go: print('hello')", stream=False)
        assert response is not None, "Agent returned None"
        assert response.content is not None, "Response content is None"
        content = response.content
        # TranslationResult has go_code field
        if hasattr(content, "go_code"):
            assert len(content.go_code.strip()) > 0, "go_code is empty"
        else:
            assert len(str(content).strip()) > 0, "Response is empty"
