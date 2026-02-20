#!/usr/bin/env python3
"""
Extract a Python source file's AST using tree-sitter and save it as JSON.

Usage:
    uv run src/scripts/extract_python_ast.py [path/to/file.py]

If no file is provided, a small sample snippet is used for demonstration.
Output is saved to src/temp/python_ast.json.
"""

import json
import sys
from pathlib import Path

import tree_sitter_python as tspython
from tree_sitter import Language, Parser

from src.config import TEMP_DIR

PY_LANGUAGE = Language(tspython.language())

SAMPLE_CODE = """\
def fibonacci(n: int) -> int:
    if n <= 1:
        return n
    a, b = 0, 1
    for _ in range(2, n + 1):
        a, b = b, a + b
    return b


class Calculator:
    def __init__(self, value: float = 0):
        self.value = value

    def add(self, x: float) -> "Calculator":
        self.value += x
        return self
"""

OUTPUT_FILE = TEMP_DIR / "python_ast.json"


def node_to_dict(node) -> dict:
    """Recursively convert a tree-sitter node into a nested dict."""
    result = {
        "type": node.type,
        "start": [node.start_point.row, node.start_point.column],
        "end": [node.end_point.row, node.end_point.column],
    }
    if node.child_count == 0:
        result["text"] = node.text.decode("utf-8")
    else:
        result["children"] = [node_to_dict(child) for child in node.children]
    return result


def print_tree(node, indent: int = 0) -> str:
    """Return a human-readable indented tree string."""
    lines = []
    prefix = "  " * indent
    if node.child_count == 0:
        lines.append(f"{prefix}{node.type}: {node.text.decode('utf-8')!r}")
    else:
        lines.append(f"{prefix}{node.type}")
        for child in node.children:
            lines.append(print_tree(child, indent + 1))
    return "\n".join(lines)


def main() -> None:
    if len(sys.argv) > 1:
        src_path = Path(sys.argv[1])
        if not src_path.exists():
            print(f"Error: {src_path} not found")
            sys.exit(1)
        code = src_path.read_text(encoding="utf-8")
        print(f"Parsing: {src_path}")
    else:
        code = SAMPLE_CODE
        print("No file provided — using built-in sample code.")

    parser = Parser(PY_LANGUAGE)
    tree = parser.parse(code.encode("utf-8"))

    # Print readable tree to stdout
    readable = print_tree(tree.root_node)
    print("\n--- AST (readable) ---\n")
    print(readable)

    # Save structured JSON
    ast_dict = {
        "source_code": code,
        "ast": node_to_dict(tree.root_node),
    }
    OUTPUT_FILE.parent.mkdir(parents=True, exist_ok=True)
    OUTPUT_FILE.write_text(json.dumps(ast_dict, indent=2, ensure_ascii=False), encoding="utf-8")
    print(f"\nJSON AST saved to: {OUTPUT_FILE}")


if __name__ == "__main__":
    main()
