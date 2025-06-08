package grpchandler

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"pms.pkg/logger"
	pb "pms.pkg/transport/grpc/services"
)

var (
	conn   *grpc.ClientConn
	client pb.AuthServiceClient
	log    *zap.SugaredLogger
)

func setupLogger() {
	logger.WithConfig(
		logger.WithCaller(true),
		logger.WithLevel("debug"),
		logger.WithDev(true),
	).Init()
	log = logger.Log
}

func setup() {
	c, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	conn = c
	client = pb.NewAuthServiceClient(conn)
}

func teardown() {
	conn.Close()
}

func TestMain(m *testing.M) {
	setupLogger()
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
