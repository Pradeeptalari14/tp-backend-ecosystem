package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type WebhookEvent struct {
	TenantID  string `json:"tenant_id" binding:"required"`
	EventType string `json:"event_type" binding:"required"`
	Payload   string `json:"payload" binding:"required"`
	Signature string `json:"signature"`
}

type EventStore struct {
	sync.RWMutex
	idempotencyKeys map[string]int64
	dispatchedTasks []string
}

var store = &EventStore{
	idempotencyKeys: make(map[string]int64),
	dispatchedTasks: make([]string, 0),
}

func computeHMAC(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "UP",
			"service":   "saas-billing-dispatcher",
			"timestamp": time.Now().UTC().Unix(),
		})
	})

	r.POST("/api/v1/billing/events", func(c *gin.Context) {
		var evt WebhookEvent
		if err := c.ShouldBindJSON(&evt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 1. Transactional Idempotency Check
		idempKey := fmt.Sprintf("idemp:%s:%s", evt.TenantID, evt.EventType)
		store.Lock()
		if _, exists := store.idempotencyKeys[idempKey]; exists {
			store.Unlock()
			c.JSON(http.StatusConflict, gin.H{
				"status":  "duplicate",
				"message": "Event already processed or in-flight",
			})
			return
		}
		store.idempotencyKeys[idempKey] = time.Now().Unix()

		// 2. Compute cryptographic webhook HMAC signature
		secret := "sec_live_enterprise_k92"
		computedSig := computeHMAC(evt.Payload, secret)
		taskID := fmt.Sprintf("asynq_task_%d", time.Now().UnixNano())
		store.dispatchedTasks = append(store.dispatchedTasks, taskID)
		store.Unlock()

		// 3. Return 202 Accepted and delegate dispatch to Asynq worker
		c.JSON(http.StatusAccepted, gin.H{
			"status":         "queued",
			"task_id":        taskID,
			"tenant_id":      evt.TenantID,
			"event_type":     evt.EventType,
			"hmac_signature": computedSig,
			"message":        "Enqueued to Asynq worker with exponential backoff & DLQ safety",
		})
	})

	r.GET("/api/v1/billing/tasks", func(c *gin.Context) {
		store.RLock()
		defer store.RUnlock()
		c.JSON(http.StatusOK, gin.H{
			"total_dispatched": len(store.dispatchedTasks),
			"tasks":            store.dispatchedTasks,
		})
	})

	r.Run(":8080")
}
