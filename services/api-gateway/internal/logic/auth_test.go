package logic

import (
	"testing"

	authclient "pms.api-gateway/internal/client/auth"
	"pms.pkg/logger"
	configgrpc "pms.pkg/transport/grpc/config"
)

func Test_AuthClientConn(t *testing.T) {
	authClient, err := authclient.New(configgrpc.ClientConfig{Host: "localhost:50051"}, logger.Log)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(authClient.Ping())
	t.Log(authClient.Ping())
	t.Log(authClient.Ping())
}
