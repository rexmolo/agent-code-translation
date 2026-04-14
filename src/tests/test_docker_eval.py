"""Tests for docker_eval module.

Unit tests run without Docker. Integration tests require a running Docker daemon.
Run unit tests:   uv run pytest src/tests/test_docker_eval.py -v -m "not integration"
Run all tests:    uv run pytest src/tests/test_docker_eval.py -v
"""

import pytest

from src.core.docker_eval import (
    _extract_declarations,
    _extract_imports,
    build_solution_file,
    build_test_file,
    build_combined_source,
    check_docker_available,
    ensure_package_declaration,
    ensure_go_mod_cache,
    run_in_docker,
    infer_test_imports,
    prepare_evaluation_sources,
    strip_markdown_fences,
    evaluate_single_task,
)


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
# infer_test_imports
# ---------------------------------------------------------------------------


class TestInferTestImports:
    def test_infers_default_imports_only(self):
        test_code = "func TestAdd(t *testing.T) { assert.True(t, Add(1, 2) == 3) }"
        assert infer_test_imports(test_code) == [
            "github.com/stretchr/testify/assert",
            "testing",
        ]

    def test_infers_known_stdlib_imports_from_prefixes(self):
        test_code = (
            "func TestHelpers(t *testing.T) {\n"
            "    _ = rand.New(rand.NewSource(42))\n"
            "    _ = time.Now()\n"
            "    _ = math.Abs(-1)\n"
            "    _ = strings.TrimSpace(\" hi \")\n"
            "}\n"
        )
        assert infer_test_imports(test_code) == [
            "github.com/stretchr/testify/assert",
            "math",
            "math/rand",
            "strings",
            "testing",
            "time",
        ]

    def test_ignores_unknown_qualified_prefixes(self):
        test_code = "func TestFoo(t *testing.T) { _ = custompkg.DoThing() }"
        assert infer_test_imports(test_code) == [
            "github.com/stretchr/testify/assert",
            "testing",
        ]


# ---------------------------------------------------------------------------
# _extract_declarations
# ---------------------------------------------------------------------------


class TestExtractDeclarations:
    def test_strips_package_and_imports(self):
        code = (
            "package main\n"
            "\n"
            'import "fmt"\n'
            "\n"
            "func Add(a, b int) int {\n"
            "    return a + b\n"
            "}"
        )
        result, _ = _extract_declarations(code)
        assert "package main" not in result
        assert 'import "fmt"' not in result
        assert "func Add(a, b int) int {" in result
        assert "return a + b" in result

    def test_strips_import_block(self):
        code = (
            "package main\n"
            "\n"
            "import (\n"
            '    "fmt"\n'
            '    "strings"\n'
            ")\n"
            "\n"
            "func Hello() string {\n"
            '    return "hello"\n'
            "}"
        )
        result, _ = _extract_declarations(code)
        assert "import" not in result
        assert "fmt" not in result
        assert "func Hello() string {" in result

    def test_strips_func_main(self):
        code = (
            "package main\n"
            "\n"
            "func Add(a, b int) int {\n"
            "    return a + b\n"
            "}\n"
            "\n"
            "func main() {\n"
            "    fmt.Println(Add(1, 2))\n"
            "}"
        )
        result, _ = _extract_declarations(code)
        assert "func main()" not in result
        assert "Println" not in result
        assert "func Add(a, b int) int {" in result

    def test_preserves_helper_functions(self):
        code = (
            "package main\n"
            "\n"
            "func helper(x int) int { return x * 2 }\n"
            "\n"
            "func Solve(n int) int {\n"
            "    return helper(n) + 1\n"
            "}\n"
            "\n"
            "func main() {\n"
            "    fmt.Println(Solve(5))\n"
            "}"
        )
        result, _ = _extract_declarations(code)
        assert "func helper(x int) int" in result
        assert "func Solve(n int) int {" in result
        assert "func main()" not in result

    def test_bare_function_no_package(self):
        """HumanEval-X ground truth format: just imports + function, no package."""
        code = (
            "import (\n"
            '    "math"\n'
            ")\n"
            "\n"
            "func HasCloseElements(numbers []float64, threshold float64) bool {\n"
            "    for i := 0; i < len(numbers); i++ {\n"
            "        for j := i + 1; j < len(numbers); j++ {\n"
            "            if math.Abs(numbers[i]-numbers[j]) < threshold {\n"
            "                return true\n"
            "            }\n"
            "        }\n"
            "    }\n"
            "    return false\n"
            "}"
        )
        result, _ = _extract_declarations(code)
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
    def test_preserves_existing_package_and_imports(self):
        generated = (
            "package main\n"
            "\n"
            'import str "strings"\n'
            "\n"
            "func Solve(x string) bool {\n"
            '    return str.HasPrefix(x, "a")\n'
            "}"
        )
        result = build_solution_file(generated)
        assert result.startswith("package main\n")
        assert 'import str "strings"' in result
        assert "func Solve(x string) bool {" in result
        assert result.count("package main") == 1

    def test_handles_bare_function_by_adding_package_only(self):
        generated = (
            "import (\n"
            '    "math"\n'
            ")\n"
            "\n"
            "func HasCloseElements(numbers []float64, threshold float64) bool {\n"
            "    return true\n"
            "}"
        )
        result = build_solution_file(generated)
        assert result.startswith("package main\n")
        assert '"math"' in result
        assert "func HasCloseElements" in result
        assert "return true" in result

    def test_strips_markdown_fences(self):
        generated = (
            "```go\npackage main\n\nfunc Add(a, b int) int { return a + b }\n```"
        )
        result = build_solution_file(generated)
        assert "```" not in result
        assert "func Add" in result

    def test_preserves_main_function(self):
        generated = (
            "func helper() int { return 1 }\n\n"
            "func main() {\n"
            "    _ = helper()\n"
            "}\n"
        )
        result = build_solution_file(generated)
        assert result.startswith("package main\n")
        assert "func main()" in result


class TestPrepareEvaluationSources:
    def test_prepare_evaluation_sources_is_near_verbatim(self):
        generated = 'import str "strings"\n\nfunc Solve(x string) bool { return str.HasPrefix(x, "a") }\n'
        solution_source, test_source = prepare_evaluation_sources(
            generated,
            "func TestSolve(t *testing.T) { assert.True(t, Solve(\"abc\")) }",
        )
        assert solution_source.startswith("package main\n")
        assert 'import str "strings"' in solution_source
        assert "str.HasPrefix" in solution_source
        assert test_source.startswith("package main\n")
        assert '"testing"' in test_source


class TestEnsurePackageDeclaration:
    def test_adds_package_when_missing(self):
        result = ensure_package_declaration("func Solve() {}\n")
        assert result.startswith("package main\n")

    def test_leaves_existing_package_unchanged(self):
        result = ensure_package_declaration("package gcd\n\nfunc Solve() {}\n")
        assert result.startswith("package gcd\n")


# ---------------------------------------------------------------------------
# build_test_file
# ---------------------------------------------------------------------------


class TestBuildTestFile:
    def test_wraps_test_with_imports(self):
        test_code = (
            "func TestAdd(t *testing.T) {\n"
            "    assert := assert.New(t)\n"
            "    assert.Equal(3, Add(1, 2))\n"
            "}"
        )
        result = build_test_file(test_code)
        assert result.startswith("package main\n")
        assert '"testing"' in result
        assert '"github.com/stretchr/testify/assert"' in result
        assert "func TestAdd" in result

    def test_detects_rand_import(self):
        test_code = (
            "func TestFoo(t *testing.T) {\n    rng := rand.New(rand.NewSource(42))\n}"
        )
        result = build_test_file(test_code)
        assert '"math/rand"' in result
        assert '"time"' not in result

    def test_detects_time_import(self):
        test_code = "func TestFoo(t *testing.T) {\n    _ = time.Now()\n}"
        result = build_test_file(test_code)
        assert '"time"' in result

    def test_detects_multiple_mapping_driven_imports(self):
        test_code = (
            "func TestFoo(t *testing.T) {\n"
            "    _ = reflect.DeepEqual([]int{1}, []int{1})\n"
            "    _ = strings.HasPrefix(\"abc\", \"a\")\n"
            "    _ = strconv.Itoa(42)\n"
            "}\n"
        )
        result = build_test_file(test_code)
        assert '"reflect"' in result
        assert '"strings"' in result
        assert '"strconv"' in result


# ---------------------------------------------------------------------------
# build_combined_source (legacy API)
# ---------------------------------------------------------------------------


class TestBuildCombinedSource:
    def test_inserts_before_main_in_test(self):
        generated = (
            "package main\n"
            "\n"
            'import "fmt"\n'
            "\n"
            "func Add(a, b int) int {\n"
            "    return a + b\n"
            "}"
        )
        test_code = (
            "package main\n"
            "\n"
            'import "fmt"\n'
            "\n"
            "func main() {\n"
            "    if Add(1, 2) != 3 {\n"
            '        panic("fail")\n'
            "    }\n"
            "}"
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
        from rich.console import Console

        console = Console()
        if not check_docker_available(console):
            pytest.skip("Docker daemon is not available")
        ensure_go_mod_cache()

    def test_check_docker_available(self):
        from rich.console import Console

        console = Console()
        assert check_docker_available(console) is True

    def test_passing_test(self):
        solution = "package main\n\nfunc Add(a, b int) int { return a + b }\n"
        test = (
            "package main\n"
            "\n"
            "import (\n"
            '\t"testing"\n'
            '\t"github.com/stretchr/testify/assert"\n'
            ")\n"
            "\n"
            "func TestAdd(t *testing.T) {\n"
            "\tassert := assert.New(t)\n"
            "\tassert.Equal(3, Add(1, 2))\n"
            "}\n"
        )
        result = run_in_docker(solution, test, "test_pass")
        assert result.compiles is True
        assert result.passes is True
        assert result.timed_out is False

    def test_failing_test(self):
        solution = "package main\n\nfunc Add(a, b int) int { return a + b }\n"
        test = (
            "package main\n"
            "\n"
            "import (\n"
            '\t"testing"\n'
            '\t"github.com/stretchr/testify/assert"\n'
            ")\n"
            "\n"
            "func TestAdd(t *testing.T) {\n"
            "\tassert := assert.New(t)\n"
            "\tassert.Equal(999, Add(1, 2))\n"
            "}\n"
        )
        result = run_in_docker(solution, test, "test_fail")
        assert result.compiles is True
        assert result.passes is False

    def test_compile_error(self):
        solution = "package main\n\nfunc Bad() { undefined_var }\n"
        test = (
            'package main\n\nimport "testing"\n\nfunc TestBad(t *testing.T) { Bad() }\n'
        )
        result = run_in_docker(solution, test, "test_compile_err")
        assert result.compiles is False
        assert result.passes is False

    def test_evaluate_single_task_pass(self):
        generated = "package main\n\nfunc Add(a, b int) int { return a + b }\n"
        test_code = (
            "func TestAdd(t *testing.T) {\n"
            "\tassert := assert.New(t)\n"
            "\tassert.Equal(3, Add(1, 2))\n"
            "}\n"
        )
        record = evaluate_single_task("Go/0", generated, test_code)
        assert record.compiles is True
        assert record.pass_at_1 is True
        assert record.dataset == "humaneval-x"

    def test_evaluate_single_task_fail(self):
        generated = "package main\n\nfunc Add(a, b int) int { return 0 }\n"
        test_code = (
            "func TestAdd(t *testing.T) {\n"
            "\tassert := assert.New(t)\n"
            "\tassert.Equal(3, Add(1, 2))\n"
            "}\n"
        )
        record = evaluate_single_task("Go/0", generated, test_code)
        assert record.compiles is True
        assert record.pass_at_1 is False
