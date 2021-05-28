package main

import (
	"encoding/json"
	"fmt"
)

type Student1 struct {
	ID     int `json:"id"`
	Gender string
	name   string // 私有不能被序列化
}

func main() {
	s1 := Student1{
		ID:     1,
		Gender: "男",
		name:   "qiaocc",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Printf("json.Marshal failed")
		return
	}
	fmt.Printf("json string : %s \n", data)
}
