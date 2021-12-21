package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-client-and-server/pkg"
	pb "grpc-client-and-server/proto"
	"log"
	"net/http"
	"strings"
	"time"
)

type SearchService struct {
	auth *Auth
}

func (s SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	// 检测是否过期
	for i := 0; i < 5; i++ {
		if ctx.Err() == context.Canceled {
			return nil, status.Errorf(codes.Canceled, "SearchService cancelled")
		}
		time.Sleep(time.Second)
	}
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	certFile := "../conf/server/server.pem"
	keyFile := "../conf/server/server.key"
	tlsServer := pkg.Server{
		CertFile: certFile,
		KeyFile:  keyFile,
	}

	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	}

	mux := GetHTTPServeMux()

	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server, &SearchService{})

	http.ListenAndServeTLS(":"+PORT,
		certFile,
		keyFile,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				server.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}))
}

func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("http: hello qiaocc"))
	})
	return mux
}

type Auth struct {
	appKey    string
	appSecret string
}

func (a Auth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "自定义认证token失败")
	}

	var (
		appKey    string
		appSecret string
	)
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}

	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return status.Errorf(codes.Unauthenticated, "自定义认证token失败")
	}

	return nil
}

func (a Auth) GetAppKey() string {
	return "qiaocc"
}

func (a Auth) GetAppSecret() string {
	return "123"
}
