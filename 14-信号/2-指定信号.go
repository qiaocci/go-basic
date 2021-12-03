package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	log.Println("启动了程序")
	s := <-c
	log.Printf("收到了信号=%v", s)
}
