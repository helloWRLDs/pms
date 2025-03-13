package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Can increase log level only by this hierarchy
//
// Debug > Info > Warn > Error > DPanic > Panic > Fatal
//
// If level set to ERROR, cannot increase it to INFO
func IncreaseLevel(log *zap.SugaredLogger, level *zapcore.Level) *zap.SugaredLogger {
	if level != nil && *level < log.Level() {
		return log.WithOptions(zap.IncreaseLevel(level))
	}
	return log
}
