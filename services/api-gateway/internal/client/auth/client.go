package authclient

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pms.api-gateway/internal/client"
	pb "pms.pkg/protobuf/services"
)

var _ client.Client = &AuthClient{}

type Config struct {
	Host string `env:"HOST"`
}

type AuthClient struct {
	pb.AuthServiceClient

	conn *grpc.ClientConn
	log  *zap.SugaredLogger
}

func New(conf Config, logger *zap.SugaredLogger) (*AuthClient, error) {
	log := logger.With(
		zap.String("func", "authclient.New"),
		zap.String("host", conf.Host),
	)
	log.Debug("New called")

	conn, err := grpc.NewClient(conf.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorw("failed to connect to auth service", "err", err)
		return nil, err
	}
	return &AuthClient{
		AuthServiceClient: pb.NewAuthServiceClient(conn),
		conn:              conn,
		log:               log,
	}, nil
}

func (c *AuthClient) Close() error {
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
