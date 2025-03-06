package data

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"pms.pkg/datastore/sqlite"
	"pms.pkg/logger"
)

const (
	dsn = "../../data/users.db"
)

var (
	repo *Repository
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
}

func setupDB() {
	db, err := sqlite.Open(dsn)
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to db")
	}
	repo = New(db, logger.Log)
}

func Test_Test(t *testing.T) {
	t.Log("migrate test")
}
