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

func LoadTotvsEnv() TotvsEnv {
	envName := strings.TrimSpace(os.Getenv("ENV"))
	envFile := "local.env"
	if envName == "prd" {
		envFile = "prd.env"
	}

	_ = godotenv.Load(envFile)

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

	abs, _ := filepath.Abs(envFile)

	return TotvsEnv{
		EnvFile:   abs,
		Hostname:  hostname,
		Username:  username,
		Password:  password,
		BasicAuth: username + ":" + password,
	}
}
