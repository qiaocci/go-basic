package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (c *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num

	return nil
}

func main() {
	rpc.Register(new(Cal))
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":8004", config)
	log.Printf("Serving rpc server on port %d\n", 8004)

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
