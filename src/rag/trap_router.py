"""Deterministic routing for the frozen CodeNet translation-trap taxonomy."""

from __future__ import annotations

import ast
import json
import re
from functools import lru_cache
from typing import Any

from src.config import TRANSLATION_TRAPS_CODENET_V1_FILE

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
    return _load_jsonl(TRANSLATION_TRAPS_CODENET_V1_FILE)


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

    return markers


def _python_markers(python_code: str, traps: list[dict]) -> set[str]:
    seen: set[str] = set()
    real_import_present = False
    try:
        root = ast.parse(python_code)
    except SyntaxError:
        root = None

    if root is not None:
        for node in ast.walk(root):
            if isinstance(node, ast.Import):
                for alias in node.names:
                    if alias.name.split(".")[0] not in _TYPE_ANNOTATION_MODULES:
                        real_import_present = True
            elif isinstance(node, ast.ImportFrom):
                module = node.module or ""
                if module.split(".")[0] not in _TYPE_ANNOTATION_MODULES:
                    real_import_present = True

    for trap in traps:
        hints = trap.get("routing_hints", {})
        for marker in hints.get("python_markers", []):
            if marker == "import ":
                if real_import_present:
                    seen.add(marker)
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


def extract_trap_signals(python_code: str, *, go_signature: str | None = None) -> dict[str, list[str]]:
    """Extract deterministic, normalized routing signals for trap lookup."""
    traps = load_translation_traps()
    ast_markers = sorted(_top_level_ast_markers(python_code))
    python_markers = sorted(_python_markers(python_code, traps))
    signature_markers = sorted(_signature_markers(go_signature))
    return {
        "ast_markers": ast_markers,
        "python_markers": python_markers,
        "signature_markers": signature_markers,
    }


def _match_trap(trap: dict[str, Any], signals: dict[str, list[str]]) -> dict[str, Any] | None:
    hints = trap.get("routing_hints", {})
    matched_ast = [m for m in hints.get("ast_markers", []) if m in signals["ast_markers"]]
    matched_python = [m for m in hints.get("python_markers", []) if m in signals["python_markers"]]
    matched_signature = [m for m in hints.get("signature_markers", []) if m in signals["signature_markers"]]

    if hints.get("ast_markers") and not matched_ast:
        return None
    if hints.get("signature_markers") and not matched_signature:
        return None

    score = (len(matched_ast) * 100) + (len(matched_python) * 10) + len(matched_signature)
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
    for key in ("ast_markers", "python_markers", "signature_markers"):
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
