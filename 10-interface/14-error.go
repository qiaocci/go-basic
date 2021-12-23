package main

import (
	"fmt"
	myerror "go-basic/10-interface/14-error-interface"
	"syscall"
)

func main() {
	fmt.Println(myerror.New("创建失败"))
	fmt.Println(myerror.New("AAA") == myerror.New("AAA")) // false

	// golang的系统调用
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}
