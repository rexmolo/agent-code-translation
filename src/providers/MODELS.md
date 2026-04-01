# Model Registry Convention

## Overview

The translation system supports multiple **providers** (e.g., Google Gemini, MiniMax), each offering multiple **model variants** (e.g., gemini-2.5-flash-lite, gemini-2.5-pro).

## Registry Structure

All models are registered in `src/models/registry.py` inside `_register_defaults()`.

Each provider entry has:
- `env_var`: API key environment variable name
- `vertex_env_vars` (optional): Vertex AI env vars (Gemini only)
- `label`: Human-readable provider name for CLI display
- `variants`: Dict of model variants

Each variant entry has:
- `model_id`: Exact model ID string passed to the Agno model class
- `label`: Human-readable name for CLI display
- `factory`: Callable returning an Agno model instance

## Adding a New Variant to an Existing Provider

1. Open `src/models/registry.py`
2. Find the provider in `_register_defaults()`
3. Add a new entry under `"variants"`:

```python
"variants": {
    # ... existing variants ...
    "2.5_flash": {
        "model_id": "gemini-2.5-flash",
        "label": "Gemini 2.5 Flash",
        "factory": _gemini_factory("gemini-2.5-flash"),
    },
},
```

## Adding a New Provider

1. If the provider needs a custom model class, create it under `src/models/<provider_name>/`
2. In `_register_defaults()`, add a factory builder and a new provider entry:

```python
def _newprovider_factory(model_id: str) -> Callable:
    def factory():
        from agno.models.newprovider import NewProvider
        return NewProvider(id=model_id)
    return factory

_REGISTRY["newprovider"] = {
    "env_var": "NEWPROVIDER_API_KEY",
    "label": "New Provider",
    "variants": {
        "v1": {
            "model_id": "newprovider-v1",
            "label": "New Provider V1",
            "factory": _newprovider_factory("newprovider-v1"),
        },
    },
}
```

3. Add the API key to your `.env` file

## Variant Key Naming Convention

The variant key is used as a **directory name** in the output path. Rules:
- Use the model version/name portion only (drop the provider prefix)
- Keep dots for version numbers (e.g., `2.5`)
- Replace hyphens and spaces with underscores
- Lowercase for Gemini-style IDs; keep original case for MiniMax-style IDs

Examples:
| Model ID | Variant Key |
|---|---|
| `gemini-2.5-flash-lite` | `2.5_flash_lite` |
| `gemini-3-pro-preview` | `3_pro_preview` |
| `MiniMax-M2.5` | `M2.5` |

## Output Directory Convention

Translated files are saved to:
```
data/translation/target/{provider}/{variant_key}/{source_subfolder}/{file}.go
```

Example:
```
source:  data/translation/source/example/extract_codenet_data.py
target:  data/translation/target/gemini/2.5_flash_lite/example/extract_codenet_data.go
```

## Environment Variables

Each provider requires one API key env var. All variants under the same provider share the same key.

| Provider | Env Var | Alternative |
|---|---|---|
| gemini | `GOOGLE_API_KEY` | Vertex AI: `GOOGLE_GENAI_USE_VERTEXAI` + `GOOGLE_CLOUD_PROJECT` + `GOOGLE_CLOUD_LOCATION` |
| minimax | `MINIMAX_API_KEY` | — |
| openai | `OPENAI_API_KEY` | — |
