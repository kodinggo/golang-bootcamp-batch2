package main

import (
	"time"
)

func main() {
	msgCh := make(chan string, 3)

	// go func() {
	// 	fmt.Println(<-msgCh)
	// }()

	msgCh <- "hello 1"
	msgCh <- "hello 2"
	msgCh <- "hello 3"

	msgCh <- "hello 4"

	time.Sleep(1 * time.Second)
}
