package suppliers

import (
	"log/slog"

	"stegia/internal/totvs/client"
)

type Service struct {
	Client *client.Client
	Log    *slog.Logger
}

func (s *Service) Create(payload map[string]any, companyId string) (map[string]any, int, error) {
	if s.Client.Hostname == "example.com" {
		s.Log.Debug("mocking supplier create", "hostname", s.Client.Hostname, "companyId", companyId)
		return s.Client.CreateSupplierMock(payload, companyId)
	}

	s.Log.Error("real HTTP calls disabled in this prototype", "hostname", s.Client.Hostname)
	return nil, 0, nil
}
