package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8001")
	log.Printf("等待连接")

	if err != nil {
		log.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			log.Println("accept failed, err:", err)
			continue
		}
		log.Printf("tcp conn success, addr=%v", conn.RemoteAddr().String())
		go process(conn) // 启动goroutine处理连接
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// 接收客户端信息
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("conn err=%v", err)
			return
		}
		log.Printf("receive data=%v", string(buffer[:n]))

		// 反馈给客户端
		bufferReturn := "我收到了"
		n2, err := conn.Write([]byte(bufferReturn))
		// 确认客户端未收到回执
		if err != nil {
			log.Printf("没有收到回执, addr=%v", conn.RemoteAddr())
			return
		}
		// 确认客户端收到回执
		n3, err := conn.Read(buffer)
		log.Printf("addr=%v, 客户端收到回执=%v, 客户收到了=%v, 实际发送了=%v", conn.RemoteAddr(), string(buffer[:n3]), n2, len(bufferReturn))

	}
}
