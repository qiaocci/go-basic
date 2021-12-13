package main

import (
	"fmt"
)

// 我们传递一个int指针类型的参数，传递的其实是这个参数的一个副本
func main() {
	i := 10
	fmt.Printf("原数据的指针是=%p\n", &i)
	ip := &i
	fmt.Printf("原始指针的地址是=%p\n", &ip)
	modifyPtr(ip)
	fmt.Printf("int的值被修改了, 新值=%v\n", i)
}

func modifyPtr(ip *int) {
	fmt.Printf("函数里接收到的指针的地址是=%p\n", &ip)
	*ip = 1
}
