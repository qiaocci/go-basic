package main

import (
	"fmt"
	"reflect"
)

//main 主程序
func main() {
	fmt.Println("你好")
	v := 0b11
	a := 3
	fmt.Printf("%d \n", a)
	// 二进制
	fmt.Printf("%b \n", a)
	fmt.Println(v)
	b := true
	fmt.Println(b)
	s := `
		你好
		qiaocc
`
	fmt.Println(s)

	var x = '中'
	//var y = 'y'
	fmt.Printf("%q的类型为: %t, 二进制为: %b\n", x, x, x)
	xType := reflect.TypeOf(x).Kind()
	fmt.Printf("x的实际类型为%s\n", xType)

	s = "hallo沙河"
	for i:=0;i<len(s);i++{
		fmt.Printf("%v(%c)", s[i], s[i])
	}

	byteS :=[]byte(s)
	byteS[1] = 'e'
	fmt.Println(string(byteS))
}


