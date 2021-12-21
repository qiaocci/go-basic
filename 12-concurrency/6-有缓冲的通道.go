package main

import "fmt"

func main() {
	c := make(chan int, 1) // 有缓冲的通道， 容量是1
	c <- 10
	fmt.Println("发送成功")
}
