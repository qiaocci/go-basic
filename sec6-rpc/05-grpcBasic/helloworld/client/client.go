package main

import (
	"context"
	"google.golang.org/grpc"
	"helloworld/pb"
	"log"
)

func main() {
	conn, _ := grpc.Dial(":8972", grpc.WithInsecure())
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	SayHello(c)

}

func SayHello(c pb.GreeterClient) error {
	r, _ := c.SayHello(context.Background(), &pb.HelloRequest{Name: "qiaocc"})
	log.Printf("greet message: %v\n", r.Message)
	return nil
}
