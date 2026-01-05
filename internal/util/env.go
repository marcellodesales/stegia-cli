package util

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type TotvsEnv struct {
	EnvFile   string
	Hostname  string
	Username  string
	Password  string
	BasicAuth string // computed "user:pass"
}

// LoadEnvFile loads the environment file based on ENV (local.env or prd.env)
// and returns its absolute path. It is safe to call multiple times.
func LoadEnvFile() string {
	envFile := envFileName()
	_ = godotenv.Load(envFile)
	abs, _ := filepath.Abs(envFile)
	return abs
}

// LogLevelFromEnv loads the .env file (if present) and returns LOG_LEVEL.
func LogLevelFromEnv() string {
	LoadEnvFile()
	return strings.TrimSpace(os.Getenv("LOG_LEVEL"))
}

func envFileName() string {
	envName := strings.TrimSpace(os.Getenv("ENV"))
	if envName == "" {
		envName = "local"
	}
	return envName + ".env"
}

func LoadTotvsEnv() TotvsEnv {
	envFile := LoadEnvFile()

	username := strings.TrimSpace(os.Getenv("TOTVS_USERNAME"))
	password := strings.TrimSpace(os.Getenv("TOTVS_PASSWORD"))
	hostname := strings.TrimSpace(os.Getenv("TOTVS_HOSTNAME"))
	if hostname == "" {
		hostname = "example.com"
	}

	if username == "" || password == "" {
		username = "admin"
		password = "admin"
	}

	return TotvsEnv{
		EnvFile:   envFile,
		Hostname:  hostname,
		Username:  username,
		Password:  password,
		BasicAuth: username + ":" + password,
	}
}
