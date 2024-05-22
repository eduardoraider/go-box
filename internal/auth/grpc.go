package auth

import (
	"context"
	pb "github.com/eduardoraider/go-box/proto/v1/auth"
)

type ServiceGRPC struct {
	pb.UnimplementedAuthServiceServer
	authHandler handler
}

func (svc *ServiceGRPC) Login(ctx context.Context, credsps *pb.Credentials) (*pb.TokenResponse, error) {
	creds := Credentials{
		Username: credsps.Username,
		Password: credsps.Password,
	}

	token, err, _ := svc.authHandler.auth(creds)
	if err != nil {
		return &pb.TokenResponse{
			Error: err.Error(),
		}, err
	}

	return &pb.TokenResponse{Token: token}, err
}

func HandleGrpcAuth(fn authenticateFunc) *ServiceGRPC {
	svc := &ServiceGRPC{
		authHandler: handler{fn},
	}
	return svc
}
