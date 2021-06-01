package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	fmt.Printf("type=%v\n", reflect.TypeOf(x))
}

func main() {
	var a float32 = 3.14
	var b *float32
	var c rune
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	p := &person{
		"qiaocc", 21,
	}
	reflectType(p)
}
