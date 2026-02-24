"""Agno Team wiring: coordinates translation and evaluation agents."""

import importlib
import os

from agno.models.litellm import LiteLLM
from agno.team.team import Team

_agents = importlib.import_module("src.lab.00_get_hands_on.agents")


def create_translation_team() -> Team:
    """Create a team that coordinates translation, testing, and evaluation."""
    translator = _agents.create_translation_agent()
    test_generator = _agents.create_test_generator_agent()
    test_translator = _agents.create_test_translator_agent()
    evaluator = _agents.create_evaluation_agent()

    return Team(
        name="Translation Pipeline",
        model=LiteLLM(
            id="minimax/MiniMax-M2.5",
            api_key=os.getenv("MINIMAX_API_KEY"),
        ),
        members=[translator, test_generator, test_translator, evaluator],
        instructions=[
            "You coordinate Python-to-Go code translation and evaluation.",
            "Step 1: Send the Python code to the Translator for translation to Go.",
            "Step 2: Check if Python tests were provided. If not, send the Python code to TestGenerator to generate tests.",
            "Step 3: Send the Python tests to the Evaluator to verify they pass against the Python source (using run_python_tests).",
            "Step 4: Send the verified Python tests and the Go source code to TestTranslator to translate tests to Go.",
            "Step 5: Send everything to the Evaluator for full evaluation: compilation, execution, I/O comparison, and Go test execution.",
            "Step 6: Return the final Go code, Go test code, and evaluation results.",
        ],
        share_member_interactions=True,
        show_members_responses=True,
        add_member_tools_to_context=False,
        markdown=True,
    )
