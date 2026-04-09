"""Docker-based evaluation for HumanEval-X Go translations.

Runs LLM-generated Go code safely inside Docker containers with
network isolation and memory limits. HumanEval-X Go tests use the
standard `testing` package with `testify/assert`, so we use `go test`.
"""

from __future__ import annotations

import re
import subprocess
import tempfile
from dataclasses import dataclass
from pathlib import Path

from rich.console import Console

from src.core.schemas import EvaluationRecord

DEFAULT_GO_IMAGE = "golang:1.26-alpine"
DEFAULT_TIMEOUT = 60

# Minimal go.mod that includes testify for HumanEval-X assertions
_GO_MOD = """\
module humaneval

go 1.26.0

require github.com/stretchr/testify v1.11.1

require (
\tgithub.com/davecgh/go-spew v1.1.1 // indirect
\tgithub.com/pmezard/go-difflib v1.0.0 // indirect
\tgopkg.in/yaml.v3 v3.0.1 // indirect
)
"""

_GO_SUM = """\
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/stretchr/testify v1.11.1 h1:7s2iGBzp5EwR7/aIZr8ao5+dra3wiQyKjjFuvgVKu7U=
github.com/stretchr/testify v1.11.1/go.mod h1:wZwfW3scLgRK+23gO65QZefKpKQRnfz6sD981Nm4B6U=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
"""


@dataclass
class DockerResult:
    """Result from running Go code inside Docker."""

    compiles: bool = False
    passes: bool = False
    stdout: str = ""
    stderr: str = ""
    timed_out: bool = False


def check_docker_available(console: Console) -> bool:
    """Verify Docker daemon is installed and running."""
    try:
        result = subprocess.run(
            ["docker", "info"],
            capture_output=True,
            text=True,
            timeout=10,
        )
        if result.returncode == 0:
            console.print("[green]OK[/green]   Docker is available")
            return True
        console.print(f"[red]FAIL[/red] Docker daemon not running: {result.stderr.strip()[:100]}")
        return False
    except FileNotFoundError:
        console.print("[red]FAIL[/red] Docker not found on PATH")
        return False
    except subprocess.TimeoutExpired:
        console.print("[red]FAIL[/red] Docker info timed out")
        return False


def ensure_go_image(console: Console, image: str = DEFAULT_GO_IMAGE) -> bool:
    """Ensure the Go Docker image is available locally, pulling if needed."""
    try:
        inspect = subprocess.run(
            ["docker", "image", "inspect", image],
            capture_output=True,
            text=True,
            timeout=10,
        )
        if inspect.returncode == 0:
            console.print(f"[green]OK[/green]   Image {image} is cached")
            return True
    except (subprocess.TimeoutExpired, FileNotFoundError):
        pass

    console.print(f"[dim]Pulling {image}...[/dim]")
    try:
        pull = subprocess.run(
            ["docker", "pull", image],
            capture_output=True,
            text=True,
            timeout=300,
        )
        if pull.returncode == 0:
            console.print(f"[green]OK[/green]   Pulled {image}")
            return True
        console.print(f"[red]FAIL[/red] Failed to pull {image}: {pull.stderr.strip()[:100]}")
        return False
    except subprocess.TimeoutExpired:
        console.print(f"[red]FAIL[/red] Pulling {image} timed out")
        return False


def strip_markdown_fences(code: str) -> str:
    """Remove markdown code fences (```go ... ```) that LLMs sometimes emit."""
    code = re.sub(r"^```(?:go|Go)?\s*\n", "", code)
    code = re.sub(r"\n```\s*$", "", code)
    return code


def _extract_declarations(code: str) -> str:
    """Extract everything from generated code except package, imports, and func main."""
    lines = code.split("\n")
    result = []
    brace_depth = 0
    in_import_block = False
    in_main = False

    i = 0
    while i < len(lines):
        line = lines[i]
        stripped = line.strip()

        # Skip package declaration
        if stripped.startswith("package "):
            i += 1
            continue

        # Skip single-line import
        if re.match(r'^import\s+"', stripped) and not stripped.endswith("("):
            i += 1
            continue

        # Skip import block
        if stripped.startswith("import (") or stripped == "import(":
            in_import_block = True
            i += 1
            continue
        if in_import_block:
            if stripped == ")":
                in_import_block = False
            i += 1
            continue

        # Skip func main()
        if re.match(r"^func\s+main\s*\(", stripped):
            in_main = True
            brace_depth = 0
            for ch in line:
                if ch == "{":
                    brace_depth += 1
                elif ch == "}":
                    brace_depth -= 1
            if brace_depth <= 0 and "{" in line:
                in_main = brace_depth > 0
            i += 1
            continue
        if in_main:
            for ch in line:
                if ch == "{":
                    brace_depth += 1
                elif ch == "}":
                    brace_depth -= 1
            if brace_depth <= 0:
                in_main = False
            i += 1
            continue

        result.append(line)
        i += 1

    return "\n".join(result).strip()


def build_solution_file(generated_code: str) -> str:
    """Build a solution.go file from LLM-generated code.

    The generated code may include package/imports/func main, or may be
    a bare function body (like HumanEval-X ground truth format).
    We extract just the declarations and wrap them in a proper Go file.
    """
    generated_code = strip_markdown_fences(generated_code)

    # Collect imports from generated code
    imports = _extract_imports(generated_code)

    # Extract function/type/var/const declarations
    declarations = _extract_declarations(generated_code)

    # Filter imports to only those actually used in the declarations.
    # Stripping func main() can orphan imports that were only used there.
    imports = [imp for imp in imports if imp.split("/")[-1] + "." in declarations]

    lines = ["package main", ""]
    if imports:
        lines.append("import (")
        for imp in sorted(imports):
            lines.append(f'\t"{imp}"')
        lines.append(")")
        lines.append("")
    lines.append(declarations)
    lines.append("")

    return "\n".join(lines)


def build_test_file(test_code: str) -> str:
    """Build a solution_test.go file from HumanEval-X test code.

    The test field from HumanEval-X is just the test function body
    (no package, no imports). We wrap it with proper package/import header.
    """
    # Standard imports needed by HumanEval-X tests
    imports = [
        "testing",
        "github.com/stretchr/testify/assert",
    ]

    # Detect additional imports needed
    if "rand." in test_code or "rand.New" in test_code:
        imports.append("math/rand")
    if "time." in test_code:
        imports.append("time")
    if "math." in test_code:
        imports.append("math")
    if "fmt." in test_code:
        imports.append("fmt")
    if "sort." in test_code:
        imports.append("sort")
    if "strings." in test_code:
        imports.append("strings")
    if "strconv." in test_code:
        imports.append("strconv")
    if "reflect." in test_code:
        imports.append("reflect")

    lines = ["package main", ""]
    lines.append("import (")
    for imp in sorted(imports):
        lines.append(f'\t"{imp}"')
    lines.append(")")
    lines.append("")
    lines.append(test_code)
    lines.append("")

    return "\n".join(lines)


def _extract_imports(code: str) -> list[str]:
    """Extract import paths from Go source code."""
    imports = []
    # Single-line: import "fmt"
    for m in re.finditer(r'^import\s+"([^"]+)"', code, re.MULTILINE):
        imports.append(m.group(1))
    # Block: import ( "fmt" \n "math" )
    for block in re.finditer(r"import\s*\((.*?)\)", code, re.DOTALL):
        for m in re.finditer(r'"([^"]+)"', block.group(1)):
            imports.append(m.group(1))
    return imports


def _insert_before_main(test_code: str, declarations: str) -> str:
    """Insert declarations before func main() in the test harness."""
    if not declarations:
        return test_code

    match = re.search(r"^func\s+main\s*\(", test_code, re.MULTILINE)
    if match:
        insert_pos = match.start()
        return test_code[:insert_pos] + declarations + "\n\n" + test_code[insert_pos:]

    return test_code + "\n\n" + declarations


def build_combined_source(generated_code: str, test_code: str) -> str:
    """Merge translated function code into the HumanEval-X test harness.

    Legacy API kept for backward compatibility with existing tests.
    """
    generated_code = strip_markdown_fences(generated_code)
    extracted_lines = _extract_declarations(generated_code)
    combined = _insert_before_main(test_code, extracted_lines)
    return combined


_GOMOD_CACHE_VOLUME = "humaneval-gomod-cache"


def ensure_go_mod_cache(image: str = DEFAULT_GO_IMAGE, timeout: int = 120) -> bool:
    """Pre-download Go modules (testify) into a Docker volume.

    This runs once with network access. Subsequent test runs reuse
    the cached modules and can run with --network=none.
    """
    with tempfile.TemporaryDirectory(prefix="humaneval_modcache_") as tmpdir:
        tmppath = Path(tmpdir)
        (tmppath / "go.mod").write_text(_GO_MOD, encoding="utf-8")
        (tmppath / "go.sum").write_text(_GO_SUM, encoding="utf-8")
        # Need a minimal .go file for go mod download to work
        (tmppath / "stub.go").write_text("package main\n", encoding="utf-8")

        try:
            result = subprocess.run(
                [
                    "docker", "run", "--rm",
                    "--memory=512m",
                    "-v", f"{tmpdir}:/app",
                    "-v", f"{_GOMOD_CACHE_VOLUME}:/go/pkg/mod",
                    "-w", "/app",
                    image,
                    "go", "mod", "download",
                ],
                capture_output=True,
                text=True,
                timeout=timeout,
            )
            return result.returncode == 0
        except (subprocess.TimeoutExpired, FileNotFoundError):
            return False


def run_in_docker(
    solution_source: str,
    test_source: str,
    task_id: str,
    image: str = DEFAULT_GO_IMAGE,
    timeout: int = DEFAULT_TIMEOUT,
) -> DockerResult:
    """Run Go solution + test in a Docker container using `go test`.

    Sets up a Go module with testify dependency, then runs:
    1. `go vet ./...` to check compilation (works on test-only packages)
    2. `go test -v -count=1 ./...` to run tests

    Uses a shared Docker volume for Go module cache (testify etc.)
    and --network=none for safety.
    """
    result = DockerResult()

    with tempfile.TemporaryDirectory(prefix=f"humaneval_{task_id}_") as tmpdir:
        tmppath = Path(tmpdir)
        (tmppath / "go.mod").write_text(_GO_MOD, encoding="utf-8")
        (tmppath / "go.sum").write_text(_GO_SUM, encoding="utf-8")
        (tmppath / "solution.go").write_text(solution_source, encoding="utf-8")
        (tmppath / "solution_test.go").write_text(test_source, encoding="utf-8")

        docker_base = [
            "docker", "run", "--rm",
            "--network=none",
            "--memory=512m",
            "-v", f"{tmpdir}:/app",
            "-v", f"{_GOMOD_CACHE_VOLUME}:/go/pkg/mod",
            "-w", "/app",
            image,
        ]

        # Phase 1: Compile check via go vet (works on test-only packages)
        try:
            comp = subprocess.run(
                [*docker_base, "go", "vet", "./..."],
                capture_output=True,
                text=True,
                timeout=timeout,
            )
            result.stderr = comp.stderr
            if comp.returncode != 0:
                return result
            result.compiles = True
        except subprocess.TimeoutExpired:
            result.timed_out = True
            return result

        # Phase 2: Run tests
        try:
            test_run = subprocess.run(
                [*docker_base, "go", "test", "-v", "-count=1", "./..."],
                capture_output=True,
                text=True,
                timeout=timeout,
            )
            result.stdout = test_run.stdout
            result.stderr = test_run.stderr
            result.passes = test_run.returncode == 0
        except subprocess.TimeoutExpired:
            result.timed_out = True

    return result


def evaluate_single_task(
    task_id: str,
    generated_code: str,
    test_code: str,
    image: str = DEFAULT_GO_IMAGE,
    timeout: int = DEFAULT_TIMEOUT,
) -> EvaluationRecord:
    """Evaluate a single HumanEval-X task: build Go files, run in Docker."""
    task_num = task_id.split("/")[1] if "/" in task_id else task_id

    record = EvaluationRecord(
        source_file=task_id,
        target_file=f"Go_{task_num}.go",
        dataset="humaneval-x",
    )

    solution_source = build_solution_file(generated_code)
    test_source = build_test_file(test_code)

    docker_result = run_in_docker(
        solution_source, test_source, task_num,
        image=image, timeout=timeout,
    )

    record.compiles = docker_result.compiles
    record.runs_successfully = docker_result.compiles and not docker_result.timed_out
    record.pass_at_1 = docker_result.passes

    if docker_result.timed_out:
        record.notes = "Timed out"
    elif not docker_result.compiles:
        record.notes = docker_result.stderr.strip()
    elif not docker_result.passes:
        record.notes = docker_result.stderr.strip()

    return record
