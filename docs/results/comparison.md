# HumanEval-X Evaluation Comparison

| Provider | Model | Strategy | Total | Compilation@1 | Pass@1 |
|----------|-------|----------|------:|--------------:|-------:|
| gemini | 2.5_pro | baseline | 164 | 28.0% | 25.0% |
| minimax | M2.5 | baseline | 164 | 32.9% | 30.5% |
| minimax | M2.5 | vec-chroma/rag-full | 164 | 47.6% | 42.7% |
| minimax | M2.5 | vec-chroma/rag-pattern-api-docs | 164 | 42.1% | 37.8% |
| minimax | M2.5 | vec-chroma/rag-pattern-only | 164 | 36.6% | 34.8% |
| minimax | M2.5 | vec-chroma/rag-pattern-samples | 164 | 26.2% | 25.0% |
| minimax | M2.5 | vec-gemini/rag-full | 164 | 41.5% | 37.8% |
| minimax | M2.5 | vec-gemini/rag-pattern-api-docs | 164 | 40.9% | 39.0% |
| minimax | M2.5 | vec-gemini/rag-pattern-only | 164 | 42.7% | 39.0% |
| minimax | M2.5 | vec-gemini/rag-pattern-samples | 164 | 28.7% | 25.0% |
