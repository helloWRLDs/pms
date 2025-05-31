package logger

import (
	"context"

	"go.uber.org/zap"
	ctxutils "pms.pkg/utils/ctx"
)

const (
	LoggerKey ctxutils.ContextKey = "logger"
)

func FromContext(ctx context.Context) *zap.SugaredLogger {
	log, ok := ctx.Value(LoggerKey).(*zap.SugaredLogger)
	if !ok {
		return zap.NewNop().Sugar()
	}
	return log
}

func ToContext(ctx context.Context, log *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, LoggerKey, log)
}
