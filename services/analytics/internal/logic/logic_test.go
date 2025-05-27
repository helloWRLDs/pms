package logic

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"pms.analytics/internal/config"
	"pms.analytics/internal/data"
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
)

var (
	logic *Logic
	log   *zap.SugaredLogger
	db    *sqlx.DB
)

func setupLogger() {
	logger.WithConfig(
		logger.WithCaller(true),
		logger.WithDev(true),
		logger.WithLevel("debug"),
	).Init()
	log = logger.Log
}

func setup() {
	var err error
	db, err = postgres.Open("postgres://postgres:postgres@127.0.0.1:5432/analytics?sslmode=disable")
	if err != nil {
		print("failed to connect to db", "err", err)
	}

	logic = New(data.New(db, log), &config.Config{}, log)
}

func terminate() {
	if err := db.Close(); err != nil {
		log.Errorw("failed to close db", "err", err)
	}
}

func TestMain(m *testing.M) {
	setupLogger()
	setup()
	code := m.Run()
	terminate()
	os.Exit(code)
}
