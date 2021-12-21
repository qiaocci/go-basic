package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-client-and-server/proto"
	"io/ioutil"
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
	// 从证书相关文件中读取和解析信息,得到证书公钥,秘钥对
	cert, err := tls.LoadX509KeyPair("../conf/server/server.pem", "../conf/server/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v\n", err)
	}

	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v\n", err)
	}
	//尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err: %v\n", err)
	}
	//构建基于 TLS 的 TransportCredentials 选项
	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(c)) // 设置服务端连接凭证
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("listen err:%v\n", err)
	}

	server.Serve(lis)
}
