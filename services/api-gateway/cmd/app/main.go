package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pms.api-gateway/internal/config"
	"pms.api-gateway/internal/logic"
	"pms.api-gateway/internal/router"
	"pms.pkg/logger"
	"pms.pkg/utils"
)

var (
	conf config.Config
)

func init() {
	path := flag.String("path", "./services/api-gateway/.env", "path to .env")
	flag.Parse()

	c, err := utils.LoadConfig[config.Config](*path)
	if err != nil {
		panic("failed to parse app config")
	}
	conf = c
	conf.Log.Init()
}

func main() {
	log := logger.Log
	logic := logic.New(conf, log)
	log.Infof("check config: %#v", conf)

	serv := router.New(conf, logic, log)
	serv.SetupWS()
	serv.SetupREST()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Infow("server started", "host", conf.Host)
		if err := serv.Start(); err != nil {
			log.Errorw("server stopped", "err", err)
		}
	}()

	<-quit
	log.Error("shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("stopping clients and connections")
	logic.CloseClients(ctx)

	if err := serv.ShutdownWithContext(ctx); err != nil {
		log.Fatalw("failed to gracefully shut down server", "err", err)
	}

	log.Info("Server stopped gracefully")
}
