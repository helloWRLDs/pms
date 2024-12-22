package ctx

import (
	"testing"

	"pms.pkg/logger"
)

func TestMain(m *testing.M) {
	logger.Init(logger.Config{
		Dev: true,
	})
	m.Run()
}
