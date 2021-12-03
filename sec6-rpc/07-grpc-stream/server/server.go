package main

import (
	"google.golang.org/grpc"
	pb "grpc-client-and-server/proto"
	"log"
	"net"
)

type StreamService struct{}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("listen err:%v\n", err)
	}

	server.Serve(lis)
}

func (s *StreamService) List(r *pb.Request, stream pb.StreamService_ListServer) error {
	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	return nil
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	return nil
}
