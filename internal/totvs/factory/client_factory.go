package factory

import (
	"log/slog"
	"net/http"
	"time"

	"stegia/internal/totvs/client"
	"stegia/internal/util"
)

type ClientFactory struct {
	Log *slog.Logger
}

func (f ClientFactory) New(env util.TotvsEnv) *client.Client {
	f.Log.Debug("creating TOTVS client", "hostname", env.Hostname, "envFile", env.EnvFile)
	return &client.Client{
		Hostname:  env.Hostname,
		BasicAuth: env.BasicAuth,
		HTTP: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
