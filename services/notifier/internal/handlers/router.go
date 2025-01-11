package handlers

import (
	"context"

	"pms.notifier/internal/config"
	"pms.notifier/internal/service"
	pb "pms.pkg/protobuf/services"
)

type Server struct {
	pb.UnimplementedNotifierServer

	service *service.NotifierService
}

func New(conf config.Config) *Server {
	return &Server{
		service: service.New(conf.Gmail),
	}
}

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	res := &pb.GreetResponse{Success: false}
	err := s.service.GreetUser(ctx, req.Name, req.Email)
	if err != nil {
		return res, err
	}
	res.Success = true
	return res, nil
}
