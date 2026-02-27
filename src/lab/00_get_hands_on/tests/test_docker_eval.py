"""Tests for docker_eval module.

Unit tests run without Docker. Integration tests require a running Docker daemon.
Run unit tests:   uv run pytest src/lab/00_get_hands_on/tests/test_docker_eval.py -v -m "not integration"
Run all tests:    uv run pytest src/lab/00_get_hands_on/tests/test_docker_eval.py -v
"""

import importlib

import pytest

_docker_eval = importlib.import_module("src.lab.00_get_hands_on.docker_eval")
strip_markdown_fences = _docker_eval.strip_markdown_fences
_extract_declarations = _docker_eval._extract_declarations
_extract_imports = _docker_eval._extract_imports
build_solution_file = _docker_eval.build_solution_file
build_test_file = _docker_eval.build_test_file
build_combined_source = _docker_eval.build_combined_source


# ---------------------------------------------------------------------------
# strip_markdown_fences
# ---------------------------------------------------------------------------

class TestStripMarkdownFences:
    def test_strips_go_fences(self):
        code = "```go\nfunc Add(a, b int) int { return a + b }\n```"
        assert strip_markdown_fences(code) == "func Add(a, b int) int { return a + b }"

    def test_strips_bare_fences(self):
        code = "```\nfunc Add(a, b int) int { return a + b }\n```"
        assert strip_markdown_fences(code) == "func Add(a, b int) int { return a + b }"

    def test_no_fences_unchanged(self):
        code = "func Add(a, b int) int { return a + b }"
        assert strip_markdown_fences(code) == code


# ---------------------------------------------------------------------------
# _extract_declarations
# ---------------------------------------------------------------------------

class TestExtractDeclarations:
    def test_strips_package_and_imports(self):
        code = (
            'package main\n'
            '\n'
            'import "fmt"\n'
            '\n'
            'func Add(a, b int) int {\n'
            '    return a + b\n'
            '}'
        )
        result = _extract_declarations(code)
        assert "package main" not in result
        assert 'import "fmt"' not in result
        assert "func Add(a, b int) int {" in result
        assert "return a + b" in result

    def test_strips_import_block(self):
        code = (
            'package main\n'
            '\n'
            'import (\n'
            '    "fmt"\n'
            '    "strings"\n'
            ')\n'
            '\n'
            'func Hello() string {\n'
            '    return "hello"\n'
            '}'
        )
        result = _extract_declarations(code)
        assert "import" not in result
        assert "fmt" not in result
        assert "func Hello() string {" in result

    def test_strips_func_main(self):
        code = (
            'package main\n'
            '\n'
            'func Add(a, b int) int {\n'
            '    return a + b\n'
            '}\n'
            '\n'
            'func main() {\n'
            '    fmt.Println(Add(1, 2))\n'
            '}'
        )
        result = _extract_declarations(code)
        assert "func main()" not in result
        assert "Println" not in result
        assert "func Add(a, b int) int {" in result

    def test_preserves_helper_functions(self):
        code = (
            'package main\n'
            '\n'
            'func helper(x int) int { return x * 2 }\n'
            '\n'
            'func Solve(n int) int {\n'
            '    return helper(n) + 1\n'
            '}\n'
            '\n'
            'func main() {\n'
            '    fmt.Println(Solve(5))\n'
            '}'
        )
        result = _extract_declarations(code)
        assert "func helper(x int) int" in result
        assert "func Solve(n int) int {" in result
        assert "func main()" not in result

    def test_bare_function_no_package(self):
        """HumanEval-X ground truth format: just imports + function, no package."""
        code = (
            'import (\n'
            '    "math"\n'
            ')\n'
            '\n'
            'func HasCloseElements(numbers []float64, threshold float64) bool {\n'
            '    for i := 0; i < len(numbers); i++ {\n'
            '        for j := i + 1; j < len(numbers); j++ {\n'
            '            if math.Abs(numbers[i]-numbers[j]) < threshold {\n'
            '                return true\n'
            '            }\n'
            '        }\n'
            '    }\n'
            '    return false\n'
            '}'
        )
        result = _extract_declarations(code)
        assert "import" not in result
        assert "func HasCloseElements" in result
        assert "return false" in result


# ---------------------------------------------------------------------------
# _extract_imports
# ---------------------------------------------------------------------------

class TestExtractImports:
    def test_single_import(self):
        code = 'import "fmt"\n\nfunc main() {}'
        assert _extract_imports(code) == ["fmt"]

    def test_import_block(self):
        code = 'import (\n    "fmt"\n    "math"\n)\n'
        result = _extract_imports(code)
        assert "fmt" in result
        assert "math" in result

    def test_no_imports(self):
        code = "func Add(a, b int) int { return a + b }"
        assert _extract_imports(code) == []


# ---------------------------------------------------------------------------
# build_solution_file
# ---------------------------------------------------------------------------

class TestBuildSolutionFile:
    def test_wraps_with_package_and_imports(self):
        generated = (
            'package main\n'
            '\n'
            'import "math"\n'
            '\n'
            'func Solve(x float64) float64 {\n'
            '    return math.Sqrt(x)\n'
            '}'
        )
        result = build_solution_file(generated)
        assert result.startswith("package main\n")
        assert '"math"' in result
        assert "func Solve(x float64) float64 {" in result
        # Only one package declaration
        assert result.count("package main") == 1

    def test_handles_bare_function(self):
        """HumanEval-X format: no package, just imports + function."""
        generated = (
            'import (\n'
            '    "math"\n'
            ')\n'
            '\n'
            'func HasCloseElements(numbers []float64, threshold float64) bool {\n'
            '    return true\n'
            '}'
        )
        result = build_solution_file(generated)
        assert result.startswith("package main\n")
        assert '"math"' in result
        assert "func HasCloseElements" in result

    def test_strips_markdown_fences(self):
        generated = (
            '```go\n'
            'package main\n'
            '\n'
            'func Add(a, b int) int { return a + b }\n'
            '```'
        )
        result = build_solution_file(generated)
        assert "```" not in result
        assert "func Add" in result


# ---------------------------------------------------------------------------
# build_test_file
# ---------------------------------------------------------------------------

class TestBuildTestFile:
    def test_wraps_test_with_imports(self):
        test_code = (
            'func TestAdd(t *testing.T) {\n'
            '    assert := assert.New(t)\n'
            '    assert.Equal(3, Add(1, 2))\n'
            '}'
        )
        result = build_test_file(test_code)
        assert result.startswith("package main\n")
        assert '"testing"' in result
        assert '"github.com/stretchr/testify/assert"' in result
        assert "func TestAdd" in result

    def test_detects_rand_import(self):
        test_code = 'func TestFoo(t *testing.T) {\n    rng := rand.New(rand.NewSource(42))\n}'
        result = build_test_file(test_code)
        assert '"math/rand"' in result
        assert '"time"' not in result

    def test_detects_time_import(self):
        test_code = 'func TestFoo(t *testing.T) {\n    _ = time.Now()\n}'
        result = build_test_file(test_code)
        assert '"time"' in result


# ---------------------------------------------------------------------------
# build_combined_source (legacy API)
# ---------------------------------------------------------------------------

class TestBuildCombinedSource:
    def test_inserts_before_main_in_test(self):
        generated = (
            'package main\n'
            '\n'
            'import "fmt"\n'
            '\n'
            'func Add(a, b int) int {\n'
            '    return a + b\n'
            '}'
        )
        test_code = (
            'package main\n'
            '\n'
            'import "fmt"\n'
            '\n'
            'func main() {\n'
            '    if Add(1, 2) != 3 {\n'
            '        panic("fail")\n'
            '    }\n'
            '}'
        )
        combined = build_combined_source(generated, test_code)
        add_pos = combined.find("func Add")
        main_pos = combined.find("func main()")
        assert add_pos != -1
        assert main_pos != -1
        assert add_pos < main_pos


# ---------------------------------------------------------------------------
# Integration tests (require Docker)
# ---------------------------------------------------------------------------

@pytest.mark.integration
class TestDockerIntegration:
    @pytest.fixture(autouse=True, scope="class")
    def _setup_mod_cache(self):
        """Ensure Go module cache is populated before running tests."""
        _docker_eval.ensure_go_mod_cache()

    def test_check_docker_available(self):
        from rich.console import Console
        console = Console()
        assert _docker_eval.check_docker_available(console) is True

    def test_passing_test(self):
        solution = (
            'package main\n'
            '\n'
            'func Add(a, b int) int { return a + b }\n'
        )
        test = (
            'package main\n'
            '\n'
            'import (\n'
            '\t"testing"\n'
            '\t"github.com/stretchr/testify/assert"\n'
            ')\n'
            '\n'
            'func TestAdd(t *testing.T) {\n'
            '\tassert := assert.New(t)\n'
            '\tassert.Equal(3, Add(1, 2))\n'
            '}\n'
        )
        result = _docker_eval.run_in_docker(solution, test, "test_pass")
        assert result.compiles is True
        assert result.passes is True
        assert result.timed_out is False

    def test_failing_test(self):
        solution = (
            'package main\n'
            '\n'
            'func Add(a, b int) int { return a + b }\n'
        )
        test = (
            'package main\n'
            '\n'
            'import (\n'
            '\t"testing"\n'
            '\t"github.com/stretchr/testify/assert"\n'
            ')\n'
            '\n'
            'func TestAdd(t *testing.T) {\n'
            '\tassert := assert.New(t)\n'
            '\tassert.Equal(999, Add(1, 2))\n'
            '}\n'
        )
        result = _docker_eval.run_in_docker(solution, test, "test_fail")
        assert result.compiles is True
        assert result.passes is False

    def test_compile_error(self):
        solution = 'package main\n\nfunc Bad() { undefined_var }\n'
        test = (
            'package main\n'
            '\n'
            'import "testing"\n'
            '\n'
            'func TestBad(t *testing.T) { Bad() }\n'
        )
        result = _docker_eval.run_in_docker(solution, test, "test_compile_err")
        assert result.compiles is False
        assert result.passes is False

    def test_evaluate_single_task_pass(self):
        generated = (
            'package main\n'
            '\n'
            'func Add(a, b int) int { return a + b }\n'
        )
        test_code = (
            'func TestAdd(t *testing.T) {\n'
            '\tassert := assert.New(t)\n'
            '\tassert.Equal(3, Add(1, 2))\n'
            '}\n'
        )
        record = _docker_eval.evaluate_single_task("Go/0", generated, test_code)
        assert record.compiles is True
        assert record.pass_at_1 is True
        assert record.dataset == "humaneval-x"

    def test_evaluate_single_task_fail(self):
        generated = (
            'package main\n'
            '\n'
            'func Add(a, b int) int { return 0 }\n'
        )
        test_code = (
            'func TestAdd(t *testing.T) {\n'
            '\tassert := assert.New(t)\n'
            '\tassert.Equal(3, Add(1, 2))\n'
            '}\n'
        )
        record = _docker_eval.evaluate_single_task("Go/0", generated, test_code)
        assert record.compiles is True
        assert record.pass_at_1 is False
