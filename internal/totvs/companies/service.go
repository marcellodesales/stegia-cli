package companies

import (
	"log/slog"

	"stegia/internal/totvs/client"
)

type Service struct {
	Client *client.Client
	Log    *slog.Logger
}

func (s *Service) List() (client.CompaniesResponse, int, error) {
	if s.Client.Hostname == "example.com" {
		s.Log.Debug("mocking companies list", "hostname", s.Client.Hostname)
		return s.Client.GetCompaniesMock()
	}

	s.Log.Error("real HTTP calls disabled in this prototype", "hostname", s.Client.Hostname)
	return client.CompaniesResponse{}, 0, nil
}
