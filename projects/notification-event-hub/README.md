# Project 2: Real-Time Notification & Event Streaming Hub

Sub-10ms broadcast and push notification engine built with Node.js, Express, WebSockets, Redis Pub/Sub, BullMQ, and PostgreSQL.

## Architectural Highlights
1. **Persistent WebSocket Pooling**: Maintains bi-directional open TCP channels with clients for immediate in-app delivery.
2. **Cluster Multi-Broadcasting**: Redis Pub/Sub ensures messages dispatched to any backend node broadcast across all other cluster nodes.
3. **Durable BullMQ Queues**: Heavy push notifications (APNS, FCM) and database history writes are processed asynchronously with automated retry backoffs.
4. **Dead Letter Queue (DLQ)**: Failed deliveries after 3 attempts are captured in a DLQ for operational review.

## Endpoints
- `GET /health`: Cluster health & connected client counts.
- `POST /api/v1/notify`: Accepts event notifications, sends immediately to active WebSockets, and enqueues worker task.
- `WS /?userId=USER_ID`: WebSocket endpoint for incoming persistent user connections.

## Quickstart
```bash
npm install
npm start
```
