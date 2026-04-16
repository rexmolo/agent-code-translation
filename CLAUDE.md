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
- The tasks will be assigned in this location .doc/Tasks/01 - tasks, activate means that task is processing and it will be moved into test after finished. then will be moved into complete after finishing test.


## Thinking
- Think if it will affect the final result whenever you are going to change files and make decitions, because we are evaluating the model output
