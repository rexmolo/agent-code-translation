"""
Shared path configuration for all scripts and experiments.

Usage:
    from src.config import REPO_ROOT, TEMP_DIR, DATA_DIR
"""

import os
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
GRAMMAR_MAPPINGS_FILE = RAG_PROCESSED_DIR / "grammar_mappings.jsonl"
API_MAPPINGS_FILE = RAG_PROCESSED_DIR / "api_mappings.jsonl"
GO_DOCS_FILE = RAG_PROCESSED_DIR / "go_docs.jsonl"

# --- Config files ---
EVAL_CONFIG_PATH = REPO_ROOT / "config" / "eval_config.yaml"
PROVIDERS_CONFIG_PATH = REPO_ROOT / "config" / "providers.yaml"

_EVAL_CONFIG_DEFAULTS = {
    "translation": {"batch_size": 5},
    "parallel": {"batch_size": 10},
    "docker": {"image": "golang:1.26-alpine", "memory_limit": "512m", "timeout": 60},
}


_providers_cfg: dict | None = None


def load_providers_config() -> dict:
    """Load provider/credential config from YAML (cached)."""
    global _providers_cfg
    if _providers_cfg is None:
        with open(PROVIDERS_CONFIG_PATH, encoding="utf-8") as f:
            _providers_cfg = yaml.safe_load(f)
    return _providers_cfg


def resolve_api_key(provider_section: dict) -> str | None:
    """Resolve an API key from a provider config section.

    Looks for ``api_key`` first (direct value), then ``api_key_env``
    (environment variable name).  Returns None if neither yields a value.
    """
    key = provider_section.get("api_key")
    if key:
        return key
    env_var = provider_section.get("api_key_env")
    if env_var:
        return os.environ.get(env_var) or None
    return None


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
