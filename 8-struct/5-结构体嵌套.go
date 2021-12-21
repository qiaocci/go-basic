package main

import "fmt"

type Address struct {
	Province string
	City     string
	Street   string
}

type User struct {
	Name    string
	Gender  string
	Address Address
}

func main() {
	u := User{
		Name:   "qiaocc",
		Gender: "男",
		Address: Address{
			Province: "河南",
			City:     "洛阳",
			Street:   "西工区",
		},
	}
	fmt.Printf("u=%#v", u)
}
