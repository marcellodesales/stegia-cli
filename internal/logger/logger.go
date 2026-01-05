package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

type Level string

const (
	LevelNone  Level = "none"
	LevelInfo  Level = "info"
	LevelDebug Level = "debug"
	LevelError Level = "error"
)

func New(level Level) *slog.Logger {
	level = Level(strings.ToLower(string(level)))

	// Default: NO logging
	if level == LevelNone || level == "" {
		return slog.New(
			slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}),
		)
	}

	var slogLevel slog.Level
	switch level {
	case LevelDebug:
		slogLevel = slog.LevelDebug
	case LevelError:
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLevel,
	})
	return slog.New(handler)
}

