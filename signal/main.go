package main

import (
	"log"
	"os"
	"os/signal"
)

// 监听所有信号
func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	log.Println("启动了程序")
	s := <-c
	log.Printf("收到信号=%v", s)
}
