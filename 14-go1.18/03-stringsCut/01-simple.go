package main

import (
	"fmt"
	"strings"
)

func main() {
	addr := "192.168.1.1:8080"
	pos := strings.Index(addr, ":")
	fmt.Println(pos) // 11
	if pos == -1 {
		panic("非法地址")
	}
	ip, port := addr[:pos], addr[pos+1:]
	fmt.Println(ip, port)

	ip, port, ok := strings.Cut(addr, ":")
	fmt.Println(ip, port, ok)

}
