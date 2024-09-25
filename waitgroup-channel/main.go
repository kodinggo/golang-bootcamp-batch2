package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("first execution")

	var wg sync.WaitGroup
	msgCh := make(chan string, 3)

	// init delta
	wg.Add(3)

	go func() {
		msgCh <- "resolve data 1"
		// decrease delta
		wg.Done()
	}()

	go func() {
		msgCh <- "resolve data 2"
		// decrease delta
		wg.Done()
	}()

	go func() {
		msgCh <- "resolve data 3"
		// decrease delta
		wg.Done()
	}()

	// wait delta until 0
	wg.Wait()
	close(msgCh)

	// next execution
	for msg := range msgCh {
		fmt.Println(msg)
	}

	time.Sleep(1 * time.Second)

}
