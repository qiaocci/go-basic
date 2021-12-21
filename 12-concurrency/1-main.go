package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello ", i)
}
func main() { // 开启主goroutine去执行main函数
	go hello(1)
	fmt.Println("main goroutine done")
	time.Sleep(time.Second)
}
