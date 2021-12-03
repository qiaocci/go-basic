package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 如何发送SIGINT信号?
	// 1. Ctrl-C发送INT信号（SIGINT）；默认情况下，这会导致进程终止。
	// 2. 执行kill -s SIGINT PID

	// 将输入信号转发到 chan c。类似于绑定信号处理程序。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("get signal: %v\n", sig)
		done <- true
	}()

	// 阻塞等待, 直到获取到指定信号, 获取之后,向done发送true, 程序就会继续往下走了
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
