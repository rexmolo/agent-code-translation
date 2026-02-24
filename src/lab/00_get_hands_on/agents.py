"""Agent definitions for the translation pipeline."""

import importlib

from agno.agent import Agent

from src.models.minimax import MiniMax

_models = importlib.import_module("src.lab.00_get_hands_on.models")

TranslationResult = _models.TranslationResult

TRANSLATION_INSTRUCTIONS = """\
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


def _get_model() -> MiniMax:
    return MiniMax(id="MiniMax-M2.5")


def create_translation_agent() -> Agent:
    """Create the Python-to-Go translation agent."""
    return Agent(
        name="Translator",
        role="Translate Python source code to Go",
        model=_get_model(),
        instructions=TRANSLATION_INSTRUCTIONS,
        output_schema=TranslationResult,
    )
