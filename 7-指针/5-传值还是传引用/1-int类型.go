package main

import "fmt"

// 我们传递一个int类型的参数，传递的其实是这个参数的一个副本
func main() {
	i := 10
	fmt.Printf("原始指针的内存地址是：%p\n", &i)
	modify(i)
	fmt.Println("int值被修改了，新值为:", i)
}

func modify(num int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &num)
	num = 2
}
