"""Tests for translation prompt construction."""

from pathlib import Path
from types import SimpleNamespace

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
        assert "Return only the Go declarations needed for that implementation." in prompt
        assert "Do not include `main()` or demo/example I/O." in prompt
        assert "Preserve the Python program's semantics and edge cases." in prompt
        assert "package and import statements only if they are required" in prompt

    def test_shared_agent_instructions_no_longer_require_standalone_programs(self):
        agents_source = Path(__file__).resolve().parents[1] / "core" / "agents.py"
        base_instructions = agents_source.read_text(encoding="utf-8")

        assert "standalone program" not in base_instructions
        assert "stdin" not in base_instructions
        assert "stdout" not in base_instructions
        assert "package declaration" not in base_instructions
        assert "You may be provided with" not in base_instructions
        assert "parallel corpus" not in base_instructions.lower()

    def test_rag_prompt_includes_retrieval_usage_contract(self):
        rag_result = SimpleNamespace(
            grammar_mappings=[
                {
                    "description": "list append pattern",
                    "python_pattern": "items.append(x)",
                    "go_pattern": "items = append(items, x)",
                }
            ],
            api_mappings=[
                {
                    "python_api": "len",
                    "go_api": "len",
                    "description": "length lookup",
                }
            ],
            parallel_corpus=[],
            documentation=[],
        )

        prompt = PromptBuilder().build_humaneval_x(
            python_code="def add(a, b):\n    return a + b\n",
            go_signature="func add(a int, b int) int",
            rag_result=rag_result,
        )

        assert "Retrieval usage contract:" in prompt
        assert (
            "Source semantics and any provided Go signature take priority over retrieved references."
            in prompt
        )
        assert (
            "Use retrieved material only when the APIs or control flow directly match the source code."
            in prompt
        )
        assert (
            "Ignore any retrieved example, mapping, or documentation that would change behavior, edge cases, or the function contract."
            in prompt
        )
        assert (
            "Treat parallel corpus code pairs as optional reference examples, not templates to copy."
            in prompt
        )

    def test_parallel_corpus_is_labeled_as_optional_reference_examples(self):
        rag_result = SimpleNamespace(
            grammar_mappings=[],
            api_mappings=[],
            parallel_corpus=[
                {
                    "python_code": "items = sorted(values)",
                    "go_code": "slices.Sort(values)",
                }
            ],
            documentation=[],
        )

        prompt = PromptBuilder().build(
            python_code="def sort_values(values):\n    return sorted(values)\n",
            rag_result=rag_result,
        )

        assert "Here are optional Python-Go reference examples" in prompt
        assert "Treat them as references, not templates:" in prompt
