package main

import "fmt"

func main() {
	var x interface{} // 定义一个空接口
	s := "你好"
	x = s
	fmt.Printf("类型: %T, value=%v\n", x, x) // 空接口: 没有任何方法的接口. 任何类型都实现了空接口

	i := 2
	x = i
	fmt.Printf("类型: %T, value=%v\n", x, x) // 空接口: 没有任何方法的接口. 任何类型都实现了空接口

	// 空接口作为map的值
	var std = make(map[string]interface{})
	std["name"] = "qiaocc"
	std["age"] = 29
	std["married"] = false
	fmt.Printf("std=%#v\n", std)
}

// 空接口作为函数参数, 代表任意类型
func show(x interface{}) {
	fmt.Printf("类型: %T, value=%v\n", x, x)
}
