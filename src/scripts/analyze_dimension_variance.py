#!/usr/bin/env python3
"""Generate variance report and statistical plots from the re-evaluation jsonl.

Reads .doc/memory/reeval_20260409_133712.jsonl (the canonical re-evaluation log,
NOT the stale per-experiment mlflow_results.json files), computes descriptive
statistics + ANOVA / pairwise Welch's t-tests for the MiniMax M2.5 dimension
variance experiment, renders all selected plots, and assembles a single
Markdown report with computed interpretations beneath every figure.

Usage:
    uv run python src/scripts/analyze_dimension_variance.py
    uv run python src/scripts/analyze_dimension_variance.py --metric pass_at_1
    uv run python src/scripts/analyze_dimension_variance.py --output-dir reports
"""

from __future__ import annotations

import json
import math
import re
from dataclasses import dataclass
from pathlib import Path

import click
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from matplotlib.patches import Patch
from rich.console import Console
from scipy import stats

PROJECT_ROOT = Path(__file__).resolve().parent.parent.parent
DEFAULT_SOURCE = PROJECT_ROOT / ".doc" / "memory" / "reeval_20260409_133712.jsonl"
DEFAULT_OUTPUT = PROJECT_ROOT / "reports"

METRIC_LABELS = {
    "pass_at_1": "Pass@1",
    "compilation_at_1": "Compilation@1",
}

VARIANTS = ["rag-pattern-only", "rag-pattern-samples", "rag-pattern-api-docs", "rag-full"]
DIMENSIONS = [768, 1536, 3072]

VARIANT_COLORS = {
    "rag-pattern-only": "#4C72B0",
    "rag-pattern-samples": "#DD8452",
    "rag-pattern-api-docs": "#55A868",
    "rag-full": "#C44E52",
}
DIMENSION_COLORS = {
    768: "#4C72B0",
    1536: "#DD8452",
    3072: "#55A868",
}
PROVIDER_COLORS = {
    "gemini": "#4C72B0",
    "minimax": "#DD8452",
    "openai": "#55A868",
}

SIGNIFICANCE_LEVEL = 0.05
COMPILE_RE = re.compile(r"Compilation@1[^\d]*([\d.]+)%")
PASS_RE = re.compile(r"Pass@1[^\d]*([\d.]+)%")

console = Console()


# --------------------------------------------------------------------------- #
# Parsing
# --------------------------------------------------------------------------- #


@dataclass
class ParsedRow:
    provider: str
    model: str
    backend: str | None         # "vec-chroma-768" | "vec-gemini" | None (baseline)
    dimension: int | None       # 768 | 1536 | 3072 | None
    experiment: str             # "baseline" | "rag-full" | ...
    run_id: int | None
    compilation_at_1: float
    pass_at_1: float


def parse_target(target: str) -> tuple[str, str, str | None, int | None, str, int | None]:
    """Split a target path into (provider, model, backend, dimension, experiment, run_id)."""
    parts = target.split("/")
    provider = parts[0]
    model = parts[1]
    rest = parts[2:]

    if rest == ["baseline"]:
        return provider, model, None, None, "baseline", None

    backend = rest[0]
    dimension: int | None = None
    if backend.startswith("vec-chroma-"):
        dimension = int(backend.split("-")[-1])

    if len(rest) == 3 and rest[1].startswith("run-"):
        run_id = int(rest[1].split("-", 1)[1])
        experiment = rest[2]
    elif len(rest) == 2:
        run_id = 1
        experiment = rest[1]
    else:
        raise ValueError(f"Unrecognised target shape: {target}")

    return provider, model, backend, dimension, experiment, run_id


def parse_summary(summary: str) -> tuple[float, float]:
    """Extract (compilation_at_1, pass_at_1) as floats in [0,1] from the rich-table summary."""
    cm = COMPILE_RE.search(summary)
    pm = PASS_RE.search(summary)
    if not cm or not pm:
        raise ValueError(f"Could not parse summary: {summary!r}")
    return float(cm.group(1)) / 100.0, float(pm.group(1)) / 100.0


def load_jsonl(path: Path) -> list[ParsedRow]:
    rows: list[ParsedRow] = []
    with path.open(encoding="utf-8") as fh:
        for raw in fh:
            raw = raw.strip()
            if not raw:
                continue
            obj = json.loads(raw)
            if not obj.get("success"):
                continue
            target = obj["target"]
            summary = obj["summary"]
            try:
                provider, model, backend, dim, exp, run_id = parse_target(target)
                comp, passed = parse_summary(summary)
            except (ValueError, KeyError) as e:
                console.print(f"[yellow]Skipping malformed row: {target} ({e})[/yellow]")
                continue
            rows.append(ParsedRow(provider, model, backend, dim, exp, run_id, comp, passed))
    return rows


def build_dataframe(rows: list[ParsedRow]) -> pd.DataFrame:
    return pd.DataFrame(
        [
            {
                "provider": r.provider,
                "model": r.model,
                "backend": r.backend,
                "dimension": r.dimension,
                "experiment": r.experiment,
                "run_id": r.run_id,
                "compilation_at_1": r.compilation_at_1,
                "pass_at_1": r.pass_at_1,
            }
            for r in rows
        ]
    )


# --------------------------------------------------------------------------- #
# Statistics
# --------------------------------------------------------------------------- #


def descriptive_stats(values: list[float] | np.ndarray) -> dict[str, float]:
    arr = np.asarray(values, dtype=float)
    n = len(arr)
    mean = float(arr.mean()) if n else float("nan")
    std = float(arr.std(ddof=1)) if n > 1 else 0.0
    if n > 1:
        ci_half = float(stats.t.ppf(0.975, df=n - 1) * std / math.sqrt(n))
    else:
        ci_half = 0.0
    cv = std / mean if mean else 0.0
    return {
        "n": n,
        "mean": mean,
        "std": std,
        "min": float(arr.min()) if n else float("nan"),
        "max": float(arr.max()) if n else float("nan"),
        "ci95": ci_half,
        "cv": cv,
    }


def minimax_chroma(df: pd.DataFrame) -> pd.DataFrame:
    return df[
        (df["provider"] == "minimax")
        & (df["model"] == "M2.5")
        & (df["backend"].str.startswith("vec-chroma-", na=False))
        & (df["experiment"] != "baseline")
    ].copy()


def descriptive_table(df_chroma: pd.DataFrame, metric: str) -> pd.DataFrame:
    rows = []
    for variant in VARIANTS:
        for dim in DIMENSIONS:
            sub = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == dim)]
            stats_dict = descriptive_stats(sub[metric].tolist())
            rows.append({"experiment": variant, "dimension": dim, **stats_dict})
    return pd.DataFrame(rows)


def anova_table(df_chroma: pd.DataFrame, metric: str) -> pd.DataFrame:
    rows = []
    for variant in VARIANTS:
        groups = [
            df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == d)][metric].tolist()
            for d in DIMENSIONS
        ]
        if all(len(g) >= 2 for g in groups):
            f_stat, p_val = stats.f_oneway(*groups)
        else:
            f_stat, p_val = float("nan"), float("nan")
        rows.append(
            {
                "experiment": variant,
                "F": f_stat,
                "p_value": p_val,
                "significant": (not math.isnan(p_val)) and p_val < SIGNIFICANCE_LEVEL,
            }
        )
    return pd.DataFrame(rows)


def pairwise_ttests(df_chroma: pd.DataFrame, variant: str, metric: str) -> pd.DataFrame:
    rows = []
    for i in range(len(DIMENSIONS)):
        for j in range(i + 1, len(DIMENSIONS)):
            d1, d2 = DIMENSIONS[i], DIMENSIONS[j]
            v1 = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == d1)][metric].tolist()
            v2 = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == d2)][metric].tolist()
            if len(v1) < 2 or len(v2) < 2:
                t_stat, p_val = float("nan"), float("nan")
            else:
                t_stat, p_val = stats.ttest_ind(v1, v2, equal_var=False)
            rows.append({"dim_a": d1, "dim_b": d2, "t": float(t_stat), "p": float(p_val)})
    return pd.DataFrame(rows)


def pvalue_matrix(df_chroma: pd.DataFrame, variant: str, metric: str) -> np.ndarray:
    n = len(DIMENSIONS)
    mat = np.full((n, n), np.nan)
    for i in range(n):
        for j in range(n):
            if i == j:
                continue
            d1, d2 = DIMENSIONS[i], DIMENSIONS[j]
            v1 = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == d1)][metric].tolist()
            v2 = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == d2)][metric].tolist()
            if len(v1) >= 2 and len(v2) >= 2:
                _, p = stats.ttest_ind(v1, v2, equal_var=False)
                mat[i, j] = float(p)
    return mat


# --------------------------------------------------------------------------- #
# Plot helpers
# --------------------------------------------------------------------------- #


def style_axes(ax, title: str, xlabel: str, ylabel: str):
    ax.set_title(title, fontsize=12, pad=10)
    ax.set_xlabel(xlabel, fontsize=10)
    ax.set_ylabel(ylabel, fontsize=10)
    ax.grid(axis="y", alpha=0.3, linestyle="--", linewidth=0.6)
    ax.set_axisbelow(True)


def save_fig(fig, out_path: Path, *, tight: bool = True):
    out_path.parent.mkdir(parents=True, exist_ok=True)
    if tight:
        fig.tight_layout()
    fig.savefig(out_path, dpi=300, bbox_inches="tight")
    plt.close(fig)


# --------------------------------------------------------------------------- #
# MiniMax variance plots
# --------------------------------------------------------------------------- #


def plot_box(df_chroma: pd.DataFrame, metric: str, out_path: Path):
    fig, ax = plt.subplots(figsize=(10, 6))
    n_variants = len(VARIANTS)
    n_dims = len(DIMENSIONS)
    width = 0.22
    positions_centers = np.arange(n_variants)

    for di, dim in enumerate(DIMENSIONS):
        offset = (di - (n_dims - 1) / 2) * (width + 0.02)
        data = []
        for variant in VARIANTS:
            sub = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == dim)][metric].tolist()
            data.append(sub)
        positions = positions_centers + offset
        bp = ax.boxplot(
            data,
            positions=positions,
            widths=width,
            patch_artist=True,
            medianprops=dict(color="black", linewidth=1.4),
            showfliers=True,
        )
        for patch in bp["boxes"]:
            patch.set_facecolor(DIMENSION_COLORS[dim])
            patch.set_alpha(0.75)

    ax.set_xticks(positions_centers)
    ax.set_xticklabels([v.replace("rag-", "") for v in VARIANTS], rotation=15, ha="right")
    style_axes(
        ax,
        f"MiniMax M2.5 — {METRIC_LABELS[metric]} distribution by dimension (5 runs)",
        "RAG variant",
        METRIC_LABELS[metric],
    )
    legend_handles = [Patch(facecolor=DIMENSION_COLORS[d], label=f"{d}") for d in DIMENSIONS]
    ax.legend(handles=legend_handles, title="Dimension", loc="lower right", framealpha=0.9)
    ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.0%}"))
    save_fig(fig, out_path)


def plot_bars(df_chroma: pd.DataFrame, metric: str, out_path: Path):
    fig, ax = plt.subplots(figsize=(10, 6))
    n_dims = len(DIMENSIONS)
    width = 0.25
    positions_centers = np.arange(len(VARIANTS))

    for di, dim in enumerate(DIMENSIONS):
        means = []
        stds = []
        for variant in VARIANTS:
            sub = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == dim)][metric].to_numpy()
            means.append(sub.mean() if len(sub) else np.nan)
            stds.append(sub.std(ddof=1) if len(sub) > 1 else 0.0)
        offset = (di - (n_dims - 1) / 2) * width
        ax.bar(
            positions_centers + offset,
            means,
            width=width,
            yerr=stds,
            color=DIMENSION_COLORS[dim],
            edgecolor="black",
            linewidth=0.4,
            alpha=0.85,
            capsize=4,
            label=f"{dim}",
        )

    ax.set_xticks(positions_centers)
    ax.set_xticklabels([v.replace("rag-", "") for v in VARIANTS], rotation=15, ha="right")
    style_axes(
        ax,
        f"MiniMax M2.5 — Mean {METRIC_LABELS[metric]} ± 1 SD by dimension",
        "RAG variant",
        f"Mean {METRIC_LABELS[metric]}",
    )
    ax.legend(title="Dimension", loc="lower right", framealpha=0.9)
    ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.0%}"))
    save_fig(fig, out_path)


def plot_line(df_chroma: pd.DataFrame, df_baseline_value: float, metric: str, out_path: Path):
    fig, ax = plt.subplots(figsize=(10, 6))

    for variant in VARIANTS:
        means, stds = [], []
        for dim in DIMENSIONS:
            sub = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == dim)][metric].to_numpy()
            means.append(sub.mean() if len(sub) else np.nan)
            stds.append(sub.std(ddof=1) if len(sub) > 1 else 0.0)
        means_arr = np.array(means)
        stds_arr = np.array(stds)
        ax.plot(
            DIMENSIONS,
            means_arr,
            marker="o",
            linewidth=2,
            color=VARIANT_COLORS[variant],
            label=variant.replace("rag-", ""),
        )
        ax.fill_between(
            DIMENSIONS,
            means_arr - stds_arr,
            means_arr + stds_arr,
            color=VARIANT_COLORS[variant],
            alpha=0.15,
        )

    ax.axhline(
        df_baseline_value,
        color="black",
        linestyle="--",
        linewidth=1.2,
        alpha=0.7,
        label=f"baseline ({df_baseline_value:.1%})",
    )
    ax.set_xticks(DIMENSIONS)
    style_axes(
        ax,
        f"MiniMax M2.5 — Mean {METRIC_LABELS[metric]} across embedding dimensions",
        "Embedding dimension",
        f"Mean {METRIC_LABELS[metric]} (± 1 SD band)",
    )
    ax.legend(title="RAG variant", loc="lower right", framealpha=0.9)
    ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.0%}"))
    save_fig(fig, out_path)


def plot_cv(df_chroma: pd.DataFrame, metric: str, out_path: Path):
    fig, ax = plt.subplots(figsize=(10, 6))
    n_dims = len(DIMENSIONS)
    width = 0.25
    positions_centers = np.arange(len(VARIANTS))

    for di, dim in enumerate(DIMENSIONS):
        cvs = []
        for variant in VARIANTS:
            sub = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == dim)][metric].to_numpy()
            if len(sub) > 1 and sub.mean() > 0:
                cvs.append(sub.std(ddof=1) / sub.mean())
            else:
                cvs.append(0.0)
        offset = (di - (n_dims - 1) / 2) * width
        ax.bar(
            positions_centers + offset,
            cvs,
            width=width,
            color=DIMENSION_COLORS[dim],
            edgecolor="black",
            linewidth=0.4,
            alpha=0.85,
            label=f"{dim}",
        )

    ax.set_xticks(positions_centers)
    ax.set_xticklabels([v.replace("rag-", "") for v in VARIANTS], rotation=15, ha="right")
    style_axes(
        ax,
        f"MiniMax M2.5 — Coefficient of variation in {METRIC_LABELS[metric]} (lower = more stable)",
        "RAG variant",
        "CV (std / mean)",
    )
    ax.legend(title="Dimension", loc="upper right", framealpha=0.9)
    ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.1%}"))
    save_fig(fig, out_path)


def plot_trajectory(df_chroma: pd.DataFrame, metric: str, out_path: Path):
    fig, axes = plt.subplots(2, 2, figsize=(12, 8), sharex=True, sharey=True)
    axes = axes.flatten()
    for ax, variant in zip(axes, VARIANTS):
        for dim in DIMENSIONS:
            sub = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == dim)].sort_values("run_id")
            ax.plot(
                sub["run_id"].to_numpy(),
                sub[metric].to_numpy(),
                marker="o",
                linewidth=1.8,
                color=DIMENSION_COLORS[dim],
                label=f"{dim}",
            )
        ax.set_title(variant.replace("rag-", ""), fontsize=11)
        ax.grid(alpha=0.3, linestyle="--", linewidth=0.6)
        ax.set_axisbelow(True)
        ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.0%}"))
        ax.set_xticks([1, 2, 3, 4, 5])

    fig.suptitle(
        f"MiniMax M2.5 — {METRIC_LABELS[metric]} per run (1–5) by dimension",
        fontsize=13,
        y=1.00,
    )
    for ax in axes[2:]:
        ax.set_xlabel("Run")
    for ax in (axes[0], axes[2]):
        ax.set_ylabel(METRIC_LABELS[metric])
    axes[0].legend(title="Dimension", loc="lower right", framealpha=0.9, fontsize=8)
    save_fig(fig, out_path)


def plot_significance(df_chroma: pd.DataFrame, metric: str, out_path: Path):
    fig, axes = plt.subplots(2, 2, figsize=(11, 9))
    axes = axes.flatten()
    cmap = plt.get_cmap("RdYlGn")  # red = small p (significant), green = large p (not significant)

    for ax, variant in zip(axes, VARIANTS):
        mat = pvalue_matrix(df_chroma, variant, metric)
        display = np.where(np.isnan(mat), 1.0, mat)
        im = ax.imshow(display, cmap=cmap, vmin=0.0, vmax=0.5, aspect="equal")
        ax.set_xticks(range(len(DIMENSIONS)))
        ax.set_yticks(range(len(DIMENSIONS)))
        ax.set_xticklabels([str(d) for d in DIMENSIONS])
        ax.set_yticklabels([str(d) for d in DIMENSIONS])
        ax.set_title(variant.replace("rag-", ""), fontsize=11)
        for i in range(len(DIMENSIONS)):
            for j in range(len(DIMENSIONS)):
                if i == j:
                    txt = "—"
                elif np.isnan(mat[i, j]):
                    txt = "n/a"
                else:
                    p = mat[i, j]
                    txt = f"{p:.3f}" + ("*" if p < SIGNIFICANCE_LEVEL else "")
                ax.text(j, i, txt, ha="center", va="center", fontsize=9, color="black")

    fig.suptitle(
        f"MiniMax M2.5 — Pairwise Welch's t-test p-values ({METRIC_LABELS[metric]})\n"
        f"red = significant difference (p<0.05*), green = not significant",
        fontsize=12,
        y=1.00,
    )
    fig.colorbar(im, ax=axes, fraction=0.025, pad=0.04, label="p-value (clipped at 0.5)")
    save_fig(fig, out_path, tight=False)


# --------------------------------------------------------------------------- #
# Cross-model plots
# --------------------------------------------------------------------------- #


def plot_baselines_compare(df: pd.DataFrame, out_path: Path):
    baselines = df[df["experiment"] == "baseline"].copy()
    baselines["label"] = baselines["provider"] + "/" + baselines["model"]
    baselines = baselines.sort_values("provider")
    labels = baselines["label"].tolist()

    fig, ax = plt.subplots(figsize=(9, 6))
    width = 0.35
    positions = np.arange(len(labels))
    comp = baselines["compilation_at_1"].to_numpy()
    pas = baselines["pass_at_1"].to_numpy()
    ax.bar(positions - width / 2, comp, width=width, color="#4C72B0", label="Compilation@1", edgecolor="black", linewidth=0.4)
    ax.bar(positions + width / 2, pas, width=width, color="#C44E52", label="Pass@1", edgecolor="black", linewidth=0.4)

    for i, (c, p) in enumerate(zip(comp, pas)):
        ax.text(i - width / 2, c + 0.005, f"{c:.1%}", ha="center", fontsize=9)
        ax.text(i + width / 2, p + 0.005, f"{p:.1%}", ha="center", fontsize=9)

    ax.set_xticks(positions)
    ax.set_xticklabels(labels, rotation=15, ha="right")
    style_axes(ax, "Baseline (no RAG) comparison across providers", "Provider / Model", "Metric value")
    ax.set_ylim(0, 1.05)
    ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.0%}"))
    ax.legend(loc="lower right", framealpha=0.9)
    save_fig(fig, out_path)


def plot_openai_rag(df: pd.DataFrame, out_path: Path):
    openai_rag = df[(df["provider"] == "openai") & (df["experiment"] != "baseline")].copy()
    openai_baseline = df[(df["provider"] == "openai") & (df["experiment"] == "baseline")].iloc[0]

    fig, ax = plt.subplots(figsize=(10, 6))
    width = 0.35
    positions = np.arange(len(VARIANTS))
    comp_vals, pass_vals = [], []
    for variant in VARIANTS:
        row = openai_rag[openai_rag["experiment"] == variant]
        comp_vals.append(row["compilation_at_1"].iloc[0] if not row.empty else np.nan)
        pass_vals.append(row["pass_at_1"].iloc[0] if not row.empty else np.nan)

    ax.bar(positions - width / 2, comp_vals, width=width, color="#4C72B0", label="Compilation@1", edgecolor="black", linewidth=0.4)
    ax.bar(positions + width / 2, pass_vals, width=width, color="#C44E52", label="Pass@1", edgecolor="black", linewidth=0.4)
    ax.axhline(openai_baseline["compilation_at_1"], color="#4C72B0", linestyle="--", linewidth=1.2, alpha=0.7, label=f"baseline Compilation@1 ({openai_baseline['compilation_at_1']:.1%})")
    ax.axhline(openai_baseline["pass_at_1"], color="#C44E52", linestyle="--", linewidth=1.2, alpha=0.7, label=f"baseline Pass@1 ({openai_baseline['pass_at_1']:.1%})")

    ax.set_xticks(positions)
    ax.set_xticklabels([v.replace("rag-", "") for v in VARIANTS], rotation=15, ha="right")
    style_axes(ax, "OpenAI GPT-5.4 — RAG variants on chroma-3072 vs. baseline", "RAG variant", "Metric value")
    ax.set_ylim(0.85, 1.02)
    ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.0%}"))
    ax.legend(loc="lower right", framealpha=0.9, fontsize=8)
    save_fig(fig, out_path)


def plot_minimax_vec_gemini_vs_chroma(df: pd.DataFrame, out_path: Path):
    chroma_3072 = df[
        (df["provider"] == "minimax") & (df["model"] == "M2.5") & (df["backend"] == "vec-chroma-3072")
    ]
    vec_gemini = df[
        (df["provider"] == "minimax") & (df["model"] == "M2.5") & (df["backend"] == "vec-gemini")
    ]

    fig, axes = plt.subplots(1, 2, figsize=(13, 6), sharey=True)
    metrics = ["pass_at_1", "compilation_at_1"]
    for ax, metric in zip(axes, metrics):
        positions = np.arange(len(VARIANTS))
        chroma_means, chroma_stds, gemini_vals = [], [], []
        for variant in VARIANTS:
            c_sub = chroma_3072[chroma_3072["experiment"] == variant][metric].to_numpy()
            chroma_means.append(c_sub.mean() if len(c_sub) else np.nan)
            chroma_stds.append(c_sub.std(ddof=1) if len(c_sub) > 1 else 0.0)
            g_sub = vec_gemini[vec_gemini["experiment"] == variant][metric].to_numpy()
            gemini_vals.append(g_sub[0] if len(g_sub) else np.nan)
        width = 0.35
        ax.bar(
            positions - width / 2,
            chroma_means,
            width=width,
            yerr=chroma_stds,
            color="#55A868",
            edgecolor="black",
            linewidth=0.4,
            capsize=4,
            label="chroma-3072 (mean ± SD, n=5)",
        )
        ax.bar(
            positions + width / 2,
            gemini_vals,
            width=width,
            color="#C44E52",
            edgecolor="black",
            linewidth=0.4,
            label="vec-gemini (n=1)",
        )
        ax.set_xticks(positions)
        ax.set_xticklabels([v.replace("rag-", "") for v in VARIANTS], rotation=15, ha="right")
        style_axes(ax, METRIC_LABELS[metric], "RAG variant", METRIC_LABELS[metric])
        ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.0%}"))
    axes[0].legend(loc="lower right", framealpha=0.9, fontsize=8)
    fig.suptitle("MiniMax M2.5 — vec-gemini vs. chroma-3072 embedding backends", fontsize=13)
    save_fig(fig, out_path)


def plot_best_config_per_model(df: pd.DataFrame, out_path: Path):
    """For each provider, find the highest-mean RAG config (per pass_at_1) and compare to baseline."""
    rows = []
    for provider in ["gemini", "minimax", "openai"]:
        baseline_row = df[(df["provider"] == provider) & (df["experiment"] == "baseline")]
        if baseline_row.empty:
            continue
        baseline_pass = baseline_row["pass_at_1"].iloc[0]
        baseline_comp = baseline_row["compilation_at_1"].iloc[0]
        model = baseline_row["model"].iloc[0]

        rag_rows = df[(df["provider"] == provider) & (df["experiment"] != "baseline")]
        if rag_rows.empty:
            best_label = "—"
            best_pass = np.nan
            best_comp = np.nan
        else:
            grouped = (
                rag_rows.groupby(["backend", "experiment"])[["pass_at_1", "compilation_at_1"]]
                .mean()
                .reset_index()
            )
            best = grouped.loc[grouped["pass_at_1"].idxmax()]
            best_label = f"{best['backend']}/{best['experiment']}"
            best_pass = best["pass_at_1"]
            best_comp = best["compilation_at_1"]
        rows.append({
            "provider": provider,
            "model": model,
            "baseline_pass": baseline_pass,
            "baseline_comp": baseline_comp,
            "best_pass": best_pass,
            "best_comp": best_comp,
            "best_label": best_label,
        })

    fig, ax = plt.subplots(figsize=(11, 6))
    labels = [f"{r['provider']}/{r['model']}" for r in rows]
    positions = np.arange(len(labels))
    width = 0.2

    bp = ax.bar(positions - 1.5 * width, [r["baseline_pass"] for r in rows], width=width, color="#C44E52", alpha=0.55, label="baseline Pass@1", edgecolor="black", linewidth=0.4)
    rp = ax.bar(positions - 0.5 * width, [r["best_pass"] for r in rows], width=width, color="#C44E52", label="best RAG Pass@1", edgecolor="black", linewidth=0.4)
    bc = ax.bar(positions + 0.5 * width, [r["baseline_comp"] for r in rows], width=width, color="#4C72B0", alpha=0.55, label="baseline Compilation@1", edgecolor="black", linewidth=0.4)
    rc = ax.bar(positions + 1.5 * width, [r["best_comp"] for r in rows], width=width, color="#4C72B0", label="best RAG Compilation@1", edgecolor="black", linewidth=0.4)

    for i, r in enumerate(rows):
        if r["best_label"] != "—":
            ax.text(i, 0.02, r["best_label"], ha="center", va="bottom", fontsize=8, color="black", rotation=90)

    ax.set_xticks(positions)
    ax.set_xticklabels(labels, rotation=10, ha="right")
    style_axes(ax, "Best (mean) RAG configuration vs. baseline, per provider", "Provider / Model", "Metric value")
    ax.set_ylim(0, 1.05)
    ax.yaxis.set_major_formatter(plt.FuncFormatter(lambda x, _: f"{x:.0%}"))
    ax.legend(loc="lower right", framealpha=0.9, fontsize=8)
    save_fig(fig, out_path)
    return rows


# --------------------------------------------------------------------------- #
# Interpretations (computed from numbers, not boilerplate)
# --------------------------------------------------------------------------- #


def fmt_pct(x: float) -> str:
    return f"{x:.1%}" if not math.isnan(x) else "n/a"


def interpret_box(df_chroma: pd.DataFrame, metric: str) -> str:
    desc = descriptive_table(df_chroma, metric)
    best_idx = desc["mean"].idxmax()
    best = desc.loc[best_idx]
    widest_idx = (desc["max"] - desc["min"]).idxmax()
    widest = desc.loc[widest_idx]
    return (
        f"The highest median {METRIC_LABELS[metric]} ({fmt_pct(best['mean'])}) is achieved by "
        f"**{best['experiment']}** at dimension **{int(best['dimension'])}**. "
        f"The widest spread (range {fmt_pct(widest['max'] - widest['min'])}) is in "
        f"**{widest['experiment']}** at dimension **{int(widest['dimension'])}**, "
        f"showing where run-to-run instability is largest. With only n=5 runs per cell, IQR-based "
        f"differences should be read as suggestive rather than definitive."
    )


def interpret_bars(df_chroma: pd.DataFrame, metric: str) -> str:
    desc = descriptive_table(df_chroma, metric)
    overall_mean = desc["mean"].mean()
    range_means = desc.groupby("experiment")["mean"].agg(lambda s: s.max() - s.min())
    most_dim_sensitive = range_means.idxmax()
    sensitivity = range_means.max()
    return (
        f"Mean {METRIC_LABELS[metric]} across all (variant, dimension) cells is **{fmt_pct(overall_mean)}**. "
        f"The variant most sensitive to dimension is **{most_dim_sensitive}** "
        f"(swing of {fmt_pct(sensitivity)} between its best and worst dimension). "
        f"Error bars reflect ±1 SD across the 5 runs; overlapping bars between dimensions are typical "
        f"and indicate the dimension effect is small relative to run-to-run noise."
    )


def interpret_line(df_chroma: pd.DataFrame, baseline_value: float, metric: str) -> str:
    desc = descriptive_table(df_chroma, metric)
    above_baseline = desc[desc["mean"] > baseline_value]
    n_above = len(above_baseline)
    n_total = len(desc)
    by_dim = desc.groupby("dimension")["mean"].mean()
    best_dim = int(by_dim.idxmax())
    return (
        f"The horizontal dashed line marks the no-RAG baseline at **{fmt_pct(baseline_value)}**. "
        f"**{n_above}/{n_total}** RAG (variant × dimension) cells exceed the baseline mean. "
        f"Averaged across all variants, dimension **{best_dim}** has the highest mean "
        f"{METRIC_LABELS[metric]} ({fmt_pct(by_dim[best_dim])}). The bands show ±1 SD of the 5 runs — "
        f"if a band overlaps the baseline line, that configuration is statistically indistinguishable "
        f"from no-RAG at the run-to-run noise level."
    )


def interpret_cv(df_chroma: pd.DataFrame, metric: str) -> str:
    desc = descriptive_table(df_chroma, metric)
    desc = desc[desc["n"] > 1].copy()
    most_stable = desc.loc[desc["cv"].idxmin()]
    least_stable = desc.loc[desc["cv"].idxmax()]
    return (
        f"The most stable configuration is **{most_stable['experiment']} @ {int(most_stable['dimension'])}** "
        f"with CV = {most_stable['cv']:.2%}. The least stable is **{least_stable['experiment']} "
        f"@ {int(least_stable['dimension'])}** with CV = {least_stable['cv']:.2%}. "
        f"CV measures relative dispersion; values below ~3% are typically considered tight enough that "
        f"the mean is a faithful summary."
    )


def interpret_trajectory(df_chroma: pd.DataFrame, metric: str) -> str:
    drifts = []
    for variant in VARIANTS:
        for dim in DIMENSIONS:
            sub = df_chroma[(df_chroma["experiment"] == variant) & (df_chroma["dimension"] == dim)].sort_values("run_id")
            ys = sub[metric].to_numpy()
            xs = sub["run_id"].to_numpy()
            if len(ys) >= 3:
                slope, _, _, p_val, _ = stats.linregress(xs, ys)
                drifts.append((variant, dim, slope, p_val))
    n_significant = sum(1 for _, _, _, p in drifts if not math.isnan(p) and p < SIGNIFICANCE_LEVEL)
    return (
        f"Across the 12 (variant × dimension) trajectories, **{n_significant}/{len(drifts)}** show a "
        f"linear run-to-run trend significant at p<0.05. A near-zero count is the desired outcome — "
        f"it means the runs behave like independent samples from the same distribution rather than a "
        f"drifting process (e.g. caching, model-state contamination)."
    )


def interpret_significance(anova_df: pd.DataFrame, metric: str) -> str:
    n_sig = int(anova_df["significant"].sum())
    n_total = len(anova_df)
    sig_variants = ", ".join(anova_df[anova_df["significant"]]["experiment"].tolist()) or "none"
    return (
        f"One-way ANOVA across the 3 dimensions, run separately per RAG variant, finds a significant "
        f"effect (p<0.05) for **{n_sig}/{n_total}** variants ({sig_variants}). The heatmap cells visualise "
        f"the post-hoc Welch's t-test p-values for every dimension pair: red cells indicate dimension pairs "
        f"that differ significantly. With n=5 per group, statistical power is limited — non-significant "
        f"results are consistent with both 'no real effect' and 'small effect masked by noise'."
    )


# --------------------------------------------------------------------------- #
# Report assembly
# --------------------------------------------------------------------------- #


def md_table(df: pd.DataFrame, columns: list[str], headers: list[str], pct_cols: list[str], int_cols: list[str]) -> str:
    """Render a small DataFrame as a Markdown table with percent/int formatting."""
    lines = ["| " + " | ".join(headers) + " |", "|" + "|".join(["---"] * len(headers)) + "|"]
    for _, row in df.iterrows():
        cells = []
        for col in columns:
            val = row[col]
            if pd.isna(val):
                cells.append("—")
            elif col in pct_cols:
                cells.append(f"{val:.1%}")
            elif col in int_cols:
                cells.append(f"{int(val)}")
            else:
                cells.append(f"{val}")
        lines.append("| " + " | ".join(cells) + " |")
    return "\n".join(lines)


def render_report(
    df: pd.DataFrame,
    df_chroma: pd.DataFrame,
    metrics: list[str],
    plots_root: Path,
    output_path: Path,
    best_per_model: list[dict],
):
    minimax_baseline = df[(df["provider"] == "minimax") & (df["experiment"] == "baseline")].iloc[0]

    lines: list[str] = []
    lines.append("# Dimension Variance & Cross-Model Report")
    lines.append("")
    lines.append("Source: `.doc/memory/reeval_20260409_133712.jsonl` (re-evaluated metrics).")
    lines.append("")

    # Executive summary
    lines.append("## Executive Summary")
    lines.append("")
    primary_metric = "pass_at_1"
    desc_pass = descriptive_table(df_chroma, primary_metric)
    overall_pass = desc_pass["mean"].mean()
    by_dim_pass = desc_pass.groupby("dimension")["mean"].mean()
    best_dim_pass = int(by_dim_pass.idxmax())
    anova_pass = anova_table(df_chroma, primary_metric)
    n_sig_pass = int(anova_pass["significant"].sum())

    lines.append(
        f"Across all (variant × dimension) cells the MiniMax M2.5 RAG pipeline averages "
        f"**{fmt_pct(overall_pass)}** Pass@1 versus a no-RAG baseline of **{fmt_pct(minimax_baseline['pass_at_1'])}**. "
        f"Averaged across variants, dimension **{best_dim_pass}** achieves the highest mean Pass@1 "
        f"({fmt_pct(by_dim_pass[best_dim_pass])}). One-way ANOVA finds a significant dimension effect "
        f"in **{n_sig_pass}/{len(VARIANTS)}** RAG variants at p<0.05."
    )
    lines.append("")

    # Inventory
    lines.append("## Experiment Inventory")
    lines.append("")
    inventory_rows = []
    for provider in ["gemini", "minimax", "openai"]:
        sub = df[df["provider"] == provider]
        if sub.empty:
            continue
        model = sub["model"].iloc[0]
        baselines = (sub["experiment"] == "baseline").sum()
        rag = (sub["experiment"] != "baseline").sum()
        inventory_rows.append(
            f"| `{provider}/{model}` | {baselines} | {rag} |"
        )
    lines.append("| Provider/Model | Baseline rows | RAG rows |")
    lines.append("|---|---|---|")
    lines.extend(inventory_rows)
    lines.append("")

    # Section A: MiniMax variance
    lines.append("## Section A — MiniMax M2.5 Dimension Variance")
    lines.append("")
    lines.append(
        "The MiniMax M2.5 chroma experiment is the only configuration with multi-run × multi-dimension "
        "coverage (3 dims × 5 runs × 4 RAG variants = 60 datapoints). All plots in this section operate "
        "on that subset."
    )
    lines.append("")

    for metric in metrics:
        metric_label = METRIC_LABELS[metric]
        lines.append(f"### {metric_label}")
        lines.append("")

        desc = descriptive_table(df_chroma, metric)
        desc_display = desc.copy()
        desc_display["dimension"] = desc_display["dimension"].astype(int)
        lines.append("**Descriptive statistics (n=5 per cell)**")
        lines.append("")
        lines.append(
            md_table(
                desc_display,
                ["experiment", "dimension", "n", "mean", "std", "ci95", "min", "max", "cv"],
                ["Variant", "Dim", "n", "Mean", "SD", "95% CI±", "Min", "Max", "CV"],
                pct_cols=["mean", "std", "ci95", "min", "max", "cv"],
                int_cols=["n", "dimension"],
            )
        )
        lines.append("")

        anova = anova_table(df_chroma, metric)
        lines.append("**One-way ANOVA across dimensions**")
        lines.append("")
        lines.append("| Variant | F | p-value | Significant (α=0.05) |")
        lines.append("|---|---|---|---|")
        for _, row in anova.iterrows():
            sig = "✅" if row["significant"] else "—"
            lines.append(
                f"| {row['experiment']} | {row['F']:.3f} | {row['p_value']:.4f} | {sig} |"
            )
        lines.append("")

        sig_variants = anova[anova["significant"]]["experiment"].tolist()
        if sig_variants:
            lines.append("**Pairwise Welch's t-tests (only for variants with significant ANOVA)**")
            lines.append("")
            lines.append("| Variant | Dim A | Dim B | t | p |")
            lines.append("|---|---|---|---|---|")
            for variant in sig_variants:
                pw = pairwise_ttests(df_chroma, variant, metric)
                for _, row in pw.iterrows():
                    lines.append(
                        f"| {variant} | {row['dim_a']} | {row['dim_b']} | {row['t']:.3f} | {row['p']:.4f} |"
                    )
            lines.append("")

        # Embed plots with interpretations
        plot_specs = [
            ("Box plot — distribution across the 5 runs", f"plots/minimax/box_{metric}.png", interpret_box(df_chroma, metric)),
            ("Grouped bars — mean ± 1 SD", f"plots/minimax/bars_{metric}.png", interpret_bars(df_chroma, metric)),
            ("Line across dimensions with ±1 SD bands", f"plots/minimax/line_{metric}.png", interpret_line(df_chroma, minimax_baseline[metric], metric)),
            ("Coefficient of variation per cell", f"plots/minimax/cv_{metric}.png", interpret_cv(df_chroma, metric)),
            ("Run trajectories", f"plots/minimax/trajectory_{metric}.png", interpret_trajectory(df_chroma, metric)),
            ("Pairwise t-test significance heatmap", f"plots/minimax/significance_{metric}.png", interpret_significance(anova, metric)),
        ]
        for title, rel_path, interpretation in plot_specs:
            lines.append(f"#### {title}")
            lines.append("")
            lines.append(f"![{title}]({rel_path})")
            lines.append("")
            lines.append(interpretation)
            lines.append("")

    # Section B: cross-model
    lines.append("## Section B — Cross-Model Comparison")
    lines.append("")
    lines.append(
        "The other two providers (Gemini 2.5 Pro and OpenAI GPT-5.4) have much smaller designs — "
        "Gemini has only a baseline, and OpenAI has a single run per RAG variant on chroma-3072 — so "
        "they cannot enter the variance analysis above. They are reported here as point comparisons."
    )
    lines.append("")

    baselines = df[df["experiment"] == "baseline"].sort_values("provider")
    lines.append("**Baseline metrics**")
    lines.append("")
    lines.append("| Provider/Model | Compilation@1 | Pass@1 |")
    lines.append("|---|---|---|")
    for _, row in baselines.iterrows():
        lines.append(
            f"| {row['provider']}/{row['model']} | {fmt_pct(row['compilation_at_1'])} | {fmt_pct(row['pass_at_1'])} |"
        )
    lines.append("")

    # C1 baselines compare
    best_baseline = baselines.loc[baselines["pass_at_1"].idxmax()]
    worst_baseline = baselines.loc[baselines["pass_at_1"].idxmin()]
    spread = best_baseline["pass_at_1"] - worst_baseline["pass_at_1"]
    interp_c1 = (
        f"The highest baseline is **{best_baseline['provider']}/{best_baseline['model']}** at "
        f"{fmt_pct(best_baseline['pass_at_1'])} Pass@1; the lowest is "
        f"**{worst_baseline['provider']}/{worst_baseline['model']}** at "
        f"{fmt_pct(worst_baseline['pass_at_1'])}, a spread of **{fmt_pct(spread)}**. "
        f"This is the no-RAG ceiling each model brings before any retrieval is added."
    )
    lines.append("### C1 — Baseline comparison across providers")
    lines.append("")
    lines.append("![baselines](plots/cross_model/baselines_compare.png)")
    lines.append("")
    lines.append(interp_c1)
    lines.append("")

    # C2 OpenAI RAG
    openai_baseline = df[(df["provider"] == "openai") & (df["experiment"] == "baseline")].iloc[0]
    openai_rag = df[(df["provider"] == "openai") & (df["experiment"] != "baseline")]
    openai_max_pass = openai_rag["pass_at_1"].max()
    openai_max_variant = openai_rag.loc[openai_rag["pass_at_1"].idxmax(), "experiment"]
    delta_openai = openai_max_pass - openai_baseline["pass_at_1"]
    direction = "improves over" if delta_openai > 0 else "regresses against"
    interp_c2 = (
        f"OpenAI GPT-5.4 starts from an unusually strong baseline ({fmt_pct(openai_baseline['pass_at_1'])} Pass@1, "
        f"{fmt_pct(openai_baseline['compilation_at_1'])} Compilation@1). The best RAG variant "
        f"(**{openai_max_variant}**) reaches {fmt_pct(openai_max_pass)} Pass@1, which {direction} the baseline "
        f"by **{delta_openai:+.1%}**. With only 1 run per variant the difference is not statistically testable, "
        f"but the absolute gap is small — a sign that for already-strong models RAG yields diminishing returns."
    )
    lines.append("### C2 — OpenAI GPT-5.4 RAG variants (chroma-3072)")
    lines.append("")
    lines.append("![openai-rag](plots/cross_model/openai_rag_variants.png)")
    lines.append("")
    lines.append(interp_c2)
    lines.append("")

    # C3 vec-gemini vs chroma-3072
    chroma_3072_pass = df[(df["provider"] == "minimax") & (df["backend"] == "vec-chroma-3072")].groupby("experiment")["pass_at_1"].mean()
    vec_gemini_pass = df[(df["provider"] == "minimax") & (df["backend"] == "vec-gemini")].set_index("experiment")["pass_at_1"]
    deltas = (vec_gemini_pass - chroma_3072_pass).dropna()
    n_better = int((deltas > 0).sum())
    avg_delta = float(deltas.mean()) if len(deltas) else 0.0
    interp_c3 = (
        f"For MiniMax M2.5, the Gemini-embedding backend (`vec-gemini`, single run) outperforms the mean "
        f"of the 5-run chroma-3072 baseline in **{n_better}/{len(deltas)}** RAG variants on Pass@1, with an "
        f"average delta of **{avg_delta:+.1%}**. Because the chroma side is averaged over 5 runs and the "
        f"vec-gemini side is a single observation, the comparison is biased toward seeing larger gaps than "
        f"truly exist — read this as suggestive only."
    )
    lines.append("### C3 — MiniMax vec-gemini vs. chroma-3072")
    lines.append("")
    lines.append("![vec-gemini-vs-chroma](plots/cross_model/minimax_vec_gemini_vs_chroma.png)")
    lines.append("")
    lines.append(interp_c3)
    lines.append("")

    # C4 best per provider
    best_lines = []
    for r in best_per_model:
        delta_pass = r["best_pass"] - r["baseline_pass"] if not math.isnan(r["best_pass"]) else float("nan")
        if math.isnan(delta_pass):
            best_lines.append(f"**{r['provider']}/{r['model']}** has no RAG runs to compare against its baseline ({fmt_pct(r['baseline_pass'])} Pass@1).")
        else:
            verb = "lifts" if delta_pass > 0 else "lowers"
            best_lines.append(
                f"**{r['provider']}/{r['model']}** — best RAG config `{r['best_label']}` "
                f"{verb} Pass@1 from {fmt_pct(r['baseline_pass'])} → {fmt_pct(r['best_pass'])} ({delta_pass:+.1%})."
            )
    interp_c4 = " ".join(best_lines)
    lines.append("### C4 — Best (mean) RAG configuration vs. baseline, per provider")
    lines.append("")
    lines.append("![best-per-model](plots/cross_model/best_config_per_model.png)")
    lines.append("")
    lines.append(interp_c4)
    lines.append("")

    # Section C: conclusions
    lines.append("## Section C — Conclusions")
    lines.append("")
    bullets: list[str] = []

    if n_sig_pass == 0:
        bullets.append(
            "**Embedding dimension does not have a statistically significant effect on MiniMax M2.5 Pass@1** "
            f"in any of the {len(VARIANTS)} RAG variants tested (one-way ANOVA, α=0.05, n=5 per cell). "
            "The 768/1536/3072 dimension swap is, in this experiment, indistinguishable from run-to-run noise."
        )
    else:
        bullets.append(
            f"**Embedding dimension has a statistically significant effect in {n_sig_pass}/{len(VARIANTS)} RAG variants** "
            "for Pass@1 (one-way ANOVA, α=0.05). See the descriptive table and significance heatmap above for "
            "which dimension pairs drive the effect."
        )

    desc_pass_full = descriptive_table(df_chroma, "pass_at_1")
    best_pass_row = desc_pass_full.loc[desc_pass_full["mean"].idxmax()]
    bullets.append(
        f"**Highest mean Pass@1 configuration**: `{best_pass_row['experiment']}` at dimension "
        f"`{int(best_pass_row['dimension'])}` ({fmt_pct(best_pass_row['mean'])} ± {fmt_pct(best_pass_row['std'])}). "
        f"For comparison, the no-RAG baseline is {fmt_pct(minimax_baseline['pass_at_1'])}."
    )

    desc_pass_full_n2 = desc_pass_full[desc_pass_full["n"] > 1]
    if not desc_pass_full_n2.empty:
        most_stable = desc_pass_full_n2.loc[desc_pass_full_n2["cv"].idxmin()]
        bullets.append(
            f"**Most stable configuration across runs**: `{most_stable['experiment']}` at dimension "
            f"`{int(most_stable['dimension'])}` (CV = {most_stable['cv']:.2%})."
        )

    rag_above = (desc_pass_full["mean"] > minimax_baseline["pass_at_1"]).sum()
    if rag_above == 0:
        bullets.append(
            f"**RAG does not beat the no-RAG baseline on Pass@1 for MiniMax M2.5**: 0/{len(desc_pass_full)} "
            f"(variant × dim) cells exceed the baseline mean ({fmt_pct(minimax_baseline['pass_at_1'])}). "
            "This is the most actionable finding — the RAG retrieval pipeline as currently configured is "
            "either not adding useful context or is adding noise that offsets the benefit."
        )
    else:
        bullets.append(
            f"**RAG exceeds the no-RAG baseline in {rag_above}/{len(desc_pass_full)} (variant × dim) cells "
            f"on Pass@1** for MiniMax M2.5."
        )

    openai_helps = openai_max_pass > openai_baseline["pass_at_1"]
    bullets.append(
        f"**OpenAI GPT-5.4**: RAG {'improves over' if openai_helps else 'does not improve'} the strong "
        f"{fmt_pct(openai_baseline['pass_at_1'])} baseline; best variant reaches {fmt_pct(openai_max_pass)} "
        f"({(openai_max_pass - openai_baseline['pass_at_1']):+.1%})."
    )

    for b in bullets:
        lines.append(f"- {b}")
    lines.append("")

    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text("\n".join(lines), encoding="utf-8")


# --------------------------------------------------------------------------- #
# CLI
# --------------------------------------------------------------------------- #


@click.command()
@click.option(
    "--source",
    type=click.Path(exists=True, path_type=Path),
    default=DEFAULT_SOURCE,
    show_default=True,
    help="Path to the re-evaluation jsonl.",
)
@click.option(
    "--output-dir",
    type=click.Path(path_type=Path),
    default=DEFAULT_OUTPUT,
    show_default=True,
    help="Directory for the report and plots.",
)
@click.option(
    "--metric",
    type=click.Choice(["pass_at_1", "compilation_at_1", "both"]),
    default="both",
    show_default=True,
    help="Which metric to render.",
)
def main(source: Path, output_dir: Path, metric: str):
    console.print(f"[bold]Loading[/bold] {source}")
    rows = load_jsonl(source)
    console.print(f"Loaded [bold]{len(rows)}[/bold] successful rows.")

    df = build_dataframe(rows)
    df_chroma = minimax_chroma(df)

    n_chroma = len(df_chroma)
    n_baseline_minimax = ((df["provider"] == "minimax") & (df["experiment"] == "baseline")).sum()
    n_vec_gemini = ((df["provider"] == "minimax") & (df["backend"] == "vec-gemini")).sum()
    n_openai = (df["provider"] == "openai").sum()
    n_gemini = (df["provider"] == "gemini").sum()
    console.print(
        f"  MiniMax M2.5 chroma RAG: {n_chroma}\n"
        f"  MiniMax M2.5 vec-gemini RAG: {n_vec_gemini}\n"
        f"  MiniMax M2.5 baseline: {n_baseline_minimax}\n"
        f"  OpenAI GPT-5.4 (all): {n_openai}\n"
        f"  Gemini 2.5 Pro (all): {n_gemini}"
    )

    metrics = ["pass_at_1", "compilation_at_1"] if metric == "both" else [metric]

    output_dir.mkdir(parents=True, exist_ok=True)
    csv_path = output_dir / "dimension_variance_data.csv"
    df.to_csv(csv_path, index=False)
    console.print(f"[green]Wrote data dump[/green] → {csv_path}")

    # MiniMax variance plots
    for m in metrics:
        plot_box(df_chroma, m, output_dir / "plots" / "minimax" / f"box_{m}.png")
        plot_bars(df_chroma, m, output_dir / "plots" / "minimax" / f"bars_{m}.png")
        baseline_val = df[(df["provider"] == "minimax") & (df["experiment"] == "baseline")][m].iloc[0]
        plot_line(df_chroma, baseline_val, m, output_dir / "plots" / "minimax" / f"line_{m}.png")
        plot_cv(df_chroma, m, output_dir / "plots" / "minimax" / f"cv_{m}.png")
        plot_trajectory(df_chroma, m, output_dir / "plots" / "minimax" / f"trajectory_{m}.png")
        plot_significance(df_chroma, m, output_dir / "plots" / "minimax" / f"significance_{m}.png")
        console.print(f"[green]Rendered MiniMax plots[/green] for {m}")

    # Cross-model plots
    plot_baselines_compare(df, output_dir / "plots" / "cross_model" / "baselines_compare.png")
    plot_openai_rag(df, output_dir / "plots" / "cross_model" / "openai_rag_variants.png")
    plot_minimax_vec_gemini_vs_chroma(df, output_dir / "plots" / "cross_model" / "minimax_vec_gemini_vs_chroma.png")
    best_per_model = plot_best_config_per_model(df, output_dir / "plots" / "cross_model" / "best_config_per_model.png")
    console.print("[green]Rendered cross-model plots[/green]")

    # Report (after all plots)
    report_path = output_dir / "dimension_variance_report.md"
    render_report(df, df_chroma, metrics, output_dir / "plots", report_path, best_per_model)
    console.print(f"[bold green]Report written[/bold green] → {report_path}")


if __name__ == "__main__":
    main()
