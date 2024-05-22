package users

import (
	"context"
	domain "github.com/eduardoraider/go-box/internal/users"
	pb "github.com/eduardoraider/go-box/proto/v1/users"
)

func (s *UserService) Create(ctx context.Context, payload *pb.UserRequest) (*pb.UserResponse, error) {
	u, err := domain.New(payload.Id, payload.Name, payload.Login, payload.Password)
	if err != nil {
		return &pb.UserResponse{
			Error: err.Error(),
		}, err
	}

	id, err := s.repo.Create(u)
	if err != nil {
		return &pb.UserResponse{
			Error: err.Error(),
		}, err
	}

	u.ID = id

	return convertToUserResponse(*u), nil
}
