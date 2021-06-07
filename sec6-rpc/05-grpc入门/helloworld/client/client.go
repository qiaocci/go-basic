package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"helloworld/pb"
)

func main() {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "qiaocc"})
	if err != nil {
		fmt.Printf("say hello failed, err:%v\n", err)
		return
	}

	fmt.Printf("greet message: %v\n", r.Message)

}
