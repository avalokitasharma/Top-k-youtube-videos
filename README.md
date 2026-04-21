## Top-K YouTube Videos System

This project implements a near real-time Top-K ranking system similar to YouTube Trending. It ingests user engagement events (views, likes, shares, watch time), processes them asynchronously, and serves low-latency ranked results.

The focus is on building a scalable streaming system using Go, Kafka, Redis, and Postgres, without relying on heavyweight stream processing frameworks.

---

## Architecture

The system consists of three services: ingestion, processor, and query.

- Ingestion service accepts events and produces them to Kafka  
- Processor consumes events, updates counters, and maintains rankings in Redis  
- Query service reads Top-K results from Redis and enriches them with metadata from Postgres  

### Key Design Decisions
1. Why Not Flink / Spark?
JVM overhead + operational complexity
Overkill for many real-time ranking systems
Instead:
Custom Go stream processor
Built on Kafka + Redis
Easier to deploy, debug, and scale

2. Kafka as the Backbone
Topic: video-events
Partitioning: video_id (ensures ordering per video)
Enables:
Horizontal scaling of processors
Backpressure handling
Replay capability

3. Redis for Real-Time Top-K

We use Redis:

HASH → per-video counters
ZSET → ranking (Top-K)

Example:

ZADD topk:global:24h <score> <video_id>

Why Redis?

O(log N) ranking updates
Sub-millisecond reads
Perfect for hot data

4. Postgres for Metadata

Stores:

Video metadata (title, category, thumbnail, etc.)

Why separate?

Redis = fast but ephemeral
Postgres = durable + queryable

5. Scoring Function (Extensible)
score = 0.6*views + 0.2*likes + 0.1*shares + 0.1*watch_time


## Project Structure
```bash
topk-youtube-system/
├── cmd/                # Service entrypoints
│   ├── ingestion/
│   ├── processor/
│   ├── query/
│
├── internal/
│   ├── kafka/          # Kafka producer/consumer
│   ├── redis/          # Redis client
│   ├── postgres/       # DB client
│   ├── handlers/       # API handlers
│   ├── models/         # Domain models
│   ├── score/          # Ranking logic
│   ├── config/         # Config loader
│
├── docker-compose.yml
├── Dockerfiles
├── init.sql
```

---

## Data Flow

- Client sends an event to the ingestion API  
- Event is published to Kafka (`video-events` topic)  
- Processor updates counters and recomputes score  
- Redis sorted sets are updated with latest rankings  
- Query API fetches Top-K and joins metadata  

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