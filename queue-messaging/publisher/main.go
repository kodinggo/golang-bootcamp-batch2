package main

import (
	"encoding/json"
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
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	// Event & Payload
	payload, _ := json.Marshal(emailDeliveryPayload{
		From:    "john@yahoo.com",
		To:      "mark@gmail.com",
		Subject: "Job Apply",
		Message: "Lorem ipsum",
	})
	task := asynq.NewTask(typeEmailDelivery, payload)

	// Mengantrikan task / Mempublish message
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
