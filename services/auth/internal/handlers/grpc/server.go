package grpchandler

import (
	"go.uber.org/zap"
	"pms.auth/internal/logic"
	pb "pms.pkg/transport/grpc/services"
)

type ServerGRPC struct {
	pb.UnimplementedAuthServiceServer

	logic *logic.Logic
	log   *zap.SugaredLogger
}

func New(logic *logic.Logic, log *zap.SugaredLogger) *ServerGRPC {
	return &ServerGRPC{
		logic: logic,
		log:   log,
	}
}
