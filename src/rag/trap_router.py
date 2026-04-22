"""Deterministic routing for the frozen CodeNet translation-trap taxonomy."""

from __future__ import annotations

import ast
import json
import re
from functools import lru_cache
from typing import Any

from src.config import TRANSLATION_TRAPS_FILE

_SLICE_COPY_RE = re.compile(r"\[[^\]]*:\s*[^\]]*\]")
_GO_SLICE_RETURN_RE = re.compile(r"\)\s*\[\][A-Za-z0-9_*\[\]]+")
_TYPE_ANNOTATION_MODULES = frozenset({
    "typing",
    "typing_extensions",
    "collections.abc",
    "abc",
})


def _load_jsonl(path) -> list[dict]:
    records: list[dict] = []
    with open(path, encoding="utf-8") as fh:
        for line in fh:
            line = line.strip()
            if line:
                records.append(json.loads(line))
    return records


@lru_cache(maxsize=1)
def load_translation_traps() -> list[dict]:
    return _load_jsonl(TRANSLATION_TRAPS_FILE)


def _is_negative_int_node(node: ast.AST | None) -> bool:
    if node is None:
        return False
    if isinstance(node, ast.Constant) and isinstance(node.value, int):
        return node.value < 0
    if (
        isinstance(node, ast.UnaryOp)
        and isinstance(node.op, ast.USub)
        and isinstance(node.operand, ast.Constant)
        and isinstance(node.operand.value, int)
    ):
        return True
    return False


def _is_full_slice_copy(slice_node: ast.AST | None) -> bool:
    return (
        isinstance(slice_node, ast.Slice)
        and slice_node.lower is None
        and slice_node.upper is None
        and slice_node.step is None
    )


def _call_name(node: ast.AST) -> str | None:
    if isinstance(node, ast.Name):
        return node.id
    if isinstance(node, ast.Attribute):
        root = _call_name(node.value)
        return f"{root}.{node.attr}" if root else node.attr
    return None


class _TrapSignalExtractor(ast.NodeVisitor):
    def __init__(self) -> None:
        self.ast_markers: set[str] = set()
        self.python_markers: set[str] = set()
        self.inferred_go_api_markers: set[str] = set()
        self._copied_slice_names: set[str] = set()
        self._dict_names: set[str] = set()

    def visit_Assign(self, node: ast.Assign) -> None:
        value = node.value
        targets = [target for target in node.targets if isinstance(target, ast.Name)]
        if isinstance(value, ast.Subscript) and _is_full_slice_copy(value.slice):
            for target in targets:
                self._copied_slice_names.add(target.id)

        if isinstance(value, (ast.Dict,)) or (
            isinstance(value, ast.Call) and _call_name(value.func) == "dict"
        ):
            for target in targets:
                self._dict_names.add(target.id)
        self.generic_visit(node)

    def visit_AugAssign(self, node: ast.AugAssign) -> None:
        if isinstance(node.target, ast.Subscript) and isinstance(node.target.value, ast.Name):
            if node.target.value.id in self._copied_slice_names:
                self.ast_markers.add("slice_copy_then_mutation")
        self.generic_visit(node)

    def visit_For(self, node: ast.For) -> None:
        if isinstance(node.iter, ast.Name) and node.iter.id in self._dict_names:
            self.ast_markers.add("dict_order_sensitive_usage")
        self.generic_visit(node)

    def visit_Return(self, node: ast.Return) -> None:
        if isinstance(node.value, ast.List) and not node.value.elts:
            self.ast_markers.add("empty_list_return")
            self.python_markers.add("return []")
        self.generic_visit(node)

    def visit_Constant(self, node: ast.Constant) -> None:
        if isinstance(node.value, float):
            self.ast_markers.add("float_literal_present")
        self.generic_visit(node)

    def visit_Subscript(self, node: ast.Subscript) -> None:
        if _is_negative_int_node(node.slice):
            self.ast_markers.add("negative_subscript")
        self.generic_visit(node)

    def visit_BinOp(self, node: ast.BinOp) -> None:
        if isinstance(node.op, ast.FloorDiv):
            self.ast_markers.add("floor_division_operator")
            self.python_markers.add("//")
        self.generic_visit(node)

    def visit_Call(self, node: ast.Call) -> None:
        call_name = _call_name(node.func) or ""
        if call_name == "sorted" and any(keyword.arg == "key" for keyword in node.keywords):
            self.ast_markers.add("sorted_call_with_key")
        if call_name.endswith(".sort") and any(keyword.arg == "key" for keyword in node.keywords):
            self.ast_markers.add("sort_call_with_key")
        if call_name in {"dict", "dict.items", "dict.keys"} or call_name.endswith(".items") or call_name.endswith(".keys"):
            self.ast_markers.add("dict_order_sensitive_usage")
        if call_name == "int":
            self.python_markers.add("int(")
            self.inferred_go_api_markers.add("strconv.Atoi")
        if call_name == "str":
            self.python_markers.add("str(")
            self.inferred_go_api_markers.update({"strconv.Itoa", "fmt.Sprintf"})
        if call_name == "print":
            self.python_markers.add("print(")
            self.inferred_go_api_markers.add("fmt.Println")
            if len(node.args) > 1:
                self.ast_markers.add("print_multiple_args")
                self.inferred_go_api_markers.add("fmt.Printf")
        if call_name.endswith(".append") and isinstance(node.func, ast.Attribute) and isinstance(node.func.value, ast.Name):
            if node.func.value.id in self._copied_slice_names:
                self.ast_markers.add("slice_copy_then_mutation")
        self.generic_visit(node)

    def visit_JoinedStr(self, node: ast.JoinedStr) -> None:
        self.python_markers.add("f\"")
        self.inferred_go_api_markers.add("fmt.Sprintf")
        self.generic_visit(node)

    def visit_Import(self, node: ast.Import) -> None:
        for alias in node.names:
            if alias.name.split(".")[0] not in _TYPE_ANNOTATION_MODULES:
                self.python_markers.add("import ")
        self.generic_visit(node)

    def visit_ImportFrom(self, node: ast.ImportFrom) -> None:
        module = node.module or ""
        if module.split(".")[0] not in _TYPE_ANNOTATION_MODULES:
            self.python_markers.add("import ")
        self.generic_visit(node)


def _top_level_ast_markers(python_code: str) -> set[str]:
    try:
        root = ast.parse(python_code)
    except SyntaxError:
        return set()

    markers: set[str] = set()
    declaration_nodes = (
        ast.FunctionDef,
        ast.AsyncFunctionDef,
        ast.ClassDef,
        ast.Import,
        ast.ImportFrom,
    )

    for node in root.body:
        if isinstance(node, ast.Expr):
            markers.add("module_level_expression")
            markers.add("module_level_statement")
        elif not isinstance(node, declaration_nodes):
            markers.add("module_level_statement")

    function_defs = sum(
        1 for node in root.body if isinstance(node, (ast.FunctionDef, ast.AsyncFunctionDef))
    )
    if function_defs > 1:
        markers.add("multiple_top_level_functions")

    return markers


def _extract_ast_signals(python_code: str) -> tuple[set[str], set[str], set[str]]:
    try:
        root = ast.parse(python_code)
        extractor = _TrapSignalExtractor()
        extractor.visit(root)
        return extractor.ast_markers, extractor.python_markers, extractor.inferred_go_api_markers
    except SyntaxError:
        return set(), set(), set()


def _python_markers(
    python_code: str,
    traps: list[dict],
    *,
    inferred_markers: set[str],
) -> set[str]:
    seen: set[str] = set(inferred_markers)

    for trap in traps:
        hints = trap.get("routing_hints", {})
        for marker in hints.get("python_markers", []):
            if marker == "import ":
                continue
            if marker and marker in python_code:
                seen.add(marker)

    if _SLICE_COPY_RE.search(python_code):
        seen.add("[:]")
    return seen


def _signature_markers(go_signature: str | None) -> set[str]:
    markers: set[str] = set()
    if not go_signature:
        return markers

    if _GO_SLICE_RETURN_RE.search(go_signature):
        markers.add("signature_returns_slice")
    for package_name in ("strings", "sort", "strconv", "math", "fmt", "unicode"):
        if re.search(rf"\b{package_name}\s+[A-Za-z0-9_\[\]*]+", go_signature):
            markers.add(f"signature_param_name_collision={package_name}")
    return markers


def _go_api_markers(python_code: str, *, inferred_markers: set[str]) -> set[str]:
    markers = set(inferred_markers)
    if "input()" in python_code or "sys.stdin" in python_code:
        markers.update({"fmt.Fscan", "fmt.Scan", "bufio.NewReader", "bufio.Scanner"})
    if "input().split()" in python_code or "map(int, input()" in python_code:
        markers.add("bufio.ScanWords")
    if "str(" in python_code:
        markers.update({"strconv.Itoa", "fmt.Sprintf"})
    if "print(" in python_code:
        markers.update({"fmt.Printf", "fmt.Println"})
    if re.search(r"\b\d+\.\d+\b", python_code):
        markers.add("float64(")
    return markers


def extract_trap_signals(python_code: str, *, go_signature: str | None = None) -> dict[str, list[str]]:
    """Extract deterministic, normalized routing signals for trap lookup."""
    traps = load_translation_traps()
    extracted_ast_markers, inferred_python_markers, inferred_go_api_markers = _extract_ast_signals(
        python_code
    )
    ast_markers = sorted(_top_level_ast_markers(python_code) | extracted_ast_markers)
    python_markers = sorted(
        _python_markers(
            python_code,
            traps,
            inferred_markers=inferred_python_markers,
        )
    )
    signature_markers = sorted(_signature_markers(go_signature))
    go_api_markers = sorted(_go_api_markers(python_code, inferred_markers=inferred_go_api_markers))
    return {
        "ast_markers": ast_markers,
        "python_markers": python_markers,
        "signature_markers": signature_markers,
        "go_api_markers": go_api_markers,
    }


def _match_trap(trap: dict[str, Any], signals: dict[str, list[str]]) -> dict[str, Any] | None:
    hints = trap.get("routing_hints", {})
    matched_ast = [m for m in hints.get("ast_markers", []) if m in signals["ast_markers"]]
    matched_python = [m for m in hints.get("python_markers", []) if m in signals["python_markers"]]
    matched_signature = [m for m in hints.get("signature_markers", []) if m in signals["signature_markers"]]
    matched_go_api = [m for m in hints.get("go_api_markers", []) if m in signals["go_api_markers"]]

    if hints.get("ast_markers") and not matched_ast:
        return None
    if hints.get("python_markers") and not matched_python:
        return None
    if hints.get("signature_markers") and not matched_signature:
        return None
    if hints.get("go_api_markers") and not matched_go_api:
        return None

    score = (
        (len(matched_ast) * 100)
        + (len(matched_python) * 10)
        + (len(matched_signature) * 5)
        + (len(matched_go_api) * 3)
    )
    if score <= 0:
        return None

    record = dict(trap)
    record["retrieval"] = {
        "backend": "none",
        "mode": "deterministic_trap_routing",
        "doc_id": trap["trap_id"],
        "merged_rank": None,
        "rrf_score": None,
        "dense_rank": None,
        "dense_distance": None,
        "sparse_rank": None,
        "sparse_score": None,
        "routing_score": score,
        "matched_signals": {
            "ast_markers": matched_ast,
            "python_markers": matched_python,
            "signature_markers": matched_signature,
            "go_api_markers": matched_go_api,
        },
    }
    return record


def route_translation_traps(
    python_code: str,
    *,
    go_signature: str | None = None,
    limit: int = 2,
) -> tuple[list[dict], dict]:
    """Return matched trap records plus a normalized source trace."""
    traps = load_translation_traps()
    signals = extract_trap_signals(python_code, go_signature=go_signature)

    candidates = []
    for trap in traps:
        matched = _match_trap(trap, signals)
        if matched is not None:
            candidates.append(matched)

    candidates.sort(
        key=lambda item: (
            -int(item["retrieval"]["routing_score"]),
            item["trap_id"],
        )
    )

    accepted = []
    for rank, item in enumerate(candidates[:limit], start=1):
        item["retrieval"]["merged_rank"] = rank
        item["retrieval"]["accepted"] = True
        accepted.append(item)

    query_parts = []
    for key in ("ast_markers", "python_markers", "signature_markers", "go_api_markers"):
        values = signals[key]
        if values:
            query_parts.append(f"{key}=" + ",".join(values))

    trace_items = []
    for item in candidates:
        trace_items.append(
            {
                "trap_id": item["trap_id"],
                "title": item.get("title", ""),
                "category": item.get("category", ""),
                "retrieval": item.get("retrieval", {}),
            }
        )

    trace = {
        "source": "translation_traps",
        "queried": True,
        "accepted": bool(accepted),
        "query_text": " | ".join(query_parts),
        "returned_count": len(candidates),
        "accepted_count": len(accepted),
        "items": trace_items,
        "acceptance": {
            "enabled": True,
            "rules": {"max_accepted_items": limit, "mode": "deterministic_trap_routing"},
            "fallback_to_no_retrieval": False,
        },
    }
    return accepted, trace
