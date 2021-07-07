package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}
type Profile struct {
	Website string `json:"website"`
	Slogan  string `json:"slogan"`
}

func main() {
	nestedStruct()
}
func nestedStruct() {
	u1 := User{
		Name:  "乔超超",
		Hobby: []string{"足球", "篮球"},
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str: %v\n", string(b))

}
