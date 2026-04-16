"""Helpers for rendering retrieved RAG snippets into prompts."""

from __future__ import annotations

import re


def sanitize_parallel_go_reference(go_code: str) -> str:
    """Strip standalone-program scaffolding from parallel corpus Go examples.

    Retrieved parallel corpus examples often contain full `package main`,
    import blocks, and `func main()` wrappers from competitive-programming
    style solutions. Those are useful retrieval artifacts, but poor prompt
    examples for HumanEval-X declaration-only generation.
    """
    lines = go_code.splitlines()
    result: list[str] = []
    in_import_block = False
    in_main = False
    brace_depth = 0

    for line in lines:
        stripped = line.strip()

        if stripped.startswith("package "):
            continue

        if re.match(r'^import\s+"', stripped) and not stripped.endswith("("):
            continue

        if stripped.startswith("import (") or stripped == "import(":
            in_import_block = True
            continue

        if in_import_block:
            if stripped == ")":
                in_import_block = False
            continue

        if re.match(r"^func\s+main\s*\(", stripped):
            in_main = True
            brace_depth = line.count("{") - line.count("}")
            if brace_depth <= 0 and "{" in line:
                in_main = False
            continue

        if in_main:
            brace_depth += line.count("{") - line.count("}")
            if brace_depth <= 0:
                in_main = False
            continue

        result.append(line)

    sanitized = "\n".join(result).strip()
    return sanitized or go_code.strip()
