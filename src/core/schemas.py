"""Pydantic schemas for the translation and evaluation pipeline."""

from pydantic import BaseModel, Field


class TranslationResult(BaseModel):
    """Structured output from the translation agent."""

    go_code: str = Field(
        ...,
        description="The translated Go source code that satisfies the task-specific output contract",
    )
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
    """Evaluation result for a single translated file/task."""

    source_file: str
    target_file: str
    dataset: str = "local"
    compiles: bool = False
    runs_successfully: bool = False
    pass_at_1: bool = False
    ast_similarity: float = 0.0
    tests_total: int = 0
    tests_passed: int = 0
    notes: str = ""
    repair_attempted: bool = False
    repair_succeeded: bool = False
    first_pass_compiles: bool | None = None
    first_pass_runs_successfully: bool | None = None
    first_pass_pass_at_1: bool | None = None
    first_pass_notes: str = ""
    repair_notes: str = ""
