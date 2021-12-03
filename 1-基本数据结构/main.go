package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数字转字符串
	res := strconv.Itoa(120) // "120"
	fmt.Println(res)

	// 修改字符串
	s := "hello"
	//s[0] = '1' // error
	tmp := []byte(s)
	tmp[0] = 'a'
	// 定义字符串，字符串是以什么形式存在于Go编译器（utf-8编码）
	fmt.Println(tmp)    // [97 101 108 108 111]
	fmt.Println(tmp[0]) // 97

	s2 := string(tmp)
	fmt.Println(s2)
}
