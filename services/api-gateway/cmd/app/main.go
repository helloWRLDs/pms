package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"pms.api-gateway/internal/config"
	"pms.api-gateway/internal/router"
	"pms.pkg/cfg"
	"pms.pkg/logger"
)

var (
	conf config.Config
)

func init() {
	path := flag.String("path", "./services/api-gateway/.env", "path to .env")
	flag.Parse()

	c, err := cfg.Load[config.Config](*path)
	if err != nil {
		logrus.Fatal("failed to parse app config", "err", err)
	}
	conf = c
	logger.Init(conf.Log)
}

func main() {
	log := logrus.WithField("func", "main")
	defer logger.Close(conf.Log)

	serv := router.New(conf)
	serv.SetupRoutes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.WithField("host", conf.Host).Info("server started")
		if err := serv.Start(); err != nil {
			log.WithError(err).Fatal("server stopped")
		}
	}()

	<-quit
	log.Error("shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("stopping clients and connections")
	time.Sleep(1 * time.Second)

	if err := serv.ShutdownWithContext(ctx); err != nil {
		log.WithError(err).Fatal("failed to gracefully shut down server")
	}

	log.Info("Server stopped gracefully")
}
