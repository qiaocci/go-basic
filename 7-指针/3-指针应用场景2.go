package main

import "fmt"

func main() {
	name := "武沛齐"
	changeData(name)
	fmt.Println(name) // 武沛齐

	changeData2(&name)
	fmt.Println(name) // 哈哈哈哈哈
}

func changeData(data string) {
	data = "哈哈哈哈哈"
}

func changeData2(ptr *string) {
	*ptr = "哈哈哈哈哈"
}
