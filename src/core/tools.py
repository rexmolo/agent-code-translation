"""Evaluation tools for compiling, running, and comparing code.

These are plain functions used by the evaluation agent via Agno's @tool decorator.
"""

import json
import os
import subprocess
import tempfile
from pathlib import Path

from agno.tools import tool


def _go_env(tmpdir: str) -> dict:
    """Return a copy of the environment with GOCACHE and GOMODCACHE inside tmpdir."""
    env = os.environ.copy()
    env["GOCACHE"] = str(Path(tmpdir) / "gocache")
    env["GOMODCACHE"] = str(Path(tmpdir) / "gomodcache")
    return env


@tool()
def compile_go_code(go_code: str) -> str:
    """Compile Go source code and report whether it compiles successfully.

    Args:
        go_code: The complete Go source code to compile.

    Returns:
        JSON string with 'success' (bool) and 'error' (str) fields.
    """
    with tempfile.TemporaryDirectory() as tmpdir:
        go_file = Path(tmpdir) / "main.go"
        go_file.write_text(go_code, encoding="utf-8")
        try:
            result = subprocess.run(
                ["go", "build", "-o", str(Path(tmpdir) / "main"), str(go_file)],
                capture_output=True,
                text=True,
                timeout=30,
                env=_go_env(tmpdir),
            )
            if result.returncode == 0:
                return json.dumps({"success": True, "error": ""})
            return json.dumps({"success": False, "error": result.stderr.strip()})
        except subprocess.TimeoutExpired:
            return json.dumps({"success": False, "error": "Compilation timed out (30s)"})
        except FileNotFoundError:
            return json.dumps({"success": False, "error": "Go compiler not found"})


@tool()
def run_go_code(go_code: str, stdin_input: str = "") -> str:
    """Run Go source code and capture its output.

    Args:
        go_code: The complete Go source code to run.
        stdin_input: Input to pass via stdin.

    Returns:
        JSON string with 'success' (bool), 'stdout' (str), and 'stderr' (str) fields.
    """
    with tempfile.TemporaryDirectory() as tmpdir:
        go_file = Path(tmpdir) / "main.go"
        go_file.write_text(go_code, encoding="utf-8")
        try:
            result = subprocess.run(
                ["go", "run", str(go_file)],
                input=stdin_input,
                capture_output=True,
                text=True,
                timeout=30,
                env=_go_env(tmpdir),
            )
            return json.dumps({
                "success": result.returncode == 0,
                "stdout": result.stdout,
                "stderr": result.stderr.strip(),
            })
        except subprocess.TimeoutExpired:
            return json.dumps({"success": False, "stdout": "", "stderr": "Execution timed out (30s)"})
        except FileNotFoundError:
            return json.dumps({"success": False, "stdout": "", "stderr": "Go runtime not found"})


@tool()
def run_python_code(python_code: str, stdin_input: str = "") -> str:
    """Run Python source code and capture its output.

    Args:
        python_code: The complete Python source code to run.
        stdin_input: Input to pass via stdin.

    Returns:
        JSON string with 'success' (bool), 'stdout' (str), and 'stderr' (str) fields.
    """
    with tempfile.TemporaryDirectory() as tmpdir:
        py_file = Path(tmpdir) / "main.py"
        py_file.write_text(python_code, encoding="utf-8")
        try:
            result = subprocess.run(
                ["python3", str(py_file)],
                input=stdin_input,
                capture_output=True,
                text=True,
                timeout=30,
            )
            return json.dumps({
                "success": result.returncode == 0,
                "stdout": result.stdout,
                "stderr": result.stderr.strip(),
            })
        except subprocess.TimeoutExpired:
            return json.dumps({"success": False, "stdout": "", "stderr": "Execution timed out (30s)"})


@tool()
def compare_outputs(python_stdout: str, go_stdout: str) -> str:
    """Compare the stdout of Python and Go executions for equivalence.

    Args:
        python_stdout: The stdout from running the Python source code.
        go_stdout: The stdout from running the translated Go code.

    Returns:
        JSON string with 'equivalent' (bool) and 'details' (str) fields.
    """
    py_normalized = python_stdout.strip()
    go_normalized = go_stdout.strip()

    if py_normalized == go_normalized:
        return json.dumps({"equivalent": True, "details": "Outputs match exactly"})

    py_lines = py_normalized.splitlines()
    go_lines = go_normalized.splitlines()
    for i, (pl, gl) in enumerate(zip(py_lines, go_lines)):
        if pl != gl:
            return json.dumps({
                "equivalent": False,
                "details": f"First difference at line {i + 1}: Python={pl!r}, Go={gl!r}",
            })

    if len(py_lines) != len(go_lines):
        return json.dumps({
            "equivalent": False,
            "details": f"Different number of lines: Python={len(py_lines)}, Go={len(go_lines)}",
        })

    return json.dumps({"equivalent": False, "details": "Outputs differ"})


@tool()
def run_python_tests(source_code: str, test_code: str) -> str:
    """Run Python tests against Python source code using pytest.

    Args:
        source_code: The Python source code to test.
        test_code: The Python test code (unittest or pytest style).

    Returns:
        JSON string with 'success' (bool), 'total' (int), 'passed' (int), and 'output' (str).
    """
    with tempfile.TemporaryDirectory() as tmpdir:
        src_file = Path(tmpdir) / "source.py"
        src_file.write_text(source_code, encoding="utf-8")
        test_file = Path(tmpdir) / "test_source.py"
        test_file.write_text(test_code, encoding="utf-8")
        try:
            result = subprocess.run(
                ["python3", "-m", "pytest", str(test_file), "-v", "--tb=short"],
                capture_output=True,
                text=True,
                timeout=30,
                cwd=tmpdir,
            )
            output = result.stdout + result.stderr
            passed = output.count(" PASSED")
            failed = output.count(" FAILED")
            errors = output.count(" ERROR")
            total = passed + failed + errors
            return json.dumps({
                "success": result.returncode == 0,
                "total": total,
                "passed": passed,
                "output": output[:500],
            })
        except subprocess.TimeoutExpired:
            return json.dumps({"success": False, "total": 0, "passed": 0, "output": "Test execution timed out (30s)"})


@tool()
def run_go_tests(go_source_code: str, go_test_code: str) -> str:
    """Run Go tests against translated Go source code.

    Args:
        go_source_code: The translated Go source code.
        go_test_code: The Go test code (using testing package).

    Returns:
        JSON string with 'success' (bool), 'total' (int), 'passed' (int), and 'output' (str).
    """
    with tempfile.TemporaryDirectory() as tmpdir:
        subprocess.run(
            ["go", "mod", "init", "testmod"],
            capture_output=True,
            text=True,
            cwd=tmpdir,
            env=_go_env(tmpdir),
        )
        source_file = Path(tmpdir) / "main.go"
        source_file.write_text(go_source_code, encoding="utf-8")
        test_file = Path(tmpdir) / "main_test.go"
        test_file.write_text(go_test_code, encoding="utf-8")
        try:
            result = subprocess.run(
                ["go", "test", "-v", "./..."],
                capture_output=True,
                text=True,
                timeout=30,
                cwd=tmpdir,
                env=_go_env(tmpdir),
            )
            output = result.stdout + result.stderr
            passed = output.count("--- PASS:")
            failed = output.count("--- FAIL:")
            total = passed + failed
            return json.dumps({
                "success": result.returncode == 0,
                "total": total,
                "passed": passed,
                "output": output[:500],
            })
        except subprocess.TimeoutExpired:
            return json.dumps({"success": False, "total": 0, "passed": 0, "output": "Test execution timed out (30s)"})
        except FileNotFoundError:
            return json.dumps({"success": False, "total": 0, "passed": 0, "output": "Go runtime not found"})
