package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var firstSigusr1 = true

func main() {
	// 忽略 Control-C (SIGINT)
	signal.Ignore(syscall.SIGINT)

	c1 := make(chan os.Signal, 2)

	// SIGHUP（挂起）默认会使得程序退出
	signal.Notify(c1, syscall.SIGHUP)
	//SIGUSR1和SIGUSR2是发送给一个进程的信号，它表示了用户定义的情况
	signal.Notify(c1, syscall.SIGUSR1)

	go func() {
		for {
			switch <-c1 {
			case syscall.SIGHUP:
				fmt.Println("sighup, reset sighup")
				signal.Reset(syscall.SIGHUP)
			case syscall.SIGUSR1:
				if firstSigusr1 {
					fmt.Println("first usr1, notify inter ")
				}

			}
		}
	}()
}
