package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// ctx, cancel := context.WithCancel(ctx)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	go doSomething(ctx)

	// time.AfterFunc(5*time.Second, func() {
	// 	cancel()
	// })

	// make sure program always run
	var s string
	fmt.Scan(&s)
}

func doSomething(ctx context.Context) {
	for {
		if isCtxCanceled(ctx) {
			fmt.Println("stop doing task!")
			break
		}
		fmt.Println("doing some task...")
		time.Sleep(1 * time.Second)
	}
}

func isCtxCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		fmt.Println("fungsi cancel() dipanggil")
		return true
	default:
		return false
	}
}
