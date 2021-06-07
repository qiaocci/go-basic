package main

import "fmt"

type Result struct {
	Num, Ans int
}

type Cal int

func (c Cal) Square(num int) *Result {
	res := &Result{
		Num: num,
		Ans: num * num,
	}
	return res
}

func main() {
	cal := new(Cal)
	result := cal.Square(3)
	fmt.Printf("%d的平方等于%d\n", result.Num, result.Ans)
}
