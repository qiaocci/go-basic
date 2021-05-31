package main

import "fmt"

type Sayer interface { // 接口
	say() string
}

type Cat struct {
}

func (c Cat) say() string { // 实现接口的方法
	return "喵喵喵"
}

type Dog struct {
}

func (d Dog) say() string {
	return "汪汪汪"
}

func main() {
	c := Cat{}
	//println(c.say())
	d := Dog{}
	println(d.say())

	var s Sayer // 声明Sayer类型的变量
	s = c
	println(s.say())
	s = d
	fmt.Println(s.say())
}
