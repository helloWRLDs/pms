package data

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"pms.pkg/datastore/sqlite"
	"pms.pkg/logger"
)

const (
	dsn = "../../data/project.db"
)

var (
	repo *Repository
	log  *zap.SugaredLogger
)

func TestMain(m *testing.M) {
	setupLogger()
	setupDB()
	code := m.Run()
	os.Exit(code)
}

func setupLogger() {
	logger.WithConfig(
		logger.WithDev(true),
		logger.WithLevel("debug"),
		logger.WithCaller(true),
	).Init()

	log = logger.Log
}

func setupDB() {
	db, err := sqlite.Open(dsn)
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	repo = New(db, logger.Log)
}

func Test_Test(t *testing.T) {
	t.Log("migrate test")
}
