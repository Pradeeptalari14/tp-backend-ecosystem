# Project 3: Enterprise Multi-Tenant Billing & Webhook Dispatcher

High-concurrency webhook and billing event dispatching engine written in Go (Gin), Redis, Asynq, and MySQL.

## Key Architectural Patterns
1. **Transactional Outbox Pattern**: Ensures event records are safely committed with billing transactions prior to external HTTP dispatch.
2. **HMAC-SHA256 Payload Signing**: Every outbound webhook request is signed with a tenant cryptographic secret in `X-Signature-SHA256`.
3. **Exponential Backoff & Jitter**: Asynq retries transient 5xx recipient endpoint failures (attempts: 1 to 5) before sending to a Dead Letter Queue.
4. **Tenant Isolation**: Multi-tenant key namespacing prevents noisy-neighbor throttling in Redis.

## Endpoints
- `GET /health`: Engine status & uptime.
- `POST /api/v1/billing/events`: Ingests billing event, validates idempotency, computes HMAC signature, and returns HTTP 202 Accepted.
- `GET /api/v1/billing/tasks`: Query list of dispatched Asynq task IDs.

## Quickstart
```bash
go run main.go
```
