package main

import "fmt"

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int, value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	default:
		fmt.Printf("unsupported type!")
	}

}

func main() {
	var x interface{}
	x = "你好"
	s, ok := x.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Printf("断言失败")
	}

	justifyType(x)

}
