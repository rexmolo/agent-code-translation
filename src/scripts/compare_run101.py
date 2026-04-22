#!/usr/bin/env python3
"""Summarize run-101 baseline vs. 5 RAG variants.

Reads the per-run `evaluation/results/summary.json` for the baseline and each
variant and prints a comparison table. Writes a Markdown table to the memory
folder so the thesis notes stay self-describing.
"""

from __future__ import annotations

import json
from pathlib import Path

import click


VARIANTS = [
    "rag-pattern-only",
    "rag-pattern-samples",
    "rag-pattern-api-docs",
    "rag-full",
    "rag-routed",
]


def _load_summary(path: Path) -> dict | None:
    if not path.is_file():
        return None
    try:
        return json.loads(path.read_text(encoding="utf-8"))
    except json.JSONDecodeError:
        return None


def _fmt_ratio(value: float) -> str:
    return f"{value*100:6.2f}%" if value is not None else "   n/a"


def _fmt_delta(value: float, baseline: float) -> str:
    if value is None or baseline is None:
        return "   n/a"
    delta = (value - baseline) * 100
    sign = "+" if delta >= 0 else ""
    return f"{sign}{delta:5.2f}pp"


@click.command()
@click.option("--root", type=click.Path(path_type=Path),
              default=Path("data/translation/target/humaneval-x/minimax/M2.5"),
              show_default=True)
@click.option("--run-id", type=int, default=101, show_default=True)
@click.option("--backend-label", default="vec-chroma-3072", show_default=True)
@click.option("--out", type=click.Path(path_type=Path),
              default=Path(".doc/Tasks/01 - active/Task-1/01 - plan/07 - run-101-variant-comparison.md"),
              show_default=True)
def main(root: Path, run_id: int, backend_label: str, out: Path) -> None:
    baseline_summary = _load_summary(
        root / "baseline" / f"run-{run_id}" / "evaluation" / "results" / "summary.json"
    )
    if baseline_summary is None:
        raise click.ClickException(f"Baseline summary missing under {root}")

    rows: list[dict] = []
    rows.append({"name": "baseline", "summary": baseline_summary})
    for variant in VARIANTS:
        path = (root / backend_label / f"run-{run_id}" / variant
                / "evaluation" / "results" / "summary.json")
        rows.append({"name": variant, "summary": _load_summary(path), "path": path})

    base_p = baseline_summary.get("pass_at_1")
    base_c = baseline_summary.get("compilation_at_1")

    # console
    header = f"{'variant':<24} {'pass@1':>8} {'Δ pass':>8}  {'compile@1':>10} {'Δ comp':>8}  {'total':>6}"
    click.echo(header)
    click.echo("-" * len(header))
    for r in rows:
        s = r["summary"] or {}
        p = s.get("pass_at_1")
        c = s.get("compilation_at_1")
        total = s.get("total_files", "-")
        click.echo(
            f"{r['name']:<24} {_fmt_ratio(p):>8} {_fmt_delta(p, base_p):>8}  "
            f"{_fmt_ratio(c):>10} {_fmt_delta(c, base_c):>8}  {total:>6}"
        )

    # markdown
    lines = [
        f"# run-{run_id} variant comparison",
        "",
        f"Backend: `{backend_label}` | Provider: `minimax/M2.5` | HumanEval-X, 164 tasks.",
        "",
        "| variant | pass@1 | Δ pass | compile@1 | Δ compile | total |",
        "|---|---:|---:|---:|---:|---:|",
    ]
    for r in rows:
        s = r["summary"] or {}
        p = s.get("pass_at_1")
        c = s.get("compilation_at_1")
        total = s.get("total_files", "-")
        lines.append(
            f"| `{r['name']}` | {_fmt_ratio(p).strip()} | "
            f"{_fmt_delta(p, base_p).strip()} | {_fmt_ratio(c).strip()} | "
            f"{_fmt_delta(c, base_c).strip()} | {total} |"
        )

    out.parent.mkdir(parents=True, exist_ok=True)
    out.write_text("\n".join(lines) + "\n", encoding="utf-8")
    click.echo(f"\nWrote: {out}")


if __name__ == "__main__":
    main()
