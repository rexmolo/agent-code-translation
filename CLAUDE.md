# CLAUDE.md

## Project Overview
- This is a thesis experiment project which is about to build a agent code translation system. 
- It focuses on the code translation Python to Go.
- All docs in the docs/ folder contain the documentation of the project


## Project Structure
- src/scripts/ contains the scripts that are used for data processing, for example, extract data and other trivial things.
- src/data/RAG/processed/ contains the processed data that will be used for embedding for RAG.
- src/lab/ contains each experiment. Each experiment has its own folder. It will start with "00_" to "99_" to indicate the order of experiments.
- src/temp/ contains the temporary files.
- data/translation/ contains two folders: source and target. The source is the source code, and the target is the translated code from our system.

## Tools
- Agno: the core framework of the system.
- LLMS: MINIMAX2.5
- Python package management with UV
- tree-sitter for code parsing
