package main

import "fmt"

func main() {
	res := intSum3(1, 2, 3)
	fmt.Println(res)

	// 函数类型
	c = add
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	println(c(1, 2))
}

func intSum(x, y int) int {
	return x + y
}

// 可变参数
func intSum2(x ...int) int {
	fmt.Println(x) // 切片类型
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}

func intSum3(x int, y ...int) int {
	for _, value := range y {
		x += value
	}
	return x
}

//我们可以使用type关键字来定义一个函数类型
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

var c calculation
