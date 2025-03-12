package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"pms.notifier/internal/config"
	mqhandler "pms.notifier/internal/handlers/mq"
	"pms.pkg/logger"
	"pms.pkg/utils"
)

var conf *config.Config

func init() {
	envPath := flag.String("path", "./services/notifier/.env", "")
	flag.Parse()

	c, err := utils.LoadConfig[config.Config](*envPath)
	if err != nil {
		panic(err)
	}
	conf = &c
	conf.Log.Init()
}

func main() {
	conf.Log.Init()
	mq, err := mqhandler.New(conf, logger.Log)
	if err != nil {
		panic("failed to setup rabbitmq handler")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Infow("server started", "host", conf.Host)
		if err = mq.Listen(ctx); err != nil {
			log.Errorw("server stopped", "err", err)
		}
	}()

	<-quit
	log.Error("shutting down the server...")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("stopping clients and connections")
	// logic.CloseClients(ctx)

	log.Info("Server stopped gracefully")
}
