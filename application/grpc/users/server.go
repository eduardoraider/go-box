package users

import (
	"github.com/eduardoraider/go-box/factories"
	domain "github.com/eduardoraider/go-box/internal/users"
	pb "github.com/eduardoraider/go-box/proto/v1/users"
	"github.com/eduardoraider/go-box/repositories"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUserService(repo repositories.UserWriteRepository, fact *factories.UserFactory) *UserService {
	return &UserService{
		repo:    repo,
		factory: fact,
	}
}

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo    repositories.UserWriteRepository
	factory *factories.UserFactory
}

func convertToUserPb(u domain.User) *pb.User {
	return &pb.User{
		Id:         u.ID,
		Name:       u.Name,
		Login:      u.Login,
		CreatedAt:  timestamppb.New(u.CreatedAt),
		ModifiedAt: timestamppb.New(u.ModifiedAt),
		LastLogin:  timestamppb.New(u.LastLogin),
	}
}

func convertToUserResponse(u domain.User) *pb.UserResponse {
	return &pb.UserResponse{
		User: convertToUserPb(u),
	}
}
