package main

import "fmt"

func main() {
	res := scoreRange(-1)
	fmt.Println(res)
}

func scoreRange(x int) string {
	if x < 0 {
		panic("分数小于0")
	}
	if x > 90 {
		return "优秀"
	}
	return "普通"
}
