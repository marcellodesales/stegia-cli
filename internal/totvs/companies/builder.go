package companies

import (
	"net/http"

	"stegia/internal/util"
)

type Builder struct{}

func (b Builder) BuildListRequest(baseURL, basicAuth string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+"/api/btb/v1/companies", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", util.BasicAuthHeader(basicAuth))
	req.Header.Set("Accept", "application/json")
	return req, nil
}
