package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go worker()
	wg.Wait()
	fmt.Println("main func over")
}

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}

	// 怎么停下来？
	wg.Done()
}
