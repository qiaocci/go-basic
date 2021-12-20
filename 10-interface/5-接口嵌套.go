package main

import "fmt"

type Sayer2 interface {
	say()
}

type Mover2 interface {
	move()
}

type animal interface {
	Sayer2
	Mover2
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}
func (c cat) move() {
	fmt.Println("猫会动")
}

func main() {
	var a animal
	a = cat{}
	a.say()
	a.move()
}
