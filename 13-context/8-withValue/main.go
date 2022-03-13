package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type TraceCode string

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "123")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
	fmt.Println("over")
}

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code!!")
	}
LOOP:
	for {
		fmt.Printf("worker trace code %v\n", traceCode)
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
