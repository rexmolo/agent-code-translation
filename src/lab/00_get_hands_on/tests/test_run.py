"""Tests for run.py helper functions (no LLM calls).

    uv run pytest src/lab/00_get_hands_on/tests/test_run.py -v
"""

import importlib
import tempfile
from pathlib import Path

import pytest

_run = importlib.import_module("src.lab.00_get_hands_on.run")
_models = importlib.import_module("src.lab.00_get_hands_on.models")

EvaluationRecord = _models.EvaluationRecord


class TestDiscoverPythonFiles:
    def test_finds_py_files(self, tmp_path):
        (tmp_path / "a.py").write_text("pass")
        (tmp_path / "b.py").write_text("pass")
        (tmp_path / "c.txt").write_text("not python")
        files = _run.discover_python_files(tmp_path)
        assert len(files) == 2
        assert all(f.suffix == ".py" for f in files)

    def test_excludes_test_files(self, tmp_path):
        (tmp_path / "app.py").write_text("pass")
        (tmp_path / "test_app.py").write_text("pass")
        (tmp_path / "app_test.py").write_text("pass")
        files = _run.discover_python_files(tmp_path)
        assert len(files) == 1
        assert files[0].name == "app.py"

    def test_recursive(self, tmp_path):
        sub = tmp_path / "pkg"
        sub.mkdir()
        (tmp_path / "a.py").write_text("pass")
        (sub / "b.py").write_text("pass")
        files = _run.discover_python_files(tmp_path)
        assert len(files) == 2

    def test_empty_dir(self, tmp_path):
        files = _run.discover_python_files(tmp_path)
        assert files == []


class TestFindTestFile:
    def test_finds_test_prefix(self, tmp_path):
        src = tmp_path / "app.py"
        src.write_text("pass")
        test = tmp_path / "test_app.py"
        test.write_text("pass")
        assert _run.find_test_file(src) == test

    def test_finds_test_suffix(self, tmp_path):
        src = tmp_path / "app.py"
        src.write_text("pass")
        test = tmp_path / "app_test.py"
        test.write_text("pass")
        assert _run.find_test_file(src) == test

    def test_finds_in_tests_subdir(self, tmp_path):
        src = tmp_path / "app.py"
        src.write_text("pass")
        tests_dir = tmp_path / "tests"
        tests_dir.mkdir()
        test = tests_dir / "test_app.py"
        test.write_text("pass")
        assert _run.find_test_file(src) == test

    def test_returns_none_when_no_tests(self, tmp_path):
        src = tmp_path / "app.py"
        src.write_text("pass")
        assert _run.find_test_file(src) is None


class TestMirrorPath:
    def test_basic_mirror(self, tmp_path):
        source_root = tmp_path / "src"
        target_root = tmp_path / "target"
        source_file = source_root / "app.py"
        result = _run.mirror_path(source_file, source_root, target_root, ".go")
        assert result == target_root / "app.go"

    def test_nested_mirror(self, tmp_path):
        source_root = tmp_path / "src"
        target_root = tmp_path / "target"
        source_file = source_root / "pkg" / "util.py"
        result = _run.mirror_path(source_file, source_root, target_root, ".go")
        assert result == target_root / "pkg" / "util.go"


class TestEvaluateFile:
    def test_go_hello_world(self, tmp_path):
        py_file = tmp_path / "hello.py"
        py_file.write_text("print('hello world')")
        go_file = tmp_path / "hello.go"
        go_file.write_text(
            'package main\nimport "fmt"\nfunc main() { fmt.Println("hello world") }\n'
        )
        record = _run.evaluate_file(py_file, go_file)
        assert record.compiles is True
        assert record.runs_successfully is True
        assert record.io_equivalent is True
        assert record.computational_accuracy is True

    def test_go_compile_failure(self, tmp_path):
        py_file = tmp_path / "bad.py"
        py_file.write_text("print('hi')")
        go_file = tmp_path / "bad.go"
        go_file.write_text("not valid go")
        record = _run.evaluate_file(py_file, go_file)
        assert record.compiles is False
        assert record.runs_successfully is False

    def test_output_mismatch(self, tmp_path):
        py_file = tmp_path / "diff.py"
        py_file.write_text("print('python')")
        go_file = tmp_path / "diff.go"
        go_file.write_text(
            'package main\nimport "fmt"\nfunc main() { fmt.Println("golang") }\n'
        )
        record = _run.evaluate_file(py_file, go_file)
        assert record.compiles is True
        assert record.runs_successfully is True
        assert record.io_equivalent is False
        assert record.computational_accuracy is False

    def test_missing_target(self, tmp_path):
        py_file = tmp_path / "src.py"
        py_file.write_text("pass")
        go_file = tmp_path / "missing.go"
        record = _run.evaluate_file(py_file, go_file)
        assert record.compiles is False
        assert "not found" in record.notes

    def test_with_go_tests(self, tmp_path):
        py_file = tmp_path / "add.py"
        py_file.write_text("def add(a, b): return a + b\nprint(add(1, 2))")
        go_file = tmp_path / "add.go"
        go_file.write_text(
            "package main\n\n"
            'import "fmt"\n\n'
            "func add(a, b int) int { return a + b }\n"
            "func main() { fmt.Println(add(1, 2)) }\n"
        )
        go_test = (
            "package main\n\n"
            'import "testing"\n\n'
            "func TestAdd(t *testing.T) {\n"
            "  if add(1, 2) != 3 {\n"
            '    t.Errorf("expected 3")\n'
            "  }\n"
            "}\n"
        )
        record = _run.evaluate_file(py_file, go_file, go_test_code=go_test)
        assert record.compiles is True
        assert record.tests_total >= 1
        assert record.tests_passed >= 1
        assert record.test_pass_rate == 1.0
