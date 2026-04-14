#!/usr/bin/env python3
"""Automated batch runner for multi-run experiments.

Runs all experiments sequentially with rate-limit-aware batching:
  1. Translate experiments (sequential, with delay between requests)
  2. Evaluate completed experiments (Docker, no API calls)
  3. Sleep until the rate limit window resets
  4. Repeat

Usage:
    # In tmux (recommended):
    tmux new -s experiments
    uv run python src/scripts/run_all_batches.py \
      --provider minimax --variant M2.5 \
      --dimensions 768,1536,3072 --runs 5 \
      --embedding-backend chromadb

    # Dry run (print schedule only):
    uv run python src/scripts/run_all_batches.py \
      --provider minimax --variant M2.5 \
      --dimensions 768,1536,3072 --runs 5 \
      --dry-run
"""

import json
import subprocess
import sys
import time
from datetime import datetime, timedelta
from pathlib import Path

import click
from src.core.humaneval_artifacts import HumanEvalRunPaths

STATE_FILE_DEFAULT = Path(__file__).resolve().parent.parent.parent / "batch_state.json"
LOG_DIR = Path(__file__).resolve().parent.parent.parent / ".doc" / "Log"
FILES_PER_EXPERIMENT = 164

RAG_EXPERIMENTS = [
    "rag-pattern-only",
    "rag-pattern-samples",
    "rag-pattern-api-docs",
    "rag-full",
]

_log_file = None


def _init_log_file() -> None:
    """Create a timestamped log file in .doc/Log/."""
    global _log_file
    LOG_DIR.mkdir(parents=True, exist_ok=True)
    ts = datetime.now().strftime("%Y%m%d_%H%M%S")
    _log_file = open(LOG_DIR / f"batch_{ts}.log", "a", encoding="utf-8")


def log(msg: str) -> None:
    """Print timestamped log message and write to log file."""
    ts = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    line = f"[{ts}] {msg}"
    print(line, flush=True)
    if _log_file is not None:
        _log_file.write(line + "\n")
        _log_file.flush()


def build_experiment_queue(
    dimensions: list[int],
    runs: int,
    include_baseline: bool = True,
) -> list[dict]:
    """Build ordered queue of experiments to run.

    Returns list of dicts with keys: experiment, dimension, run_id, needs_rag
    """
    queue = []

    # Baseline runs first (no embedding needed)
    if include_baseline:
        for run_id in range(1, runs + 1):
            queue.append({
                "experiment": "baseline",
                "dimension": None,
                "run_id": run_id,
                "needs_rag": False,
            })

    # Then each dimension, all runs
    for dim in dimensions:
        for run_id in range(1, runs + 1):
            for exp in RAG_EXPERIMENTS:
                queue.append({
                    "experiment": exp,
                    "dimension": dim,
                    "run_id": run_id,
                    "needs_rag": True,
                })

    return queue


def experiment_key(item: dict) -> str:
    """Unique key for tracking completion."""
    dim = item["dimension"] or "none"
    return f"{item['experiment']}/dim-{dim}/run-{item['run_id']}"


def load_state(state_file: Path) -> set[str]:
    """Load completed experiment keys from state file."""
    if state_file.exists():
        try:
            data = json.loads(state_file.read_text(encoding="utf-8"))
            return set(data.get("completed", []))
        except (json.JSONDecodeError, OSError):
            pass
    return set()


def save_state(state_file: Path, completed: set[str]) -> None:
    """Save completed experiment keys to state file."""
    state_file.write_text(
        json.dumps({"completed": sorted(completed)}, indent=2),
        encoding="utf-8",
    )


def run_translate(
    provider: str,
    variant: str,
    experiment: str,
    embedding_backend: str,
    run_id: int,
    dimension: int | None,
    delay: float,
    sample: int | None = None,
) -> bool:
    """Run a single translation experiment via CLI subprocess.

    Returns True if successful.
    """
    cmd = [
        sys.executable, "-m", "src.cli", "translate",
        "-p", provider,
        "-v", variant,
        "--dataset", "humaneval-x",
        "-e", experiment,
        "--embedding-backend", embedding_backend,
        "--run", str(run_id),
        "--skip-preflight",
    ]
    if dimension is not None:
        cmd.extend(["--dimension", str(dimension)])
    if sample is not None:
        cmd.extend(["-n", str(sample)])

    n_files = sample or FILES_PER_EXPERIMENT
    log(f"  Translating: {experiment} dim={dimension} run-{run_id} ({n_files} files, {delay}s delay)")

    try:
        result = subprocess.run(
            cmd,
            cwd=Path(__file__).resolve().parent.parent.parent,
            capture_output=False,
            timeout=3600,  # 1 hour max per experiment
        )
        if result.returncode == 0:
            log(f"  ✓ {experiment} dim={dimension} run-{run_id} complete")
            return True
        else:
            log(f"  ✗ {experiment} dim={dimension} run-{run_id} failed (exit {result.returncode})")
            return False
    except subprocess.TimeoutExpired:
        log(f"  ✗ {experiment} dim={dimension} run-{run_id} timed out")
        return False
    except Exception as e:
        log(f"  ✗ {experiment} dim={dimension} run-{run_id} error: {e}")
        return False


def verify_and_retry(
    provider: str,
    variant: str,
    experiment: str,
    embedding_backend: str,
    run_id: int,
    dimension: int | None,
    target_dir: Path,
    expected: int,
    max_retries: int = 3,
) -> bool:
    """Check that all expected Go files exist; retry missing ones.

    Returns True if all files are present after retries.
    """
    for attempt in range(1, max_retries + 1):
        expected_nums = set(range(0, expected))
        run_paths = HumanEvalRunPaths(target_dir)
        existing = {
            int(task_paths.task_name.split("_")[1])
            for task_paths in run_paths.iter_task_dirs()
            if task_paths.translation_go.exists()
        }
        missing = sorted(expected_nums - existing)
        if not missing:
            return True

        log(f"  ⚠ {len(missing)} missing file(s) in {target_dir.name}: {missing[:10]}{'...' if len(missing) > 10 else ''}")
        log(f"    Retry {attempt}/{max_retries}: re-translating missing problems")

        problems_str = ",".join(str(n) for n in missing)
        cmd = [
            sys.executable, "-m", "src.cli", "translate",
            "-p", provider,
            "-v", variant,
            "--dataset", "humaneval-x",
            "-e", experiment,
            "--embedding-backend", embedding_backend,
            "--run", str(run_id),
            "--skip-preflight",
            "--problems", problems_str,
        ]
        if dimension is not None:
            cmd.extend(["--dimension", str(dimension)])

        try:
            subprocess.run(
                cmd,
                cwd=Path(__file__).resolve().parent.parent.parent,
                capture_output=False,
                timeout=1800,
            )
        except (subprocess.TimeoutExpired, Exception) as e:
            log(f"    Retry failed: {e}")

    # Final check
    run_paths = HumanEvalRunPaths(target_dir)
    existing = {
        int(task_paths.task_name.split("_")[1])
        for task_paths in run_paths.iter_task_dirs()
        if task_paths.translation_go.exists()
    }
    still_missing = sorted(set(range(0, expected)) - existing)
    if still_missing:
        log(f"  ✗ Still missing {len(still_missing)} file(s) after {max_retries} retries: {still_missing}")
        return False
    return True


def get_target_dir(
    provider: str,
    variant: str,
    experiment: str,
    embedding_backend: str,
    run_id: int,
    dimension: int | None,
) -> Path:
    """Compute the target directory for an experiment (for evaluation)."""
    from src.config import HUMANEVAL_X_DIR
    from src.core.humaneval_artifacts import humaneval_run_root

    backend_label = None
    if experiment != "baseline":
        backend_label = "vec-gemini" if embedding_backend == "gemini" else f"vec-chroma-{dimension}"

    return humaneval_run_root(HUMANEVAL_X_DIR, provider, variant, experiment, backend_label, run_id)


def run_evaluate(target_dir: Path) -> bool:
    """Run evaluation on a translated experiment directory.

    Returns True if successful.
    """
    if not target_dir.exists():
        log(f"  ⚠ Target dir not found: {target_dir}")
        return False

    task_count = sum(1 for task in HumanEvalRunPaths(target_dir).iter_task_dirs() if task.translation_go.exists())
    if task_count == 0:
        log(f"  ⚠ No Go files in {target_dir}")
        return False

    cmd = [
        sys.executable, "-m", "src.cli", "evaluate",
        "--dataset", "humaneval-x",
        "--target-dir", str(target_dir),
    ]

    log(f"  Evaluating: {target_dir.relative_to(target_dir.parents[5])}")

    try:
        result = subprocess.run(
            cmd,
            cwd=Path(__file__).resolve().parent.parent.parent,
            capture_output=False,
            timeout=1800,  # 30 min max per evaluation
        )
        if result.returncode == 0:
            log(f"  ✓ Evaluation complete")
            return True
        else:
            log(f"  ✗ Evaluation failed (exit {result.returncode})")
            return False
    except subprocess.TimeoutExpired:
        log(f"  ✗ Evaluation timed out")
        return False
    except Exception as e:
        log(f"  ✗ Evaluation error: {e}")
        return False


def print_schedule(
    queue: list[dict],
    completed: set[str],
    batch_size: int,
    window_hours: float,
) -> None:
    """Print the planned execution schedule."""
    remaining = [q for q in queue if experiment_key(q) not in completed]
    total_batches = -(-len(remaining) // batch_size)  # ceiling division

    print(f"\n{'='*60}")
    print(f"EXPERIMENT SCHEDULE")
    print(f"{'='*60}")
    print(f"Total experiments:  {len(queue)}")
    print(f"Already completed:  {len(completed)}")
    print(f"Remaining:          {len(remaining)}")
    print(f"Batch size:         {batch_size}")
    print(f"Window:             {window_hours}h")
    print(f"Batches needed:     {total_batches}")

    est_hours = total_batches * window_hours
    est_done = datetime.now() + timedelta(hours=est_hours)
    print(f"Estimated time:     {est_hours:.0f} hours")
    print(f"Estimated done:     {est_done.strftime('%a %Y-%m-%d %H:%M')}")
    print(f"{'='*60}\n")

    now = datetime.now()
    for batch_idx in range(total_batches):
        start = batch_idx * batch_size
        end = min(start + batch_size, len(remaining))
        batch = remaining[start:end]
        batch_start = now + timedelta(hours=batch_idx * window_hours)
        batch_end = batch_start + timedelta(hours=window_hours)

        print(f"Batch {batch_idx + 1}/{total_batches}  "
              f"({batch_start.strftime('%a %H:%M')} - {batch_end.strftime('%a %H:%M')})")
        for item in batch:
            dim_str = f"dim={item['dimension']}" if item['dimension'] else "no-dim"
            print(f"  - {item['experiment']} {dim_str} run-{item['run_id']}")
        print()


@click.command()
@click.option("-p", "--provider", required=True, help="Provider key (e.g. minimax).")
@click.option("-v", "--variant", required=True, help="Model variant (e.g. M2.5).")
@click.option(
    "--dimensions", required=True,
    help="Comma-separated dimensions (e.g. 768,1536,3072).",
)
@click.option("--runs", type=int, default=5, show_default=True, help="Runs per configuration.")
@click.option(
    "--embedding-backend", type=click.Choice(["chromadb", "gemini"]),
    default="chromadb", show_default=True,
)
@click.option("--batch-size", type=int, default=9, show_default=True, help="Experiments per batch.")
@click.option("--window-hours", type=float, default=5.0, show_default=True, help="Rate limit window (hours).")
@click.option("--delay", type=float, default=3.0, show_default=True, help="Delay between API requests (seconds).")
@click.option("--state-file", type=click.Path(path_type=Path), default=None, help="State file path.")
@click.option("--dry-run", is_flag=True, default=False, help="Print schedule without running.")
@click.option("--no-baseline", is_flag=True, default=False, help="Skip baseline runs.")
@click.option("--no-evaluate", is_flag=True, default=False, help="Skip evaluation phase.")
@click.option("--sample", type=int, default=None, help="Translate only N files per experiment (for smoke testing).")
def main(
    provider: str,
    variant: str,
    dimensions: str,
    runs: int,
    embedding_backend: str,
    batch_size: int,
    window_hours: float,
    delay: float,
    state_file: Path | None,
    dry_run: bool,
    no_baseline: bool,
    no_evaluate: bool,
    sample: int | None,
):
    """Run all experiments in automated batches with rate limit management."""
    _init_log_file()
    dims = [int(d.strip()) for d in dimensions.split(",")]
    sf = state_file or STATE_FILE_DEFAULT

    queue = build_experiment_queue(dims, runs, include_baseline=not no_baseline)
    completed = load_state(sf)

    # Print schedule
    print_schedule(queue, completed, batch_size, window_hours)

    if dry_run:
        print("Dry run — exiting without running experiments.")
        return

    # Filter remaining
    remaining = [q for q in queue if experiment_key(q) not in completed]
    if not remaining:
        log("All experiments already completed!")
        return

    total_batches = -(-len(remaining) // batch_size)

    for batch_idx in range(total_batches):
        start = batch_idx * batch_size
        end = min(start + batch_size, len(remaining))
        batch = remaining[start:end]
        batch_start_time = time.time()

        log(f"{'='*60}")
        log(f"Batch {batch_idx + 1}/{total_batches} ({len(batch)} experiments)")
        log(f"{'='*60}")

        # Phase 1: Translate
        translated_dirs = []
        for item in batch:
            key = experiment_key(item)
            if key in completed:
                log(f"  Skipping (already done): {key}")
                continue

            success = run_translate(
                provider=provider,
                variant=variant,
                experiment=item["experiment"],
                embedding_backend=embedding_backend,
                run_id=item["run_id"],
                dimension=item["dimension"],
                delay=delay,
                sample=sample,
            )

            if success:
                target_dir = get_target_dir(
                    provider=provider,
                    variant=variant,
                    experiment=item["experiment"],
                    embedding_backend=embedding_backend,
                    run_id=item["run_id"],
                    dimension=item["dimension"],
                )
                n_expected = sample or FILES_PER_EXPERIMENT
                verify_and_retry(
                    provider=provider,
                    variant=variant,
                    experiment=item["experiment"],
                    embedding_backend=embedding_backend,
                    run_id=item["run_id"],
                    dimension=item["dimension"],
                    target_dir=target_dir,
                    expected=n_expected,
                )
                translated_dirs.append((item, target_dir))

            # Delay between experiments
            if item != batch[-1]:
                time.sleep(delay)

        # Phase 2: Evaluate
        if not no_evaluate and translated_dirs:
            log(f"\n--- Evaluating {len(translated_dirs)} experiments ---")
            for item, target_dir in translated_dirs:
                success = run_evaluate(target_dir)
                key = experiment_key(item)
                if success:
                    completed.add(key)
                    save_state(sf, completed)
                    log(f"  ✓ Saved state ({len(completed)} total completed)")
        elif no_evaluate:
            # Mark as completed even without evaluation
            for item, _ in translated_dirs:
                completed.add(experiment_key(item))
            save_state(sf, completed)

        # Phase 3: Sleep until window resets (unless last batch)
        is_last_batch = (batch_idx == total_batches - 1)
        if not is_last_batch:
            elapsed = time.time() - batch_start_time
            window_seconds = window_hours * 3600
            sleep_seconds = max(0, window_seconds - elapsed)

            if sleep_seconds > 0:
                wake_time = datetime.now() + timedelta(seconds=sleep_seconds)
                log(f"\nSleeping until {wake_time.strftime('%H:%M:%S')} "
                    f"({sleep_seconds / 3600:.1f}h remaining)")

                # Print countdown every 30 minutes
                while sleep_seconds > 0:
                    chunk = min(sleep_seconds, 1800)  # 30 min chunks
                    time.sleep(chunk)
                    sleep_seconds -= chunk
                    if sleep_seconds > 0:
                        log(f"  ... {sleep_seconds / 3600:.1f}h remaining")

    log(f"\n{'='*60}")
    log(f"ALL DONE! {len(completed)} experiments completed.")
    log(f"{'='*60}")
    log(f"Run statistics analysis:")
    log(f"  uv run python src/scripts/analyze_statistics.py")


if __name__ == "__main__":
    main()
