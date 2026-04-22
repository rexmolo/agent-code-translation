from src.rag.trap_router import extract_trap_signals, route_translation_traps


def test_extract_trap_signals_detects_module_level_and_stdio_markers():
    source = (
        "import sys\n"
        "n = int(input())\n"
        "print(n)\n"
    )

    signals = extract_trap_signals(source)

    assert "module_level_statement" in signals["ast_markers"]
    assert "module_level_expression" in signals["ast_markers"]
    assert "input()" in signals["python_markers"]
    assert "print(" in signals["python_markers"]


def test_typing_only_import_does_not_trigger_unused_import_trap():
    source = (
        "from typing import List\n\n"
        "def f(xs: List[int]) -> int:\n"
        "    return len(xs)\n"
    )

    signals = extract_trap_signals(source, go_signature="func F(xs []int) int")
    traps, _ = route_translation_traps(source, go_signature="func F(xs []int) int", limit=5)
    trap_ids = {trap["trap_id"] for trap in traps}

    assert "import " not in signals["python_markers"]
    assert "trap_unused_imports_are_compile_errors" not in trap_ids


def test_route_translation_traps_matches_top_level_program_traps():
    source = (
        "import sys\n"
        "n = int(input())\n"
        "print(n)\n"
    )

    traps, trace = route_translation_traps(source, limit=3)
    trap_ids = {trap["trap_id"] for trap in traps}

    assert "trap_top_level_statements_outside_main" in trap_ids
    assert trace["queried"] is True
    assert trace["accepted"] is True


def test_route_translation_traps_distinguishes_last_from_second_last_index():
    source = (
        "def f(xs):\n"
        "    return xs[-2]\n"
    )

    traps, _ = route_translation_traps(source, limit=5)
    trap_ids = {trap["trap_id"] for trap in traps}

    assert "trap_negative_index_rewrite_second_last" in trap_ids
    assert "trap_negative_index_rewrite_last" not in trap_ids


def test_route_translation_traps_matches_signature_shadowing():
    source = (
        "def concatenate(items):\n"
        "    return ''.join(items)\n"
    )
    signature = "func Concatenate(strings []string) string"

    traps, _ = route_translation_traps(source, go_signature=signature, limit=5)
    trap_ids = {trap["trap_id"] for trap in traps}

    assert "trap_signature_param_shadows_strings_package" in trap_ids


def test_extract_trap_signals_detects_multiple_helpers_and_multi_arg_print():
    source = (
        "def helper(x):\n"
        "    return x + 1\n\n"
        "def solve(x):\n"
        "    print(x, helper(x))\n"
        "    return helper(x)\n"
    )

    signals = extract_trap_signals(source)

    assert "multiple_top_level_functions" in signals["ast_markers"]
    assert "print_multiple_args" in signals["ast_markers"]


def test_route_translation_traps_matches_sorted_key_and_floor_division():
    source = (
        "def f(xs):\n"
        "    ys = sorted(xs, key=lambda x: x % 3)\n"
        "    return ys[0] // 2\n"
    )

    traps, _ = route_translation_traps(source, limit=10)
    trap_ids = {trap["trap_id"] for trap in traps}

    assert "trap_sorted_with_key_needs_stable_sort" in trap_ids
    assert "trap_floor_division_negative_operands" in trap_ids


def test_route_translation_traps_matches_math_gcd_source_pattern():
    source = (
        "import math\n\n"
        "def f(a, b):\n"
        "    return math.gcd(a, b)\n"
    )

    traps, _ = route_translation_traps(source, limit=10)
    trap_ids = {trap["trap_id"] for trap in traps}

    assert "trap_fake_math_gcd_api" in trap_ids
