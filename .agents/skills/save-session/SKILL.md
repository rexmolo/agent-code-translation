---
name: save-session
description: Save current project session progress to experiments/.doc/memory/. Trigger when user invokes /save-session or wants to persist current work state. Creates YYYY-MM-DD folders automatically and names files as 01-title-time.md, 02-title-time.md with proper incrementing.
---

# Save Session Skill

## When to Use
- User invokes `/save-session`
- User asks to "save progress", "save session", "checkpoint current work"
- User wants to persist what was done today

## Workflow

### Step 1: Determine Target Directory
- Base path: `{project_root}/.doc/memory/`
- Today's folder: `{project_root}/.doc/memory/YYYY-MM-DD/`
- Use the actual current local date from the environment or `date +%F`; do not hardcode dates.

### Step 2: Create Date Folder if Needed
- If `YYYY-MM-DD` folder doesn't exist, create it

### Step 3: Determine File Name
- List existing `.md` files in the date folder
- Find the highest sequence number (01, 02, etc.)
- New file: `{next_number}-{slug-from-title}-{HHMMSS}.md`
- If no existing files, start with `01`
- Sequence is per-date folder

### Step 4: Write Content
- Summarize the session for future thesis work after long LLM conversations.
- Keep it concise, factual, and practical. Capture:
  - Design rationale and key decisions
  - Experiment setup, benchmark/task, model/provider, prompt/RAG mode, and relevant config
  - Commands, scripts, evaluations, or runs performed
  - Metrics, results, failures, and where outputs/logs are stored
  - Dataset, embedding, retrieval, and target-output versions or paths used
  - Files modified, created, or intentionally left unchanged
  - Caveats, risks, blocked items, and next steps
  - Thesis-ready takeaways: what this implies for the research question

## File Naming Convention
```
01-agent-translation-103045.md
02-rag-pipeline-fix-104520.md
```
- Format: `{seq:02d}-{slug}-{HHMMSS}.md`
- Slug is derived from session title/topic (kebab-case, max 30 chars)
- Time is current timestamp (24hr, no separators)

## Example File Content
```markdown
---
date: 2026-04-21
time: 10:30:45
---

# Session: Agent Translation System - RAG Pipeline

## Summary
- Goal: evaluate whether the updated RAG context improves HumanEval-X Python-to-Go translation.
- Rationale: keep raw LLM translation outputs unchanged so evaluation reflects model behavior.
- Setup: MiniMax 2.5, HumanEval-X, RAG mode `grammar_mappings`, dataset path `src/data/RAG/processed/...`.

## Runs and Results
- Command: `uv run python -m src.cli evaluate ...`
- Outputs/logs: `data/translation/target/...`
- Metrics: pass@1 changed from X to Y; failures concentrated in parsing edge cases.

## Files and Data
- Modified: `src/scripts/...`
- Dataset/version: `src/data/RAG/processed/...`
- Left unchanged: model outputs before evaluation.

## Caveats and Next Steps
- Caveat: sample size was limited to N tasks.
- Next: run full benchmark and compare against non-RAG baseline.

## Thesis Takeaway
- The result suggests whether retrieved grammar examples help translation correctness under the current setup.
```
