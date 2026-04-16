"""Agent definitions for the translation pipeline."""

import re

from agno.agent import Agent

from src.core.schemas import TranslationResult

_BASE_INSTRUCTIONS = """\
You translate Python source code to Go.

Preserve the original semantics.
Follow the exact task contract given in the user prompt.
Return only the Go code required by that contract.
"""

_PLAINTEXT_INSTRUCTIONS = """\
You translate Python source code to Go.

Preserve the original semantics.
Follow the exact task contract given in the user prompt.
Return ONLY the Go code inside a single fenced code block like:
```go
<your Go code here>
```
Do not include prose before or after the code block.
"""

_GO_FENCE_RE = re.compile(r"```(?:go)?\s*\n(.*?)```", re.DOTALL | re.IGNORECASE)


def _extract_go_code(text: str) -> str:
    match = _GO_FENCE_RE.search(text)
    if match:
        return match.group(1).rstrip()
    return text.strip()


def _supports_structured_outputs(model) -> bool:
    """Return whether the model can reliably honor Agno structured outputs."""
    checker = getattr(model, "_supports_structured_outputs", None)
    if callable(checker):
        try:
            return bool(checker())
        except Exception:
            return True
    return True


class _PlainTextTranslationAgent:
    """Wraps an Agno Agent whose model cannot reliably emit structured output.

    The inner agent returns plain text containing a fenced Go code block;
    this wrapper parses it and hands back a TranslationResult so the
    pipeline sees the same type it always sees.
    """

    def __init__(self, inner: Agent):
        self._inner = inner

    def run(self, *args, **kwargs):
        response = self._inner.run(*args, **kwargs)
        content = getattr(response, "content", None)
        status = str(getattr(response, "status", "") or "").upper()
        is_error = "ERROR" in status or "FAIL" in status
        if isinstance(content, str) and not is_error:
            response.content = TranslationResult(
                go_code=_extract_go_code(content),
                explanation="",
            )
        return response

    def __getattr__(self, name):
        return getattr(self._inner, name)


def create_translation_agent(model, kb_toggles: dict | None = None):
    """Create the Python-to-Go translation agent.

    Args:
        model:       An Agno model instance (e.g. MiniMax, Gemini).
        kb_toggles:  Dict of active knowledge bases, e.g.
                     {"grammar": True, "api_mappings": True, ...}.
                     Reserved for compatibility; task-specific retrieval wording
                     is assembled in PromptBuilder.
    """

    if getattr(model, "provider", None) == "LMStudio" or not _supports_structured_outputs(model):
        inner = Agent(
            name="Translator",
            role="Translate Python source code to Go",
            model=model,
            instructions=_PLAINTEXT_INSTRUCTIONS,
        )
        return _PlainTextTranslationAgent(inner)

    return Agent(
        name="Translator",
        role="Translate Python source code to Go",
        model=model,
        instructions=_BASE_INSTRUCTIONS,
        output_schema=TranslationResult,
    )
