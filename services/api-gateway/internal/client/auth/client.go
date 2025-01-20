package authclient

import (
	"github.com/sirupsen/logrus"
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
	pb.AuthClient

	conn *grpc.ClientConn
	log  *logrus.Entry
}

func New(conf Config) (*AuthClient, error) {
	log := logrus.WithField("client", "authClient")
	conn, err := grpc.NewClient(conf.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithError(err).Error("failed to connect to auth service")
		return nil, err
	}
	return &AuthClient{
		AuthClient: pb.NewAuthClient(conn),
		conn:       conn,
		log:        log,
	}, nil
}

func (c *AuthClient) Close() error {
	log := c.log.WithField("func", "Close")
	log.Debug("Close func called")

	err := c.conn.Close()
	if err != nil {
		log.WithError(err).Error("failed to close connection")
		return err
	}
	log.Debug("connection closed")
	return nil
}
