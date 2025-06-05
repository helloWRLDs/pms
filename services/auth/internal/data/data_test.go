package data

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
)

const (
	dsn = "postgres://postgres:postgres@127.0.0.1:5432/auth?sslmode=disable"
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
	db, err := postgres.Open(dsn)
	if err != nil {
		log.Fatalw("failed to connect to db", "err", err)
	}
	repo = New(db, logger.Log)
}

func Test_Test(t *testing.T) {
	t.Log("migrate test")
}

func Test_MigrateAdminRole(t *testing.T) {
	err := repo.MigrateAdminRole()
	if err != nil {
		t.Fatal(err)
	}
}
