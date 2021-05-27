package main

import "fmt"

func main() {
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("a=%d 指针: %p\n", a, &a)
	fmt.Printf("b=%p 类型： %T\n", b, b)
	fmt.Println(&b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("c=%v, 类型: %T\n", c, c)

	num := 0xc00001c0c0
	fmt.Printf("num=%v, 类型: %T, 指针: %v\n", num, num, &num)
	// num=824633835712, 类型: int, 指针: 0xc0000b8030
	// num打印的时候，转成十进制了

	// 指针传值
	n := 10
	modify1(n)
	fmt.Printf("n=%v\n", n) // 传递副本，不修改原始值
	modify2(&n)
	fmt.Printf("n=%v\n", n) // 如果想修改原始值，需要用指针
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
