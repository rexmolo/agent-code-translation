Done. ChromaDB is running:

  - Container: chromadb - Up 4 seconds
  - Port: 8000 (exposed to host)
  - Data: Persists in thesis/chroma/

 # API is running
  curl http://localhost:8000/api/v2/version
  # Returns: "1.0.0"

  You interact with ChromaDB via:
  - Python client: pip install chromadb
  - JavaScript/TypeScript client: npm install chromadb
  - REST API: http://localhost:8000/api/v2/