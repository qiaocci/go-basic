package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept() // 阻塞, 等待新的连接
		if err != nil {
			log.Println(err) // eg, connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(time.RFC3339)+"\n")
		if err != nil {
			return // eg, client disconnected¬
		}
		time.Sleep(time.Second)
	}
}
