package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, _ := tls.Dial("tcp", "127.0.0.1:8004", config)
	defer conn.Close()

	client := rpc.NewClient(conn)

	var result Result
	// 同步调用
	if err := client.Call("Cal.Square", 6, &result); err != nil {
		log.Fatalf("call failed, err:%v\n", err)
		return
	}

	//// 异步调用
	//asyncCall := client.Go("Cal.Square", 12, &result, nil)
	//fmt.Printf("%d的平方等于%d\n", result.Num, result.Ans)
	//
	//<-asyncCall.Done
	fmt.Printf("%d的平方等于%d\n", result.Num, result.Ans)
}
