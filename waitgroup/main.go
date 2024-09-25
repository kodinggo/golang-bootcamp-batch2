package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("first execution")

	var wg sync.WaitGroup

	// init delta
	wg.Add(3)

	go func() {
		fmt.Println("resolve data 1")
		// decrease delta
		wg.Done()
	}()

	go func() {
		fmt.Println("resolve data 2")
		// decrease delta
		wg.Done()
	}()

	go func() {
		fmt.Println("resolve data 3")
		// decrease delta
		wg.Done()
	}()

	// wait delta until 0
	wg.Wait()

	fmt.Println("next execution")

	time.Sleep(1 * time.Second)

}
