package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "122.51.107.111:8004")
	if err != nil {
		log.Fatalf("dial failed, err:%v\n", err)
		return
	}
	var result Result
	// 同步调用
	//if err := client.Call("Cal.Square", 6, &result); err != nil {
	//	log.Fatalf("call failed, err:%v\n", err)
	//	return
	//}

	// 异步调用
	asyncCall := client.Go("Cal.Square", 12, &result, nil)
	fmt.Printf("%d的平方等于%d\n", result.Num, result.Ans)

	<-asyncCall.Done
	fmt.Printf("%d的平方等于%d\n", result.Num, result.Ans)
}
