package main

import "fmt"

type Sayer4 interface {
	say()
}

type Mover4 interface {
	move()
}

type animal4 interface {
	Sayer4
	Mover4
}

type cat4 struct {
}

func (c cat4) say() {
	fmt.Println("喵喵喵")
}
func (c cat4) move() {
	fmt.Println("猫会动")
}

func hello(a Sayer4) {
	fmt.Println(a)
}

func main() {
	//从上面的代码中我们可以发现，使用值接收者实现接口之后，
	//不管是cat3结构体还是结构体指针*cat3类型的变量都可以赋值给该接口变量。
	var a animal4
	a = cat4{}
	a.say()
	a.move()

	// 函数的接收的参数子接口
	// 但是我们传递了父接口，也是可以的
	hello(a)
}
