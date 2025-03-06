package transaction

import (
	"os"
	"testing"

	"pms.pkg/logger"
)

func TestMain(m *testing.M) {
	setupLogger()
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
