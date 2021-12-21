package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	std1 := Student{
		Name:  "qiaocc",
		Score: 99,
	}
	t := reflect.TypeOf(std1)
	fmt.Println(t.Name(), t.Kind()) // Student struct

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name: %s, index: %d, type: %v, json tag: %v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

}
