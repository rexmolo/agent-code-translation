# Vertex AI Vector Search Integration Bugs and Fixes

During the development and ingestion of RAG embeddings using Google Vertex AI Vector Search (Matching Engine), several bugs and limits were encountered. 

This document details the issues and their respective solutions for future reference.

## 1. Index Creation Metadata (`algorithmConfig`) Error
**Issue:**
When calling `MatchingEngineIndex.create_tree_ah_index()`, the API returned a `FailedPrecondition 400` error:
`algorithmConfig is required but missing from the metadata.`

**Cause:**
Newer versions of the `google-cloud-aiplatform` SDK (e.g., `v1.141.0`) require explicit configuration for the approximate nearest neighbor (ANN) algorithm, specifically the number of leaf nodes and search percentage. Simply passing `approximate_neighbors_count` is no longer sufficient.

**Fix:**
Added `leaf_node_embedding_count` and `leaf_nodes_to_search_percent` to the `create_tree_ah_index` parameter list in `src/rag/vertex_store.py`:
```python
index = MatchingEngineIndex.create_tree_ah_index(
    display_name=display_name,
    dimensions=dimensions,
    approximate_neighbors_count=50,
    leaf_node_embedding_count=500,        # Explicitly required now
    leaf_nodes_to_search_percent=7,       # Explicitly required now
    distance_measure_type="COSINE_DISTANCE",
    index_update_method="STREAM_UPDATE",
    shard_size="SHARD_SIZE_SMALL",
)
```

## 2. Quota Exceeded during Data Ingestion (`ResourceExhausted 429`)
**Issue:**
The ingestion loop (`ingest_rag_gemini.py`) crashed with the following error when upserting large amounts of RAG data:
`429 Quota exceeded for quota metric 'Matching Engine stream update throughput' and limit 'Matching Engine stream update throughput per minute per region'.`

**Cause:**
Vertex AI's `STREAM_UPDATE` indexes have strict rate limits on payload size and throughput per minute. Uploading chunks of 500 documents rapidly blasted past the provisioned capacity.

**Fix:**
1. Reduced the maximum `batch_size` constraint from `500` to `100`.
2. Wrapped the `index.upsert_datapoints()` call in a standard retry loop with exponential backoff inside `src/rag/vertex_store.py`.

```python
import time
from google.api_core.exceptions import ResourceExhausted

batch_size = min(batch_size, 100)  # Ceiling limit

for start in range(0, total, batch_size):
    # ... generate datapoints list ...
    retries = 0
    while True:
        try:
            index.upsert_datapoints(datapoints=datapoints)
            time.sleep(1.0)  # Gentle pause
            break
        except ResourceExhausted as e:
            retries += 1
            if retries > 10:
                raise e
            wait_time = 10 * retries
            time.sleep(wait_time)
```

## 3. Empty RAG Context due to Query Client Initialization
**Issue:**
When evaluating translations and retrieving the context with `build_translation_context`, Vertex AI queries errored out with:
`AttributeError: 'MatchingEngineIndexEndpoint' object has no attribute '_public_match_client'`

**Cause:**
In `src/rag/vertex_store.py`, `endpoints = MatchingEngineIndexEndpoint.list()` is used to check if the endpoint already exists. However, objects returned by `list()` are lightweight wrappers and do not fully initialize the underlying gRPC `_public_match_client` used for `find_neighbors` queries. 

**Fix:**
Instead of returning the retrieved object directly, we re-instantiate it using the correct keyword argument `index_endpoint_name`. (Note: in early SDKs this was `endpoint_name`, but it has been renamed in v1.141+).

```python
endpoints = MatchingEngineIndexEndpoint.list()
for ep in endpoints:
    if ep.display_name == display_name:
        # DO NOT `return ep` directly.
        # MUST re-instantiate using `index_endpoint_name` to bind the match client.
        return MatchingEngineIndexEndpoint(index_endpoint_name=ep.resource_name)
```
