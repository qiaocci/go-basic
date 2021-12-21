package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Name   string
	Gender string
}

type Class struct {
	Title   string
	Student []*Student
}

func main() {
	c := &Class{
		Title:   "一班",
		Student: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		std := Student{
			ID:     i,
			Name:   fmt.Sprintf("std_%v", i),
			Gender: "男",
		}
		c.Student = append(c.Student, &std)
	}

	// 序列化
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json: %s \n", data)

	// 反序列化
	c1 := &Class{}
	err = json.Unmarshal([]byte(data), c1)
	if err != nil {
		fmt.Printf("json Unmarshal failed")
		return
	}
	fmt.Printf("c1=%#v\n", c1)
}
