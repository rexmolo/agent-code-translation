"""Tests for translation prompt construction."""

from pathlib import Path

from src.core.prompt_builder import PromptBuilder


class TestPromptBuilder:
    def test_humaneval_x_prompt_is_declaration_only(self):
        prompt = PromptBuilder().build_humaneval_x(
            python_code="def add(a, b):\n    return a + b\n",
            go_signature="func add(a int, b int) int",
        )

        assert "func add(a int, b int) int" in prompt
        assert "HumanEval-X instructions:" in prompt
        assert "standalone program" not in prompt
        assert "Implement the provided Go signature only." in prompt
        assert "Return only the Go code needed for the function implementation." in prompt
        assert "Do not include `main()` or demo/example I/O." in prompt
        assert "package and import statements only if they are required" in prompt

    def test_shared_agent_instructions_no_longer_require_standalone_programs(self):
        agents_source = Path(__file__).resolve().parents[1] / "core" / "agents.py"
        base_instructions = agents_source.read_text(encoding="utf-8")

        assert "standalone program" not in base_instructions
        assert "stdin" not in base_instructions
        assert "stdout" not in base_instructions
        assert "package declaration" not in base_instructions
