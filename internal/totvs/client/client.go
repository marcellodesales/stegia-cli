package client

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Hostname  string
	BasicAuth string
	HTTP      *http.Client
}

func (c *Client) BaseURL() string {
	return fmt.Sprintf("https://%s", c.Hostname)
}

// In this prototype, example.com is mocked; real calls can be enabled later.
func (c *Client) GetCompaniesMock() (CompaniesResponse, int, error) {
	resp := CompaniesResponse{
		Items: []Company{
			{CompanyId: "01", CompanyCode: "TOTVS-BR", CompanyName: "TOTVS BRASIL MATRIZ", Country: "BR", State: "SP", City: "São Paulo", Status: "ACTIVE"},
			{CompanyId: "02", CompanyCode: "TOTVS-GO", CompanyName: "TOTVS GOIÁS", Country: "BR", State: "GO", City: "Goiânia", Status: "ACTIVE"},
			{CompanyId: "03", CompanyCode: "TOTVS-RJ", CompanyName: "TOTVS RIO DE JANEIRO", Country: "BR", State: "RJ", City: "Rio de Janeiro", Status: "INACTIVE"},
		},
		Count: 3,
	}
	_ = time.Now()
	return resp, 200, nil
}

func (c *Client) CreateSupplierMock(payload map[string]any, companyId string) (map[string]any, int, error) {
	resp := map[string]any{
		"supplierId":   "SUP-902341",
		"supplierCode": "FORN-000902341",
		"companyId":    companyId,
		"status":       "CREATED",
		"createdAt":    time.Now().UTC().Format(time.RFC3339),
		"links": map[string]any{
			"self": "/api/cdp/v1/suppliers/SUP-902341",
		},
		"echoRequest": payload,
		"message":     "Mocked response (example.com); no real Datasul call executed.",
	}
	return resp, 201, nil
}
