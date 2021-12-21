package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定义字符串，字符串是以什么形式存在于go编译器（utf8）编码
	name := "乔超超"

	// 其他语言： 0 乔, 1 超, 3 超
	// go语言中，表示字符串底层的每一个字节
	// go语言： 乔													   十进制  二进制
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2)) // 228 11100100
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2)) // 185 10111001
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2)) // 148 10010100

	// 超
	fmt.Println(name[3])
	fmt.Println(name[4])
	fmt.Println(name[5])

	// 超
	fmt.Println(name[6])
	fmt.Println(name[7])
	fmt.Println(name[8])
}
