package server

import (
	"context"

	pb "pms.pkg/protobuf/services"
)

type Server struct {
	pb.UnimplementedAuthServer
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	return nil, nil
}
