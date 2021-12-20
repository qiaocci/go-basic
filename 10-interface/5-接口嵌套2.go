package main

import "fmt"

type Sayer3 interface {
	say()
}

type Mover3 interface {
	move()
}

type animal3 interface {
	Sayer3
	Mover3
}

type cat3 struct {
}

func (c cat3) say() {
	fmt.Println("喵喵喵")
}
func (c cat3) move() {
	fmt.Println("猫会动")
}

func main() {
	//从上面的代码中我们可以发现，使用值接收者实现接口之后，
	//不管是cat3结构体还是结构体指针*cat3类型的变量都可以赋值给该接口变量。
	var a animal3
	a = cat3{}
	a.say()
	a.move()
}
