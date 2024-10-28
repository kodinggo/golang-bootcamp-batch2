package main

import (
	"fmt"
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

	// subscribe event
	js.QueueSubscribe("PRODUCTS.create", "worker-group", func(msg *nats.Msg) {
		log.Printf("subcribe data %s", string(msg.Data))
	})

	var s string
	fmt.Scan(&s)
}
