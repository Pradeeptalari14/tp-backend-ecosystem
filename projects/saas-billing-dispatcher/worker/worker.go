package main

import (
	"fmt"
	"time"
)

// DispatchWebhook simulates delivery of an signed HMAC event to client endpoints
func DispatchWebhook(tenantID string, eventType string, payload string, attempt int) error {
	fmt.Printf("[Asynq Worker] Dispatching event '%s' for tenant '%s' (Attempt: %d)\n", eventType, tenantID, attempt)
	// Simulate outbound HTTPS post with timeout
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("[Asynq Worker] Delivered successfully. Response: HTTP 200 OK\n")
	return nil
}

func main() {
	fmt.Println("Asynq Billing Webhook Worker initialized in simulation mode...")
	_ = DispatchWebhook("tenant_acme_corp", "invoice.payment_succeeded", `{"amount": 12900}`, 1)
}
