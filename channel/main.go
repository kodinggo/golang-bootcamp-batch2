package main

import "fmt"

func main() {
	msgCH := make(chan string)

	go func() {
		sender(msgCH, "Hello from GR 1")
	}()

	go func() {
		sender(msgCH, "Hello from GR 2")
	}()

	go func() {
		sender(msgCH, "Hello from GR 3")
	}()

	fmt.Println(receiver(msgCH))
	fmt.Println(receiver(msgCH))
	fmt.Println(receiver(msgCH))
}

func sender(ch chan<- string, msg string) {
	ch <- msg
}

func receiver(ch <-chan string) string {
	return <-ch
}
