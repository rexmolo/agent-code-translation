"""SQLite persistence for evaluation errors.

Stores full error context (stderr, provider, model, etc.) for later analysis.
Uses lazy connection singleton — table is auto-created on first call.
"""

import sqlite3
import threading
from datetime import datetime, timezone

from src.config import ERROR_DB_PATH

_CREATE_TABLE = """\
CREATE TABLE IF NOT EXISTS evaluation_errors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_name TEXT NOT NULL,
    err_type TEXT NOT NULL,
    err_log_context TEXT,
    provider TEXT NOT NULL,
    model TEXT NOT NULL,
    variant TEXT NOT NULL,
    time TEXT NOT NULL,
    experiment_type TEXT NOT NULL
)
"""

_lock = threading.Lock()
_conn: sqlite3.Connection | None = None


def _get_conn() -> sqlite3.Connection:
    """Return (and cache) the SQLite connection, creating the table if needed."""
    global _conn
    if _conn is None:
        with _lock:
            if _conn is None:
                ERROR_DB_PATH.parent.mkdir(parents=True, exist_ok=True)
                _conn = sqlite3.connect(str(ERROR_DB_PATH), check_same_thread=False)
                _conn.execute(_CREATE_TABLE)
                _conn.commit()
    return _conn


def save_error(
    file_name: str,
    err_type: str,
    err_log_context: str | None,
    provider: str,
    model: str,
    variant: str,
    experiment_type: str,
) -> None:
    """Insert one evaluation error record."""
    conn = _get_conn()
    with _lock:
        conn.execute(
            """\
            INSERT INTO evaluation_errors
                (file_name, err_type, err_log_context, provider, model, variant, time, experiment_type)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?)
            """,
            (
                file_name,
                err_type,
                err_log_context,
                provider,
                model,
                variant,
                datetime.now(timezone.utc).isoformat(),
                experiment_type,
            ),
        )
        conn.commit()
