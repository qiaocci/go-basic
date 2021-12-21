package main

import "fmt"

// 我们传递一个指针类型的参数，传递的其实是这个参数的一个副本
func main() {
	i := 10
	ip := &i
	fmt.Printf("原始变量: %d, 指针: %p\n", i, &i)
	fmt.Printf("原始指针的值是: %p\n", ip)
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	modify(ip)
	fmt.Println("int值被修改了，新值为:", i)
}

func modify(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	*ip = 2
}
