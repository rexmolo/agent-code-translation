"""Re-evaluate all HumanEval-X experiments after pipeline fix.

Runs evaluations 2 at a time (Docker memory constraint: 16GB limit,
each eval uses ~5GB with batch_size=10).

Usage:
    uv run python src/scripts/reeval_all.py
    uv run python src/scripts/reeval_all.py --dry-run
"""

from __future__ import annotations

import argparse
import subprocess
import sys
import time
from concurrent.futures import ThreadPoolExecutor, as_completed
from pathlib import Path

BASE = Path("data/translation/target/humaneval-x")

def discover_targets() -> list[Path]:
    """Find all leaf experiment directories that contain Go_*.go files."""
    targets = []
    for go_file in BASE.rglob("Go_0.go"):
        targets.append(go_file.parent)
    return sorted(targets)


def run_eval(target: Path, idx: int, total: int) -> tuple[Path, bool, str]:
    """Run a single evaluation and return (path, success, summary)."""
    label = str(target.relative_to(BASE))
    print(f"[{idx}/{total}] Starting: {label}", flush=True)
    t0 = time.time()

    result = subprocess.run(
        [
            sys.executable, "-m", "src.cli", "evaluate",
            "-d", "humaneval-x",
            "--target-dir", str(target),
        ],
        capture_output=True,
        text=True,
        timeout=600,
    )

    elapsed = time.time() - t0
    # Extract metrics from output
    summary = ""
    for line in result.stdout.split("\n"):
        if "Compilation@1" in line or "Pass@1" in line:
            summary += line.strip() + "  "

    success = result.returncode == 0
    status = "OK" if success else "FAIL"
    print(f"[{idx}/{total}] {status} ({elapsed:.0f}s): {label}  {summary}", flush=True)

    if not success and result.stderr:
        # Print last 3 lines of stderr for diagnosis
        err_lines = result.stderr.strip().split("\n")[-3:]
        for line in err_lines:
            print(f"         {line}", flush=True)

    return target, success, summary


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--dry-run", action="store_true", help="List targets without running")
    parser.add_argument("--parallel", type=int, default=2, help="Max parallel evaluations")
    args = parser.parse_args()

    targets = discover_targets()
    print(f"Found {len(targets)} evaluation targets (parallelism: {args.parallel})\n")

    if args.dry_run:
        for i, t in enumerate(targets, 1):
            print(f"  {i:3d}. {t.relative_to(BASE)}")
        return

    results = []
    t_start = time.time()

    with ThreadPoolExecutor(max_workers=args.parallel) as pool:
        futures = {
            pool.submit(run_eval, t, i, len(targets)): t
            for i, t in enumerate(targets, 1)
        }
        for future in as_completed(futures):
            target, success, summary = future.result()
            results.append((target, success, summary))

    elapsed = time.time() - t_start
    ok = sum(1 for _, s, _ in results if s)
    print(f"\n{'='*60}")
    print(f"Done: {ok}/{len(results)} succeeded in {elapsed/60:.1f} minutes")

    failed = [(t, s) for t, s, _ in results if not s]
    if failed:
        print(f"\nFailed ({len(failed)}):")
        for t, _ in failed:
            print(f"  - {t.relative_to(BASE)}")


if __name__ == "__main__":
    main()
