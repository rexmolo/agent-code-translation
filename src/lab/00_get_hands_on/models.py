"""Pydantic schemas for the translation and evaluation pipeline."""

from pydantic import BaseModel, Field


class TranslationResult(BaseModel):
    """Structured output from the translation agent."""

    go_code: str = Field(..., description="The translated Go source code, complete and compilable")
    explanation: str = Field(
        default="", description="Brief notes on key translation decisions"
    )


class TestGenerationResult(BaseModel):
    """Structured output for generated Python tests."""

    test_code: str = Field(..., description="Complete Python test code using unittest or pytest")
    explanation: str = Field(
        default="", description="Brief notes on what the tests cover"
    )


class TestTranslationResult(BaseModel):
    """Structured output for translated Go tests."""

    test_code: str = Field(..., description="Complete Go test code using testing package")
    explanation: str = Field(
        default="", description="Brief notes on test translation decisions"
    )


class EvaluationRecord(BaseModel):
    """Evaluation result for a single translated file."""

    source_file: str
    target_file: str
    compiles: bool = False
    runs_successfully: bool = False
    io_equivalent: bool = False
    computational_accuracy: bool = False
    test_pass_rate: float = 0.0
    tests_total: int = 0
    tests_passed: int = 0
    notes: str = ""
