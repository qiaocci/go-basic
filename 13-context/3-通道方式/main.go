package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 使用管道的问题:
// 使用全局变量在跨包调用时,不容易统一和规范, 需要维护一个共同的channel
func main() {
	var exitChan = make(chan struct{})

	wg.Add(1)
	go worker(exitChan)

	time.Sleep(time.Second * 3)
	exitChan <- struct{}{} // 给子goroutine发送退出信号
	close(exitChan)
	wg.Wait()
	fmt.Println("main func over")
}

func worker(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:

		}
	}

	// 怎么停下来？
	wg.Done()
}
