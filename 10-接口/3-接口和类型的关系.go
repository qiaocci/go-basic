package main

import "fmt"

type Sayer1 interface {
	say()
}

type Mover1 interface {
	move()
}

type dog1 struct {
	name string
}

type cat1 struct {
	name string
}

// 一个类型, 可以实现多个接口
func (d dog1) say() {
	fmt.Printf("%s 汪汪叫\n", d.name)
}
func (d dog1) move() {
	fmt.Printf("%s 动起来了\n", d.name)
}

// 一个接口, 可以被多个类型实现
func (c cat1) say() {
	fmt.Printf("%s 喵喵叫\n", c.name)
}

func main() {
	d := dog1{
		name: "旺财",
	}
	var s Sayer1
	var m Mover1
	s = d
	m = d
	s.say()
	m.move()

	c := cat1{
		"猫",
	}
	s = c
	s.say()
}
