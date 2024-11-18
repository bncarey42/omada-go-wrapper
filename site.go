package main

type SiteService struct {
	client *OmadaClient
}

type Site struct {
	SiteId    string   `json:"siteId"`
	Name      string   `json:"name"`
	TagIds    []string `json:"tagIds"`
	Region    string   ``
	Timezone  string
	Scenario  string
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
  sites := {Site{SiteId: "123455", Name: "noosphere", Region: "america/Chicago", Scenario: "Home",  ,}}  
  return 
}
