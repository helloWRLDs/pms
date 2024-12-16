package ctx

import (
	"testing"

	"pms.pkg/logger"
)

func TestMain(m *testing.M) {
	cfg := logger.Config{
		Dev: true,
	}
	cfg.Init()
	m.Run()
}
