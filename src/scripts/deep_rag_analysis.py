#!/usr/bin/env python3
"""Deep analysis of RAG failure modes for the thesis Discussion chapter.

Reads existing per-task result.json, retrieval.json, and prompt.json files
across 5 variants x 6 runs x 164 tasks. Produces CSV tables and a markdown
report in final/results/deep_analysis/.
"""

import json
import math
import re
import sys
from collections import Counter, defaultdict
from pathlib import Path

try:
    import numpy as np
    import pandas as pd
except ImportError:
    print("ERROR: numpy and pandas are required. Install with: pip install numpy pandas")
    sys.exit(1)

REPO_ROOT = Path(__file__).resolve().parents[2]
THESIS_ROOT = REPO_ROOT.parent / "final"
OUT_DIR = THESIS_ROOT / "results" / "deep_analysis"

TARGET_ROOT = REPO_ROOT / "data" / "translation" / "target" / "humaneval-x" / "minimax" / "M2.5"
BASELINE_ROOT = TARGET_ROOT / "baseline"
RAG_ROOT = TARGET_ROOT / "vec-chroma-3072"

VARIANTS = ["baseline", "rag-pattern-only", "rag-pattern-samples", "rag-pattern-api-docs", "rag-full"]
RAG_VARIANTS = [v for v in VARIANTS if v != "baseline"]
RUNS = list(range(1, 7))
TASKS = list(range(164))

KB_PROCESSED = REPO_ROOT / "data" / "RAG" / "processed"


def task_dir(variant: str, run: int, task_num: int) -> Path:
    if variant == "baseline":
        return BASELINE_ROOT / f"run-{run}" / "tasks" / f"Go_{task_num}"
    return RAG_ROOT / f"run-{run}" / variant / "tasks" / f"Go_{task_num}"


# ── Data Loading ──────────────────────────────────────────────────────────

def load_all_outcomes() -> pd.DataFrame:
    rows = []
    for variant in VARIANTS:
        for run in RUNS:
            for t in TASKS:
                rfile = task_dir(variant, run, t) / "evaluation" / "result.json"
                if not rfile.exists():
                    continue
                r = json.loads(rfile.read_text())
                rows.append({
                    "task_id": f"Go_{t}",
                    "task_num": t,
                    "variant": variant,
                    "run": run,
                    "compiles": bool(r.get("compiles", False)),
                    "pass_at_1": bool(r.get("pass_at_1", False)),
                    "notes": r.get("notes", ""),
                })
    return pd.DataFrame(rows)


def load_all_retrieval() -> pd.DataFrame:
    rows = []
    run = 1
    for variant in RAG_VARIANTS:
        for t in TASKS:
            rfile = task_dir(variant, run, t) / "retrieval.json"
            if not rfile.exists():
                continue
            r = json.loads(rfile.read_text())
            summary = r.get("summary", {})
            counts = r.get("retrieval_counts", {})
            ctx = r.get("rendered_context", "")
            rows.append({
                "task_id": f"Go_{t}",
                "task_num": t,
                "variant": variant,
                "has_usable_items": bool(summary.get("has_usable_items", False)),
                "total_accepted": summary.get("total_accepted", 0),
                "total_candidates": summary.get("total_candidates", 0),
                "grammar_accepted": counts.get("grammar_mappings", 0),
                "parallel_accepted": counts.get("parallel_corpus", 0),
                "api_accepted": counts.get("api_mappings", 0),
                "docs_accepted": counts.get("documentation", 0),
                "traps_accepted": counts.get("translation_traps", 0),
                "rendered_context_len": len(ctx) if ctx else 0,
            })
    return pd.DataFrame(rows)


def load_all_items() -> pd.DataFrame:
    rows = []
    run = 1
    for variant in RAG_VARIANTS:
        for t in TASKS:
            rfile = task_dir(variant, run, t) / "retrieval.json"
            if not rfile.exists():
                continue
            r = json.loads(rfile.read_text())
            items_dict = r.get("items", {})
            for source, items in items_dict.items():
                if not isinstance(items, list):
                    continue
                for item in items:
                    retr = item.get("retrieval", {})
                    rows.append({
                        "task_id": f"Go_{t}",
                        "task_num": t,
                        "variant": variant,
                        "source": source,
                        "doc_id": item.get("_id", ""),
                        "dense_distance": retr.get("dense_distance"),
                        "dense_rank": retr.get("dense_rank"),
                        "merged_rank": retr.get("merged_rank"),
                        "accepted": bool(retr.get("accepted", False)),
                    })
    return pd.DataFrame(rows)


# ── Analysis 1: Task Difficulty Profile ───────────────────────────────────

def analyze_task_difficulty(outcomes: pd.DataFrame) -> dict:
    print("\n═══ Analysis 1: Task Difficulty Profile ═══")

    task_pass = outcomes.groupby("task_num")["pass_at_1"].mean().reset_index()
    task_pass.columns = ["task_num", "overall_pass_rate"]

    def classify(rate):
        if rate == 1.0:
            return "always_easy"
        elif rate == 0.0:
            return "always_hard"
        else:
            return "swing"

    task_pass["category"] = task_pass["overall_pass_rate"].apply(classify)

    counts = task_pass["category"].value_counts()
    print(f"  Always-easy (30/30 pass): {counts.get('always_easy', 0)}")
    print(f"  Always-hard (0/30 pass):  {counts.get('always_hard', 0)}")
    print(f"  Swing (1-29/30):          {counts.get('swing', 0)}")

    per_variant = outcomes.groupby(["task_num", "variant"])["pass_at_1"].mean().reset_index()
    per_variant.columns = ["task_num", "variant", "pass_rate"]
    pivot = per_variant.pivot(index="task_num", columns="variant", values="pass_rate")
    task_profile = task_pass.merge(pivot, on="task_num")
    task_profile = task_profile.sort_values("task_num")

    task_profile.to_csv(OUT_DIR / "task_difficulty_profile.csv", index=False)

    summary_rows = []
    for cat in ["always_easy", "swing", "always_hard"]:
        cat_tasks = task_profile[task_profile["category"] == cat]
        row = {"category": cat, "count": len(cat_tasks)}
        for v in VARIANTS:
            if v in cat_tasks.columns:
                row[f"{v}_mean_pass"] = f"{cat_tasks[v].mean() * 100:.1f}%"
        summary_rows.append(row)
    summary_df = pd.DataFrame(summary_rows)
    summary_df.to_csv(OUT_DIR / "difficulty_summary.csv", index=False)

    print(f"\n  Per-variant pass rates on SWING tasks only:")
    swing_tasks = task_profile[task_profile["category"] == "swing"]
    for v in VARIANTS:
        if v in swing_tasks.columns:
            print(f"    {v:30s}: {swing_tasks[v].mean() * 100:.1f}%")

    return {
        "counts": counts.to_dict(),
        "swing_count": counts.get("swing", 0),
        "always_easy_count": counts.get("always_easy", 0),
        "always_hard_count": counts.get("always_hard", 0),
        "task_profile": task_profile,
    }


# ── Analysis 2: Retrieval Coverage ────────────────────────────────────────

def analyze_retrieval_coverage(retrieval: pd.DataFrame, outcomes: pd.DataFrame) -> dict:
    print("\n═══ Analysis 2: Retrieval Coverage & Effective Fallback ═══")

    coverage_rows = []
    for v in RAG_VARIANTS:
        vr = retrieval[retrieval["variant"] == v]
        n_total = len(vr)
        n_with = (vr["total_accepted"] > 0).sum()
        n_empty = (vr["total_accepted"] == 0).sum()
        coverage_rows.append({
            "variant": v,
            "tasks_with_items": int(n_with),
            "tasks_empty": int(n_empty),
            "coverage_pct": f"{n_with / n_total * 100:.1f}%" if n_total > 0 else "N/A",
            "mean_accepted": f"{vr['total_accepted'].mean():.2f}",
            "median_accepted": f"{vr['total_accepted'].median():.1f}",
        })
        print(f"  {v:30s}: {n_empty:3d} empty, {n_with:3d} with items, coverage={n_with/n_total*100:.1f}%")

    coverage_df = pd.DataFrame(coverage_rows)
    coverage_df.to_csv(OUT_DIR / "retrieval_coverage.csv", index=False)

    print(f"\n  Pass rate: items vs no-items (majority of 6 runs):")
    majority = outcomes.groupby(["task_num", "variant"])["pass_at_1"].mean().reset_index()
    majority.columns = ["task_num", "variant", "pass_rate"]
    majority["majority_pass"] = majority["pass_rate"] > 0.5

    cov_outcome_rows = []
    for v in RAG_VARIANTS:
        vr = retrieval[retrieval["variant"] == v][["task_num", "total_accepted"]].copy()
        vr["has_items"] = vr["total_accepted"] > 0
        vm = majority[majority["variant"] == v][["task_num", "pass_rate", "majority_pass"]]
        merged = vr.merge(vm, on="task_num")

        with_items = merged[merged["has_items"]]
        no_items = merged[~merged["has_items"]]

        row = {
            "variant": v,
            "with_items_count": len(with_items),
            "with_items_pass_rate": f"{with_items['pass_rate'].mean() * 100:.1f}%" if len(with_items) > 0 else "N/A",
            "no_items_count": len(no_items),
            "no_items_pass_rate": f"{no_items['pass_rate'].mean() * 100:.1f}%" if len(no_items) > 0 else "N/A",
        }
        cov_outcome_rows.append(row)
        with_str = f"{with_items['pass_rate'].mean()*100:.1f}%" if len(with_items) > 0 else "N/A"
        no_str = f"{no_items['pass_rate'].mean()*100:.1f}%" if len(no_items) > 0 else "N/A"
        print(f"    {v:30s}: with_items={with_str} ({len(with_items)}), no_items={no_str} ({len(no_items)})")

    cov_outcome_df = pd.DataFrame(cov_outcome_rows)
    cov_outcome_df.to_csv(OUT_DIR / "coverage_vs_outcome.csv", index=False)

    return {"coverage": coverage_df, "coverage_vs_outcome": cov_outcome_df}


# ── Analysis 3: Retrieval Relevance vs Outcome ────────────────────────────

def analyze_retrieval_relevance(items: pd.DataFrame, retrieval: pd.DataFrame, outcomes: pd.DataFrame) -> dict:
    print("\n═══ Analysis 3: Retrieval Relevance vs Outcome ═══")

    if items.empty or "dense_distance" not in items.columns:
        print("  No item-level data available, skipping.")
        return {}

    accepted = items[items["accepted"] == True].copy()
    if accepted.empty:
        print("  No accepted items found, skipping.")
        return {}

    task_quality = accepted.groupby(["task_num", "variant"]).agg(
        mean_distance=("dense_distance", "mean"),
        min_distance=("dense_distance", "min"),
        item_count=("doc_id", "count"),
    ).reset_index()

    majority = outcomes.groupby(["task_num", "variant"])["pass_at_1"].mean().reset_index()
    majority.columns = ["task_num", "variant", "pass_rate"]
    majority["majority_pass"] = majority["pass_rate"] > 0.5

    merged = task_quality.merge(majority, on=["task_num", "variant"])

    merged.to_csv(OUT_DIR / "retrieval_quality_vs_outcome.csv", index=False)

    for v in RAG_VARIANTS:
        vm = merged[merged["variant"] == v]
        if len(vm) < 5:
            continue
        valid = vm.dropna(subset=["mean_distance"])
        if len(valid) < 5:
            continue
        corr = valid["mean_distance"].corr(valid["pass_rate"])
        print(f"  {v:30s}: distance-vs-pass correlation = {corr:+.3f} (n={len(valid)})")

    valid_all = merged.dropna(subset=["mean_distance"])
    if len(valid_all) > 10:
        q1 = valid_all["mean_distance"].quantile(0.25)
        q3 = valid_all["mean_distance"].quantile(0.75)
        best_q = valid_all[valid_all["mean_distance"] <= q1]
        worst_q = valid_all[valid_all["mean_distance"] >= q3]
        print(f"\n  Quartile comparison (all variants pooled):")
        print(f"    Best quartile  (distance <= {q1:.3f}): pass_rate = {best_q['pass_rate'].mean()*100:.1f}% (n={len(best_q)})")
        print(f"    Worst quartile (distance >= {q3:.3f}): pass_rate = {worst_q['pass_rate'].mean()*100:.1f}% (n={len(worst_q)})")

        quartile_df = pd.DataFrame([
            {"quartile": "best (Q1)", "threshold": f"<= {q1:.3f}", "mean_pass_rate": f"{best_q['pass_rate'].mean()*100:.1f}%", "n": len(best_q)},
            {"quartile": "worst (Q4)", "threshold": f">= {q3:.3f}", "mean_pass_rate": f"{worst_q['pass_rate'].mean()*100:.1f}%", "n": len(worst_q)},
        ])
        quartile_df.to_csv(OUT_DIR / "quality_quartile_summary.csv", index=False)

    return {"merged": merged}


# ── Analysis 4: Error Signature Clustering ────────────────────────────────

def classify_error(notes: str) -> str:
    if not notes:
        return "unknown"
    if "imported and not used" in notes:
        return "unused_import"
    if "declared and not used" in notes or "declared but not used" in notes:
        return "unused_var"
    if "Timed out" in notes or "timed out" in notes:
        return "timeout"
    if "signal: killed" in notes:
        return "oom"
    if "undefined:" in notes:
        return "undefined_id"
    if "invalid operation" in notes or "mismatched types" in notes:
        return "type_error"
    if "cannot use" in notes or "cannot convert" in notes:
        return "type_error"
    if "assignment mismatch" in notes:
        return "assignment_mismatch"
    if "expected" in notes and ("declaration" in notes or "}" in notes or ";" in notes):
        return "syntax_error"
    if "has no field or method" in notes:
        return "wrong_method"
    if "too many arguments" in notes or "not enough arguments" in notes:
        return "arg_mismatch"
    return "other"


def analyze_error_signatures(outcomes: pd.DataFrame) -> dict:
    print("\n═══ Analysis 4: Error Signature Clustering ═══")

    outcomes = outcomes.copy()
    outcomes["outcome"] = "pass"
    outcomes.loc[outcomes["pass_at_1"] == False, "outcome"] = "test_fail"
    outcomes.loc[outcomes["compiles"] == False, "outcome"] = "compile_fail"

    outcomes["error_category"] = "none"
    compile_mask = outcomes["outcome"] == "compile_fail"
    outcomes.loc[compile_mask, "error_category"] = outcomes.loc[compile_mask, "notes"].apply(classify_error)

    summary = outcomes.groupby("variant")["outcome"].value_counts().unstack(fill_value=0)
    print(f"\n  Outcome counts (across 6 runs):")
    print(summary.to_string())

    error_cross = outcomes[compile_mask].groupby(["variant", "error_category"]).size().unstack(fill_value=0)
    print(f"\n  Compile error categories (across 6 runs):")
    print(error_cross.to_string())

    summary.to_csv(OUT_DIR / "error_categories.csv")
    error_cross.to_csv(OUT_DIR / "error_category_detail.csv")

    return {"outcome_summary": summary, "error_cross": error_cross}


# ── Analysis 5: Task-Level Flip Analysis ──────────────────────────────────

def analyze_task_flips(outcomes: pd.DataFrame, retrieval: pd.DataFrame) -> dict:
    print("\n═══ Analysis 5: Task-Level Flip Analysis (Multi-Run) ═══")

    baseline_passes = outcomes[outcomes["variant"] == "baseline"].groupby("task_num")["pass_at_1"].sum().reset_index()
    baseline_passes.columns = ["task_num", "baseline_passes"]

    flip_rows = []
    for v in RAG_VARIANTS:
        rag_passes = outcomes[outcomes["variant"] == v].groupby("task_num")["pass_at_1"].sum().reset_index()
        rag_passes.columns = ["task_num", "rag_passes"]
        merged = baseline_passes.merge(rag_passes, on="task_num")
        merged["delta"] = merged["rag_passes"] - merged["baseline_passes"]
        merged["variant"] = v

        def classify_flip(row):
            if row["delta"] > 0:
                return "recovery"
            elif row["delta"] < 0:
                return "regression"
            else:
                return "no_change"

        merged["flip_category"] = merged.apply(classify_flip, axis=1)
        flip_rows.append(merged)

    flip_df = pd.concat(flip_rows, ignore_index=True)
    flip_df["task_id"] = flip_df["task_num"].apply(lambda x: f"Go_{x}")
    flip_df.to_csv(OUT_DIR / "task_flip_summary.csv", index=False)

    for v in RAG_VARIANTS:
        vf = flip_df[flip_df["variant"] == v]
        cats = vf["flip_category"].value_counts()
        print(f"  {v:30s}: recovery={cats.get('recovery',0):3d}, regression={cats.get('regression',0):3d}, no_change={cats.get('no_change',0):3d}")

    top_reg = flip_df[flip_df["delta"] < 0].sort_values("delta").head(15).copy()
    top_rec = flip_df[flip_df["delta"] > 0].sort_values("delta", ascending=False).head(15).copy()

    def enrich_with_retrieval(df: pd.DataFrame) -> pd.DataFrame:
        ret_info = []
        for _, row in df.iterrows():
            v = row["variant"]
            t = int(row["task_num"])
            rfile = task_dir(v, 1, t) / "retrieval.json"
            info = {"grammar": 0, "parallel": 0, "api": 0, "docs": 0, "total": 0}
            if rfile.exists():
                r = json.loads(rfile.read_text())
                counts = r.get("retrieval_counts", {})
                info["grammar"] = counts.get("grammar_mappings", 0)
                info["parallel"] = counts.get("parallel_corpus", 0)
                info["api"] = counts.get("api_mappings", 0)
                info["docs"] = counts.get("documentation", 0)
                info["total"] = r.get("summary", {}).get("total_accepted", 0)
            ret_info.append(info)
        ret_df = pd.DataFrame(ret_info)
        return pd.concat([df.reset_index(drop=True), ret_df], axis=1)

    top_reg = enrich_with_retrieval(top_reg)
    top_rec = enrich_with_retrieval(top_rec)

    top_reg.to_csv(OUT_DIR / "top_regressions.csv", index=False)
    top_rec.to_csv(OUT_DIR / "top_recoveries.csv", index=False)

    print(f"\n  Top 5 regressions (baseline_passes - rag_passes):")
    for _, row in top_reg.head(5).iterrows():
        print(f"    {row['task_id']:8s} {row['variant']:30s}: {int(row['baseline_passes'])}→{int(row['rag_passes'])} (delta={int(row['delta'])}), retrieved: g={int(row['grammar'])} p={int(row['parallel'])} a={int(row['api'])} d={int(row['docs'])}")

    print(f"\n  Top 5 recoveries:")
    for _, row in top_rec.head(5).iterrows():
        print(f"    {row['task_id']:8s} {row['variant']:30s}: {int(row['baseline_passes'])}→{int(row['rag_passes'])} (delta=+{int(row['delta'])}), retrieved: g={int(row['grammar'])} p={int(row['parallel'])} a={int(row['api'])} d={int(row['docs'])}")

    return {"flip_df": flip_df, "top_regressions": top_reg, "top_recoveries": top_rec}


# ── Analysis 6: Prompt Size vs Outcome ────────────────────────────────────

def analyze_prompt_size(outcomes: pd.DataFrame, retrieval: pd.DataFrame) -> dict:
    print("\n═══ Analysis 6: Prompt Size vs Outcome ═══")

    rows = []
    run = 1
    for v in RAG_VARIANTS:
        for t in TASKS:
            pfile = task_dir(v, run, t) / "prompt.json"
            if not pfile.exists():
                continue
            p = json.loads(pfile.read_text())
            user_prompt_len = len(p.get("user_prompt", ""))
            system_prompt_len = len(p.get("system_prompt", ""))
            rows.append({
                "task_id": f"Go_{t}",
                "task_num": t,
                "variant": v,
                "user_prompt_len": user_prompt_len,
                "system_prompt_len": system_prompt_len,
                "total_prompt_len": user_prompt_len + system_prompt_len,
            })

    prompt_df = pd.DataFrame(rows)

    ctx_col = retrieval[["task_num", "variant", "rendered_context_len"]].copy()
    prompt_df = prompt_df.merge(ctx_col, on=["task_num", "variant"], how="left")

    majority = outcomes.groupby(["task_num", "variant"])["pass_at_1"].mean().reset_index()
    majority.columns = ["task_num", "variant", "pass_rate"]
    prompt_df = prompt_df.merge(majority, on=["task_num", "variant"], how="left")

    prompt_df.to_csv(OUT_DIR / "prompt_size_analysis.csv", index=False)

    print(f"\n  Correlation (total_prompt_len vs pass_rate) per variant:")
    for v in RAG_VARIANTS:
        vp = prompt_df[prompt_df["variant"] == v]
        if len(vp) < 5:
            continue
        corr = vp["total_prompt_len"].corr(vp["pass_rate"])
        print(f"    {v:30s}: r = {corr:+.3f} (n={len(vp)})")

    print(f"\n  Median-split comparison (all variants pooled):")
    median_ctx = prompt_df["rendered_context_len"].median()
    above = prompt_df[prompt_df["rendered_context_len"] > median_ctx]
    below = prompt_df[prompt_df["rendered_context_len"] <= median_ctx]
    print(f"    Context <= {median_ctx:.0f} chars: pass_rate = {below['pass_rate'].mean()*100:.1f}% (n={len(below)})")
    print(f"    Context >  {median_ctx:.0f} chars: pass_rate = {above['pass_rate'].mean()*100:.1f}% (n={len(above)})")

    return {"prompt_df": prompt_df}


# ── Analysis 7: Source-Specific Hit Rate & KB Utilization ─────────────────

def analyze_source_hit_rates(retrieval: pd.DataFrame, items: pd.DataFrame, outcomes: pd.DataFrame) -> dict:
    print("\n═══ Analysis 7: Source-Specific Hit Rate & KB Utilization ═══")

    kb_sizes = {}
    for name, filename in [
        ("grammar_mappings", "grammar_mappings.jsonl"),
        ("api_mappings", "api_mappings.jsonl"),
        ("documentation", "go_docs.jsonl"),
        ("parallel_corpus", "translation_traps_codenet_v3.jsonl"),
    ]:
        f = KB_PROCESSED / filename
        if f.exists():
            kb_sizes[name] = sum(1 for _ in f.open())

    parallel_file = KB_PROCESSED / "translation_traps_codenet_v3.jsonl"
    for candidate in ["python_go_pairs.jsonl"]:
        pf = KB_PROCESSED / "parallel_corpus" / "codeNet" / candidate
        if pf.exists():
            kb_sizes["parallel_corpus"] = sum(1 for _ in pf.open())
            break
    if "parallel_corpus" not in kb_sizes:
        kb_sizes["parallel_corpus"] = 1668

    print(f"  KB sizes: {kb_sizes}")

    full_retr = retrieval[retrieval["variant"] == "rag-full"]
    source_cols = {
        "grammar_mappings": "grammar_accepted",
        "api_mappings": "api_accepted",
        "documentation": "docs_accepted",
        "parallel_corpus": "parallel_accepted",
    }

    hit_rows = []
    for source, col in source_cols.items():
        if col not in full_retr.columns:
            continue
        queried = len(full_retr)
        with_items = (full_retr[col] > 0).sum()
        total_items = full_retr[col].sum()
        hit_rows.append({
            "source": source,
            "kb_size": kb_sizes.get(source, "?"),
            "tasks_queried": queried,
            "tasks_with_items": int(with_items),
            "hit_rate_pct": f"{with_items/queried*100:.1f}%",
            "total_items_retrieved": int(total_items),
            "mean_items_per_task": f"{total_items/queried:.2f}",
        })
        print(f"  {source:20s}: {with_items:3d}/{queried} tasks ({with_items/queried*100:.1f}%), total_items={int(total_items)}, KB_size={kb_sizes.get(source, '?')}")

    hit_df = pd.DataFrame(hit_rows)
    hit_df.to_csv(OUT_DIR / "source_hit_rates.csv", index=False)

    if not items.empty:
        full_items = items[items["variant"] == "rag-full"]
        accepted_items = full_items[full_items["accepted"] == True]
        util_rows = []
        for source in source_cols:
            src_items = accepted_items[accepted_items["source"] == source]
            unique_ids = src_items["doc_id"].nunique()
            kb_size = kb_sizes.get(source, 0)
            util_rows.append({
                "source": source,
                "kb_size": kb_size,
                "unique_items_retrieved": unique_ids,
                "utilization_pct": f"{unique_ids/kb_size*100:.1f}%" if kb_size > 0 else "N/A",
            })
            print(f"    {source:20s}: {unique_ids} unique items from {kb_size} KB records ({unique_ids/kb_size*100:.1f}% utilization)" if kb_size > 0 else f"    {source}: {unique_ids} unique items")

        util_df = pd.DataFrame(util_rows)
        util_df.to_csv(OUT_DIR / "kb_utilization.csv", index=False)

    majority = outcomes.groupby(["task_num", "variant"])["pass_at_1"].mean().reset_index()
    majority.columns = ["task_num", "variant", "pass_rate"]
    full_majority = majority[majority["variant"] == "rag-full"]

    print(f"\n  Pass rate by source presence (rag-full):")
    for source, col in source_cols.items():
        src_tasks = full_retr[["task_num", col]].copy()
        src_tasks["has_source"] = src_tasks[col] > 0
        m = src_tasks.merge(full_majority[["task_num", "pass_rate"]], on="task_num")
        with_src = m[m["has_source"]]
        without_src = m[~m["has_source"]]
        w_str = f"{with_src['pass_rate'].mean()*100:.1f}%" if len(with_src) > 0 else "N/A"
        wo_str = f"{without_src['pass_rate'].mean()*100:.1f}%" if len(without_src) > 0 else "N/A"
        print(f"    {source:20s}: with={w_str} (n={len(with_src)}), without={wo_str} (n={len(without_src)})")

    return {"hit_rates": hit_df}


# ── Report Synthesis ──────────────────────────────────────────────────────

def write_report(r1, r2, r3, r4, r5, r6, r7):
    lines = []
    lines.append("# Deep RAG Analysis Report")
    lines.append("")
    lines.append("This report examines 4,920 translation outcomes (164 tasks x 6 runs x 5 conditions)")
    lines.append("to understand WHY RAG did not improve Python-to-Go translation on HumanEval-X.")
    lines.append("")

    lines.append("## 1. Task Difficulty Profile")
    lines.append("")
    if r1:
        lines.append(f"- **Always-easy** (pass in all 30 outcomes): {r1['always_easy_count']} tasks")
        lines.append(f"- **Always-hard** (fail in all 30 outcomes): {r1['always_hard_count']} tasks")
        lines.append(f"- **Swing** (pass rate 0-100%): {r1['swing_count']} tasks")
        lines.append("")
        lines.append("RAG can only matter for swing tasks. The always-easy tasks are ceiling effects")
        lines.append("where retrieval cannot improve the outcome.")
        lines.append("")
        lines.append("See `task_difficulty_profile.csv` and `difficulty_summary.csv` for full data.")
    lines.append("")

    lines.append("## 2. Retrieval Coverage")
    lines.append("")
    lines.append("When retrieval returns zero items, the prompt is identical to the baseline.")
    lines.append("")
    if r2 and "coverage" in r2:
        lines.append("| Variant | Empty | With Items | Coverage |")
        lines.append("|---------|-------|------------|----------|")
        for _, row in r2["coverage"].iterrows():
            lines.append(f"| {row['variant']} | {row['tasks_empty']} | {row['tasks_with_items']} | {row['coverage_pct']} |")
    lines.append("")
    lines.append("See `retrieval_coverage.csv` and `coverage_vs_outcome.csv` for full data.")
    lines.append("")

    lines.append("## 3. Retrieval Relevance vs Outcome")
    lines.append("")
    lines.append("See `retrieval_quality_vs_outcome.csv` and `quality_quartile_summary.csv`.")
    lines.append("")

    lines.append("## 4. Error Signature Clustering")
    lines.append("")
    if r4 and "outcome_summary" in r4:
        lines.append("### Outcome distribution across variants (6 runs)")
        lines.append("")
        lines.append(r4["outcome_summary"].to_string())
        lines.append("")
    if r4 and "error_cross" in r4:
        lines.append("### Compile error categories")
        lines.append("")
        lines.append(r4["error_cross"].to_string())
    lines.append("")

    lines.append("## 5. Task-Level Flip Analysis")
    lines.append("")
    lines.append("See `task_flip_summary.csv`, `top_regressions.csv`, `top_recoveries.csv`.")
    lines.append("")

    lines.append("## 6. Prompt Size vs Outcome")
    lines.append("")
    lines.append("See `prompt_size_analysis.csv`.")
    lines.append("")

    lines.append("## 7. Source-Specific Hit Rate & KB Utilization")
    lines.append("")
    lines.append("See `source_hit_rates.csv` and `kb_utilization.csv`.")
    lines.append("")

    (OUT_DIR / "deep_analysis_report.md").write_text("\n".join(lines))
    print(f"\n  Report written to {OUT_DIR / 'deep_analysis_report.md'}")


# ── Main ──────────────────────────────────────────────────────────────────

def main():
    OUT_DIR.mkdir(parents=True, exist_ok=True)
    print(f"Output directory: {OUT_DIR}")

    print("\nLoading outcomes...")
    outcomes = load_all_outcomes()
    print(f"  Loaded {len(outcomes)} outcome rows")

    print("Loading retrieval data (run-1 only, deterministic)...")
    retrieval = load_all_retrieval()
    print(f"  Loaded {len(retrieval)} retrieval rows")

    print("Loading item-level data...")
    items = load_all_items()
    print(f"  Loaded {len(items)} item rows")

    r1 = analyze_task_difficulty(outcomes)
    r2 = analyze_retrieval_coverage(retrieval, outcomes)
    r3 = analyze_retrieval_relevance(items, retrieval, outcomes)
    r4 = analyze_error_signatures(outcomes)
    r5 = analyze_task_flips(outcomes, retrieval)
    r6 = analyze_prompt_size(outcomes, retrieval)
    r7 = analyze_source_hit_rates(retrieval, items, outcomes)

    write_report(r1, r2, r3, r4, r5, r6, r7)

    print("\n═══ DONE ═══")
    print(f"All outputs in: {OUT_DIR}")


if __name__ == "__main__":
    main()
