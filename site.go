package main

type SiteService struct {
	client *OmadaClient
}

type Site struct {
	SiteId    string   `json:"siteId"`
	Name      string   `json:"name"`
	TagIds    []string `json:"tagIds"`
	Region    string   ``
	Timezone  Timezone
	Scenario  Scenario
	Longitude float32
	Latitude  float64
	Address   string
	Type      int
}

func NewSiteService(omadaClient OmadaClient) *SiteService {
	return &SiteService{
		client: &omadaClient,
	}
}

func (c SiteSiteService) ListSites() []Site {
}
