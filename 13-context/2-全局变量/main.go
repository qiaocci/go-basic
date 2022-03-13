package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

// 使用全局变量的问题
// 1. 在跨包调用时不容易统一
// 2. 如果在worker中再启动goroutine, 就不太好控制了

func main() {
	wg.Add(1)
	go worker()
	time.Sleep(time.Second * 3)
	exit = true // 修改全局变量， 实现子goroutine退出

	wg.Wait()
	fmt.Println("main func over")
}

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}

	// 怎么停下来？
	wg.Done()
}
