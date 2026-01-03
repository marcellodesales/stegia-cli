package client

type Company struct {
	CompanyId   string `json:"companyId"`
	CompanyCode string `json:"companyCode"`
	CompanyName string `json:"companyName"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	Status      string `json:"status"`
}

type CompaniesResponse struct {
	Items []Company `json:"items"`
	Count int       `json:"count"`
}
