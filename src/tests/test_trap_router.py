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


def test_route_translation_traps_matches_top_level_program_traps():
    source = (
        "import sys\n"
        "n = int(input())\n"
        "print(n)\n"
    )

    traps, trace = route_translation_traps(source, limit=3)
    trap_ids = {trap["trap_id"] for trap in traps}

    assert "trap_top_level_statements_outside_main" in trap_ids
    assert "trap_output_formatting_matches_python" in trap_ids
    assert trace["queried"] is True
    assert trace["accepted"] is True
