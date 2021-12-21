package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello1(i int) {
	defer wg.Done() // goroutine结束登记 -1
	fmt.Println("hello gorountine", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello1(i)
	}

	wg.Wait()
}
