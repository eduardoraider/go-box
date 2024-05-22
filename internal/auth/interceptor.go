package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

var bypass = map[string]string{}

func AddBypassInterceptor(svc, method string) {
	bypass[svc] = method
}

func Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// bypass token validation
	parts := strings.Split(info.FullMethod, "/")
	if len(parts) == 3 {
		svcName := parts[1]
		methodName := parts[2]

		if v, ok := bypass[svcName]; ok {
			if v == methodName {
				return handler(ctx, req)
			}
		}
	}

	// token validation
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md["authorization"]) > 0 {
			token := md["authorization"][0]
			token = strings.TrimPrefix(token, "Bearer ")

			claims, err, _ := validate(token)
			if err != nil {
				return nil, err
			}

			ctx := context.WithValue(ctx, "user_id", claims.UserID)
			ctx = context.WithValue(ctx, "user_name", claims.UserName)

			return handler(ctx, req)
		}
	}

	return nil, status.Error(codes.Unauthenticated, "authentication failed")
}
