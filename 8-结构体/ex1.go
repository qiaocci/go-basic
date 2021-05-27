package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stds := []student{
		{name: "tom", age: 22},
		{name: "linda", age: 23},
		{name: "jerry", age: 24},
	}
	for _, std := range stds {
		fmt.Printf("std=%#v, %p\n", std, &std)
		m[std.name] = &std
	}
	for k, v := range m {
		fmt.Println(k, v.name)
	}
	//fmt.Println(m)
}
