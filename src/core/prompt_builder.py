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

import re

from src.rag.schema import rag_result_has_usable_items
from src.rag.rendering import sanitize_parallel_go_reference


_HUMANEVAL_X_INSTRUCTIONS = (
    "HumanEval-X output contract:\n"
    "- Emit a complete Go file in this exact order:\n"
    "    1. `package main`\n"
    "    2. A single `import (...)` block listing every standard-library package the body calls "
    "(e.g. `strings`, `sort`, `strconv`, `math`, `slices`, `regexp`, `unicode`, `fmt`).\n"
    "    3. The function implementing the provided Go signature.\n"
    "- Do not add `main()` or demo/example I/O.\n"
    "- Preserve the Python program's semantics and edge cases.\n"
    "- Retrieved snippets below are reference only — do not copy their package layout or omit "
    "imports just because a snippet does."
)

_PARALLEL_REFERENCE_HEADER = (
    "// Reference only — do NOT copy package/imports; add your own based on the APIs you use."
)

_GO_STDLIB_IMPORT_RE = re.compile(r"^([a-z][a-z0-9_/]*)\.")


def _derive_import_path(go_api: str) -> str | None:
    """Return the stdlib import path for a qualified Go API like `strings.Join`."""
    if not go_api:
        return None
    text = go_api.strip().lstrip("`").lstrip("*&")
    match = _GO_STDLIB_IMPORT_RE.match(text)
    if not match:
        return None
    pkg = match.group(1)
    if "/" in pkg or pkg in {
        "strings", "strconv", "sort", "math", "slices", "regexp", "unicode",
        "fmt", "errors", "bytes", "bufio", "os", "io", "time", "utf8", "utf16",
    }:
        return pkg
    return pkg


def _import_suffix(go_api: str) -> str:
    pkg = _derive_import_path(go_api)
    return f" (import `\"{pkg}\"`)" if pkg else ""

_RETRIEVAL_USAGE_CONTRACT = (
    "Retrieval usage contract:\n"
    "- Source semantics and any provided Go signature take priority over retrieved references.\n"
    "- Use retrieved material only when the APIs or control flow directly match the source code.\n"
    "- Ignore any retrieved example, mapping, or documentation that would change behavior, edge cases, or the function contract.\n"
    "- Treat parallel corpus code pairs as optional reference examples, not templates to copy."
)


def _load_prompt_config() -> dict:
    from src.rag.embeddings import load_rag_config

    return load_rag_config().get("retrieval", {})


def _compact_code_snippet(text: str, *, limit: int = 96) -> str:
    collapsed = " ".join(text.split())
    if len(collapsed) <= limit:
        return collapsed
    return collapsed[: limit - 3].rstrip() + "..."


class PromptBuilder:
    """Assembles the translation prompt from Python source and optional RAG results."""

    def __init__(
        self,
        *,
        prompt_format: str | None = None,
        retrieval_contract: bool | None = None,
    ) -> None:
        retrieval_cfg = _load_prompt_config()
        self._prompt_format = prompt_format or retrieval_cfg.get("prompt_format", "verbose")
        if retrieval_contract is None:
            retrieval_contract = retrieval_cfg.get("retrieval_contract", True)
        self._retrieval_contract = bool(retrieval_contract)

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

        # --- RAG sections (only if rag_result is provided and has content) ---
        if rag_result_has_usable_items(rag_result):
            self._stamp_prompt_metadata(rag_result, includes_retrieval=True)
            if self._retrieval_contract:
                parts.append(_RETRIEVAL_USAGE_CONTRACT)
            parts.extend(self._rag_sections(rag_result, prompt_format=self._prompt_format))
        elif rag_result is not None:
            self._stamp_prompt_metadata(rag_result, includes_retrieval=False)

        # --- HumanEval-X output contract LAST, so it's most salient to the model ---
        if go_signature:
            parts.append(_HUMANEVAL_X_INSTRUCTIONS)

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

    def _stamp_prompt_metadata(self, rag_result, *, includes_retrieval: bool) -> None:
        prompt_metadata = getattr(rag_result, "prompt_metadata", None)
        if isinstance(prompt_metadata, dict):
            prompt_metadata["format"] = self._prompt_format
            prompt_metadata["retrieval_contract"] = "on" if self._retrieval_contract else "off"
            prompt_metadata["includes_retrieval"] = includes_retrieval

    def _rag_sections(self, rag_result, *, prompt_format: str) -> list[str]:
        if prompt_format == "compact":
            return self._compact_rag_sections(rag_result)
        return self._verbose_rag_sections(rag_result)

    def _verbose_rag_sections(self, rag_result) -> list[str]:
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
                f"- `{m['python_api']}` → `{m['go_api']}`{_import_suffix(m['go_api'])}: {m['description']}"
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
                go_code = sanitize_parallel_go_reference(p["go_code"])
                blocks.append(
                    f"Python:\n```python\n{p['python_code']}\n```\n"
                    f"Go:\n```go\n{_PARALLEL_REFERENCE_HEADER}\n{go_code}\n```"
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
                entry = f"- **{d['api']}**{_import_suffix(d['api'])}: {d['description']}"
                if d.get("example"):
                    entry += f"\n  ```go\n  {d['example']}\n  ```"
                lines.append(entry)
            sections.append(
                "Here is Go documentation for the relevant APIs:\n\n"
                + "\n\n".join(lines)
            )

        if getattr(rag_result, "api_sequences", None):
            lines = [
                f"- `{record['sequence_text']}`"
                for record in rag_result.api_sequences
            ]
            sections.append(
                "Here are relevant Go API usage sequences:\n\n"
                + "\n".join(lines)
            )

        return sections

    def _compact_rag_sections(self, rag_result) -> list[str]:
        sections: list[str] = []

        if rag_result.grammar_mappings:
            lines = []
            for mapping in rag_result.grammar_mappings:
                go_idiom = " ".join(mapping["go_pattern"].split())
                lines.append(
                    f"- {mapping['description']}\n"
                    f"  Go idiom: `{go_idiom}`"
                )
            sections.append(
                "Relevant Go grammar evidence:\n\n" + "\n".join(lines)
            )

        if rag_result.api_mappings:
            lines = [
                f"- `{mapping['python_api']}` -> `{mapping['go_api']}`{_import_suffix(mapping['go_api'])}: {mapping['description']}"
                for mapping in rag_result.api_mappings
            ]
            sections.append(
                "Relevant API evidence:\n\n" + "\n".join(lines)
            )

        if rag_result.parallel_corpus:
            lines = []
            for pair in rag_result.parallel_corpus:
                python_code = _compact_code_snippet(pair["python_code"])
                go_code = _compact_code_snippet(sanitize_parallel_go_reference(pair["go_code"]))
                lines.append(
                    f"- Python `{python_code}` -> Go `{go_code}`"
                )
            sections.append(
                "Optional reference pairs (reference only — do NOT copy package/imports):\n\n"
                + "\n".join(lines)
            )

        if rag_result.documentation:
            lines = []
            for doc in rag_result.documentation:
                usage_note = doc["description"]
                if doc.get("example"):
                    example = " ".join(doc["example"].split())
                    usage_note = f"{usage_note} Usage: `{example}`"
                lines.append(
                    f"- API `{doc['api']}`{_import_suffix(doc['api'])}: {usage_note}"
                )
            sections.append(
                "Relevant Go documentation:\n\n" + "\n".join(lines)
            )

        if getattr(rag_result, "api_sequences", None):
            lines = [
                f"- Sequence `{record['sequence_text']}`"
                for record in rag_result.api_sequences
            ]
            sections.append(
                "Relevant Go API usage sequences:\n\n" + "\n".join(lines)
            )

        return sections
