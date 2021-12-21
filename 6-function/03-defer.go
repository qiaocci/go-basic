package main

import "fmt"

func main() {
	data := double(2)
	fmt.Println("data:", data)
}

func demo1() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func double(x int) (res int) {
	defer func() {
		fmt.Printf("doubled(%d) = %d\n", x, res)
		res = res * 2 // defer能修改函数的返回值, 神奇~!
	}()
	return x + x
}
