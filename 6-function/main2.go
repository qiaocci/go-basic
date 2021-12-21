package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	res := calc(1, 2, sum)
	fmt.Println(res)
}

func sum(x, y int) int {
	return x + y

}

// 函数作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return sum, nil
	default:
		err := errors.New("无法识别操作符")
		return nil, err

	}

}
