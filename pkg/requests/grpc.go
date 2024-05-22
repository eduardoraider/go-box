package requests

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"os"
)

func GetGRPCConn() *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf(":%s", os.Getenv("SERVER_GRPC_PORT")))
	if err != nil {
		log.Fatalf("erro connection to gRPC server: %v", err)
	}
	return conn
}

func GetGRPCWithTokenConn() *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf(":%s", os.Getenv("SERVER_GRPC_PORT")), grpc.WithUnaryInterceptor(tokenInterceptor))
	if err != nil {
		log.Fatalf("erro connection to gRPC server: %v", err)
	}
	return conn
}

func tokenInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	token, err := readCacheToken()
	if err != nil {
		log.Printf("Failed to read token from cache: %s", err)
		return err
	}

	opts = append(opts, grpc.PerRPCCredentials(oauth.TokenSource{
		TokenSource: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: token,
		}),
	}))

	return invoker(ctx, method, req, reply, cc, opts...)
}
