"""
Shared path configuration for all scripts and experiments.

Usage:
    from src.config import REPO_ROOT, TEMP_DIR, DATA_DIR
"""

from pathlib import Path

import yaml

# Project root (the 'experiments' directory)
REPO_ROOT = Path(__file__).resolve().parent.parent

# --- Source directories ---
SCRIPTS_DIR = REPO_ROOT / "src" / "scripts"
TEMP_DIR = REPO_ROOT / "src" / "temp"

# --- Data directories ---
DATA_DIR = REPO_ROOT / "data"
TRANSLATION_SOURCE_DIR = DATA_DIR / "translation" / "source"
TRANSLATION_TARGET_DIR = DATA_DIR / "translation" / "target"
LOCAL_TARGET_DIR = TRANSLATION_TARGET_DIR / "local"
HUMANEVAL_X_TARGET_DIR = TRANSLATION_TARGET_DIR / "humaneval-x"
RAG_PROCESSED_DIR = DATA_DIR / "RAG" / "processed"
ERROR_DB_PATH = DATA_DIR / "errors.db"

# --- RAG ---
PARALLEL_CORPUS_FILE = RAG_PROCESSED_DIR / "parallel_corpus" / "codeNet" / "python_go_pairs.jsonl"
API_MAPPINGS_FILE = RAG_PROCESSED_DIR / "api_mappings.jsonl"
GO_DOCS_FILE = RAG_PROCESSED_DIR / "go_docs.jsonl"

# --- Config files ---
EVAL_CONFIG_PATH = REPO_ROOT / "config" / "eval_config.yaml"

_EVAL_CONFIG_DEFAULTS = {
    "translation": {"batch_size": 5},
    "parallel": {"batch_size": 10},
    "docker": {"image": "golang:1.26-alpine", "memory_limit": "512m", "timeout": 60},
}


def load_eval_config() -> dict:
    """Load evaluation config from YAML, falling back to defaults."""
    if EVAL_CONFIG_PATH.exists():
        with open(EVAL_CONFIG_PATH, encoding="utf-8") as f:
            cfg = yaml.safe_load(f) or {}
        # Merge defaults for missing keys
        for section, defaults in _EVAL_CONFIG_DEFAULTS.items():
            if section not in cfg:
                cfg[section] = defaults
            elif isinstance(defaults, dict):
                for k, v in defaults.items():
                    cfg[section].setdefault(k, v)
        return cfg
    return dict(_EVAL_CONFIG_DEFAULTS)
