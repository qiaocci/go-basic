package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-client-and-server/pkg"
	pb "grpc-client-and-server/proto"
	"log"
)

const PORT = "9001"

func main() {
	tlsClient := pkg.Client{
		ServerName: "go-grpc-example",
		CertFile:   "../conf/server/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
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
