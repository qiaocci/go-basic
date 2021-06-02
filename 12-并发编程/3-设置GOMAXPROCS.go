package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg2 sync.WaitGroup

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	wg2.Done()
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
	wg2.Done()
}

func main() {
	runtime.GOMAXPROCS(8)
	wg2.Add(2)
	go a()
	go b()
	wg2.Wait()
}
