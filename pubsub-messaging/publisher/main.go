package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// konek ke nats server
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Panicf("connect to nats server: %v", err)
	}

	// daftarkan stream (topic) dan subject (event)
	js, err := nc.JetStream()
	if err != nil {
		log.Panicf("init jetstream: %v", err)
	}
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "PRODUCTS",
		Subjects: []string{"PRODUCTS.*"},
	})
	if err != nil {
		log.Panicf("create stream: %v", err)
	}

	// publishe event
	m := map[string]interface{}{
		"product_name": "Apple MacBoook Pro",
		"price":        123,
	}
	payload, _ := json.Marshal(m)
	js.Publish("PRODUCTS.create", payload)
}
