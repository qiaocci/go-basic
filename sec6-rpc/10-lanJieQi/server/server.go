package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-client-and-server/proto"
	"log"
	"net"
	"runtime/debug"
)

type SearchService struct {
}

func (s SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			LoggingInterceptor,
			RecoveryInterceptor,
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("listen err:%v\n", err)
	}

	server.Serve(lis)
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("grpc method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("grpc method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()
	return handler(ctx, req)
}
