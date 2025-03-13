package logger

import (
	"testing"

	"go.uber.org/zap"
)

func Test_CompareLogLevels(t *testing.T) {
	l1 := zap.ErrorLevel
	l2 := zap.InfoLevel
	t.Log(l2 < l1)
}
