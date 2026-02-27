"""
Shared path configuration for all scripts and experiments.

Usage:
    from src.config import REPO_ROOT, TEMP_DIR, DATA_DIR
"""

from pathlib import Path

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
RAG_PROCESSED_DIR = REPO_ROOT / "src" / "data" / "RAG" / "processed"
