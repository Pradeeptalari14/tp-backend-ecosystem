# Project 1: E-Commerce High-Volume Order & Inventory Engine

Production-grade asynchronous order processing architecture built with Python FastAPI, Redis, PostgreSQL, Celery, and Stripe integration.

## Key Architectural Patterns
1. **Idempotency Key Guard**: Redis key checks prevent duplicate charging when users double-click submit or network timeouts occur.
2. **Sub-2ms Inventory Caching**: Real-time stock counts checked in Redis before committing to heavier PostgreSQL disk transactions.
3. **Decoupled Asynchronous Tasks**: Heavy payment charging and notification dispatch offloaded to Celery message queue workers.
4. **Resilient Retries**: Celery tasks configured with exponential backoff and jitter.

## Endpoints
- `GET /health`: Component health check & readiness probe.
- `POST /api/v1/orders`: Accepts order payload, validates JWT auth header, checks stock, and queues fulfillment worker.
- `GET /api/v1/orders`: Lists placed orders.

## Running Locally
```bash
pip install -r requirements.txt
uvicorn main:app --reload --port 8000
```
