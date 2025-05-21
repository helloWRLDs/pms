package main

import (
	"flag"
	"net"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"pms.auth/internal/config"
	"pms.auth/internal/data"
	grpchandler "pms.auth/internal/handlers/grpc"
	"pms.auth/internal/logic"
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/utils"
)

var conf config.Config

func init() {
	path := flag.String("path", "./services/auth/.env", "path to .env")
	flag.Parse()

	c, err := utils.LoadConfig[config.Config](*path)
	if err != nil {
		panic("failed to parse app config")
	}
	conf = c
	conf.Log.Init()
}

func main() {
	log := logger.Log.With(
		zap.String("func", "main"),
	)
	log.Debugw("check dsn", "dsn", conf.DB.DSN())
	db, err := postgres.Open(conf.DB.DSN())
	if err != nil {
		log.Errorw("failed to open db conn", "err", err)
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", conf.Host)
	if err != nil {
		log.Errorw("failed to create listener", "host", conf.Host, "err", err)
		os.Exit(1)
	}
	serv := grpc.NewServer()

	data := data.New(db, logger.Log)
	logic := logic.New(data, &conf, logger.Log)
	grpcHandler := grpchandler.New(logic, logger.Log)

	pb.RegisterAuthServiceServer(serv, grpcHandler)

	log.Infow("server started", "addr", conf.Host)
	if err := serv.Serve(lis); err != nil {
		log.Fatal("failed to serve", "err", err)
	}
}
