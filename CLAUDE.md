# CLAUDE.md

## Project Overview
- This is a thesis experiment project building an agent-based Python to Go translation system.
- The active benchmark workflow is HumanEval-X.
- Read the relevant Markdown files in each folder before working in that area.


## Project Structure
- src/scripts/ contains data processing and experiment orchestration scripts.
- src/data/RAG/processed/ contains the processed data used for RAG embeddings.
- src/lab/ contains experiments. Each experiment has its own folder and uses a numeric prefix for ordering.
- src/temp/ contains temporary files.
- data/translation/ contains translated output artifacts under the HumanEval-X target tree.

## Tools
- Agno: the core framework of the system.
- LLMs: MiniMax 2.5 and other configured providers.
- Python package management with UV.
- tree-sitter for code parsing.

## rules
- don't read the files inside .doc folder except memory.
