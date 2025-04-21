package logic

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"pms.auth/internal/config"
	"pms.auth/internal/data"
	"pms.pkg/datastore/sqlite"
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
	db, err = sqlite.Open("../../data/users.db")
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
