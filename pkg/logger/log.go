package logger

import (
	"context"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l     *log.Logger
	attrs map[string]any
	group string
}

func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	return &PrettyHandler{
		l:       log.New(out, "", 0),
		attrs:   make(map[string]any, 0),
		group:   "",
		Handler: slog.NewJSONHandler(out, &opts.HandlerOptions),
	}
}

func (h PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()
	switch r.Level {
	case slog.LevelInfo:
		level = color.GreenString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelDebug:
		level = color.HiBlueString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}
	h.l.Printf("[%s] %s", r.Time.String(), level)
	return nil
}
