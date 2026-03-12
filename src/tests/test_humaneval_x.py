"""Tests for humaneval_x.py data loader.

    uv run pytest src/tests/test_humaneval_x.py -v

Unit tests mock the network call. The integration test hits HuggingFace.
    uv run pytest src/tests/test_humaneval_x.py -v -m integration
"""

from unittest.mock import patch, MagicMock

import pytest

from src.data.humaneval_x import load_humaneval_x, _HUMANEVAL_X_JSONL


# ---------------------------------------------------------------------------
# Unit tests (no network)
# ---------------------------------------------------------------------------

class TestLoadHumanEvalX:
    def test_returns_correct_structure(self):
        """Verify returned dicts have all required keys."""
        fake_py = [
            {"task_id": "Python/0", "declaration": "def f():", "prompt": "def f():\n", "canonical_solution": "  return 1\n"},
            {"task_id": "Python/1", "declaration": "def g():", "prompt": "def g():\n", "canonical_solution": "  return 2\n"},
        ]
        fake_go = [
            {"task_id": "Go/0", "declaration": "func f() int", "prompt": "func f() int {\n", "canonical_solution": "  return 1\n}", "test": "func TestF(t *testing.T) {}"},
            {"task_id": "Go/1", "declaration": "func g() int", "prompt": "func g() int {\n", "canonical_solution": "  return 2\n}", "test": "func TestG(t *testing.T) {}"},
        ]

        with patch("datasets.load_dataset") as mock_load:
            mock_load.side_effect = [fake_py, fake_go]
            pairs = load_humaneval_x()

        assert len(pairs) == 2
        required_keys = {"task_id", "declaration", "py_declaration", "py_solution", "go_solution", "test"}
        for pair in pairs:
            assert set(pair.keys()) == required_keys

    def test_task_id_comes_from_go(self):
        """task_id should be the Go task_id (e.g. Go/0), not Python."""
        fake_py = [{"task_id": "Python/0", "declaration": "def f():", "prompt": "", "canonical_solution": ""}]
        fake_go = [{"task_id": "Go/0", "declaration": "func f()", "prompt": "", "canonical_solution": "", "test": ""}]

        with patch("datasets.load_dataset") as mock_load:
            mock_load.side_effect = [fake_py, fake_go]
            pairs = load_humaneval_x()

        assert pairs[0]["task_id"] == "Go/0"

    def test_py_solution_concatenates_prompt_and_solution(self):
        fake_py = [{"task_id": "Python/0", "declaration": "def f():", "prompt": "def f():\n", "canonical_solution": "  return 1\n"}]
        fake_go = [{"task_id": "Go/0", "declaration": "func f()", "prompt": "", "canonical_solution": "", "test": ""}]

        with patch("datasets.load_dataset") as mock_load:
            mock_load.side_effect = [fake_py, fake_go]
            pairs = load_humaneval_x()

        assert pairs[0]["py_solution"] == "def f():\n  return 1\n"

    def test_go_solution_concatenates_prompt_and_solution(self):
        fake_py = [{"task_id": "Python/0", "declaration": "", "prompt": "", "canonical_solution": ""}]
        fake_go = [{"task_id": "Go/0", "declaration": "", "prompt": "func f() {\n", "canonical_solution": "  return 1\n}", "test": ""}]

        with patch("datasets.load_dataset") as mock_load:
            mock_load.side_effect = [fake_py, fake_go]
            pairs = load_humaneval_x()

        assert pairs[0]["go_solution"] == "func f() {\n  return 1\n}"

    def test_url_template_has_lang_placeholder(self):
        assert "{lang}" in _HUMANEVAL_X_JSONL

    def test_empty_dataset(self):
        with patch("datasets.load_dataset") as mock_load:
            mock_load.side_effect = [[], []]
            pairs = load_humaneval_x()

        assert pairs == []


# ---------------------------------------------------------------------------
# Integration test (hits HuggingFace)
# ---------------------------------------------------------------------------

@pytest.mark.integration
class TestLoadHumanEvalXIntegration:
    def test_loads_164_problems(self):
        pairs = load_humaneval_x()
        assert len(pairs) == 164

    def test_first_problem_structure(self):
        pairs = load_humaneval_x()
        p = pairs[0]
        assert p["task_id"] == "Go/0"
        assert "func" in p["declaration"]
        assert "def" in p["py_declaration"]
        assert len(p["test"]) > 0
