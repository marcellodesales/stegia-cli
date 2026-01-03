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
	getStr := func(k string) string {
		v, ok := doc[k]
		if !ok {
			return ""
		}
		s, _ := v.(string)
		return strings.TrimSpace(s)
	}

	addr := map[string]any{}
	if a, ok := doc["address"].(map[string]any); ok {
		addr = a
	}

	supplierName := firstNonEmpty(getStr("supplierName"), getStr("name"), "COCA-COLA INDUSTRIAS LTDA")
	tradeName := firstNonEmpty(getStr("tradeName"), "COCA-COLA BRASIL")
	cnpj := firstNonEmpty(getStr("cnpj"), "00000000000000")

	return map[string]any{
		"supplierType": "JURIDICAL",
		"supplierName": supplierName,
		"tradeName":    tradeName,
		"taxId": map[string]any{
			"cnpj": cnpj,
		},
		"status":  "ACTIVE",
		"country": "BR",
		"address": map[string]any{
			"city":       firstNonEmpty(util.StrAny(addr["city"]), "Goiânia"),
			"state":      firstNonEmpty(util.StrAny(addr["state"]), "GO"),
			"street":     firstNonEmpty(util.StrAny(addr["street"]), "Av. Anhanguera"),
			"number":     firstNonEmpty(util.StrAny(addr["number"]), "5000"),
			"district":   firstNonEmpty(util.StrAny(addr["district"]), "Setor Central"),
			"zipCode":    firstNonEmpty(util.StrAny(addr["zipCode"]), util.StrAny(addr["zip"]), "74043010"),
			"complement": util.StrAny(addr["complement"]),
		},
		"integration": map[string]any{
			"externalId":   firstNonEmpty(getStr("externalId"), "toon:coca-cola-br-go"),
			"sourceSystem": "stegia",
		},
		"source": doc,
	}
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
