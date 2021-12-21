package main

import "fmt"

func main() {
	//var a *int
	//// 指针赋值： 方法1
	////n := 100
	////a = &n
	//
	//// 方法2
	//*a = 100
	//fmt.Println(a, *a)

	// 错误写法
	//var b map[string]int
	//b["欧阳娜娜"] = 18
	//fmt.Println(b)

	// 以上都是错误写法.
	// 引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。

	a := new(int) // new生成指针类型 *int
	b := new(bool)
	*a = 10
	fmt.Printf("a=%v, 类型: %T, *a=%v\n", a, a, *a)
	fmt.Printf("b=%v, 类型: %T, *b=%v\n", b, b, *b)

	var c map[string]int
	c = make(map[string]int, 10)
	c["欧阳娜娜"] = 18
	fmt.Println(c)

}
