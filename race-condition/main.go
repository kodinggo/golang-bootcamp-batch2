package main

import (
	"fmt"
	"sync"
)

func main() {
	totalProcess := 1000

	var wg sync.WaitGroup
	wg.Add(totalProcess)

	var mu sync.Mutex

	counter := 0

	for i := 1; i <= totalProcess; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
