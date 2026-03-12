# Note: RAG Architecture and Scaling on Cloud/VPS

## The Problem: Memory Limits
In the current thesis experiment setup, the RAG retrieved documents (Parallel Corpus, Go documentation, and API mappings) are stored in `.jsonl` files and loaded directly into a Python dictionary (`self._id_to_doc`) in local RAM (e.g., `src/rag/retriever.py`). 

When deployed to a small Cloud VPS with limited memory (e.g., 1-2 GB RAM), loading massive datasets into a Python dictionary is not feasible and will result in Out-of-Memory (OOM) crashes.

## Current Research Architecture vs Production 
### 1. Why is it loading into memory right now?
The current dataset is relatively small (totaling ~2000 text snippets). When loaded into a Python dictionary, all of this text takes up less than 10 Megabytes of RAM. 

Because this is a research code-translation pipeline that runs via CLI on a local machine (or a CI runner) rather than a persistent web server, reading the `.jsonl` files directly into memory is the fastest and simplest way to operate. It avoids the overhead of managing a separate relational database just for 10MB of text.

### 2. How do the Vectors Match the Text?
The Vertex AI Vector Search database acts as a highly optimized, mathematical "search engine". It does not store the text, it only stores two things:
- The 3072-dimensional vector.
- A string ID (e.g., `corpus_p00001`).

The IDs are generated deterministically during the ingestion phase (`ingest_rag_gemini.py`). Since the identical IDs are present in the local `.jsonl` dataset (e.g., `{"_id": "corpus_p00001"}`), the local dictionary acts as the lookup table to map the mathematical match back to the real text.

### 3. How to Scale for Production Deployments
To deploy this as a real web app (like a coding assistant API) with millions of documents, the database layer must be split into two partnered systems:

1. **The Vector Database (Vertex AI / Pinecone / Milvus):** Stores the embeddings and the string IDs for lightning-fast mathematical similarity searches.
2. **The Metadata Database (PostgreSQL / MongoDB / Redis):** Stores the actual heavy text documents on disk.

**The Production Flow (No Local RAM Dictionary):**
1. The user submits a query.
2. The query is embedded and sent to Vertex AI.
3. Vertex AI returns a list of matched IDs (e.g., `["doc_4091", "doc_8842"]`).
4. The backend server queries PostgreSQL: `SELECT text FROM documents WHERE id IN ('doc_4091', 'doc_8842')`.
5. PostgreSQL retrieves the exact text strings rapidly from disk.
6. The text is injected into the prompt.

By treating the local `.jsonl` files as a mock relational database, the pipeline remains simple for experimentation but models the correct architectural flow for larger systems.
