package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) GetProjectStats(ctx context.Context, companyID string) ([]*dto.UserTaskStats, error) {
	log := l.log.Named("GetProjectStats").With(
		zap.String("companyID", companyID),
	)
	log.Debug("GetProjectStats called")

	stats, err := l.analyticsClient.GetUserTaskStats(ctx, &pb.GetUserTaskStatsRequest{
		CompanyId: companyID,
	})
	if err != nil {
		return nil, err
	}
	return stats.Stats, nil
}
