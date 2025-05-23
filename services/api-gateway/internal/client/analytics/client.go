package analyticsclient

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pms.api-gateway/internal/client"
	configgrpc "pms.pkg/transport/grpc/config"
	pb "pms.pkg/transport/grpc/services"
)

var _ client.Client = &AnalyticsClient{}

type AnalyticsClient struct {
	pb.AnalyticsServiceClient

	conn *grpc.ClientConn
	log  *zap.SugaredLogger
}

func New(conf configgrpc.ClientConfig, logger *zap.SugaredLogger) (*AnalyticsClient, error) {
	log := new(zap.SugaredLogger)
	{
		if conf.DisableLog {
			log = zap.NewNop().Sugar()
		} else {
			log = logger.Named("analyticsclient.New").With(
				zap.String("host", conf.Host),
			)
		}
	}
	log.Info("analyticsclient.New called")

	conn, err := grpc.NewClient(conf.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorw("failed to connect to auth service", "err", err)
		return nil, err
	}
	return &AnalyticsClient{
		AnalyticsServiceClient: pb.NewAnalyticsServiceClient(conn),
		conn:                   conn,
		log:                    log,
	}, nil
}

func (c *AnalyticsClient) Close() error {
	log := c.log.With(zap.String("func", "Close"))
	log.Debug("Close func called")

	err := c.conn.Close()
	if err != nil {
		log.Errorw("failed to close connection", "err", err)
		return err
	}
	log.Debug("connection closed")
	return nil
}
