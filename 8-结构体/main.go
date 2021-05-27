package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	type MyInt int    // 定义一个新类型MyInt
	type MyInt2 = int // 类型别名, 不是新的类型

	var p1 person
	p1.name = "qiaocc"
	p1.city = "上海"
	p1.age = 29
	fmt.Printf("p1=%v\n", p1)  // p1={qiaocc 上海 29}
	fmt.Printf("p1=%#v\n", p1) // p1=main.person{name:"qiaocc", city:"上海", age:29}

	// 匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "qiaocc"
	user.Age = 29
	fmt.Printf("%#v\n", user)

	// 创建指针类型的结构体, 方法1
	var p2 = new(person)
	fmt.Printf("p2类型:%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)
	p2.name = "qiaocc"
	p2.age = 29
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2)

	// 方法2
	p3 := &person{}
	fmt.Printf("p3类型:%T, p3=%#v\n", p3, p3)
	p3.name = "tom"
	p3.age = 21
	fmt.Printf("p3类型:%T, p3=%#v\n", p3, p3)

	// 零值
	var p4 person
	fmt.Printf("p4=%#v\n", p4) // p4=main.person{name:"", city:"", age:0}

	// 初始化
	p5 := person{
		name: "linda",
		city: "北京",
		age:  23,
	}
	fmt.Printf("p5=%#v\n", p5)

	// 结构体在内存中是连续的
	// int8: 一个字节 int16: 2个字节 int32: 4个字节 int64: 8个字节
	type test struct {
		a int8 // 占用一个字节
		b int8
		c int8
		d int8
	}
	n := test{1, 2, 3, 4}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	//
	p9 := NewPerson("tom", "shanghai", 21)
	fmt.Printf("p9=%#v\n", p9)
}

func NewPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
