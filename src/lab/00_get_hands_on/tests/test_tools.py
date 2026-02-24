"""Tests for evaluation tools (compile, run, compare).

These tests are offline — no LLM calls, just subprocess execution.
We call .entrypoint on each tool because Agno's @tool() decorator wraps
functions into Function objects that aren't directly callable.

    uv run pytest src/lab/00_get_hands_on/tests/test_tools.py -v
"""

import importlib
import json

import pytest

_tools_mod = importlib.import_module("src.lab.00_get_hands_on.tools")

# Unwrap Agno @tool() wrappers to get the raw callables
compile_go_code = _tools_mod.compile_go_code.entrypoint
run_go_code = _tools_mod.run_go_code.entrypoint
run_python_code = _tools_mod.run_python_code.entrypoint
compare_outputs = _tools_mod.compare_outputs.entrypoint
run_python_tests = _tools_mod.run_python_tests.entrypoint
run_go_tests = _tools_mod.run_go_tests.entrypoint


class TestCompileGoCode:
    def test_valid_go_compiles(self):
        code = 'package main\nimport "fmt"\nfunc main() { fmt.Println("hi") }\n'
        result = json.loads(compile_go_code(code))
        assert result["success"] is True
        assert result["error"] == ""

    def test_invalid_go_fails(self):
        code = "this is not go code"
        result = json.loads(compile_go_code(code))
        assert result["success"] is False
        assert len(result["error"]) > 0


class TestRunGoCode:
    def test_hello_world(self):
        code = 'package main\nimport "fmt"\nfunc main() { fmt.Println("hello") }\n'
        result = json.loads(run_go_code(code))
        assert result["success"] is True
        assert "hello" in result["stdout"]

    def test_with_stdin(self):
        code = (
            'package main\nimport ("bufio"; "fmt"; "os")\n'
            "func main() {\n"
            "  scanner := bufio.NewScanner(os.Stdin)\n"
            "  scanner.Scan()\n"
            '  fmt.Println("got:", scanner.Text())\n'
            "}\n"
        )
        result = json.loads(run_go_code(code, stdin_input="test input"))
        assert result["success"] is True
        assert "got: test input" in result["stdout"]

    def test_runtime_error(self):
        code = (
            "package main\n"
            "func main() {\n"
            "  var s []int\n"
            "  _ = s[100]\n"
            "}\n"
        )
        result = json.loads(run_go_code(code))
        assert result["success"] is False


class TestRunPythonCode:
    def test_hello_world(self):
        code = "print('hello')"
        result = json.loads(run_python_code(code))
        assert result["success"] is True
        assert "hello" in result["stdout"]

    def test_syntax_error(self):
        code = "def ("
        result = json.loads(run_python_code(code))
        assert result["success"] is False


class TestCompareOutputs:
    def test_identical_outputs(self):
        result = json.loads(compare_outputs("hello\nworld\n", "hello\nworld\n"))
        assert result["equivalent"] is True

    def test_identical_after_strip(self):
        result = json.loads(compare_outputs("hello\n", "  hello  \n"))
        assert result["equivalent"] is True

    def test_different_outputs(self):
        result = json.loads(compare_outputs("hello", "goodbye"))
        assert result["equivalent"] is False
        assert "line 1" in result["details"]

    def test_different_line_count(self):
        result = json.loads(compare_outputs("a\nb\n", "a\n"))
        assert result["equivalent"] is False


class TestRunGoTests:
    def test_passing_go_tests(self):
        source = (
            "package testmod\n\n"
            "func Add(a, b int) int { return a + b }\n"
        )
        tests = (
            "package testmod\n\n"
            'import "testing"\n\n'
            "func TestAdd(t *testing.T) {\n"
            "  if Add(1, 2) != 3 {\n"
            '    t.Errorf("expected 3")\n'
            "  }\n"
            "}\n"
        )
        result = json.loads(run_go_tests(source, tests))
        assert result["success"] is True
        assert result["passed"] >= 1

    def test_failing_go_tests(self):
        source = (
            "package testmod\n\n"
            "func Add(a, b int) int { return a - b }\n"
        )
        tests = (
            "package testmod\n\n"
            'import "testing"\n\n'
            "func TestAdd(t *testing.T) {\n"
            "  if Add(1, 2) != 3 {\n"
            '    t.Errorf("expected 3, got %d", Add(1, 2))\n'
            "  }\n"
            "}\n"
        )
        result = json.loads(run_go_tests(source, tests))
        assert result["success"] is False


class TestRunPythonTests:
    def test_passing_python_tests(self):
        source = "def add(a, b): return a + b\n"
        tests = "from source import add\n\ndef test_add():\n    assert add(1, 2) == 3\n"
        result = json.loads(run_python_tests(source, tests))
        assert result["success"] is True
        assert result["passed"] >= 1

    def test_failing_python_tests(self):
        source = "def add(a, b): return a - b\n"
        tests = "from source import add\n\ndef test_add():\n    assert add(1, 2) == 3\n"
        result = json.loads(run_python_tests(source, tests))
        assert result["success"] is False
