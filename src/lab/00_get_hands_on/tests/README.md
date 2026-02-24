# Tests for Experiment 00: Get Hands On

## Quick Reference

```bash
# From project root: /Volumes/MyZhiTai/DEV/www/thesis/experiments/

# Run ALL offline tests (no API key needed, ~5s)
uv run pytest src/lab/00_get_hands_on/tests/ -v --ignore=src/lab/00_get_hands_on/tests/test_preflight.py

# Run preflight only (needs valid MINIMAX_API_KEY in .env)
uv run pytest src/lab/00_get_hands_on/tests/test_preflight.py -v

# Run everything (offline + preflight)
uv run pytest src/lab/00_get_hands_on/tests/ -v

# Run a single test file
uv run pytest src/lab/00_get_hands_on/tests/test_tools.py -v

# Run a single test class
uv run pytest src/lab/00_get_hands_on/tests/test_tools.py::TestCompileGoCode -v

# Run a single test
uv run pytest src/lab/00_get_hands_on/tests/test_tools.py::TestCompileGoCode::test_valid_go_compiles -v
```

## Test Files

### test_preflight.py — Environment & API Connection

**Requires**: Live API key, Go compiler, Python3.

Run this first when setting up or after changing API keys. If these fail, the pipeline
will not work.

| Test | What it checks |
|------|----------------|
| `TestEnvironment::test_minimax_api_key_is_set` | `MINIMAX_API_KEY` exists in `.env` and is non-trivial |
| `TestEnvironment::test_go_compiler_available` | `go version` runs successfully |
| `TestEnvironment::test_python3_available` | `python3 --version` runs successfully |
| `TestAPIConnection::test_minimax_api_responds` | Creates a translation agent, sends a minimal prompt, verifies non-empty `go_code` response |

### test_tools.py — Evaluation Tools (Offline)

**Requires**: Go compiler, Python3. No API key needed.

Tests the 6 `@tool()`-decorated functions in `tools.py`. These are the subprocess
wrappers that compile, run, and compare code. Tests call `.entrypoint` on each tool
because Agno's `@tool()` decorator wraps functions into `Function` objects that are
not directly callable.

| Class | Tools tested | Key cases |
|-------|-------------|-----------|
| `TestCompileGoCode` | `compile_go_code` | Valid Go compiles, invalid Go fails |
| `TestRunGoCode` | `run_go_code` | Hello world, stdin input, runtime panic |
| `TestRunPythonCode` | `run_python_code` | Hello world, syntax error |
| `TestCompareOutputs` | `compare_outputs` | Exact match, whitespace stripping, line diff, line count mismatch |
| `TestRunGoTests` | `run_go_tests` | Passing tests, failing tests (with `go mod init` + `go test -v`) |
| `TestRunPythonTests` | `run_python_tests` | Passing pytest, failing pytest |

### test_metrics.py — Metric Computation & Display (Offline)

**Requires**: Nothing (pure Python).

Tests `compute_summary()` and the Rich table display functions. Uses fabricated
`EvaluationRecord` objects — no file I/O or subprocesses.

| Class | What it covers |
|-------|---------------|
| `TestComputeSummary` | Empty records, all pass, all fail, mixed results, test_pass_rate only averages files that have tests |
| `TestDisplayTables` | Smoke tests — verifies `display_summary_table` and `display_per_file_table` don't crash |

### test_run.py — Pipeline Helper Functions (Offline)

**Requires**: Go compiler, Python3. No API key needed.

Tests the utility functions and `evaluate_file()` from `run.py`. Uses `tmp_path`
(pytest fixture) to create temporary files for each test.

| Class | Functions tested | Key cases |
|-------|-----------------|-----------|
| `TestDiscoverPythonFiles` | `discover_python_files` | Finds .py, excludes test_*.py / *_test.py, recursive, empty dir |
| `TestFindTestFile` | `find_test_file` | test_ prefix, _test suffix, tests/ subdir, no tests found |
| `TestMirrorPath` | `mirror_path` | Flat path, nested path with extension change |
| `TestEvaluateFile` | `evaluate_file` | Go hello world (full pass), compile failure, output mismatch, missing target file, Go test execution with pass rate |

## Recommended Test Order

1. **test_preflight.py** — Run once to verify environment is set up correctly.
2. **test_tools.py** — Validates the core execution tools work on your machine.
3. **test_metrics.py** — Validates metric math and display.
4. **test_run.py** — Validates the pipeline logic end-to-end (without LLM calls).

If steps 2-4 pass but the pipeline still fails, the issue is in the LLM response
parsing (`extract_go_code`, `extract_go_test_code`) or the Team coordination.

## Adding New Tests

- Place new test files in this directory following the `test_<module>.py` convention.
- Use `importlib.import_module("src.lab.00_get_hands_on.<module>")` for imports
  (the `00_` prefix prevents normal Python import syntax).
- For Agno tools, call `.entrypoint` to get the raw function:
  ```python
  _tools = importlib.import_module("src.lab.00_get_hands_on.tools")
  compile_go_code = _tools.compile_go_code.entrypoint
  ```
- Use `tmp_path` (pytest built-in fixture) for any file I/O tests.
- Keep tests independent — each test should set up its own data.
