package github

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"pms.pkg/logger"
)

var (
	log *zap.SugaredLogger

	client = Client{
		conf: Config{
			ClientID:     "",
			ClientSecret: "",
			RedirectURL:  "",
			HOST:         "https://api.github.com",
			Scopes:       []string{"user"},
		},
		accessToken: "token",
	}
)

func TestMain(m *testing.M) {
	checkClient()
	setupLogger()
	code := m.Run()

	os.Exit(code)
}

func setupLogger() {
	logger.WithConfig(
		logger.WithLevel("debug"),
		logger.WithDev(true),
	).Init()
	log = logger.Log
}

func checkClient() {
	if client.accessToken == "" {
		panic("access_token required")
	}
}
