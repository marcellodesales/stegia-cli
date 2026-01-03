package factory

import (
	"log/slog"

	"stegia/internal/totvs/client"
	companies "stegia/internal/totvs/companies"
	suppliers "stegia/internal/totvs/suppliers"
)

type ServiceFactory struct {
	Log *slog.Logger
}

func (f ServiceFactory) CompaniesService(c *client.Client) *companies.Service {
	return &companies.Service{Client: c, Log: f.Log}
}

func (f ServiceFactory) SuppliersService(c *client.Client) *suppliers.Service {
	return &suppliers.Service{Client: c, Log: f.Log}
}
