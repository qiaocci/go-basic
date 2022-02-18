package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	// 指定的src复制到dst，直到在src上达到EOF或引发错误为止。
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatalln(err)
	}
}
