package logic

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
	"pms.project/internal/config"
	"pms.project/internal/data"
)

const dsn = "postgres://postgres:postgres@127.0.0.1:5432/project?sslmode=disable"

var (
	log   *zap.SugaredLogger
	logic *Logic
	db    *sqlx.DB
)

func TestMain(m *testing.M) {
	setupLogger()
	setupLogic()

	code := m.Run()

	terminate()

	os.Exit(code)
}

func setupLogger() {
	logger.WithConfig(
		logger.WithDev(true),
		logger.WithCaller(true),
		logger.WithFile(false, ""),
	).Init()
	log = logger.Log
}

func setupLogic() {
	var err error
	db, err = postgres.Open(dsn)
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
