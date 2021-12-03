package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User              // 匿名嵌套
	Password *struct{} `json:"password,omitempty"` // 匿名结构体指针类型
}

func main() {
	omitPassword()
}

func omitPassword() {
	u1 := User{
		Name:     "乔超超",
		Password: "123",
	}
	b, err := json.Marshal(PublicUser{User: &u1})
	if err != nil {
		fmt.Printf("json marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str: %v\n", string(b))

}
