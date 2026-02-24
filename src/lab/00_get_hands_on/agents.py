"""Agent definitions for the translation and evaluation pipeline."""

import importlib
import os

from agno.agent import Agent
from agno.models.litellm import LiteLLM

_models = importlib.import_module("src.lab.00_get_hands_on.models")
_tools = importlib.import_module("src.lab.00_get_hands_on.tools")

TranslationResult = _models.TranslationResult
TestGenerationResult = _models.TestGenerationResult
TestTranslationResult = _models.TestTranslationResult

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

EVALUATION_INSTRUCTIONS = """\
You are a code evaluation specialist. Your job is to evaluate whether a Go translation
of Python code is correct.

Given the original Python code and the translated Go code, you must:
1. Use compile_go_code to check if the Go code compiles.
2. If it compiles, use run_python_code and run_go_code to execute both programs.
3. Use compare_outputs to check if they produce the same output.
4. If Go test code is provided, use run_go_tests to run the tests.
5. Report your findings clearly.

Always run all available checks systematically.
"""

TEST_GENERATION_INSTRUCTIONS = """\
You are an expert Python test writer. Given Python source code, generate comprehensive
tests using pytest.

Rules:
1. Import from the source module using: from source import *
2. Test all public functions and classes.
3. Include edge cases and boundary conditions.
4. Use clear test names that describe what is being tested.
5. Each test should be independent and self-contained.
6. Use pytest assertions (assert, pytest.raises, etc.).
"""

TEST_TRANSLATION_INSTRUCTIONS = """\
You are an expert at translating Python tests to Go tests.

Given Python test code and the corresponding Go source code, translate the tests to Go.

Rules:
1. Use the Go testing package (import "testing").
2. Test function names must start with Test (e.g., TestAdd).
3. Use the same package as the source code.
4. Preserve the test logic and assertions exactly.
5. Use t.Errorf or t.Fatalf for assertions.
6. The test file must be compilable alongside the source code.
"""


def _get_model() -> LiteLLM:
    return LiteLLM(
        id="minimax/MiniMax-M2.5",
        api_key=os.getenv("MINIMAX_API_KEY"),
    )


def create_translation_agent() -> Agent:
    """Create the Python-to-Go translation agent."""
    return Agent(
        name="Translator",
        role="Translate Python source code to Go",
        model=_get_model(),
        instructions=TRANSLATION_INSTRUCTIONS,
        output_schema=TranslationResult,
    )


def create_evaluation_agent() -> Agent:
    """Create the evaluation agent with code execution tools."""
    return Agent(
        name="Evaluator",
        role="Evaluate the correctness of translated Go code",
        model=_get_model(),
        instructions=EVALUATION_INSTRUCTIONS,
        tools=[
            _tools.compile_go_code,
            _tools.run_go_code,
            _tools.run_python_code,
            _tools.compare_outputs,
            _tools.run_python_tests,
            _tools.run_go_tests,
        ],
    )


def create_test_generator_agent() -> Agent:
    """Create the Python test generation agent."""
    return Agent(
        name="TestGenerator",
        role="Generate Python tests for source code",
        model=_get_model(),
        instructions=TEST_GENERATION_INSTRUCTIONS,
        output_schema=TestGenerationResult,
    )


def create_test_translator_agent() -> Agent:
    """Create the test translator agent (Python tests -> Go tests)."""
    return Agent(
        name="TestTranslator",
        role="Translate Python tests to Go tests",
        model=_get_model(),
        instructions=TEST_TRANSLATION_INSTRUCTIONS,
        output_schema=TestTranslationResult,
    )
