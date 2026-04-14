"""Agent definitions for the translation pipeline."""

from agno.agent import Agent

from src.core.schemas import TranslationResult

_BASE_INSTRUCTIONS = """\
You translate Python source code to Go.

Preserve the original semantics.
Follow the exact task contract given in the user prompt.
Return only the Go code required by that contract.
"""


def create_translation_agent(model, kb_toggles: dict | None = None) -> Agent:
    """Create the Python-to-Go translation agent.

    Args:
        model:       An Agno model instance (e.g. MiniMax, Gemini).
        kb_toggles:  Dict of active knowledge bases, e.g.
                     {"grammar": True, "api_mappings": True, ...}.
                     Reserved for compatibility; task-specific retrieval wording
                     is assembled in PromptBuilder.
    """

    return Agent(
        name="Translator",
        role="Translate Python source code to Go",
        model=model,
        instructions=_BASE_INSTRUCTIONS,
        output_schema=TranslationResult,
    )
