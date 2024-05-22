package users

import (
	"context"
	pb "github.com/eduardoraider/go-box/proto/v1/users"
)

func (s *UserService) Get(ctx context.Context, payload *pb.UserRequest) (*pb.UserResponse, error) {
	u, err := s.factory.RestoreOne(payload.Id)
	if err != nil {
		return &pb.UserResponse{
			Error: err.Error(),
		}, err
	}

	return convertToUserResponse(*u), nil
}
