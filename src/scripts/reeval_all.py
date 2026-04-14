"""Re-evaluate all HumanEval-X experiments after pipeline fix.

Runs evaluations sequentially to avoid Docker resource contention
(batch_size=10 containers per eval × 512MB each).

Usage:
    uv run python src/scripts/reeval_all.py
    uv run python src/scripts/reeval_all.py --dry-run
"""

from __future__ import annotations

import argparse
import json
import os
import signal
import subprocess
import sys
import time
from datetime import datetime, timezone
from pathlib import Path

from src.core.humaneval_artifacts import is_humaneval_run_root

BASE = Path("data/translation/humaneval-x")
RESULTS_DIR = Path(".doc/memory")

# Track active child so we can clean up on SIGTERM/SIGINT
_active_proc: subprocess.Popen | None = None


def _cleanup_handler(signum, frame):
    """Kill child process group when this script is terminated."""
    if _active_proc and _active_proc.poll() is None:
        os.killpg(_active_proc.pid, signal.SIGTERM)
    sys.exit(1)


signal.signal(signal.SIGTERM, _cleanup_handler)
signal.signal(signal.SIGINT, _cleanup_handler)


def discover_targets() -> list[Path]:
    """Find all HumanEval-X run roots."""
    return sorted(path for path in BASE.rglob("*") if is_humaneval_run_root(path))


def run_eval(target: Path, idx: int, total: int, timeout: int = 600) -> tuple[Path, bool, str]:
    """Run a single evaluation and return (path, success, summary)."""
    label = str(target.relative_to(BASE))
    print(f"[{idx}/{total}] Starting: {label}", flush=True)
    t0 = time.time()

    global _active_proc
    proc = subprocess.Popen(
        [
            sys.executable, "-m", "src.cli", "evaluate",
            "-d", "humaneval-x",
            "--target-dir", str(target),
        ],
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        text=True,
        start_new_session=True,
    )
    _active_proc = proc
    try:
        stdout, stderr = proc.communicate(timeout=timeout)
    except subprocess.TimeoutExpired:
        # Kill the entire process group so child Docker processes die too
        os.killpg(proc.pid, signal.SIGTERM)
        proc.wait()
        elapsed = time.time() - t0
        print(f"[{idx}/{total}] TIMEOUT ({elapsed:.0f}s): {label}", flush=True)
        return target, False, ""
    finally:
        _active_proc = None  # Clear so signal handler won't kill a stale process

    elapsed = time.time() - t0
    # Extract metrics from output
    summary = ""
    for line in stdout.split("\n"):
        if "Compilation@1" in line or "Pass@1" in line:
            summary += line.strip() + "  "

    success = proc.returncode == 0
    status = "OK" if success else "FAIL"
    print(f"[{idx}/{total}] {status} ({elapsed:.0f}s): {label}  {summary}", flush=True)

    if not success and stderr:
        # Print last 3 lines of stderr for diagnosis
        err_lines = stderr.strip().split("\n")[-3:]
        for line in err_lines:
            print(f"         {line}", flush=True)

    return target, success, summary


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--dry-run", action="store_true", help="List targets without running")
    parser.add_argument("--timeout", type=int, default=600, help="Timeout per eval in seconds (default: 600)")
    args = parser.parse_args()

    targets = discover_targets()
    print(f"Found {len(targets)} evaluation targets\n")

    if args.dry_run:
        for i, t in enumerate(targets, 1):
            print(f"  {i:3d}. {t.relative_to(BASE)}")
        return

    results = []
    t_start = time.time()
    ts = datetime.now(timezone.utc).strftime("%Y%m%d_%H%M%S")
    results_file = RESULTS_DIR / f"reeval_{ts}.jsonl"
    results_file.parent.mkdir(parents=True, exist_ok=True)
    print(f"Results will be saved to: {results_file}\n")

    for i, target in enumerate(targets, 1):
        target, success, summary = run_eval(target, i, len(targets), timeout=args.timeout)
        results.append((target, success, summary))
        # Write incrementally so partial results survive crashes
        record = {
            "target": str(target.relative_to(BASE)),
            "success": success,
            "summary": summary.strip(),
            "timestamp": datetime.now(timezone.utc).isoformat(),
        }
        with open(results_file, "a") as f:
            f.write(json.dumps(record) + "\n")

    elapsed = time.time() - t_start
    ok = sum(1 for _, s, _ in results if s)
    print(f"\n{'='*60}")
    print(f"Done: {ok}/{len(results)} succeeded in {elapsed/60:.1f} minutes")
    print(f"Results saved to: {results_file}")

    failed = [(t, s) for t, s, _ in results if not s]
    if failed:
        print(f"\nFailed ({len(failed)}):")
        for t, _ in failed:
            print(f"  - {t.relative_to(BASE)}")


if __name__ == "__main__":
    main()
