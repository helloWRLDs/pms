package grpc

import (
	"context"

	pb "pms.pkg/protobuf/services"
)

type ServerGRPC struct {
	pb.UnimplementedAuthServer
}

func (s *ServerGRPC) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	return nil, nil
}
