# Agent Code Translation

An agent-based code translation system that translates Python to Go, built for thesis experiments. Uses [Agno](https://github.com/agno-agi/agno) as the agent framework with multi-provider LLM support.

## Overview

The system translates Python source code into idiomatic Go, then evaluates the results using compilation checks, output comparison, and test execution. The active thesis workflow uses **HumanEval-X**, evaluated inside Docker containers with `go test` and testify assertions.

Multiple LLM providers are supported through a registry (Google Gemini, MiniMax, OpenAI), each with several model variants.

## Requirements

- Python 3.13+
- [uv](https://docs.astral.sh/uv/) for package management
- Docker (for HumanEval-X evaluation)
- ChromaDB Docker container (for RAG vector store)

## Setup

```bash
# Install dependencies
uv sync

# Configure provider credentials
cp config/providers.yaml.example config/providers.yaml
# Then edit config/providers.yaml with your API keys
```

Credentials are managed in `config/providers.yaml`. Each provider supports two methods — set the key directly (`api_key`) or point to an environment variable (`api_key_env`):

| Provider | Credential | Notes |
|---|---|---|
| MiniMax | `api_key` or `MINIMAX_API_KEY` env | Required for MiniMax models |
| Google Gemini | `api_key` / `GOOGLE_API_KEY` env, or Vertex AI mode | See Vertex AI section below |
| OpenAI | `api_key` or `OPENAI_API_KEY` env | Required for OpenAI/GPT models and OpenAI RAG embeddings |

**Vertex AI mode** (for Google Gemini): If using Vertex AI instead of a standard API key, set the following environment variables and authenticate via `gcloud auth application-default login`:

| Variable | Description |
|---|---|
| `GOOGLE_GENAI_USE_VERTEXAI` | Set to `true` to activate Vertex AI mode |
| `GOOGLE_CLOUD_PROJECT` | Your GCP project ID |
| `GOOGLE_CLOUD_LOCATION` | Region (default: `us-central1`) |

## Usage

### Interactive mode (recommended)

```bash
uv run python -m src.cli
```

Walks through dataset selection, action, sample size, and model choice with arrow-key menus.

### Subcommand mode

```bash
# Translate with defaults (baseline experiment)
uv run python -m src.cli translate

# Translate HumanEval-X with a specific provider/model and experiment name
uv run python -m src.cli translate -d humaneval-x -p minimax -v M2.5 -e rag -n 10 --skip-preflight

# Translate first 10 items only
uv run python -m src.cli translate -d humaneval-x -p gemini -v 2.5_pro -n 10

# Evaluate existing translations
uv run python -m src.cli evaluate -d humaneval-x --target-dir data/translation/target/humaneval-x/gemini/2.5_pro/baseline

# Evaluate with custom parallelism (default batch size from config/eval_config.yaml)
uv run python -m src.cli evaluate -d humaneval-x --target-dir data/translation/target/humaneval-x/minimax/M2.5/baseline -b 20
```

Subcommand options for `evaluate`:

| Option | Description |
|---|---|
| `-d`, `--dataset` | Dataset: `humaneval-x` (active workflow; `local` is legacy) |
| `--source-dir` | Override source directory |
| `--target-dir` | Path to translated output folder to evaluate |
| `-b`, `--batch-size` | Number of parallel Docker evaluations (overrides `config/eval_config.yaml`) |
| `-V`, `--verbose` | Enable verbose step-by-step logging |

Subcommand options for `translate`:

| Option | Description |
|---|---|
| `-d`, `--dataset` | Dataset: `humaneval-x` (active workflow; `local` is legacy) |
| `-p`, `--provider` | Provider key (e.g. `minimax`, `gemini`) |
| `-v`, `--variant` | Model variant key (e.g. `M2.5`, `2.5_pro`) |
| `-e`, `--experiment` | Experiment name: `baseline`, `rag-pattern-only`, `rag-pattern-samples`, `rag-pattern-api-docs`, `rag-full` (default: `baseline`) |
| `--embedding-backend` | Embedding backend for RAG: `chromadb` (default) or `gemini` (Vertex AI Vector Search) |
| `-n`, `--sample` | Translate only the first N items |
| `--run` | Run number for multi-run experiments (e.g. `--run 1`) |
| `--skip-preflight` | Skip API/environment checks |

### Automated Batch Runner

For multi-run experiments (statistical significance testing), the batch runner translates and evaluates all experiments automatically with rate limit management:

```bash
# Preview the schedule without running:
uv run python src/scripts/run_all_batches.py \
  --provider minimax --variant M2.5 \
  --dimensions 768,1536,3072 --runs 5 \
  --embedding-backend chromadb --no-baseline --dry-run

# Run in tmux (recommended — runs unattended for ~35 hours):
tmux new -s experiments
uv run python src/scripts/run_all_batches.py \
  --provider minimax --variant M2.5 \
  --dimensions 768,1536,3072 --runs 5 \
  --embedding-backend chromadb --no-baseline \
  --batch-size 9 --window-hours 5 --delay 3
# Ctrl+B, D to detach; tmux attach -t experiments to reconnect
```

Each batch cycle: translates 9 experiments (sequential, 3s delay) → evaluates them locally via Docker → sleeps until the rate limit window resets. State is saved to `batch_state.json` so it resumes from where it left off if interrupted.

### Statistical Analysis

After multi-run experiments complete, analyze significance across embedding dimensions:

```bash
# ANOVA + pairwise t-tests on Pass@1
uv run python src/scripts/analyze_statistics.py

# Analyze Compilation@1 instead
uv run python src/scripts/analyze_statistics.py --metric compilation_at_1
```

### Batch Evaluation

A headless evaluation script discovers and evaluates all HumanEval-X experiments at once, generating a Markdown comparison table and bar chart:

```bash
uv run python src/scripts/ci_evaluate_all.py
uv run python src/scripts/ci_evaluate_all.py --batch-size 3
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
| `grammar_mappings` | ~45 | Python-Go structural syntax patterns & paradigms |
| `api_mappings` | ~190 | Python → Go API equivalences (e.g., `json.loads()` → `json.Unmarshal()`) |
| `go_docs` | ~165 | Go standard library docs + error handling & API sequence patterns |

### RAG Pipeline

```
Python code
  → tree-sitter extracts API calls + detects try/except
  → query api_mappings (Hybrid) with extracted API names
  → query go_docs (Hybrid) with matched Go APIs + error patterns
  → query grammar_mappings (Dense) using raw Python code to find structural idiom mappings
  → formatted context → LLM prompt
```

Retrieval uses **True Hybrid Search** (BM25 exact keyword matching + Dense embeddings for semantic similarity, merged via Reciprocal Rank Fusion) for API mappings and Go docs. The `grammar_mappings` operate via pure **Dense Search** to structurally match Python patterns based on context.

### Embedding Backends

The system supports two embedding backends for comparing embedding model performance:

| Backend | Embedding Model | Vector Store | Dimensions |
|---|---|---|---|
| `chromadb` (default) | all-MiniLM-L6-v2 (local) or OpenAI | ChromaDB (local Docker) | 384 / 3072 |
| `gemini` | Gemini Embedding 001 | Vertex AI Vector Search (Google Cloud) | 3072 |

Both backends use the same hybrid retrieval strategy (BM25 + dense + RRF). Select the backend via the interactive CLI or the `--embedding-backend` flag:

```bash
uv run python -m src.cli translate -d humaneval-x -p gemini -v 2.5_pro -e rag --embedding-backend gemini -n 10
```

### Ablation Experiments

The system supports additive ablation experiments to measure the contribution of each RAG knowledge base component. The experiment name controls which knowledge bases are active — no manual config editing needed.

| Experiment | Grammar | Parallel Corpus | API Mappings | Go Docs |
|---|---|---|---|---|
| `baseline` | — | — | — | — |
| `rag-pattern-only` | ON | OFF | OFF | OFF |
| `rag-pattern-samples` | ON | ON | OFF | OFF |
| `rag-pattern-api-docs` | ON | OFF | ON | ON |
| `rag-full` | ON | ON | ON | ON |

Select an experiment via the interactive CLI or the `-e` flag:

```bash
uv run python -m src.cli translate -d humaneval-x -p minimax -v M2.5 -e rag-full -n 10
```

For multi-run experiments, add the `--run` flag:

```bash
uv run python -m src.cli translate -d humaneval-x -p minimax -v M2.5 -e rag-full --run 1
```

Before translation or evaluation, the active KB configuration is displayed:

```
── Model: minimax/M2.5 ──
   Experiment: rag-full
   Run: 1
   RAG: Grammar Patterns: ON | Parallel Corpus: ON | API Mappings: ON | Go Docs: ON
   Embedding: ChromaDB
   Output: data/translation/target/humaneval-x/minimax/M2.5/vec-chroma-768/run-1/rag-full
   Parallel batch size: 5
```

### Setup & Configuration

ChromaDB runs as a Docker container. Ensure it's running before ingesting or querying:

```bash
# Example: start ChromaDB via docker-compose (from your Docker project)
docker compose up -d chromadb
```

Pipeline settings (parallelism, Docker timeout) are in `config/eval_config.yaml`:

```yaml
translation:
  batch_size: 5     # Parallel LLM requests for translation

parallel:
  batch_size: 10    # Concurrent Docker containers for evaluation

docker:
  image: "golang:1.26-alpine"
  memory_limit: "512m"
  timeout: 60       # Per-container timeout in seconds
```

Both translation and evaluation run in parallel using `ThreadPoolExecutor`. Each translation thread creates its own agent instance for thread safety.

Provider credentials are in `config/providers.yaml` (see [Setup](#setup)). This file is gitignored; copy from `config/providers.yaml.example`.

Connection, embedding, and knowledge base settings are in `config/rag_config.yaml`:

```yaml
chromadb:
  host: "localhost"
  port: 8000

embedding:
  provider: "default"   # "default" (free, local), "openai", or "gemini"

knowledge_bases:
  code_snippets: true    # Grammar mappings (Python-Go structural syntax patterns)
  api_mappings: true     # Python -> Go API equivalences
  documentation: true    # Go standard library docs & patterns
```

The `knowledge_bases` toggles serve as a manual fallback for custom experiment names. For the built-in experiment presets, toggles are applied automatically (see [Ablation Experiments](#ablation-experiments)).

**Option 1: Free local embeddings (default, no API key needed)**

Uses ChromaDB's built-in `all-MiniLM-L6-v2` model (384 dims, runs via ONNX locally). Good for testing the pipeline.

```bash
# Ingest all data into ChromaDB (uses free local model)
uv run python src/scripts/ingest_rag.py
```

**Option 2: OpenAI embeddings (higher quality)**

Uses `text-embedding-3-large` (3072 dims). Better retrieval quality for code.

```bash
# 1. Set your API key in config/providers.yaml under openai.api_key

# 2. Switch provider in config/rag_config.yaml
#    provider: "openai"

# 3. Re-ingest (old collections with different dimensions will be overwritten)
uv run python src/scripts/ingest_rag.py
```

**Option 3: Gemini embeddings + Vertex AI Vector Search**

Uses Google's `gemini-embedding-001` model (3072 dims) with Vertex AI Vector Search as the vector store. Requires GCP project and Vertex AI access.

```bash
# 1. Ensure Vertex AI credentials are set (see Setup section)

# 2. Ingest data into Vertex AI (first run creates index + endpoint, ~20-30 min)
uv run python src/scripts/ingest_rag_gemini.py

# 3. Use --embedding-backend gemini when translating
uv run python -m src.cli translate -e rag --embedding-backend gemini
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
├── config.py             # Shared path constants and config loaders
├── core/                 # Core logic
│   ├── agents.py         # Agno translation agent definition
│   ├── docker_eval.py    # Docker-based HumanEval-X evaluation
│   ├── error_db.py       # SQLite persistence for evaluation errors (thread-safe)
│   ├── evaluation.py     # File discovery and HumanEval-X evaluation helpers
│   ├── logger.py         # Verbose pipeline logger with Rich (thread-safe)
│   ├── pipeline.py       # High-level orchestration (translate / evaluate, parallel)
│   ├── reporting.py      # Rich summary tables and metric computation
│   ├── schemas.py        # Pydantic data models (TranslationResult, EvaluationRecord, etc.)
│   └── tools.py          # Agno @tool functions (compile, run, compare)
├── data/                 # Dataset loaders
│   └── humaneval_x.py    # HumanEval-X loader from HuggingFace datasets
├── providers/            # LLM provider adapters
│   ├── minimax/          # MiniMax (Anthropic-compatible API)
│   ├── openai/           # OpenAI GPT (via Agno OpenAIChat)
│   └── registry.py       # Multi-provider model registry with lazy factories
├── rag/                  # RAG retrieval system
│   ├── api_extractor.py  # Tree-sitter Python API/call extraction
│   ├── embeddings.py     # Embedding function factory (default, OpenAI, Gemini)
│   ├── retriever.py      # Hybrid retrieval (BM25 + dense + RRF), ChromaDB & Vertex AI backends
│   ├── store.py          # ChromaDB HttpClient & collection management
│   └── vertex_store.py   # Vertex AI Vector Search resource management
├── scripts/              # One-off data processing scripts
│   ├── analyze_statistics.py      # ANOVA/t-test analysis of multi-run results
│   ├── extract_codenet_data.py    # Extract CodeNet parallel corpus
│   ├── generate_api_mappings.py   # Generate Python→Go API mappings
│   ├── generate_go_docs.py        # Generate Go std library docs
│   ├── ingest_rag.py             # Ingest JSONL data into ChromaDB
│   ├── ingest_rag_gemini.py      # Ingest JSONL data into Vertex AI Vector Search
│   └── run_all_batches.py        # Automated batch runner with rate limiting
└── tests/                # All tests

data/
├── RAG/
│   ├── processed/                 # JSONL data for RAG
│   │   ├── api_mappings.jsonl     # Python→Go API mappings (~190 entries)
│   │   ├── go_docs.jsonl          # Go docs + patterns (~165 entries)
│   │   └── grammar_mappings.jsonl # Python-Go structural patterns (~45 entries)
└── translation/
    └── target/           # Translated Go output (gitignored)
        ├── humaneval-x/<provider>/<variant>/baseline/                       # Baseline (flat)
        └── humaneval-x/<provider>/<variant>/<backend>/run-<N>/<experiment>/ # RAG with multi-run
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
2. **KB configuration** — apply knowledge base toggles based on experiment name and display active status
3. **File discovery** — load HumanEval-X problems from HuggingFace
4. **Parallel translation** — files are translated concurrently using a thread pool (batch size from `config/eval_config.yaml`). Each thread:
   - Retrieves RAG context (API mappings, documentation, code snippets) based on active KB toggles
   - Creates its own Agno agent instance and sends the prompt to the LLM
   - Receives a `TranslationResult` with structured Go output
5. **Output storage** — translated Go files are saved to `data/translation/target/<dataset>/<provider>/<variant>/<experiment>/`

### Evaluation pipeline

Evaluation differs by dataset because the correctness signals differ:

**HumanEval-X** (Docker-sandboxed, parallel):
1. Build `solution.go` from the LLM output and the provided Go signature
2. Build `solution_test.go` from HumanEval-X test harness (add package header, detect needed imports like testify)
3. Run in parallel batches (configurable via `config/eval_config.yaml` or `-b` flag) inside Docker (`golang:1.26-alpine`) with `--network=none` and `--memory=512m`:
   - `go vet ./...` for compile check
   - `go test -v -count=1 ./...` for test execution
4. Aggregate results

Docker isolation is necessary for HumanEval-X because we run untrusted LLM-generated code. A shared Docker volume caches Go modules (testify) so individual test runs don't need network access.

### Why this evaluation path?

HumanEval-X provides proper test suites using `testing` + `testify/assert`, so evaluation uses `go test`. These tests exercise function signatures and edge cases, giving a more rigorous correctness signal. They require Docker for safe execution and dependency management (testify).

### Multi-provider model registry

The registry (`providers/registry.py`) uses a **lazy factory pattern**: model constructors are registered as closures but only instantiated when `get_enabled_models()` is called. This avoids importing provider SDKs at startup and allows enabling/disabling models at runtime through the CLI.

```
providers/
├── registry.py          # register(provider, variant, factory)
│                        # enable_model() / get_enabled_models()
├── minimax/
│   └── minimax.py       # Custom Anthropic-compatible adapter
└── openai/
    └── openai.py        # GPT wrapper (extends Agno OpenAIChat)
```

Currently supported:
- **Google Gemini** — 6 variants (2.5 Flash Lite/Flash/Pro, 3 Flash/Pro Preview, 3.1 Pro Preview)
- **MiniMax** — 3 variants (M2, M2.1, M2.5) via Anthropic-compatible API
- **OpenAI** — GPT-5.4 via Agno's built-in OpenAIChat

Adding a new provider requires only registering factory functions — no changes to the translation or evaluation logic.

## Evaluation Metrics

| Metric | Datasets | Description |
|---|---|---|
| Compilation@1 | HumanEval-X | Fraction of translations that compile successfully on first attempt |
| Runs Successfully | HumanEval-X | Not used in the active HumanEval-X-only workflow |
| Pass@1 | HumanEval-X | Fraction that pass all tests on first attempt |
