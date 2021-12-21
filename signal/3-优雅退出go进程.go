package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(
		c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1,
		syscall.SIGUSR2,
	)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				log.Println("退出:", s)
				exitFunc()
			case syscall.SIGUSR1:
				log.Println("usr1:", s)
			case syscall.SIGUSR2:
				log.Println("usr2:", s)
			default:
				log.Println("其他信号:", s)
			}
		}
	}()
}

func exitFunc() {
	log.Println("开始退出")
	log.Println("执行清理")
	log.Println("结束退出")
	os.Exit(0)
}
