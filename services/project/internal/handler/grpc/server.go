package grpchandler

import (
	"go.uber.org/zap"
	pb "pms.pkg/transport/grpc/services"
	"pms.project/internal/logic"
)

type ServerGRPC struct {
	pb.UnimplementedProjectServiceServer

	logic *logic.Logic
	log   *zap.SugaredLogger
}

func New(logic *logic.Logic, log *zap.SugaredLogger) *ServerGRPC {
	return &ServerGRPC{
		logic: logic,
		log:   log,
	}
}
