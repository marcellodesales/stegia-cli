package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"stegia/internal/util"
)

type Level string

const (
	LevelNone  Level = "none"
	LevelInfo  Level = "info"
	LevelDebug Level = "debug"
	LevelError Level = "error"
)

var levelOverride string

// SetLevelOverride allows CLI flags to override LOG_LEVEL from .env.
func SetLevelOverride(level string) {
	levelOverride = strings.TrimSpace(level)
}

func resolveLevel() Level {
	envLevel := strings.ToLower(util.LogLevelFromEnv())
	effective := envLevel

	if strings.TrimSpace(levelOverride) != "" {
		effective = strings.ToLower(levelOverride)
	}
	if strings.TrimSpace(effective) == "" {
		effective = string(LevelNone)
	}
	return Level(effective)
}

// New builds a slog.Logger using LOG_LEVEL from the loaded .env file,
// optionally overridden by SetLevelOverride (CLI flag).
func New() *slog.Logger {
	level := resolveLevel()

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
