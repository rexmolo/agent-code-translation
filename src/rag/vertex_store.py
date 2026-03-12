"""Vertex AI Vector Search resource management.

Handles index creation, endpoint deployment, data upsert, and querying
for the Gemini embedding backend.
"""

from __future__ import annotations

import os
import time

from google.cloud import aiplatform
from google.cloud.aiplatform.matching_engine import (
    MatchingEngineIndex,
    MatchingEngineIndexEndpoint,
)
from google.cloud.aiplatform.matching_engine.matching_engine_index_endpoint import (
    Namespace,
)
from rich.console import Console

from src.rag.embeddings import load_rag_config

_console = Console()


def _init_aiplatform() -> None:
    """Ensure aiplatform SDK is initialised with project/location from env."""
    from src.config import load_providers_config
    vertex_cfg = load_providers_config()["google"]["vertex_ai"]

    project = os.environ.get(vertex_cfg["project_env"], "")
    location = os.environ.get(
        vertex_cfg["location_env"], vertex_cfg["default_location"],
    )
    if not project:
        raise ValueError(
            f"{vertex_cfg['project_env']} is not set. "
            "Set it in .env to use Vertex AI Vector Search."
        )
    aiplatform.init(project=project, location=location)


def _get_vertex_config() -> dict:
    cfg = load_rag_config()
    return cfg["vertex_ai"]


# ---------------------------------------------------------------------------
# Index management
# ---------------------------------------------------------------------------

def get_or_create_index(
    dimensions: int | None = None,
    display_name: str | None = None,
) -> MatchingEngineIndex:
    """Return existing index by display name, or create a new STREAM_UPDATE index."""
    _init_aiplatform()
    vcfg = _get_vertex_config()
    display_name = display_name or vcfg["index_display_name"]
    dimensions = dimensions or vcfg["dimensions"]

    # Search for existing index
    indexes = MatchingEngineIndex.list()
    for idx in indexes:
        if idx.display_name == display_name:
            _console.print(f"  Found existing index: [green]{idx.resource_name}[/green]")
            return idx

    _console.print(f"  Creating new index [bold]{display_name}[/bold] (dimensions={dimensions})...")
    index = MatchingEngineIndex.create_tree_ah_index(
        display_name=display_name,
        dimensions=dimensions,
        approximate_neighbors_count=50,
        leaf_node_embedding_count=500,
        leaf_nodes_to_search_percent=7,
        distance_measure_type="COSINE_DISTANCE",
        index_update_method="STREAM_UPDATE",
        shard_size="SHARD_SIZE_SMALL",
    )
    _console.print(f"  Index created: [green]{index.resource_name}[/green]")
    return index


# ---------------------------------------------------------------------------
# Endpoint management
# ---------------------------------------------------------------------------

def get_or_create_endpoint(
    display_name: str | None = None,
) -> MatchingEngineIndexEndpoint:
    """Return existing endpoint by display name, or create a new public endpoint."""
    _init_aiplatform()
    vcfg = _get_vertex_config()
    display_name = display_name or vcfg["endpoint_display_name"]

    endpoints = MatchingEngineIndexEndpoint.list()
    for ep in endpoints:
        if ep.display_name == display_name:
            _console.print(f"  Found existing endpoint: [green]{ep.resource_name}[/green]")
            return MatchingEngineIndexEndpoint(index_endpoint_name=ep.resource_name)

    _console.print(f"  Creating new endpoint [bold]{display_name}[/bold]...")
    endpoint = MatchingEngineIndexEndpoint.create(
        display_name=display_name,
        public_endpoint_enabled=True,
    )
    _console.print(f"  Endpoint created: [green]{endpoint.resource_name}[/green]")
    return endpoint


# ---------------------------------------------------------------------------
# Deployment
# ---------------------------------------------------------------------------

def ensure_deployed(
    index: MatchingEngineIndex,
    endpoint: MatchingEngineIndexEndpoint,
    deployed_index_id: str = "thesis_rag_deployed",
) -> str:
    """Deploy the index to the endpoint if not already deployed.

    Returns the deployed_index_id.
    """
    # Check if already deployed
    for di in endpoint.deployed_indexes:
        if di.id == deployed_index_id:
            _console.print(f"  Index already deployed as [green]{deployed_index_id}[/green]")
            return deployed_index_id

    _console.print(f"  Deploying index to endpoint (this may take ~20 min)...")
    endpoint.deploy_index(
        index=index,
        deployed_index_id=deployed_index_id,
        display_name=deployed_index_id,
        machine_type="e2-standard-2",
        min_replica_count=1,
        max_replica_count=1,
    )
    _console.print(f"  [green]Deployed![/green] deployed_index_id={deployed_index_id}")
    return deployed_index_id


# ---------------------------------------------------------------------------
# Data operations
# ---------------------------------------------------------------------------

def upsert_datapoints(
    index: MatchingEngineIndex,
    datapoint_ids: list[str],
    embeddings: list[list[float]],
    collection_name: str,
    batch_size: int = 500,
) -> None:
    """Batch-upsert datapoints with a collection restrict filter."""
    from google.cloud.aiplatform_v1.types import (
        IndexDatapoint,
        UpsertDatapointsRequest,
    )
    from google.cloud.aiplatform_v1.types.index import IndexDatapoint as IDPType

    import time
    from google.api_core.exceptions import ResourceExhausted

    # Vertex AI stream updates have strict throughput quotas. Cap batch size.
    batch_size = min(batch_size, 100)
    total = len(datapoint_ids)
    
    for start in range(0, total, batch_size):
        end = min(start + batch_size, total)
        datapoints = []
        for i in range(start, end):
            dp = IndexDatapoint(
                datapoint_id=datapoint_ids[i],
                feature_vector=embeddings[i],
                restricts=[
                    IndexDatapoint.Restriction(
                        namespace="collection",
                        allow_list=[collection_name],
                    )
                ],
            )
            datapoints.append(dp)
            
        retries = 0
        while True:
            try:
                index.upsert_datapoints(datapoints=datapoints)
                _console.print(f"    Upserted {end}/{total}")
                time.sleep(1.0)  # Gentle pause to respect rate limits
                break
            except ResourceExhausted as e:
                retries += 1
                if retries > 10:
                    raise e
                wait_time = 10 * retries
                _console.print(f"    [yellow]Quota exceeded (429). Waiting {wait_time}s before retrying...[/yellow]")
                time.sleep(wait_time)


# ---------------------------------------------------------------------------
# Query
# ---------------------------------------------------------------------------

def query_neighbors(
    endpoint: MatchingEngineIndexEndpoint,
    deployed_index_id: str,
    embedding: list[float],
    n_results: int = 5,
    collection_name: str | None = None,
) -> list[str]:
    """Query nearest neighbors, returning a list of datapoint IDs.

    If *collection_name* is given, filters results to that collection.
    """
    num_neighbors = n_results * 3  # over-fetch for RRF

    filter_restricts = None
    if collection_name:
        filter_restricts = [
            Namespace(name="collection", allow_tokens=[collection_name]),
        ]

    response = endpoint.find_neighbors(
        deployed_index_id=deployed_index_id,
        queries=[embedding],
        num_neighbors=num_neighbors,
        filter=filter_restricts,
    )

    if not response or not response[0]:
        return []
    return [neighbor.id for neighbor in response[0]]
