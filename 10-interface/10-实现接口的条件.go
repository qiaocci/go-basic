package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	fmt.Println(w)

	// *bytes.Buffer must satisfy io.Writer
	// 断言: *bytes.Buffer必须实现了io.Writer接口, 用来做判断
	var _ io.Writer = (*bytes.Buffer)(nil)
	fmt.Println((*bytes.Buffer)(nil))

	//var p Person
	//fmt.Println(p == nil) // 报错
	var m map[string]int
	fmt.Println(m == nil) // true
	m = map[string]int{}
	m["a"] = 1
	fmt.Println(m)
}

type Person struct {
	name string
	age  int
}
