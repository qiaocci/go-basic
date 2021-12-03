package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-client-and-server/proto"
	"log"
	"net"
)

type SearchService struct {
}

func (s SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	// 根据证书文件，构造tls凭证
	c, err := credentials.NewServerTLSFromFile("../conf/server.pem", "../conf/server.key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v\n", err)
	}

	server := grpc.NewServer(grpc.Creds(c)) // 设置服务端连接凭证
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("listen err:%v\n", err)
	}

	server.Serve(lis)
}
