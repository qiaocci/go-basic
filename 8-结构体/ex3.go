package main

import (
	"fmt"
	"sort"
)

type Student2 struct {
	ID    int
	Name  string
	Age   int
	Score float64
}
type Class2 struct {
	Students map[int]*Student2
}

func (c Class2) AddStudent(id int, name string, age int) {
	std := &Student2{
		ID:   id,
		Name: name,
		Age:  age,
	}
	c.Students[id] = std
	fmt.Printf("保存成功, stds=%#v\n", c.Students)
}

func (c *Class2) DeleteStudent(id int) {
	for k := range c.Students {
		if k == id {
			delete(c.Students, k)
		}
	}
	fmt.Printf("删除std: id=%v\n", id)
}

func (c Class2) UpdateStudent(targetId int, name string, age int) {
	std := &Student2{
		ID:   targetId,
		Name: name,
		Age:  age,
	}
	c.Students[targetId] = std
	fmt.Println("更新成功")

}

func (c *Class2) ListStudent() {
	sortID := make([]int, 0)
	for k := range c.Students {
		sortID = append(sortID, k)
	}
	sort.Ints(sortID)
	for _, v := range sortID {
		fmt.Printf("std: id=%v, name=%v, age=%v\n", v, c.Students[v].Name, c.Students[v].Age)
	}
}

func main() {
	c := &Class2{
		Students: make(map[int]*Student2, 3),
	}
	//c.ListStudent()
	c.AddStudent(1, "qiaocc", 29)
	c.AddStudent(2, "tom", 28)
	c.AddStudent(3, "jerry", 27)
	c.ListStudent()
	c.DeleteStudent(2)
	c.ListStudent()

	c.UpdateStudent(3, "linda", 21)
	c.ListStudent()
}
