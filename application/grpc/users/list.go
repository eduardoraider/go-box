package users

import (
	"context"
	pb "github.com/eduardoraider/go-box/proto/v1/users"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) List(ctx context.Context, payload *emptypb.Empty) (*pb.ListUserResponse, error) {
	us, err := s.factory.RestoreAll()
	if err != nil {
		return &pb.ListUserResponse{
			Error: err.Error(),
		}, err
	}
	data := make([]*pb.User, 0, len(us))
	for k, u := range us {
		data[k] = convertToUserPb(u)
	}

	return &pb.ListUserResponse{
		Users: data,
	}, nil

}
