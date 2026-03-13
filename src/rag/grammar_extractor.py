"""Extract Python grammar patterns from source code using tree-sitter.

Detects 14 grammar categories that map directly to tree-sitter node types,
extracts a representative code fragment for each, and deduplicates by category.

Used by the RAG retriever to issue focused dense search queries per category
instead of passing raw source code as a single query.
"""

from __future__ import annotations

import tree_sitter_python as tspython
from tree_sitter import Language, Parser

_PY_LANGUAGE = Language(tspython.language())
_PARSER = Parser(_PY_LANGUAGE)

_MAX_FRAGMENT_LEN = 500

# Direct node type -> category mappings
_NODE_TYPE_TO_CATEGORY: dict[str, str] = {
    "try_statement": "Error Handling",
    "list_comprehension": "List Comprehensions",
    "class_definition": "Classes and Constructors",
    "conditional_expression": "Ternary Operator",
    "with_statement": "Context Managers (with)",
    "lambda": "Lambda Functions",
    "assert_statement": "Assertions",
    "dictionary_comprehension": "Dictionary Comprehensions",
    "global_statement": "Global Variables",
    "pass_statement": "Pass Statement",
    "default_parameter": "Default Function Arguments",
}

# Categories that need special detection logic (not in the dict above)
_CAT_DECORATORS = "Decorators"
_CAT_GENERATORS = "Generator Functions (yield)"
_CAT_MULTIPLE_INHERITANCE = "Multiple Inheritance"


def _find_ancestor(node, target_type: str):
    """Walk up the parent chain to find a node of the given type."""
    current = node.parent
    while current is not None:
        if current.type == target_type:
            return current
        current = current.parent
    return None


def _has_multiple_bases(class_node) -> bool:
    """Check if a class_definition has more than one base class."""
    for child in class_node.children:
        if child.type == "argument_list":
            bases = [c for c in child.children if c.type not in ("(", ")", ",")]
            return len(bases) > 1
    return False


def _truncate(text: str) -> str:
    if len(text) <= _MAX_FRAGMENT_LEN:
        return text
    return text[:_MAX_FRAGMENT_LEN] + "..."


def _walk(node, found: dict[str, str]) -> None:
    """Populate found with {category: code_fragment}, first occurrence wins."""
    ntype = node.type

    # --- Direct node type mappings ---
    if ntype in _NODE_TYPE_TO_CATEGORY:
        category = _NODE_TYPE_TO_CATEGORY[ntype]

        # Special: class_definition may also be Multiple Inheritance
        if ntype == "class_definition":
            if _has_multiple_bases(node) and _CAT_MULTIPLE_INHERITANCE not in found:
                found[_CAT_MULTIPLE_INHERITANCE] = _truncate(
                    node.text.decode("utf-8")
                )
            # Still record "Classes and Constructors" if not yet seen
            if category not in found:
                found[category] = _truncate(node.text.decode("utf-8"))

        # Special: default_parameter -> use enclosing function as fragment
        elif ntype == "default_parameter":
            if category not in found:
                func_node = _find_ancestor(node, "function_definition")
                if func_node:
                    found[category] = _truncate(func_node.text.decode("utf-8"))
                else:
                    found[category] = _truncate(node.text.decode("utf-8"))

        # General case
        elif category not in found:
            found[category] = _truncate(node.text.decode("utf-8"))

    # --- Special: decorator -> use parent decorated_definition as fragment ---
    if ntype == "decorator" and _CAT_DECORATORS not in found:
        parent = node.parent
        if parent and parent.type == "decorated_definition":
            found[_CAT_DECORATORS] = _truncate(parent.text.decode("utf-8"))
        else:
            found[_CAT_DECORATORS] = _truncate(node.text.decode("utf-8"))

    # --- Special: yield inside function -> Generator Functions ---
    if ntype == "yield" and _CAT_GENERATORS not in found:
        func_node = _find_ancestor(node, "function_definition")
        if func_node:
            found[_CAT_GENERATORS] = _truncate(func_node.text.decode("utf-8"))

    for child in node.children:
        _walk(child, found)


def extract_grammar_patterns(python_code: str) -> list[dict]:
    """Extract grammar patterns from Python source code.

    Returns:
        List of {"category": str, "fragment": str} dicts, deduplicated by category.
        At most 14 entries (one per supported grammar category).
    """
    tree = _PARSER.parse(python_code.encode("utf-8"))
    found: dict[str, str] = {}
    _walk(tree.root_node, found)
    return [{"category": cat, "fragment": frag} for cat, frag in found.items()]
