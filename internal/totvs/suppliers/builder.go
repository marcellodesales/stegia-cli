package suppliers

import (
	"bytes"
	"net/http"
	"strings"

	"stegia/internal/util"
)

type Builder struct{}

func (b Builder) BuildCreateRequest(baseURL, basicAuth, companyId string, payload map[string]any) (*http.Request, []byte, error) {
	body := util.JSONPretty(payload)

	req, err := http.NewRequest(http.MethodPost, baseURL+"/api/cdp/v1/suppliers", bytes.NewReader(body))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Authorization", util.BasicAuthHeader(basicAuth))
	req.Header.Set("companyId", companyId)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req, body, nil
}

func (b Builder) BuildPayloadFromTOON(doc map[string]any) map[string]any {
	// The TOON document is the canonical payload.
	// Apply safe defaults only when missing.
	payload := doc

	// Ensure integration.sourceSystem defaults to "stegia" if present/missing
	integ, _ := payload["integration"].(map[string]any)
	if integ == nil {
		integ = map[string]any{}
		payload["integration"] = integ
	}
	if util.StrAny(integ["sourceSystem"]) == "" {
		integ["sourceSystem"] = "stegia"
	}

	// Default status/country if omitted
	if util.StrAny(payload["status"]) == "" {
		payload["status"] = "ACTIVE"
	}
	if util.StrAny(payload["country"]) == "" {
		payload["country"] = "BR"
	}

	return payload
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
