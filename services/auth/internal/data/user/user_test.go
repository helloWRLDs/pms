package userdata

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"pms.pkg/datastore/sqlite"
	"pms.pkg/logger"
)

const (
	dsn = "../../../data/users.db"
)

var (
	repo Repository
)

func TestMain(m *testing.M) {
	setupLogger()
	setupDB()
	code := m.Run()
	os.Exit(code)
}

func setupLogger() {
	conf := logger.Config{
		Dev:  true,
		Path: "",
	}
	logger.Init(conf)
}

func setupDB() {
	db, err := sqlite.Open(dsn)
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to db")
	}
	repo = *New(db)
}
