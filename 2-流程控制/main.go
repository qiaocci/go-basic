package main

import "fmt"

func main() {
	//ifDemo1()
	switchDemo1()
}

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("优秀")
	} else if score > 75 {
		fmt.Println("良好")
	} else {
		fmt.Println("及格")
	}

}

func switchDemo1() {
	finger := 31
	switch finger {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	default:
		fmt.Println("无效的输入")
	}
}
