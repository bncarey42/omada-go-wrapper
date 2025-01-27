package main

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

type CreateSiteEntity struct {
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
