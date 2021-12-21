package main

import "fmt"

type Result struct {
	Num, Ans int
}

type Cal int

func (c *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	cal := new(Cal)
	var result Result
	err := cal.Square(3, &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d的平方等于%d\n", result.Num, result.Ans)
}
