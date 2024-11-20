package main

import "fmt"

type SiteSumaryInfo struct {
	SiteID    string   `json:"siteId"`
	Name      string   `json:"name"`
	TagIds    []string `json:"tagIds"`
	Region    string   `json:"region"`
	TimeZone  string   `json:"timeZone"`
	Scenario  string   `json:"scenario"`
	Longitude int      `json:"longitude"`
	Latitude  int      `json:"latitude"`
	Address   string   `json:"address"`
	Type      int      `json:"type"`
	SupportES bool     `json:"supportES"`
	SupportL2 bool     `json:"supportL2"`
}

type SiteEntity struct {
	SiteID     string   `json:"siteId"`
	Name       string   `json:"name"`
	Type       int32    `json:"type"`
	TagIds     []string `json:"tagIds"`
	Region     string   `json:"region"`
	TimeZone   string   `json:"timeZone"`
	NtpEnable  bool     `json:"ntpEnable"`
	NtpServers []string `json:"ntpServers"`
	Dst        struct {
		Enable bool `json:"enable"`
		Start  struct {
			Month  int `json:"month"`
			Serial int `json:"serial"`
			Day    int `json:"day"`
			Hour   int `json:"hour"`
			Minute int `json:"minute"`
		} `json:"start"`
		End struct {
			Month  int `json:"month"`
			Serial int `json:"serial"`
			Day    int `json:"day"`
			Hour   int `json:"hour"`
			Minute int `json:"minute"`
		} `json:"end"`
		Status    bool   `json:"status"`
		StartTime int    `json:"startTime"`
		EndTime   int    `json:"endTime"`
		Offset    int    `json:"offset"`
		NextStart int    `json:"nextStart"`
		NextEnd   int    `json:"nextEnd"`
		TimeZone  string `json:"timeZone"`
		LastStart int    `json:"lastStart"`
		LastEnd   int    `json:"lastEnd"`
	} `json:"dst"`
	Scenario  string `json:"scenario"`
	Longitude int    `json:"longitude"`
	Latitude  int    `json:"latitude"`
	Address   string `json:"address"`
	SupportES bool   `json:"supportES"`
	SupportL2 bool   `json:"supportL2"`
}
type CreateSiteEntity struct {
	Name                 string   `json:"name"`
	Type                 int      `json:"type"`
	Region               string   `json:"region"`
	TimeZone             string   `json:"timeZone"`
	Scenario             string   `json:"scenario"`
	TagIds               []string `json:"tagIds"`
	Longitude            int      `json:"longitude"`
	Latitude             int      `json:"latitude"`
	Address              string   `json:"address"`
	DeviceAccountSetting struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"deviceAccountSetting"`
	SupportES bool `json:"supportES"`
	SupportL2 bool `json:"supportL2"`
}

type SiteService struct {
	omadaClient *OmadaClient
}

func NewSiteService(client *OmadaClient) SiteService {
	return SiteService{omadaClient: client}
}

func (ss *SiteService) GetSiteList(page int32, pageSize int32) ([]SiteSumaryInfo, error) {
	url := ss.omadaClient.BuildApiURL("sites")
	query := map[string]string{"page": fmt.Sprint(page), "pageSize": fmt.Sprint(pageSize)}
	sites, err := HttpRequest[PaginatedApiData[SiteSumaryInfo]]("GET", url, query, nil, ss.omadaClient)
	if err != nil {
		return nil, fmt.Errorf("error getting Site Sumaries %s", err.Error())
	}
	return sites.Data, nil
}

func (ss *SiteService) Create
