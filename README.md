## Top-K YouTube Videos System

This project implements a near real-time Top-K ranking system similar to YouTube Trending. It ingests user engagement events (views, likes, shares, watch time), processes them asynchronously, and serves low-latency ranked results.

The focus is on building a scalable streaming system using Go, Kafka, Redis, and Postgres, without relying on heavyweight stream processing frameworks.

---

## Architecture

The system consists of three services: ingestion, processor, and query.

- Ingestion service accepts events and produces them to Kafka  
- Processor consumes events, updates counters, and maintains rankings in Redis  
- Query service reads Top-K results from Redis and enriches them with metadata from Postgres  

Kafka acts as the event backbone, Redis handles real-time computation, and Postgres stores durable metadata.

---

## Data Flow

- Client sends an event to the ingestion API  
- Event is published to Kafka (`video-events` topic)  
- Processor updates counters and recomputes score  
- Redis sorted sets are updated with latest rankings  
- Query API fetches Top-K and joins metadata  

---

## Storage

Redis is used for fast updates and reads:

- `counters:<video_id>` → engagement counts  
- `topk:global:<window>` → rankings  

Postgres stores video metadata such as title and category.

---

### Running Locally

```bash
docker compose up --build
```
Ingest event:

```bash
curl -X POST http://localhost/v1/ingest \
  -H "Content-Type: application/json" \
  -d '{"video_id":"00000000-0000-0000-0000-000000000001","event_type":"view","value":1}'
```

Query Top-K:
```bash
curl "http://localhost/v1/topk?k=10&window=24h"
```

### Scaling
- Stateless ingestion and query services scale via replication
- Kafka partitions enable parallel processing
- Processor instances scale with partitions
- Redis can be clustered for higher throughput

### Tradeoffs
- Eventual consistency instead of strict guarantees
- No built-in exactly-once processing
- Redis memory usage grows with data

These are acceptable for ranking systems where slight delays are tolerable.

### Future Work
- Time-decay scoring
- Sliding window aggregation
- Per-region or per-category rankings
- Streaming updates (WebSockets/SSE)

#### Notes

This project emphasizes simple, production-friendly design using core primitives. It is easy to run locally and scales cleanly to a distributed setup.