package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

// Dream 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好go语言\n", p.name)
}

// 结构体： 传值， 传递副本。如果不用指针，只会修改副本, 无法修改接收者本身
func (p *Person) setAge(newAge int) {
	p.age = newAge // 本质上是(*p).age = newAge
}

func main() {
	p1 := newPerson("qiaocc", 21)
	println(p1.name, p1.age)
	p1.Dream()
	p1.setAge(22)
	fmt.Println(p1.age)
}
