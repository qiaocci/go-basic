package main

type Mover interface {
	move()
}

type dog struct {
}

//func (d dog) move() { // 值类型
//
//}

func (d *dog) move() { // 指针类型

}

// 值类型的接收者, 结构体是值dog还是指针*dog类型都可以赋值调用
// 指针类型的接收者, 结构体必须是指针
func main() {
	//wangcai := dog{}
	//wangcai.move()
	var m Mover
	//m = wangcai
	//m.move()

	fugui := &dog{}
	m = fugui
	m.move()
}
