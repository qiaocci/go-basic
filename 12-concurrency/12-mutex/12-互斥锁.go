package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.Mutex
var x int

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + i
		lock.Unlock() // 解锁
		time.Sleep(time.Millisecond * 1)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x) // 这次算出来是正确的： 24995000
}
