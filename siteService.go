package main

import (
	"encoding/json"
	"fmt"
)

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

type Date struct {
	Month  int `json:"month"`
	Serial int `json:"serial"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type DST struct {
	Enable    bool   `json:"enable"`
	Start     Date   `json:"start"`
	End       Date   `json:"end"`
	Status    bool   `json:"status"`
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
	Offset    int    `json:"offset"`
	NextStart int    `json:"nextStart"`
	NextEnd   int    `json:"nextEnd"`
	TimeZone  string `json:"timeZone"`
	LastStart int    `json:"lastStart"`
	LastEnd   int    `json:"lastEnd"`
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
	Dst        DST      `json:"dst"`
	Scenario   string   `json:"scenario"`
	Longitude  int      `json:"longitude"`
	Latitude   int      `json:"latitude"`
	Address    string   `json:"address"`
	SupportES  bool     `json:"supportES"`
	SupportL2  bool     `json:"supportL2"`
}

type DeviceAccountSetting struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type createSiteEntity struct {
	Name                 string               `json:"name"`
	Type                 *int                 `json:"type,omitempty"`
	Region               string               `json:"region"`
	TimeZone             string               `json:"timeZone"`
	Scenario             string               `json:"scenario"`
	TagIds               *[]string            `json:"tagIds,omitempty"`
	Longitude            *int                 `json:"longitude,omitempty"`
	Latitude             *int                 `json:"latitude,omitempty"`
	Address              *string              `json:"address,omitempty"`
	DeviceAccountSetting DeviceAccountSetting `json:"deviceAccountSetting"`
	SupportES            *bool                `json:"supportES,omitempty"`
	SupportL2            *bool                `json:"supportL2,omitempty"`
}

type SiteService struct {
	omadaClient *OmadaClient
	baseUrl     string
}

func NewSiteService(client *OmadaClient) SiteService {
	return SiteService{omadaClient: client, baseUrl: "sites"}
}

func (ss *SiteService) GetSiteList(page int32, pageSize int32) ([]SiteSumaryInfo, error) {
	url := ss.omadaClient.BuildApiURL(ss.baseUrl)
	query := map[string]string{"page": fmt.Sprint(page), "pageSize": fmt.Sprint(pageSize)}
	sites, err := HttpRequest[PaginatedApiData[SiteSumaryInfo]]("GET", url, query, nil, ss.omadaClient)
	if err != nil {
		return nil, fmt.Errorf("error getting Site Sumaries %s", err.Error())
	}
	return sites.Data, nil
}

func (ss *SiteService) GetSiteInfo(siteId string) (*SiteEntity, error) {
	url := ss.omadaClient.BuildApiURL(fmt.Sprintf("%s/%s", ss.baseUrl, siteId))
	site, err := HttpRequest[SiteEntity]("GET", url, nil, nil, ss.omadaClient)
	if err != nil {
		return nil, fmt.Errorf("error getting Site Sumaries %s", err.Error())
	}
	return site, nil
}

func (ss *SiteService) CreateNewSite(name string, region string, timeZone string, senario string, deviceUserName string, devicePassword string, extra ...any) error {
	siteEntity := createSiteEntity{Name: name, Region: region, TimeZone: timeZone, Scenario: senario, DeviceAccountSetting: DeviceAccountSetting{Username: deviceUserName, Password: devicePassword}}
	url := ss.omadaClient.BuildApiURL(ss.baseUrl)
	jsonBodyStr, err := json.Marshal(siteEntity)
	if err != nil {
		return fmt.Errorf("failed to marshal new site entity :: %s", err.Error())
	}
	body := []byte(jsonBodyStr)

	_, err = HttpRequest[struct{}]("POST", url, nil, body, ss.omadaClient)
	return err
}

func (ss *SiteService) DeleteSite(siteId string) error {
	url := ss.omadaClient.BuildApiURL(fmt.Sprintf("%s/%s", ss.baseUrl, siteId))
	_, err := HttpRequest[struct{}]("DELETE", url, nil, nil, ss.omadaClient)
	return err
}
