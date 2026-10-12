FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o worker_bin worker/worker.go

FROM alpine:3.19
RUN apk --no-cache add ca-certificates curl
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/worker_bin .

EXPOSE 8080

HEALTHCHECK --interval=15s --timeout=3s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

CMD ["./server"]
