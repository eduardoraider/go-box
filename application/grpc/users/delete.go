package users

import (
	"context"
	pb "github.com/eduardoraider/go-box/proto/v1/users"
)

func (s *UserService) Delete(ctx context.Context, payload *pb.UserRequest) (*pb.UserResponse, error) {
	err := s.repo.Delete(payload.Id)
	if err != nil {
		return &pb.UserResponse{
			Error: err.Error(),
		}, err
	}

	return &pb.UserResponse{}, nil
}
