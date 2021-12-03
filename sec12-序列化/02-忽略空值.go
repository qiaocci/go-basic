package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Email string   `json:"cus_email,omitempty"`
	Hobby []string `json:"cus_hobby,omitempty"`
}

func main() {
	omitEmpty()
}

func omitEmpty() {
	u1 := User{
		Name: "乔超超",
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str: %v\n", string(b))
}
