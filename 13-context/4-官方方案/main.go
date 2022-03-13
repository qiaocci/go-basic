package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx)

	time.Sleep(time.Second * 3)
	cancel()

	wg.Wait()
	fmt.Println("main func over")
}

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待商机通知
			break LOOP
		default:

		}
	}

	// 怎么停下来？
	wg.Done()
}
