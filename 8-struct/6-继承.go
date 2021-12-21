package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%s can move \n", a.name)
}

type Dog struct {
	Feet    int
	*Animal //通过嵌套匿名结构体实现继承
}

func (d Dog) wang() {
	fmt.Printf("%s can wang~\n", d.name)
}

func main() {
	d1 := Dog{
		Feet: 4,
		Animal: &Animal{
			"旺财",
		},
	}
	d1.wang()
	d1.move()
}
