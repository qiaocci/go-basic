package main

import "fmt"

type MyInt int

func (m MyInt) sayHello() {
	fmt.Println("hello")
}

func main() {
	var m MyInt
	m = 1
	m.sayHello()
	fmt.Printf("m=%#v, 类型:%T\n", m, m)
}
