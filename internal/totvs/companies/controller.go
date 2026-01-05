package companies

import (
    "context"
	"fmt"
	"log/slog"
	"strings"

	"stegia/internal/totvs/client"
	"stegia/internal/util"
)

type Controller struct {
	Service *Service
	Builder Builder
	Log     *slog.Logger
}

func (c *Controller) ListAndPrint() error {
	req, err := c.Builder.BuildListRequest(c.Service.Client.BaseURL(), c.Service.Client.BasicAuth)
	if err != nil {
		c.Log.Error("build companies request failed", "error", err)
		return err
	}

    if c.Log.Enabled(context.Background(), slog.LevelDebug) {
    	util.PrintHTTPRequest(req, nil)
    }

	res, status, err := c.Service.List()
	if err != nil {
		c.Log.Error("companies list failed", "error", err)
		return err
	}

	body := util.JSONPretty(res)

    if c.Log.Enabled(context.Background(), slog.LevelDebug) {
    	util.PrintHTTPResponse(status, map[string]string{"Content-Type": "application/json"}, body)
    }

	for _, it := range res.Items {
		if strings.EqualFold(it.Status, "ACTIVE") {
			fmt.Printf("- companyId=%s code=%s name=%s (%s/%s)\n",
				it.CompanyId, it.CompanyCode, it.CompanyName, it.City, it.State)
		}
	}
	return nil
}

// Used by suppliers flow
func (c *Controller) ListForSelection() (client.CompaniesResponse, int, error) {
	return c.Service.List()
}
