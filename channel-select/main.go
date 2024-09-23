package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "hello i'm goroutine 1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "hello i'm goroutine 2"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msgCh1 := <-ch1:
			fmt.Println(msgCh1)
		case msgCh2 := <-ch2:
			fmt.Println(msgCh2)
		}
	}
}
