---
date: 2026-04-21
time: 11:38:45
---

# Session: Trap-KB Design, Router Iteration, and Held-Out Results

## Summary
- Goal: redesign the Python-to-Go translation-trap knowledge base so it can improve MiniMax M2.5 HumanEval-X translation without post-processing LLM outputs.
- Core decision: separate trap data by lifecycle phase instead of putting all mined failures into one prompt-time RAG file.
- Current thesis framing: CodeNet + the 67 manually labeled bug programs should be treated as the discovery/development set; HumanEval-X should remain the held-out test set.
- Important rule maintained: generated LLM Go code is fed into evaluation as-is. We did not patch generated translations to improve scores.

## Dataset Design
- `translation_traps_prompt_v4.jsonl`: active prompt-time traps only. These are inserted into the translation prompt through deterministic routing from Python source and/or Go signature.
- `translation_traps_validator_v4.jsonl`: post-generation classifiers only. These classify compiler/vet/runtime/test failures after translation; they are not shown to the LLM.
- `translation_trap_candidates_v4.jsonl`: staging area for raw or semi-structured mined hypotheses from CodeNet/manual failures. Candidates are not active until they have evidence and deterministic routing.

## Why Three Files
- Prompt-time traps directly affect model output and therefore must be precise, routeable, and low-noise.
- Validator rules are useful after generation because many failures are only visible in generated Go, such as unused imports, missing helper definitions, fake APIs, third-party imports, and malformed assignments.
- Candidate records preserve mined observations without immediately contaminating the prompt. This avoids repeating the earlier `v1` mistake where a broad semantic trap fired on too many unrelated HumanEval-X tasks.

## Evidence and Promotion Rule
- A candidate needs repeated evidence from CodeNet/manual data, not HumanEval-X, before promotion to prompt-time use.
- Evidence means multiple independent failures with the same failure mode, a deterministic pre-generation signal, a specific correction pattern, and acceptable false-positive risk.
- Suggested thresholds: compile traps need at least 3 independent examples; semantic traps need at least 5; broad semantic traps require manual review.
- Deterministic routing should use AST markers, lexical markers, and Go-signature markers. It should not use held-out task IDs, generated Go, compiler stderr, or LLM judgment as pre-generation routing signals.

## Methodology Decision
- Correct loop:
  1. Translate CodeNet/manual discovery set.
  2. Compile/vet/test outputs.
  3. Classify failures with `validator_v4`.
  4. Aggregate recurring failure modes.
  5. Promote safe candidates into `prompt_v4`.
  6. Re-run CodeNet/manual focused subsets or full discovery set to validate the design.
  7. Freeze the prompt dataset.
  8. Evaluate HumanEval-X once or a small fixed number of times.
- HumanEval-X can be inspected for diagnosis after a run, but should not be used to create or tune final active traps if the thesis needs a clean held-out benchmark.

## Trap Dataset Iterations
- `translation_traps_codenet_v1.jsonl`: first mined trap set. It contained a broad trap, `trap_api_semantic_mismatch_after_replace`, that was too high-recall and low-precision.
- `translation_traps_codenet_v2.jsonl`: split broad semantic traps into narrower categories.
- `translation_traps_codenet_v3.jsonl`: made traps more atomic and added router support for granular markers.
- `v3` remains a useful draft taxonomy, but `v4` should be rebuilt from CodeNet/manual evidence using the three-file design above.

## Router and Pipeline Changes
- Updated `src/rag/trap_router.py` to use the active trap source through `TRANSLATION_TRAPS_FILE`, currently pointing at `translation_traps_codenet_v3.jsonl`.
- Added granular deterministic signals including:
  - `negative_subscript`
  - `multiple_top_level_functions`
  - `print_multiple_args`
  - `float_literal_present`
  - `slice_copy_then_mutation`
  - `dict_order_sensitive_usage`
  - `sorted_call_with_key`
  - `sort_call_with_key`
  - `floor_division_operator`
  - `empty_list_return`
  - `signature_param_name_collision=<package>`
- Tightened matching so traps with multiple populated hint groups require all relevant groups to match. This prevents broad AST-only matches from satisfying a trap that also requires a specific lexical marker, such as distinguishing `[-1]` from `[-2]`.
- Fixed a false-positive import regression: `from typing import List` no longer routes to `trap_unused_imports_are_compile_errors`.
- Removed or reworked non-routeable active traps, including the dead unused-local-variable and no-third-party-module prompt traps.
- Added `rag-traps-codenet-v3` experiment preset and `rule-traps` backend labeling.
- Added `--skip-existing` support to HumanEval-X evaluation so evaluation can run in multiple passes during long translation jobs. A later bug was found and patched so summaries rebuild from per-task `evaluation/result.json` files rather than relying on overwritten `per_task.jsonl`.

## Tests
- Targeted router/prompt/artifact/multi-run tests passed after the router updates:
  - `55 passed`
- Covered regressions include:
  - typing-only import does not trigger unused-import trap
  - `[-2]` routes to second-last trap but not last-element trap
  - signature parameter shadowing routes correctly
  - sorted-with-key and floor-division markers route correctly
  - `math.gcd(` source pattern can route to the fake `math.Gcd` warning

## Experiments and Results
- Baseline reference: `minimax/M2.5`, HumanEval-X, `baseline/run-101`.
  - Compilation@1: `155/164 = 94.51%`
  - Pass@1: `144/164 = 87.80%`
  - Summary path: `data/translation/target/humaneval-x/minimax/M2.5/baseline/run-101/evaluation/results/summary.json`
- First trap run: `rag-traps-codenet-v1`, HumanEval-X, `run-11`.
  - Compilation@1: `150/164 = 91.46%`
  - Pass@1: `139/164 = 84.76%`
  - Result: worse than baseline by 5 pass tasks.
  - Main reason: over-broad `trap_api_semantic_mismatch_after_replace` fired too often and injected irrelevant advice.
  - Regression analysis saved previously at `.doc/memory/2026-04-20/01-trap-run11-regression-220048.md`.
- Failed-subset rerun: `rag-traps-codenet-v3`, HumanEval-X failed subset, `run-123`.
  - Input subset: 25 tasks that failed in the old trap run.
  - Compilation@1: `22/25 = 88.0%`
  - Pass@1: `17/25 = 68.0%`
  - Interpretation: strong targeted improvement over the old failure subset, where these same 25 tasks were all failures.
  - Remaining failed tasks after subset run: `Go_50`, `Go_87`, `Go_93`, `Go_96`, `Go_103`, `Go_104`, `Go_160`, `Go_163`.
- Full v3 held-out run: `rag-traps-codenet-v3`, HumanEval-X, `run-124`.
  - Translation completed: `164/164`.
  - Reliable final metrics computed from all per-task `tasks/*/evaluation/result.json` files:
    - Compilation@1: `152/164 = 92.7%`
    - Pass@1: `142/164 = 86.6%`
    - Compile-only failures: `10`
    - Failed-to-compile: `12`
  - Remaining failed tasks: `Go_7`, `Go_10`, `Go_28`, `Go_30`, `Go_32`, `Go_38`, `Go_51`, `Go_62`, `Go_87`, `Go_94`, `Go_97`, `Go_103`, `Go_104`, `Go_105`, `Go_116`, `Go_117`, `Go_144`, `Go_149`, `Go_151`, `Go_153`, `Go_162`, `Go_163`.
  - Caveat: `run-124/evaluation/results/summary.json` is stale and shows `144` files because evaluation ran in multiple passes before the summary merge bug was fully corrected. The per-task `evaluation/result.json` files cover all `164` tasks and are the reliable source.

## Performance Trajectory
- `v1` trap RAG harmed performance: `139/164` pass vs baseline `144/164`.
- `v3` improved the old failure subset substantially: `17/25` previously failed tasks passed after redesign.
- Full `v3` still did not beat baseline: `142/164` pass vs baseline `144/164`.
- The direction is promising but not finalized. The design reduced earlier failure modes but introduced or failed to prevent some remaining compile and semantic errors.

## Remaining Failure Profile from Run-124
- Total failed tasks: `22`.
- Failure categories from per-task results:
  - wrong output or runtime/test failure: `10`
  - unused import compile failure: `5`
  - undefined identifier compile failure: `4`
  - assignment mismatch compile failure: `2`
  - type mismatch compile failure: `1`
- Traps present among failures included:
  - `trap_missing_helper_function_definition`
  - `trap_sorted_with_key_needs_stable_sort`
  - `trap_int_to_string_requires_strconv`
  - `trap_atoi_requires_value_and_error`
  - `trap_mixed_int_float_requires_explicit_widening`
  - `trap_unused_imports_are_compile_errors`
- Thesis interpretation: some prompt-time traps are still too generic or not actionable enough, especially conversion and sorting traps. Some compile failures are post-generation hygiene and belong in validator/candidate analysis, not active prompt RAG.

## Files Modified or Created
- Updated:
  - `src/config.py`
  - `src/rag/trap_router.py`
  - `src/rag/retriever.py`
  - `src/core/pipeline.py`
  - `src/cli/__init__.py`
  - `src/tests/test_trap_router.py`
  - `src/tests/test_multi_run.py`
  - `.agents/skills/save-session/SKILL.md`
- Created/updated trap datasets:
  - `data/RAG/processed/translation_traps_codenet_v2.jsonl`
  - `data/RAG/processed/translation_traps_codenet_v3.jsonl`
- Important artifacts:
  - `data/translation/target/humaneval-x/minimax/M2.5/rule-traps/run-123/rag-traps-codenet-v3`
  - `data/translation/target/humaneval-x/minimax/M2.5/rule-traps/run-124/rag-traps-codenet-v3`

## Save-Session Skill Update
- Updated project skill `.agents/skills/save-session/SKILL.md`.
- The skill now explicitly instructs future agents to save thesis-ready session context: rationale, setup, runs, commands, metrics, dataset versions, modified files, caveats, and research takeaways.
- Removed stale hardcoded date guidance and clarified that saves belong under `.doc/memory/YYYY-MM-DD`.

## Caveats
- HumanEval-X was used during this conversation to diagnose and validate router behavior. For final thesis claims, active `v4` prompt traps should be frozen from CodeNet/manual evidence before the held-out HumanEval-X evaluation.
- `run-124` confirms `v3` is better than the initial harmful `v1` design on the earlier failure subset, but the full held-out result is still slightly below baseline.
- Because full translation/evaluation runs are time-consuming, no additional full CodeNet or HumanEval-X reruns were performed after the final `v4` design discussion.

## Thesis-Ready Takeaway
- A small, high-precision translation-trap KB is more defensible than broad semantic RAG for Python-to-Go translation.
- The central design lesson is phase separation: use CodeNet/manual failures to mine and classify errors, promote only routeable traps into prompt-time RAG, and reserve HumanEval-X for held-out evaluation.
- The results show that broad traps can degrade performance, while granular deterministic routing can recover many previously failed tasks. However, performance gains are not guaranteed until the dataset is rebuilt and validated through the CodeNet/manual discovery loop.

## Recommended Next Steps
- Build `translation_traps_validator_v4.jsonl` from CodeNet/manual compiler, vet, runtime, and sample-I/O failures.
- Build `translation_trap_candidates_v4.jsonl` from recurring classified failures with evidence counts and possible deterministic routes.
- Promote only safe candidates into `translation_traps_prompt_v4.jsonl`.
- Validate `prompt_v4` on CodeNet/manual focused subsets first, then full CodeNet if feasible.
- Freeze `prompt_v4` before a final HumanEval-X held-out run.
