# Agent Code Translation

An agent-based code translation system that translates Python to Go, built for thesis experiments. Uses [Agno](https://github.com/agno-agi/agno) as the agent framework with multi-provider LLM support.

## Overview

The system translates Python source code into idiomatic Go, then evaluates the results using compilation checks, output comparison, and test execution. It supports two datasets:

- **Local** — custom Python source files in `data/translation/source/`
- **HumanEval-X** — the multilingual benchmark, evaluated inside Docker containers with `go test` and testify assertions

Multiple LLM providers are supported through a registry (Google Gemini, MiniMax), each with several model variants.

## Requirements

- Python 3.13+
- [uv](https://docs.astral.sh/uv/) for package management
- Go compiler (for local evaluation)
- Docker (for HumanEval-X evaluation)

## Setup

```bash
# Install dependencies
uv sync

# Configure environment variables
cp .env.example .env   # then fill in your API keys
```

Required environment variables (depending on provider):

| Variable | Provider | Required |
|---|---|---|
| `MINIMAX_API_KEY` | MiniMax | For MiniMax models |
| `GOOGLE_API_KEY` | Google Gemini | For Gemini models |
| `OPENAI_API_KEY` | OpenAI | Only if using OpenAI embeddings for RAG (optional) |

## Usage

### Interactive mode (recommended)

```bash
uv run python -m src.cli
```

Walks through dataset selection, action, sample size, and model choice with arrow-key menus.

### Subcommand mode

```bash
# Translate with defaults
uv run python -m src.cli translate

# Translate HumanEval-X dataset, skip preflight
uv run python -m src.cli translate -d humaneval-x --skip-preflight

# Evaluate existing translations
uv run python -m src.cli evaluate -d humaneval-x --target-dir data/translation/target/humaneval-x/gemini/2.5_flash
```

### Run tests

```bash
# All tests
uv run pytest src/tests/ -v

# Specific test file
uv run pytest src/tests/test_run.py -v

# Skip Docker-dependent integration tests
uv run pytest src/tests/ -v -m "not integration"
```

## RAG (Retrieval-Augmented Generation)

The system uses RAG to provide the LLM with relevant context before translation: similar code examples, Python-to-Go API mappings, Go documentation, and error handling patterns.

### RAG Data Sources

| Collection | Entries | Description |
|---|---|---|
| `parallel_corpus` | 1,668 | Python-Go code pairs from IBM CodeNet (few-shot examples) |
| `api_mappings` | 180 | Python → Go API equivalences (e.g., `json.loads()` → `json.Unmarshal()`) |
| `go_docs` | 157 | Go standard library docs + error handling & API sequence patterns |

### RAG Pipeline

```
Python code
  → tree-sitter extracts API calls + detects try/except
  → query api_mappings with extracted API names
  → query go_docs with matched Go APIs + error patterns
  → query parallel_corpus for similar full examples
  → formatted context → LLM prompt
```

Retrieval uses **Hybrid Search**: BM25 (exact keyword matching, good for API names) + Dense embeddings (semantic similarity), merged with Reciprocal Rank Fusion.

### Setup & Configuration

The embedding model is configured in `src/rag/rag_config.yaml`:

```yaml
embedding:
  provider: "default"   # "default" (free, local) or "openai" (requires API key)
```

**Option 1: Free local embeddings (default, no API key needed)**

Uses ChromaDB's built-in `all-MiniLM-L6-v2` model (384 dims, runs via ONNX locally). Good for testing the pipeline.

```bash
# Ingest all data into ChromaDB (uses free local model)
uv run python src/scripts/ingest_rag.py
```

**Option 2: OpenAI embeddings (higher quality)**

Uses `text-embedding-3-large` (3072 dims). Better retrieval quality for code.

```bash
# 1. Set your API key in .env
#    OPENAI_API_KEY=sk-...

# 2. Switch provider in src/rag/rag_config.yaml
#    provider: "openai"

# 3. Clear old embeddings (different dimensions) and re-ingest
rm -rf data/RAG/chromadb/
uv run python src/scripts/ingest_rag.py
```

### Ingest Options

```bash
# Ingest all collections
uv run python src/scripts/ingest_rag.py

# Ingest a specific collection
uv run python src/scripts/ingest_rag.py --collection parallel_corpus
uv run python src/scripts/ingest_rag.py --collection api_mappings
uv run python src/scripts/ingest_rag.py --collection go_docs
```

### Expanding RAG Data

```bash
# Generate additional Python→Go API mappings (curated, appends to api_mappings.jsonl)
uv run python src/scripts/generate_api_mappings.py

# Generate additional Go standard library docs (curated, appends to go_docs.jsonl)
uv run python src/scripts/generate_go_docs.py

# After expanding data, re-ingest the affected collection(s)
uv run python src/scripts/ingest_rag.py --collection api_mappings
uv run python src/scripts/ingest_rag.py --collection go_docs
```

## Project Structure

```
src/
├── cli/                  # Click + Questionary interactive CLI
│   ├── __init__.py       # CLI entry point, interactive & subcommand modes
│   └── __main__.py       # python -m src.cli support
├── config.py             # Shared path constants
├── core/                 # Core logic
│   ├── agents.py         # Agno translation agent definition
│   ├── docker_eval.py    # Docker-based HumanEval-X evaluation
│   ├── evaluation.py     # File discovery, mirroring, local Go evaluation
│   ├── pipeline.py       # High-level orchestration (translate / evaluate)
│   ├── reporting.py      # Rich summary tables and metric computation
│   ├── schemas.py        # Pydantic data models (TranslationResult, EvaluationRecord, etc.)
│   └── tools.py          # Agno @tool functions (compile, run, compare)
├── data/                 # Dataset loaders
│   └── humaneval_x.py    # HumanEval-X loader from HuggingFace datasets
├── providers/            # LLM provider adapters
│   ├── minimax/          # MiniMax (Anthropic-compatible API)
│   └── registry.py       # Multi-provider model registry with lazy factories
├── rag/                  # RAG retrieval system
│   ├── api_extractor.py  # Tree-sitter Python API/call extraction
│   ├── embeddings.py     # Embedding function factory (reads rag_config.yaml)
│   ├── rag_config.yaml   # Embedding provider & retrieval settings
│   ├── retriever.py      # Hybrid retrieval (BM25 + dense + RRF)
│   └── store.py          # ChromaDB client & collection management
├── scripts/              # One-off data processing scripts
│   ├── extract_codenet_data.py    # Extract CodeNet parallel corpus
│   ├── generate_api_mappings.py   # Generate Python→Go API mappings
│   ├── generate_go_docs.py        # Generate Go std library docs
│   └── ingest_rag.py             # Ingest JSONL data into ChromaDB
└── tests/                # All tests (64 test cases)

data/
├── RAG/
│   ├── processed/                 # JSONL data for RAG
│   │   ├── api_mappings.jsonl     # Python→Go API mappings (180 entries)
│   │   ├── go_docs.jsonl          # Go docs + patterns (157 entries)
│   │   └── parallel_corpus/       # CodeNet Python-Go pairs (1,668 entries)
│   └── chromadb/                  # ChromaDB persistent storage (gitignored)
└── translation/
    ├── source/           # Python source files (local dataset)
    └── target/           # Translated Go output (provider/variant subdirs)
```

## Design

### Why agent-based translation?

Traditional rule-based transpilers (e.g. c2go, py2many) work at the syntax level — they map language constructs one-to-one. This produces correct but unidiomatic output and struggles with language-specific patterns (Python's duck typing vs Go's static interfaces, error handling conventions, goroutines vs asyncio). LLMs understand both languages semantically and can produce idiomatic target code, but raw LLM output is unpredictable in structure.

We use the Agno agent framework to solve this: the translation agent enforces a **structured output schema** (`TranslationResult`) so every LLM response is guaranteed to contain a `go_code` field with compilable Go code, plus an optional `explanation`. This eliminates post-processing and makes the pipeline fully automated.

### Architecture flow

```
                         ┌─────────────┐
                         │   CLI       │  Interactive (Questionary) or
                         │  src/cli/   │  subcommand (Click)
                         └──────┬──────┘
                                │
                                ▼
                      ┌──────────────────┐
                      │    Pipeline      │  Orchestration layer
                      │ core/pipeline.py │  (translate / evaluate)
                      └───────┬──┬───────┘
                 ┌────────────┘  └────────────┐
                 ▼                             ▼
        ┌─────────────────┐          ┌─────────────────┐
        │   Translation   │          │   Evaluation     │
        │  core/agents.py │          │ core/evaluation.py│
        │                 │          │ core/docker_eval.py│
        └────────┬────────┘          └────────┬─────────┘
                 │                             │
                 ▼                             ▼
        ┌─────────────────┐          ┌─────────────────┐
        │   LLM Provider  │          │  Go Toolchain   │
        │  providers/     │          │  go build/run/  │
        │  registry.py    │          │  test (+ Docker)│
        └─────────────────┘          └─────────────────┘
```

### Translation pipeline

1. **Preflight checks** — verify API keys, Go compiler, and LLM connectivity before starting
2. **File discovery** — find Python source files (local) or load HumanEval-X problems from HuggingFace
3. **Agent translation** — for each file, the Agno agent sends the Python code to the selected LLM and receives a `TranslationResult` with structured Go output
4. **Output storage** — translated Go files are saved to `data/translation/target/<dataset>/<provider>/<variant>/`

### Evaluation pipeline

Evaluation differs by dataset because the correctness signals differ:

**Local dataset** (direct Go toolchain):
1. `go build` — does it compile?
2. `go run` — does it execute without runtime errors?
3. **Output comparison** — run both the Python source and Go translation, compare stdout. If a Go test file exists, run `go test` instead
4. Aggregate results into per-file and summary tables

**HumanEval-X** (Docker-sandboxed):
1. Build `solution.go` from LLM output (strip markdown fences, extract declarations, reconstruct package/imports)
2. Build `solution_test.go` from HumanEval-X test harness (add package header, detect needed imports like testify)
3. Run inside Docker (`golang:1.26-alpine`) with `--network=none` and `--memory=512m`:
   - `go vet ./...` for compile check
   - `go test -v -count=1 ./...` for test execution
4. Aggregate results

Docker isolation is necessary for HumanEval-X because we run untrusted LLM-generated code. A shared Docker volume caches Go modules (testify) so individual test runs don't need network access.

### Why two evaluation paths?

The local dataset uses simple stdout comparison — the Python and Go programs should produce identical output. This works for standalone programs but not for library functions.

HumanEval-X provides proper test suites using `testing` + `testify/assert`, so evaluation uses `go test`. These tests exercise function signatures and edge cases, giving a more rigorous correctness signal. But they require Docker for safe execution and dependency management (testify).

### Multi-provider model registry

The registry (`providers/registry.py`) uses a **lazy factory pattern**: model constructors are registered as closures but only instantiated when `get_enabled_models()` is called. This avoids importing provider SDKs at startup and allows enabling/disabling models at runtime through the CLI.

```
providers/
├── registry.py          # register(provider, variant, factory)
│                        # enable_model() / get_enabled_models()
└── minimax/
    └── minimax.py       # Custom Anthropic-compatible adapter
```

Currently supported:
- **Google Gemini** — 6 variants (2.5 Flash Lite/Flash/Pro, 3 Flash/Pro Preview, 3.1 Pro Preview)
- **MiniMax** — 3 variants (M2, M2.1, M2.5) via Anthropic-compatible API

Adding a new provider requires only registering factory functions — no changes to the translation or evaluation logic.

## Evaluation Metrics

| Metric | Description |
|---|---|
| Compilation@1 | Fraction of translations that compile successfully on first attempt |
| Runs Rate | Fraction that execute without runtime errors |
| Pass@1 | Fraction that produce correct output or pass all tests on first attempt |
| AST Similarity | Structural similarity between source and target ASTs |
| Test Pass Rate | Fraction of unit tests passed (when test suites are available) |
