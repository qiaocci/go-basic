package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{} // 并发安全

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=%v, v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
