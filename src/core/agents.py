"""Agent definitions for the translation pipeline."""

from agno.agent import Agent

from src.core.schemas import TranslationResult

TRANSLATION_INSTRUCTIONS = """\
You are an expert code translator specializing in Python to Go translation.

Given Python source code, translate it to idiomatic Go code.

You may be provided with:
- Similar Python→Go translation examples for reference
- API mapping equivalences between Python and Go standard libraries
- Go standard library documentation

Use these references to inform your translation choices (e.g. correct Go API names, \
idiomatic patterns), but do not copy examples blindly — adapt them to the specific code.

Rules:
1. Preserve the exact logic and behavior of the original code.
2. Follow Go conventions: proper error handling, naming (camelCase), etc.
3. Include all necessary import statements and package declaration.
4. The code must be compilable and runnable as a standalone program.
5. If the Python code reads from stdin and writes to stdout, the Go code must do the same.
6. Use the standard library where possible.
"""


def create_translation_agent(model) -> Agent:
    """Create the Python-to-Go translation agent.

    Args:
        model: An Agno model instance (e.g. MiniMax, Gemini).
    """
    return Agent(
        name="Translator",
        role="Translate Python source code to Go",
        model=model,
        instructions=TRANSLATION_INSTRUCTIONS,
        output_schema=TranslationResult,
    )
