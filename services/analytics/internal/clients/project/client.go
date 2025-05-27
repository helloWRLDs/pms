package projectclient

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	configgrpc "pms.pkg/transport/grpc/config"
	pb "pms.pkg/transport/grpc/services"
)

type ProjectClient struct {
	pb.ProjectServiceClient

	conn *grpc.ClientConn
	log  *zap.SugaredLogger
}

func New(conf configgrpc.ClientConfig, logger *zap.SugaredLogger) (*ProjectClient, error) {
	var log *zap.SugaredLogger
	{
		if conf.DisableLog {
			log = zap.NewNop().Sugar()
		} else {
			log = logger.Named("projectclient.New").With(
				zap.String("host", conf.Host),
			)
		}
	}
	log.Debug("New called")

	conn, err := grpc.NewClient(conf.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorw("failed to connect to auth service", "err", err)
		return nil, err
	}
	return &ProjectClient{
		ProjectServiceClient: pb.NewProjectServiceClient(conn),
		conn:                 conn,
		log:                  log,
	}, nil
}

func (c *ProjectClient) Close() error {
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
