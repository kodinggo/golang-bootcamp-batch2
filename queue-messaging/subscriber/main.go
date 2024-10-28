package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

const typeEmailDelivery = "EMAIL_DELIVERY"

type emailDeliveryPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"default": 2,
			},
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	// Handle subscriptions
	mux.HandleFunc(typeEmailDelivery, func(ctx context.Context, t *asynq.Task) error {
		var p emailDeliveryPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		}

		log.Println("Event Received", p.From, p.To)

		return nil
	})

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
