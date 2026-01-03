package logger

import (
	"log/slog"
	"os"
	"strings"
)

func New() *slog.Logger {
	lvl := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL")))
	var slogLevel slog.Level
	switch lvl {
	case "debug":
		slogLevel = slog.LevelDebug
	case "error":
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slogLevel})
	return slog.New(h)
}
