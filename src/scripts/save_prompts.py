"""Save and compare LLM prompts for each RAG variant across embedding backends.

Generates prompts for Gemini (Vertex AI) and ChromaDB backends using
HumanEval-X problem 0, saves them to .doc/Notes/, then prints a diff
showing what changed between backends per variant.

Usage:
    uv run python src/scripts/save_prompts.py
"""

import difflib
from pathlib import Path

from src.data.humaneval_x import load_humaneval_x
from src.core.agents import _BASE_INSTRUCTIONS, _build_rag_section
from src.core.prompt_builder import PromptBuilder
from src.rag.retriever import (
    build_translation_context,
    configure_kb_for_experiment,
    get_active_kb_toggles,
)

OUTPUT_DIR = Path(__file__).resolve().parents[2] / ".doc" / "Notes"

VARIANTS = [
    "rag-pattern-only",
    "rag-pattern-samples",
    "rag-pattern-api-docs",
    "rag-full",
]

BACKENDS = [
    ("gemini",   "vec-gemini"),
    ("chromadb", "vec-chroma"),
]


def _build_full_prompt(python_code: str, go_signature: str, variant: str, backend: str) -> tuple[str, str]:
    """Return (system_prompt, user_prompt) for a given variant + backend."""
    # Reset retriever cache between backends
    from src.rag import retriever as _ret
    _ret._retrievers.clear()

    configure_kb_for_experiment(variant)
    kb_toggles = get_active_kb_toggles(variant)
    rag_result = build_translation_context(python_code, embedding_backend=backend)

    rag_section = _build_rag_section(kb_toggles) if kb_toggles else ""
    system_prompt = _BASE_INSTRUCTIONS + ("\n" + rag_section if rag_section else "")
    user_prompt = PromptBuilder().build_humaneval_x(
        python_code, go_signature=go_signature, rag_result=rag_result
    )
    return system_prompt, user_prompt


def main() -> None:
    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)

    print("Loading HumanEval-X problem 0...")
    pairs = load_humaneval_x()
    pair = pairs[0]
    python_code = pair["py_solution"]
    go_signature = pair["declaration"]
    task_id = pair["task_id"]
    print(f"Problem: {task_id}\n")

    # Store prompts for comparison
    prompts: dict[tuple[str, str], tuple[str, str]] = {}

    for backend, label in BACKENDS:
        print(f"\n=== Backend: {label} ===")
        for variant in VARIANTS:
            print(f"  Building prompt for: {variant} ...")
            system_prompt, user_prompt = _build_full_prompt(
                python_code, go_signature, variant, backend
            )
            prompts[(variant, backend)] = (system_prompt, user_prompt)

            out_path = OUTPUT_DIR / f"prompt_{variant}_{label}.md"
            out_path.write_text(
                f"# Prompt — {variant} [{label}]\n\n"
                f"**Source problem:** {task_id}\n\n"
                f"---\n\n"
                f"## System Prompt\n\n"
                f"```\n{system_prompt}\n```\n\n"
                f"---\n\n"
                f"## User Prompt\n\n"
                f"```\n{user_prompt}\n```\n",
                encoding="utf-8",
            )
            print(f"  Saved → {out_path.name}")

    # Compare vec-gemini vs vec-chroma for each variant
    print("\n\n" + "=" * 60)
    print("COMPARISON: vec-gemini vs vec-chroma")
    print("=" * 60)

    comparison_lines = [
        "# Prompt Comparison: vec-gemini vs vec-chroma",
        "",
        f"**Source problem:** {task_id}",
        "",
        "System prompts are identical across backends (same KB toggles).",
        "Differences below are in the **User Prompt** only (retrieved RAG content).",
        "",
    ]

    for variant in VARIANTS:
        _, gemini_user = prompts[(variant, "gemini")]
        _, chroma_user = prompts[(variant, "chromadb")]

        gemini_lines = gemini_user.splitlines(keepends=True)
        chroma_lines = chroma_user.splitlines(keepends=True)

        diff = list(difflib.unified_diff(
            gemini_lines, chroma_lines,
            fromfile=f"{variant} [vec-gemini]",
            tofile=f"{variant} [vec-chroma]",
            lineterm="",
        ))

        comparison_lines.append(f"## {variant}")
        comparison_lines.append("")

        if not diff:
            comparison_lines.append("✅ **Identical** — same content retrieved from both backends.")
            print(f"\n{variant}: ✅ IDENTICAL")
        else:
            added   = sum(1 for l in diff if l.startswith("+") and not l.startswith("+++"))
            removed = sum(1 for l in diff if l.startswith("-") and not l.startswith("---"))
            comparison_lines.append(f"⚠️ **Different** — {added} lines added, {removed} lines removed.")
            comparison_lines.append("")
            comparison_lines.append("```diff")
            comparison_lines.extend(l.rstrip() for l in diff)
            comparison_lines.append("```")
            print(f"\n{variant}: ⚠️  DIFFERENT (+{added} -{removed} lines)")
            for line in diff[:20]:
                print(" ", line.rstrip())
            if len(diff) > 20:
                print(f"  ... ({len(diff) - 20} more lines)")

        comparison_lines.append("")

    compare_path = OUTPUT_DIR / "prompt_comparison_gemini_vs_chroma.md"
    compare_path.write_text("\n".join(comparison_lines), encoding="utf-8")
    print(f"\n\nComparison saved → {compare_path.name}")
    print("\nDone.")


if __name__ == "__main__":
    main()
