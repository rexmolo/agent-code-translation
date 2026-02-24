# Experiment 00: Get Hands On — Python-to-Go Translation Pipeline

## Purpose

This is the first experiment in the thesis project. It establishes a baseline agent-based
pipeline for translating Python source code to Go, and evaluating the quality of those
translations using programmatic metrics. The goal is to have a simple, working starting
point that can be extended with more sophisticated translation strategies in later
experiments.

There is **no ground truth Go code** in this experiment. We are targeting real-world
project translation, so all evaluation metrics are computed by comparing the behavior of
the original Python code against the translated Go code directly.

## Architecture

The pipeline uses a single **Translator** agent that directly translates Python code to
Go. There is no team coordinator — the pipeline is a simple sequential loop that calls
the translator once per file, producing exactly **1 LLM call per file**.

Evaluation (compilation, execution, I/O comparison) is available as a separate step and
can be run manually after translation.

```
    ┌──────────────┐     ┌───────────────┐     ┌──────────────┐
    │  Python file  │ ──▶ │  Translator   │ ──▶ │   Go file    │
    │  (source)     │     │  (1 LLM call) │     │  (target)    │
    └──────────────┘     └───────────────┘     └──────────────┘
```

### Agent

| Agent | Role | Output Schema | Tools |
|-------|------|---------------|-------|
| **Translator** | Translates a Python file to idiomatic Go | `TranslationResult` (go_code, explanation) | None (LLM only) |

## Evaluation Metrics

All metrics are computed without ground truth, by comparing the translated Go code's
behavior against the original Python source. Evaluation is run separately from translation.

| Metric | What It Measures | How It Is Computed |
|--------|------------------|--------------------|
| **Compilation Success Rate** | Does the Go code compile? | `go build` on translated `.go` file; % that return exit code 0 |
| **Successful Translation Rate** | Does it compile AND run? | Subset of compiled files that also execute via `go run` without crashing |
| **Computational Accuracy (CA)** | Does Go produce same output as Python? | Run both with empty stdin, compare stdout (strip + exact match) |
| **I/O Equivalence Rate** | Same as CA in this experiment | Run both programs, compare stdout. Equivalent to CA when there is no separate ground truth |
| **Test Pass Rate** | Do Go tests pass? | Run `go test -v` on translated test files; average of (passed / total) across files with tests |

### Why CA and I/O Equivalence Are the Same Here

In the literature, CA compares against **ground truth Go code**, while I/O equivalence
compares against the **original Python code**. Since this experiment has no ground truth,
both metrics compare against Python source output. They will diverge in future experiments
when ground truth data (e.g., from CodeNet parallel corpus) is introduced.

### Metrics Display

Results are rendered as two Rich tables in the terminal:

1. **Per-file table**: Shows each file with Y/N for each metric plus test counts.
2. **Summary table**: Shows aggregate percentages across all files.

## File Structure

```
00_get_hands_on/
├── README.md          # This document
├── __init__.py        # Package init (empty)
├── __main__.py        # Enables: python -c "import importlib; ..."
├── models.py          # Pydantic schemas for structured LLM output and records
├── tools.py           # @tool-decorated functions for code execution (subprocess)
├── agents.py          # Agent definition (Translator only)
├── metrics.py         # Metric aggregation and Rich table rendering
└── run.py             # Entry point script
```

### File Responsibilities

**`models.py`** — Defines all data structures:
- `TranslationResult`: LLM returns `go_code` + `explanation`.
- `TestGenerationResult`: LLM returns `test_code` + `explanation`.
- `TestTranslationResult`: LLM returns Go `test_code` + `explanation`.
- `EvaluationRecord`: Per-file result with booleans for each metric, test counts, and notes.

**`tools.py`** — Subprocess wrappers exposed as Agno tools (`@tool()` decorator):
- `compile_go_code`: Writes code to temp dir, runs `go build`, returns success/error.
- `run_go_code` / `run_python_code`: Runs code via `go run` / `python3` with stdin, captures stdout/stderr.
- `compare_outputs`: Strips and compares two stdout strings, reports first diff line.
- `run_python_tests`: Runs pytest in a temp dir, parses PASSED/FAILED counts.
- `run_go_tests`: Inits a Go module in temp dir, runs `go test -v`, parses PASS/FAIL counts.

All tools use 30-second timeouts and clean up temp directories automatically.

**`agents.py`** — Creates the Translator agent with its model, instructions, and output
schema. Uses a `_get_model()` factory so the LLM configuration is defined once.

**`metrics.py`** — Pure computation + display:
- `compute_summary()`: Aggregates `EvaluationRecord` list into percentage metrics.
  Test pass rate is averaged only over files that have tests (avoids dilution by files
  without tests).
- `display_summary_table()`: Rich table with metric names and values.
- `display_per_file_table()`: Rich table with per-file Y/N and test counts.

**`run.py`** — The entry point that orchestrates everything:
1. Loads `.env` for the `MINIMAX_API_KEY`.
2. Discovers `.py` files (excluding `test_*.py` and `*_test.py`).
3. Creates the Translator agent **once** (never inside the loop).
4. Sends each file to the Translator directly (1 LLM call per file).
5. Saves `.go` files to the target directory.
6. Displays a summary of translated files.

Evaluation (`evaluate_file()`) is available in `run.py` for manual use but is not
called automatically during translation.

## Usage

```bash
# From the project root (experiments/):

# 1. Run tests first (always do this before running the pipeline)
uv run pytest src/lab/00_get_hands_on/tests/ -v

# 2. Run only preflight tests (API connection, Go compiler, env vars)
uv run pytest src/lab/00_get_hands_on/tests/test_preflight.py -v

# 3. Run the translation pipeline
uv run python src/lab/00_get_hands_on/run.py

# 4. Skip preflight check if you already verified the API
uv run python src/lab/00_get_hands_on/run.py --skip-preflight

# Custom source/target directories
uv run python src/lab/00_get_hands_on/run.py --source-dir /path/to/python/code
uv run python src/lab/00_get_hands_on/run.py --target-dir /path/to/output
```

The pipeline runs a **preflight check** before processing files. It verifies:
1. `MINIMAX_API_KEY` is set and non-empty.
2. `go` compiler is on PATH.
3. The MiniMax API accepts the key and returns a response.

If any check fails, the pipeline terminates immediately with a clear error message
instead of producing empty results.

## Tests

Tests use **pytest** and are organized by scope:

```
tests/
├── __init__.py
├── test_preflight.py    # Environment + API connection (requires live API)
├── test_tools.py        # All 6 evaluation tools (offline, uses subprocess)
├── test_metrics.py      # Metric computation + Rich table rendering
└── test_run.py          # File discovery, path mirroring, evaluate_file()
```

| Test file | What it covers | Requires API? |
|-----------|----------------|---------------|
| `test_preflight.py` | Env vars, Go/Python availability, MiniMax API connection | Yes |
| `test_tools.py` | compile_go_code, run_go_code, run_python_code, compare_outputs, run_go_tests, run_python_tests | No |
| `test_metrics.py` | compute_summary, display_summary_table, display_per_file_table | No |
| `test_run.py` | discover_python_files, find_test_file, mirror_path, evaluate_file (with Go compilation + I/O + tests) | No |

Run offline tests only (fast, no API calls):
```bash
uv run pytest src/lab/00_get_hands_on/tests/ -v --ignore=src/lab/00_get_hands_on/tests/test_preflight.py
```

### Prerequisites

- **Python 3.13+** with `uv` package manager.
- **Go** installed and on PATH (used for compilation and execution).
- **MINIMAX_API_KEY** set in `.env` at the project root.
- Dependencies: `agno`, `litellm`, `python-dotenv`, `rich`, `pydantic`, `pytest` (all
  in `pyproject.toml`).

## Technical Note: importlib

The directory name `00_get_hands_on` starts with a digit, which is invalid in Python's
import syntax (`from src.lab.00_get_hands_on...` is a `SyntaxError`). All cross-module
references use `importlib.import_module("src.lab.00_get_hands_on.<module>")` instead.
This is a deliberate trade-off to preserve the numbered experiment naming convention
(`00_` through `99_`) described in the project's CLAUDE.md.

## LLM Configuration

All agents use **MiniMax-M2.5** via a custom MiniMax model class:
```python
MiniMax(id="MiniMax-M2.5")
```
To switch models, modify the `_get_model()` function in `agents.py`.

## Known Limitations and Future Work

- **No ground truth**: CA and I/O equivalence are currently identical. Introducing
  CodeNet parallel corpus (`src/data/RAG/processed/parallel_corpus/codeNet/python_go_pairs.jsonl`,
  1,668 pairs) as ground truth would differentiate them.
- **Empty stdin only**: Programs that require stdin input will fail the I/O equivalence
  check. Future work: parse test cases from problem descriptions or provide sample inputs.
- **No retry loop**: If translation fails or produces non-compilable code, the pipeline
  records the failure but does not retry. A feedback loop where evaluation findings
  are sent back to the Translator for correction would improve success rates.
- **Single-file scope**: Each file is translated independently. Cross-file dependencies
  (imports between modules) are not handled. Project-level translation is a separate
  challenge.
- **Evaluation is manual**: The `evaluate_file()` function exists in `run.py` but is not
  called automatically. A separate evaluation step can be added in future experiments.
