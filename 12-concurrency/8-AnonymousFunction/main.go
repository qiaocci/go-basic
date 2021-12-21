package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func hello(i int) {
	fmt.Println("hello ", i)
	wg1.Done()
}
func main() { // 开启主goroutine去执行main函数

	wg1.Add(10000) // 计数牌
	for i := 0; i < 10000; i++ {
		//go func() {
		//	fmt.Println("hello ", i)  // 问题： 打印出来数字都是10000？闭包，包含外部函数的引用
		//	wg1.Done()
		//}()
		go func(i int) {
			fmt.Println("hello ", i) // 解决：使用函数传参
			wg1.Done()
		}(i)
	}
	fmt.Println("hello main")

	wg1.Wait() // 阻塞，等所有小弟干完活才收兵
}
