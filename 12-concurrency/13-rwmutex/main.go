package main

import (
	"fmt"
	"sync"
	"time"
)

// 适用于 读多写少
// 当goroutine获取读锁,其他goroutine可以继续获取读锁, 获取写锁就会等待
// 当goroutine获取写锁,其他goroutine获取读锁和写锁都会等待

var x int
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func write() {
	//lock.Lock()
	//rwlock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10) // 假设写操作耗时10毫秒
	//rwlock.Unlock()
	//lock.Unlock()
	wg.Done()
}
func read() {
	//lock.Lock()
	//rwlock.Lock()
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	//rwlock.Unlock()
	//lock.Unlock()
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}

	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}

	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))

	// 不加锁 10ms
	// 加互斥锁 1.2s
	// 加读写互斥所 1.2s
}
