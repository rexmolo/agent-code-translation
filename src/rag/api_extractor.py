"""Extract Python API calls and patterns from source code using tree-sitter.

Extracts:
- Function/method calls (e.g., json.loads, str.split, open)
- Attribute chains (e.g., sys.stdin, os.path)
- API call sequences preserving order (e.g., open -> read -> json.loads)
- Error handling patterns (try/except blocks)
"""

from __future__ import annotations

import tree_sitter_python as tspython
from tree_sitter import Language, Parser

_PY_LANGUAGE = Language(tspython.language())
_PARSER = Parser(_PY_LANGUAGE)


def _collect_calls(node, calls: list[str]) -> None:
    """Walk the AST and collect function/method call names in order."""
    if node.type == "call":
        func = node.child_by_field_name("function")
        if func:
            name = func.text.decode("utf-8")
            calls.append(name)

    for child in node.children:
        _collect_calls(child, calls)


def _has_try_except(node) -> bool:
    """Check if the AST contains any try/except blocks."""
    if node.type == "try_statement":
        return True
    return any(_has_try_except(child) for child in node.children)


def _collect_imports(node, imports: list[str]) -> None:
    """Collect imported module names."""
    if node.type == "import_statement":
        for child in node.children:
            if child.type == "dotted_name":
                imports.append(child.text.decode("utf-8"))
    elif node.type == "import_from_statement":
        module = node.child_by_field_name("module_name")
        if module:
            imports.append(module.text.decode("utf-8"))
        for child in node.children:
            if child.type == "dotted_name" and child != module:
                mod_prefix = module.text.decode("utf-8") + "." if module else ""
                imports.append(mod_prefix + child.text.decode("utf-8"))

    for child in node.children:
        _collect_imports(child, imports)


def extract_api_info(python_code: str) -> dict:
    """Extract structured API information from Python source code.

    Returns:
        {
            "calls": ["open", "f.read", "json.loads", ...],  # ordered call sequence
            "imports": ["json", "sys", ...],
            "has_error_handling": True/False,
            "query_apis": "open read json.loads ...",          # flat string for retrieval
            "query_imports": "json sys ...",
        }
    """
    tree = _PARSER.parse(python_code.encode("utf-8"))
    root = tree.root_node

    calls: list[str] = []
    _collect_calls(root, calls)

    imports: list[str] = []
    _collect_imports(root, imports)

    has_error_handling = _has_try_except(root)

    # Deduplicate while preserving order for the query string
    seen = set()
    unique_calls = []
    for c in calls:
        if c not in seen:
            seen.add(c)
            unique_calls.append(c)

    return {
        "calls": unique_calls,
        "imports": imports,
        "has_error_handling": has_error_handling,
        "query_apis": " ".join(unique_calls),
        "query_imports": " ".join(imports),
    }
