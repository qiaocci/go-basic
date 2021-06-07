package main

import (
	"log"
	"net/http"
	"net/rpc"
	"time"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (c *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num

	time.Sleep(time.Second * 3) //模拟耗时任务
	return nil
}

func main() {
	err := rpc.Register(new(Cal)) // 发布rpc方法
	if err != nil {
		log.Fatalf("rpc register failed, err:%v\n", err)
	}
	rpc.HandleHTTP() // 注册用于处理rpc消息的handler

	log.Printf("Serving rpc server on port %d\n", 8004)
	// 监听8000端口,等待rpc请求
	if err := http.ListenAndServe(":8004", nil); err != nil {
		log.Fatalf("serving failed, err:%v\n", err)
	}
}
