package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "122.51.107.111:20000")
	if err != nil {
		fmt.Println("dail err:", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}

		_, err = conn.Write([]byte(inputInfo)) // 发数据
		if err != nil {
			fmt.Println("write error", err)
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read error", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
