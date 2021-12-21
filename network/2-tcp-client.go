package main

import (
	"log"
	"net"
)

func main() {
	server, err := net.ResolveTCPAddr("tcp", "118.24.4.110:8001")
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		log.Println("dail err:", err)
		return
	}
	log.Println("conn success")
	defer conn.Close()

	words := "hello server"
	msgBack, err := conn.Write([]byte(words))
	if err != nil {
		log.Printf("addr=%v, fatal err=%v", conn.RemoteAddr(), err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	log.Printf("addr=%v, 服务器反馈=%v. msgBack=%v, 实际发送了=%v", conn.RemoteAddr(), string(buffer[:n]), msgBack, len(words))
	conn.Write([]byte("ok"))
	conn.Write([]byte("hhh"))
}
