package main

import (
	"fmt"
	"runtime"
	"time"
)

// nil channel
func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan int
	go func() {
		<-ch
	}()

	time.Sleep(time.Second)
}
