package grpchandler

import (
	"go.uber.org/zap"
	"pms.analytics/internal/logic"
	pb "pms.pkg/transport/grpc/services"
)

type ServerGRPC struct {
	pb.UnimplementedAnalyticsServiceServer

	logic *logic.Logic
	log   *zap.SugaredLogger
}

func New(logic *logic.Logic, log *zap.SugaredLogger) *ServerGRPC {
	return &ServerGRPC{
		logic: logic,
		log:   log,
	}
}
