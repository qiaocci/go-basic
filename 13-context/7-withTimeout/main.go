package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)

	go worker(ctx)
	cancel() // 通知子goroutine结束
	wg.Wait()

}

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}

	wg.Done()
	fmt.Println("worker done!")

}
