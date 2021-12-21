package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"helloworld/pb"
	"net"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + r.Name}, nil
}

func main() {
	gServer := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(gServer, &server{}) // 在gRPC服务端注册服务
	reflection.Register(gServer)                 //在给定的gRPC服务器上注册服务器反射服务

	// Serve方法在listener上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。

	listener, _ := net.Listen("tcp", ":8972")
	gServer.Serve(listener)
}
