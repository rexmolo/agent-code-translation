"""HumanEval-X data loader for Python→Go translation evaluation.

Loads the THUDM/humaneval-x dataset from HuggingFace and returns
paired Python/Go problems for translation and evaluation.

Each problem has:
    task_id:        e.g. "Go/0"
    declaration:    Go function signature (target)
    py_declaration: Python function signature (source)
    py_solution:    Full Python code (prompt + canonical_solution)
    go_solution:    Full Go code (ground truth)
    test:           Go test suite for evaluation
"""

from __future__ import annotations


_HUMANEVAL_X_JSONL = (
    "https://huggingface.co/datasets/THUDM/humaneval-x"
    "/resolve/main/data/{lang}/data/humaneval.jsonl"
)


def load_humaneval_x() -> list[dict]:
    """Load HumanEval-X Python→Go translation pairs.

    Loads JSONL files directly from the THUDM/humaneval-x repository
    to avoid the deprecated dataset loading script.

    Returns 164 problems, each a dict with keys:
        task_id, declaration, py_declaration, py_solution, go_solution, test
    """
    from datasets import load_dataset

    py_ds = load_dataset(
        "json",
        data_files=_HUMANEVAL_X_JSONL.format(lang="python"),
        split="train",
    )
    go_ds = load_dataset(
        "json",
        data_files=_HUMANEVAL_X_JSONL.format(lang="go"),
        split="train",
    )

    pairs = []
    for py_item, go_item in zip(py_ds, go_ds):
        pairs.append({
            "task_id": go_item["task_id"],
            "declaration": go_item["declaration"],
            "py_declaration": py_item["declaration"],
            "py_solution": py_item["prompt"] + py_item["canonical_solution"],
            "go_solution": go_item["prompt"] + go_item["canonical_solution"],
            "test": go_item["test"],
        })
    return pairs
