package main

import "fmt"

type Skills []string
type Human struct {
	name  string
	age   int
	phone string
}

type Student3 struct {
	Human
	Skills
	int
	classroom string
	phone     string
}

func main() {
	tom := Student3{
		Human{name: "tom", age: 29, phone: "11111"},
		Skills{"python", "go"},
		10,
		"101",
		"22222",
	}
	// 访问字段
	fmt.Println(tom.classroom)
	fmt.Println(tom.Human.name) // 通过Human获取
	fmt.Println(tom.name)       // 直接获取

	// 修改字段
	tom.classroom = "102"
	tom.age = 30
	fmt.Println("修改后")
	fmt.Println(tom.classroom)
	fmt.Println(tom.age)

	// Student访问属性age和name的时候，就像访问自己所有用的字段一样，对，匿名字段就是这样，能够实现字段的继承。

	// 如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，那么该怎么办呢？
	//Go里面很简单的解决了这个问题，最外层的优先访问，也就是当你通过student.phone访问的时候，是访问student里面的字段，而不是human里面的字段。
	fmt.Println(tom.phone)       // 访问Student.phone
	fmt.Println(tom.Human.phone) // 访问Human.phone
}
