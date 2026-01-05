package suppliers

import (
	"fmt"
	"log/slog"
	"strings"

	"stegia/internal/totvs/client"
	"stegia/internal/util"
)

type CompaniesLister interface {
	ListForSelection() (client.CompaniesResponse, int, error)
}

type Controller struct {
	Service   *Service
	Builder   Builder
	Companies CompaniesLister
	Log       *slog.Logger
}

func (c *Controller) AddFromTOON(filePath string, explicitCompanyId string) error {
	doc, err := util.ParseTOONFile(filePath)
	if err != nil {
		c.Log.Error("parse TOON failed", "file", filePath, "error", err)
		return err
	}
	c.Log.Info("parsed TOON file", "file", filePath)

	companies, _, err := c.Companies.ListForSelection()
	if err != nil {
		c.Log.Error("list companies failed", "error", err)
		return err
	}

	companyId, reason := selectCompanyId(companies, explicitCompanyId, doc)
	if companyId == "" {
		c.Log.Error("no company could be selected (set --company-id)", "file", filePath)
		return fmt.Errorf("no company selected; provide --company-id")
	}
	c.Log.Info("selected company", "companyId", companyId, "reason", reason)

	payload := c.Builder.BuildPayloadFromTOON(doc)
	req, body, err := c.Builder.BuildCreateRequest(c.Service.Client.BaseURL(), c.Service.Client.BasicAuth, companyId, payload)
	if err != nil {
		c.Log.Error("build supplier create request failed", "error", err)
		return err
	}

	util.PrintHTTPRequest(req, body)

	resp, status, err := c.Service.Create(payload, companyId)
	if err != nil {
		c.Log.Error("supplier create failed", "companyId", companyId, "error", err)
		return err
	}

	respBytes := util.JSONPretty(resp)
	util.PrintHTTPResponse(status, map[string]string{"Content-Type": "application/json"}, respBytes)

	// Cache only in mock mode (example.com) OR always cache, your choice.
	if c.Service.Client.Hostname == "example.com" {
		supplierId := util.StrAny(resp["supplierId"])
		if supplierId == "" {
			c.Log.Error("mock response missing supplierId; cannot cache")
			return nil
		}

		cacheBase := "examples"
		cachePath := util.SupplierCachePath(cacheBase, supplierId)

		if err := util.WriteFileAtomic(cachePath, respBytes); err != nil {
			c.Log.Error("failed to cache supplier", "path", cachePath, "error", err)
			return nil // do not fail the command just because caching failed
		}

		c.Log.Info("cached supplier", "path", cachePath)
	}
	return nil
}

func selectCompanyId(companies client.CompaniesResponse, explicit string, toonDoc map[string]any) (string, string) {
	if strings.TrimSpace(explicit) != "" {
		for _, it := range companies.Items {
			if it.CompanyId == explicit && strings.EqualFold(it.Status, "ACTIVE") {
				return it.CompanyId, "explicit --company-id (validated ACTIVE)"
			}
		}
		return explicit, "explicit --company-id (not found/ACTIVE in list)"
	}

	addr, _ := toonDoc["address"].(map[string]any)
	city := strings.ToLower(strings.TrimSpace(util.StrAny(addr["city"])))
	state := strings.ToLower(strings.TrimSpace(util.StrAny(addr["state"])))

	for _, it := range companies.Items {
		if !strings.EqualFold(it.Status, "ACTIVE") {
			continue
		}
		if strings.ToLower(it.State) == state && strings.ToLower(it.City) == city {
			return it.CompanyId, "auto-match by TOON address city/state"
		}
	}

	for _, it := range companies.Items {
		if strings.EqualFold(it.Status, "ACTIVE") {
			return it.CompanyId, "fallback to first ACTIVE company"
		}
	}
	return "", "no ACTIVE companies"
}

func (c *Controller) ViewFromCache(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("missing --id")
	}

	cachePath := util.SupplierCachePath("examples", id)

	c.Log.Info("loading cached supplier", "id", id, "path", cachePath)

	b, err := c.Service.LoadCached(cachePath)
	if err != nil {
		c.Log.Error("failed to read cached supplier", "path", cachePath, "error", err)
		return err
	}

	fmt.Println("\n=== CACHED SUPPLIER ===")
	fmt.Println(string(b))
	return nil
}

