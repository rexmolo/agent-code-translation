"""Tests for translation prompt construction."""

from pathlib import Path
from types import SimpleNamespace

from src.core.prompt_builder import PromptBuilder
from src.rag.rendering import sanitize_parallel_go_reference


class TestPromptBuilder:
    def test_humaneval_x_prompt_is_declaration_only(self):
        prompt = PromptBuilder().build_humaneval_x(
            python_code="def add(a, b):\n    return a + b\n",
            go_signature="func add(a int, b int) int",
        )

        assert "func add(a int, b int) int" in prompt
        assert "HumanEval-X output contract:" in prompt
        assert "standalone program" not in prompt
        assert "Emit a complete Go file in this exact order:" in prompt
        assert "The function implementing the provided Go signature." in prompt
        assert "Do not add `main()` or demo/example I/O." in prompt
        assert "Preserve the Python program's semantics and edge cases." in prompt
        assert "A single `import (...)` block listing every standard-library package the body calls" in prompt

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

        prompt = PromptBuilder(
            prompt_format="verbose",
            retrieval_contract=True,
        ).build_humaneval_x(
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
                    "go_code": 'package main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("demo")\n}\n\nfunc sortValues(values []int) {\n    slices.Sort(values)\n}\n',
                }
            ],
            documentation=[],
        )

        prompt = PromptBuilder(
            prompt_format="verbose",
            retrieval_contract=True,
        ).build(
            python_code="def sort_values(values):\n    return sorted(values)\n",
            rag_result=rag_result,
        )

        assert "Here are optional Python-Go reference examples" in prompt
        assert "Treat them as references, not templates:" in prompt
        assert "package main" not in prompt
        assert 'import "fmt"' not in prompt
        assert "func main()" not in prompt
        assert "func sortValues(values []int)" in prompt

    def test_empty_rag_result_uses_baseline_prompt_path(self):
        rag_result = SimpleNamespace(
            grammar_mappings=[],
            api_mappings=[],
            parallel_corpus=[],
            documentation=[],
        )

        prompt = PromptBuilder().build_humaneval_x(
            python_code="def add(a, b):\n    return a + b\n",
            go_signature="func add(a int, b int) int",
            rag_result=rag_result,
        )

        assert "Retrieval usage contract:" not in prompt

    def test_compact_prompt_format_uses_bullets_instead_of_fences_for_docs(self):
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
                    "python_api": "strip",
                    "go_api": "strings.TrimSpace",
                    "description": "trim whitespace",
                }
            ],
            parallel_corpus=[],
            documentation=[
                {
                    "api": "strings.TrimSpace",
                    "description": "Removes leading and trailing whitespace.",
                    "example": "strings.TrimSpace(value)",
                }
            ],
            prompt_metadata={},
        )

        prompt = PromptBuilder(
            prompt_format="compact",
            retrieval_contract=True,
        ).build(
            python_code="def clean(value):\n    return value.strip()\n",
            rag_result=rag_result,
        )

        assert "Relevant Go grammar evidence:" in prompt
        assert "Go idiom: `items = append(items, x)`" in prompt
        assert "Relevant Go documentation:" in prompt
        assert "```go" not in prompt
        assert rag_result.prompt_metadata["format"] == "compact"
        assert rag_result.prompt_metadata["retrieval_contract"] == "on"

    def test_retrieval_contract_can_be_disabled(self):
        rag_result = SimpleNamespace(
            grammar_mappings=[{"description": "x", "python_pattern": "x", "go_pattern": "x"}],
            api_mappings=[],
            parallel_corpus=[],
            documentation=[],
            api_sequences=[],
            prompt_metadata={},
        )

        prompt = PromptBuilder(
            prompt_format="compact",
            retrieval_contract=False,
        ).build(
            python_code="def f():\n    return 1\n",
            rag_result=rag_result,
        )

        assert "Retrieval usage contract:" not in prompt
        assert "Relevant Go grammar evidence:" in prompt
        assert rag_result.prompt_metadata["retrieval_contract"] == "off"

    def test_compact_parallel_corpus_truncates_long_snippets(self):
        rag_result = SimpleNamespace(
            grammar_mappings=[],
            api_mappings=[],
            parallel_corpus=[
                {
                    "python_code": "def sort_values(values):\n    cleaned = [v.strip() for v in values if v]\n    cleaned.sort()\n    return cleaned\n",
                    "go_code": "func sortValues(values []string) []string {\n    cleaned := make([]string, 0, len(values))\n    for _, v := range values {\n        if v != \"\" {\n            cleaned = append(cleaned, strings.TrimSpace(v))\n        }\n    }\n    slices.Sort(cleaned)\n    return cleaned\n}\n",
                }
            ],
            documentation=[],
            api_sequences=[],
            prompt_metadata={},
        )

        prompt = PromptBuilder(
            prompt_format="compact",
            retrieval_contract=True,
        ).build(
            python_code="def sort_values(values):\n    return sorted(v.strip() for v in values if v)\n",
            rag_result=rag_result,
        )

        assert "Optional reference pairs (reference only" in prompt
        assert "..." in prompt

    def test_sanitize_parallel_go_reference_strips_program_scaffolding(self):
        go_code = (
            "package main\n\n"
            "import (\n"
            '    "fmt"\n'
            '    "strings"\n'
            ")\n\n"
            "func helper() string {\n"
            '    return strings.TrimSpace(" x ")\n'
            "}\n\n"
            "func main() {\n"
            '    fmt.Println(helper())\n'
            "}\n"
        )

        sanitized = sanitize_parallel_go_reference(go_code)

        assert "package main" not in sanitized
        assert "import (" not in sanitized
        assert "func main()" not in sanitized
        assert "func helper() string" in sanitized

    def test_prompt_builder_noops_metadata_stamp_when_attr_missing(self):
        rag_result = SimpleNamespace(
            grammar_mappings=[{"description": "x", "python_pattern": "x", "go_pattern": "x"}],
            api_mappings=[],
            parallel_corpus=[],
            documentation=[],
            api_sequences=[],
        )

        prompt = PromptBuilder(
            prompt_format="compact",
            retrieval_contract=False,
        ).build(
            python_code="def f():\n    return 1\n",
            rag_result=rag_result,
        )

        assert "Relevant Go grammar evidence:" in prompt
        assert not hasattr(rag_result, "prompt_metadata")

    def test_compact_prompt_renders_api_sequences(self):
        rag_result = SimpleNamespace(
            grammar_mappings=[],
            api_mappings=[],
            parallel_corpus=[],
            documentation=[],
            api_sequences=[
                {"sequence_text": "strings.TrimSpace -> strings.Split -> strings.Join"}
            ],
            prompt_metadata={},
        )

        prompt = PromptBuilder(
            prompt_format="compact",
            retrieval_contract=True,
        ).build(
            python_code="def normalize(value):\n    return '|'.join(value.strip().split(','))\n",
            rag_result=rag_result,
        )

        assert "Relevant Go API usage sequences:" in prompt
        assert "strings.TrimSpace -> strings.Split -> strings.Join" in prompt
