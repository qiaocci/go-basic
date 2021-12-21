package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-client-and-server/proto"
	"log"
)

const PORT = "9001"

func main() {
	// 根据证书文件, 构造tls证书凭证
	c, err := credentials.NewClientTLSFromFile("../conf/server.pem", "go-grpc-example")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v\n", err)
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c)) // 配置连接的option选项
	if err != nil {
		log.Fatalf("grpc.Dial err: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "qiaocc",
	})
	if err != nil {
		log.Fatalf("client search err:%v\n", err)
	}
	log.Printf("response: %v\n", resp.Response)
}
