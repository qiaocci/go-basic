package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-client-and-server/proto"
	"log"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
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
