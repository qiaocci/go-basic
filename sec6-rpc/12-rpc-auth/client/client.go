package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-client-and-server/pkg"
	pb "grpc-client-and-server/proto"
	"log"
)

const PORT = "9001"

type Auth struct {
	AppKey    string
	AppSecret string
}

//GetRequestMetadata 获取当前请求认证所需的元数据（metadata）
func (a Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (a Auth) RequireTransportSecurity() bool {
	return true
}

func main() {
	tlsClient := pkg.Client{
		ServerName: "go-grpc-example",
		CertFile:   "../conf/server/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	auth := Auth{
		AppKey:    "qiaocc",
		AppSecret: "123",
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c), grpc.WithPerRPCCredentials(&auth))
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
