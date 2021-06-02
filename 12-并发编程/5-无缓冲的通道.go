package main

import (
	"fmt"
	"time"
)

func recv(c chan int) {
	val := <-c
	fmt.Println("接收成功", val)
}

func main() {
	c := make(chan int)
	go recv(c)
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("发送成功")

}
