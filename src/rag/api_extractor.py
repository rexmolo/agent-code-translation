"""Extract Python API calls and Go API sequences using tree-sitter.

Python-side extraction is used to build focused retrieval queries for:
- API mappings
- API-sequence retrieval

Go-side extraction is used by the API-sequence corpus builder to produce
ordered target-language API usage sequences.
"""

from __future__ import annotations

from pathlib import Path

_MAX_SEQUENCE_FUNCTION_CODE = 2000
_PYTHON_PARSER = None
_GO_PARSER = None

# Modules that are pure type annotations — no Go API equivalent exists.
_TYPE_ANNOTATION_MODULES = frozenset({
    "typing",
    "typing_extensions",
    "collections.abc",
    "abc",
})


def _load_python_parser():
    global _PYTHON_PARSER
    if _PYTHON_PARSER is not None:
        return _PYTHON_PARSER

    from tree_sitter import Language, Parser
    import tree_sitter_python as tspython

    _PYTHON_PARSER = Parser(Language(tspython.language()))
    return _PYTHON_PARSER


def _load_go_parser():
    global _GO_PARSER
    if _GO_PARSER is not None:
        return _GO_PARSER

    from tree_sitter import Language, Parser
    try:
        import tree_sitter_go as tsgo
    except ModuleNotFoundError as exc:
        raise ModuleNotFoundError(
            "tree_sitter_go is required for Go API sequence extraction. "
            "Add tree-sitter-go to the environment before building the corpus."
        ) from exc

    _GO_PARSER = Parser(Language(tsgo.language()))
    return _GO_PARSER


def _node_text(node) -> str:
    return node.text.decode("utf-8")


def _collect_calls(node, calls: list[str]) -> None:
    """Walk the AST and collect function/method call names in order."""
    if node.type == "call":
        func = node.child_by_field_name("function")
        if func:
            calls.append(_node_text(func))

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
                imports.append(_node_text(child))
    elif node.type == "import_from_statement":
        module = node.child_by_field_name("module_name")
        if module:
            imports.append(_node_text(module))
        for child in node.children:
            if child.type == "dotted_name" and child != module:
                mod_prefix = _node_text(module) + "." if module else ""
                imports.append(mod_prefix + _node_text(child))

    for child in node.children:
        _collect_imports(child, imports)


def extract_api_info(python_code: str) -> dict:
    """Extract structured API information from Python source code."""
    parser = _load_python_parser()
    tree = parser.parse(python_code.encode("utf-8"))
    root = tree.root_node

    calls: list[str] = []
    _collect_calls(root, calls)

    imports: list[str] = []
    _collect_imports(root, imports)

    has_error_handling = _has_try_except(root)

    seen = set()
    unique_calls = []
    for call in calls:
        if call not in seen:
            seen.add(call)
            unique_calls.append(call)

    query_imports = [
        imported for imported in imports
        if imported.split(".")[0] not in _TYPE_ANNOTATION_MODULES
    ]

    return {
        "calls": unique_calls,
        "imports": imports,
        "has_error_handling": has_error_handling,
        "query_apis": " ".join(unique_calls),
        "query_imports": " ".join(query_imports),
    }


def _collect_go_imports(root) -> list[str]:
    imports: list[str] = []

    def walk(node) -> None:
        if node.type == "import_spec":
            path_node = node.child_by_field_name("path")
            if path_node:
                imports.append(_node_text(path_node).strip('"'))
        for child in node.children:
            walk(child)

    walk(root)
    return imports


def _selector_to_api_name(node, imports: set[str]) -> str | None:
    if node.type != "selector_expression":
        return None

    operand = node.child_by_field_name("operand")
    field = node.child_by_field_name("field")
    if operand is None or field is None:
        return None

    operand_text = _node_text(operand)
    field_text = _node_text(field)
    if operand_text in imports:
        return f"{operand_text}.{field_text}"
    return None


def _call_node_to_api_name(call_node, imports: set[str]) -> str | None:
    function = call_node.child_by_field_name("function")
    if function is None:
        return None

    if function.type == "selector_expression":
        return _selector_to_api_name(function, imports)
    return None


def _collect_go_call_sequence(node, imports: set[str], apis: list[str]) -> None:
    for child in node.children:
        _collect_go_call_sequence(child, imports, apis)

    if node.type == "call_expression":
        api_name = _call_node_to_api_name(node, imports)
        if api_name:
            apis.append(api_name)


def _truncate_function_code(text: str, *, limit: int = _MAX_SEQUENCE_FUNCTION_CODE) -> str:
    if len(text) <= limit:
        return text
    return text[:limit].rstrip() + "\n..."


def _iter_function_nodes(root):
    for child in root.children:
        if child.type == "function_declaration":
            yield child


def extract_go_api_sequences(go_code: str, *, file_path: str = "") -> list[dict]:
    """Extract ordered package-qualified API sequences from Go functions."""
    parser = _load_go_parser()
    tree = parser.parse(go_code.encode("utf-8"))
    root = tree.root_node
    imports = set(_collect_go_imports(root))

    records: list[dict] = []
    seen_sequences: set[str] = set()
    for function_node in _iter_function_nodes(root):
        name_node = function_node.child_by_field_name("name")
        function_name = _node_text(name_node) if name_node else ""
        apis: list[str] = []
        _collect_go_call_sequence(function_node, imports, apis)
        if len(apis) < 2:
            continue

        sequence_text = " -> ".join(apis)
        if sequence_text in seen_sequences:
            continue
        seen_sequences.add(sequence_text)

        records.append(
            {
                "file_path": file_path,
                "function_name": function_name,
                "sequence_text": sequence_text,
                "apis": apis,
                "imports": sorted(imports),
                "function_code": _truncate_function_code(_node_text(function_node)),
            }
        )

    return records


def extract_go_api_sequences_from_file(path: str | Path) -> list[dict]:
    file_path = Path(path)
    go_code = file_path.read_text(encoding="utf-8", errors="replace")
    return extract_go_api_sequences(go_code, file_path=str(file_path))
