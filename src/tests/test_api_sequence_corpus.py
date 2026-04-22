import json

from src.scripts.build_api_sequence_corpus import build_api_sequence_records, write_jsonl


def test_build_api_sequence_records_dedupes_by_sequence(monkeypatch, tmp_path):
    problem_a = tmp_path / "p00001" / "Go"
    problem_b = tmp_path / "p00002" / "Go"
    problem_a.mkdir(parents=True)
    problem_b.mkdir(parents=True)
    file_a = problem_a / "a.go"
    file_b = problem_b / "b.go"
    file_a.write_text("package main\n", encoding="utf-8")
    file_b.write_text("package main\n", encoding="utf-8")

    def fake_extract(path):
        if path == file_a:
            return [
                {
                    "function_name": "Normalize",
                    "sequence_text": "strings.TrimSpace -> strings.Split",
                    "apis": ["strings.TrimSpace", "strings.Split"],
                    "imports": ["strings"],
                    "function_code": "func Normalize() {}",
                }
            ]
        return [
            {
                "function_name": "NormalizeAgain",
                "sequence_text": "strings.TrimSpace -> strings.Split",
                "apis": ["strings.TrimSpace", "strings.Split"],
                "imports": ["strings"],
                "function_code": "func NormalizeAgain() {}",
            },
            {
                "function_name": "JoinParts",
                "sequence_text": "strings.Split -> strings.Join",
                "apis": ["strings.Split", "strings.Join"],
                "imports": ["strings"],
                "function_code": "func JoinParts() {}",
            },
        ]

    monkeypatch.setattr(
        "src.scripts.build_api_sequence_corpus.extract_go_api_sequences_from_file",
        fake_extract,
    )

    records = build_api_sequence_records(tmp_path)

    assert len(records) == 2
    assert records[0]["_id"] == "api_seq_go_000001"
    assert records[0]["sequence_text"] == "strings.TrimSpace -> strings.Split"
    assert records[1]["_id"] == "api_seq_go_000002"
    assert records[1]["sequence_text"] == "strings.Split -> strings.Join"


def test_write_jsonl_writes_records(tmp_path):
    output_path = tmp_path / "go_api_sequences.jsonl"
    records = [
        {
            "_id": "api_seq_go_000001",
            "language": "go",
            "source_corpus": "project_codenet_go",
            "file_path": "p00001/Go/a.go",
            "function_name": "Normalize",
            "sequence_text": "strings.TrimSpace -> strings.Split",
            "apis": ["strings.TrimSpace", "strings.Split"],
            "imports": ["strings"],
            "function_code": "func Normalize() {}",
        }
    ]

    write_jsonl(records, output_path)

    written = [json.loads(line) for line in output_path.read_text(encoding="utf-8").splitlines()]
    assert written == records
