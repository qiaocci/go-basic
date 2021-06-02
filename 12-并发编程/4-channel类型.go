package main

import "fmt"

func main() {
	var c1 chan int
	var c2 chan string
	var c3 chan []int
	fmt.Println(c1, c2, c3) // <nil> 指针类型， 需要初始化

	c4 := make(chan int)
	c5 := make(chan bool)
	c6 := make(chan []int)
	fmt.Println(c4, c5, c6) // 0xc000094060 0xc0000940c0 0xc000094120
}
