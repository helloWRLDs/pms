package notifierclient

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pms.api-gateway/internal/client"
	pb "pms.pkg/protobuf/services"
)

var _ client.Client = &NotifierClient{}

type Config struct {
	Host string `env:"HOST"`
}

type NotifierClient struct {
	pb.NotifierClient

	conn *grpc.ClientConn
	log  *logrus.Entry
}

func New(conf Config) (*NotifierClient, error) {
	log := logrus.WithField("client", "notifierClient")
	conn, err := grpc.NewClient(conf.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithError(err).Error("failed to connect to notifier service")
		return nil, err
	}
	return &NotifierClient{
		NotifierClient: pb.NewNotifierClient(conn),
		conn:           conn,
		log:            log,
	}, nil
}

func (c *NotifierClient) Close() error {
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
