"""Builds the final LLM prompt for Python-to-Go translation.

Usage:
    builder = PromptBuilder()

    # Baseline (no RAG)
    prompt = builder.build(python_code)

    # With RAG context
    prompt = builder.build(python_code, rag_result=rag_result)

    # HumanEval-X (includes a required Go function signature)
    prompt = builder.build(python_code, rag_result=rag_result, go_signature=declaration)
"""

from __future__ import annotations


_HUMANEVAL_X_INSTRUCTIONS = (
    "HumanEval-X instructions:\n"
    "- Implement the provided Go signature only.\n"
    "- Return only the Go declarations needed for that implementation.\n"
    "- Do not include `main()` or demo/example I/O.\n"
    "- Preserve the Python program's semantics and edge cases.\n"
    "- Include package and import statements only if they are required by the implementation."
)

_RETRIEVAL_USAGE_CONTRACT = (
    "Retrieval usage contract:\n"
    "- Source semantics and any provided Go signature take priority over retrieved references.\n"
    "- Use retrieved material only when the APIs or control flow directly match the source code.\n"
    "- Ignore any retrieved example, mapping, or documentation that would change behavior, edge cases, or the function contract.\n"
    "- Treat parallel corpus code pairs as optional reference examples, not templates to copy."
)


class PromptBuilder:
    """Assembles the translation prompt from Python source and optional RAG results."""

    def build(
        self,
        python_code: str,
        rag_result=None,
        go_signature: str | None = None,
    ) -> str:
        """Build the final prompt.

        Args:
            python_code:  The Python source code to translate.
            rag_result:   A RAGResult instance, or None for baseline.
            go_signature: Optional Go function signature (used for HumanEval-X).

        Returns:
            The complete prompt string to pass to the LLM.
        """
        parts: list[str] = []

        # --- Core instruction ---
        parts.append(
            "Translate the Python code below to Go."
        )

        # --- Python source ---
        parts.append(f"Python code:\n```python\n{python_code}\n```")

        # --- Go function signature (HumanEval-X only) ---
        if go_signature:
            parts.append(
                f"Use this Go function signature:\n```go\n{go_signature}\n```"
            )
            parts.append(_HUMANEVAL_X_INSTRUCTIONS)

        # --- RAG sections (only if rag_result is provided and has content) ---
        if rag_result is not None:
            parts.append(_RETRIEVAL_USAGE_CONTRACT)
            parts.extend(self._rag_sections(rag_result))

        return "\n\n".join(parts)

    def build_humaneval_x(
        self,
        python_code: str,
        go_signature: str,
        rag_result=None,
    ) -> str:
        return self.build(python_code, rag_result=rag_result, go_signature=go_signature)

    def build_local(self, python_code: str, rag_result=None) -> str:
        return self.build(python_code, rag_result=rag_result)

    # ------------------------------------------------------------------
    # Private helpers
    # ------------------------------------------------------------------

    def _rag_sections(self, rag_result) -> list[str]:
        sections: list[str] = []

        # Grammar patterns
        if rag_result.grammar_mappings:
            lines = []
            for m in rag_result.grammar_mappings:
                lines.append(
                    f"- {m['description']}\n"
                    f"  Python pattern:\n  ```python\n  {m['python_pattern']}\n  ```\n"
                    f"  Go pattern:\n  ```go\n  {m['go_pattern']}\n  ```"
                )
            sections.append(
                "Here are similar Python patterns and their Go equivalents that you can refer to:\n\n"
                + "\n\n".join(lines)
            )

        # API mappings
        if rag_result.api_mappings:
            lines = [
                f"- `{m['python_api']}` → `{m['go_api']}`: {m['description']}"
                for m in rag_result.api_mappings
            ]
            sections.append(
                "Here are Python APIs extracted from the source code and their Go equivalents:\n\n"
                + "\n".join(lines)
            )

        # Parallel corpus
        if rag_result.parallel_corpus:
            blocks = []
            for p in rag_result.parallel_corpus:
                blocks.append(
                    f"Python:\n```python\n{p['python_code']}\n```\n"
                    f"Go:\n```go\n{p['go_code']}\n```"
                )
            sections.append(
                "Here are optional Python-Go reference examples to help you understand "
                "how similar Python code is often expressed in Go. Treat them as references, not templates:\n\n"
                + "\n\n".join(blocks)
            )

        # Go documentation
        if rag_result.documentation:
            lines = []
            for d in rag_result.documentation:
                entry = f"- **{d['api']}**: {d['description']}"
                if d.get("example"):
                    entry += f"\n  ```go\n  {d['example']}\n  ```"
                lines.append(entry)
            sections.append(
                "Here is Go documentation for the relevant APIs:\n\n"
                + "\n\n".join(lines)
            )

        return sections
