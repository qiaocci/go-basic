package main

import "fmt"

func main() {
	var username string
	fmt.Println("请输入用户名:")

	// 提示用户输入，当用户输入后，把输入的值复制给内存地址对应的区域中
	fmt.Scanf("%s", &username)
	if username == "武沛齐" {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}
}
