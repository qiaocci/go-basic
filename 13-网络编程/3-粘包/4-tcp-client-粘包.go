package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "122.51.107.111:20000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		conn.Write([]byte(msg))
	}
}
