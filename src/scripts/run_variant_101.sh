#!/usr/bin/env bash
# Run a single vector-RAG variant for run-101 on MiniMax M2.5:
#   1. translate (164 tasks, chromadb backend)
#   2. validate translations; retry any invalid task
#   3. evaluate with Docker
#
# Usage:
#   src/scripts/run_variant_101.sh <experiment>
# where <experiment> is one of: rag-pattern-only, rag-pattern-samples,
#   rag-pattern-api-docs, rag-full, rag-routed
#
# Trap-KB runs are a separate top-level category under `rule-traps/run-N/...`
# and are intentionally not part of this run-101 vector comparison workflow.

set -u

EXPERIMENT="${1:-}"
if [[ -z "$EXPERIMENT" ]]; then
  echo "Usage: $0 <experiment>" >&2
  exit 2
fi

PROVIDER="minimax"
VARIANT="M2.5"
RUN_ID=101
BACKEND="chromadb"
TARGET_DIR="data/translation/target/humaneval-x/${PROVIDER}/${VARIANT}/vec-chroma-3072/run-${RUN_ID}/${EXPERIMENT}"
LOG_DIR=".doc/Log/run-101-variants"
TRANSLATE_LOG="${LOG_DIR}/${EXPERIMENT}.log"
VALIDATE_JSON="${LOG_DIR}/${EXPERIMENT}-validate.json"
EVAL_LOG="${LOG_DIR}/${EXPERIMENT}-eval.log"

mkdir -p "$LOG_DIR"

echo "=== $(date) START variant: $EXPERIMENT ==="

# 1. translate -------------------------------------------------------------
echo "--- translate ---"
uv run python -m src.cli translate \
  -p "$PROVIDER" -v "$VARIANT" \
  --dataset humaneval-x -e "$EXPERIMENT" \
  --embedding-backend "$BACKEND" \
  --run "$RUN_ID" \
  --skip-preflight \
  > "$TRANSLATE_LOG" 2>&1
TR_EXIT=$?
echo "translate exit=$TR_EXIT"

# 2. validate + retry up to 3 times ---------------------------------------
for ATTEMPT in 1 2 3; do
  echo "--- validate attempt $ATTEMPT ---"
  set +e
  RETRY_OUTPUT=$(uv run python src/scripts/validate_translations.py \
    --target-dir "$TARGET_DIR" \
    --expected 164 \
    --report-json "$VALIDATE_JSON" 2>&1)
  V_EXIT=$?
  set -e
  echo "$RETRY_OUTPUT" | tail -20
  if [[ $V_EXIT -eq 0 ]]; then
    echo "all translations valid"
    break
  fi

  RETRY_IDS=$(echo "$RETRY_OUTPUT" | awk -F': ' '/^Retry list: /{print $2}')
  if [[ -z "$RETRY_IDS" ]]; then
    echo "validate failed but no retry list; aborting"
    break
  fi

  echo "--- retranslating problems: $RETRY_IDS (attempt $ATTEMPT) ---"
  uv run python -m src.cli translate \
    -p "$PROVIDER" -v "$VARIANT" \
    --dataset humaneval-x -e "$EXPERIMENT" \
    --embedding-backend "$BACKEND" \
    --run "$RUN_ID" \
    --skip-preflight \
    --problems "$RETRY_IDS" \
    >> "$TRANSLATE_LOG" 2>&1
done

# 3. evaluate -------------------------------------------------------------
echo "--- evaluate ---"
uv run python -m src.cli evaluate \
  --dataset humaneval-x \
  --target-dir "$TARGET_DIR" \
  > "$EVAL_LOG" 2>&1
E_EXIT=$?
echo "evaluate exit=$E_EXIT"

# summary -----------------------------------------------------------------
SUMMARY_FILE="${TARGET_DIR}/evaluation/results/summary.json"
if [[ -f "$SUMMARY_FILE" ]]; then
  echo "--- summary ---"
  cat "$SUMMARY_FILE"
fi

echo "=== $(date) END variant: $EXPERIMENT ==="
