package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"cus_name"` // 自定义字段名称
	Age    int
	Weight float64 `json:"-"` // 忽略该字段
}

func main() {
	p1 := Person{
		Name:   "乔超超",
		Age:    29,
		Weight: 71.5,
	}

	// struct -> json string
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json marshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("str: %v\n", string(b))

	// json string -> struct
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json unmarshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("p2: %#v\n", p2)
}
