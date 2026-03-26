"""Agent definitions for the translation pipeline."""

from agno.agent import Agent

from src.core.schemas import TranslationResult

_BASE_INSTRUCTIONS = """\
You are an expert code translator specializing in Python to Go translation.

Given Python source code, translate it to idiomatic Go code.

Rules:
1. Preserve the exact logic and behavior of the original code.
2. Follow Go conventions: proper error handling, naming (camelCase), etc.
3. Include all necessary import statements and package declaration.
4. The code must be compilable and runnable as a standalone program.
5. If the Python code reads from stdin and writes to stdout, the Go code must do the same.
6. Use the standard library where possible.
"""

_KB_DESCRIPTIONS = {
    "grammar":        "- Similar Python→Go structural patterns for reference",
    "parallel_corpus": "- Python→Go code pairs showing how Python patterns translate to Go",
    "api_mappings":   "- API mapping equivalences between Python and Go standard libraries",
    "documentation":  "- Go standard library documentation",
}

_RAG_GUIDANCE = (
    "Use these references to inform your translation choices (e.g. correct Go API names, "
    "idiomatic patterns), but do not copy examples blindly — adapt them to the specific code."
)


def _build_rag_section(kb_toggles: dict) -> str:
    """Build the 'You may be provided with' paragraph for the active KBs."""
    active = [
        _KB_DESCRIPTIONS[key]
        for key in ("grammar", "parallel_corpus", "api_mappings", "documentation")
        if kb_toggles.get(key, False) and key in _KB_DESCRIPTIONS
    ]
    if not active:
        return ""
    lines = ["You may be provided with:"] + active + ["", _RAG_GUIDANCE]
    return "\n".join(lines)


def create_translation_agent(model, kb_toggles: dict | None = None) -> Agent:
    """Create the Python-to-Go translation agent.

    Args:
        model:       An Agno model instance (e.g. MiniMax, Gemini).
        kb_toggles:  Dict of active knowledge bases, e.g.
                     {"grammar": True, "api_mappings": True, ...}.
                     Pass None for baseline (no RAG guidance in instructions).
    """
    instructions = _BASE_INSTRUCTIONS
    if kb_toggles:
        rag_section = _build_rag_section(kb_toggles)
        if rag_section:
            instructions = _BASE_INSTRUCTIONS + "\n" + rag_section

    return Agent(
        name="Translator",
        role="Translate Python source code to Go",
        model=model,
        instructions=instructions,
        output_schema=TranslationResult,
    )
