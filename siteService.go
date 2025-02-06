package main

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
	SiteID               *string              `json:"siteId,omitempty"`
	Name                 string               `json:"name"`
	Type                 int32                `json:"type"`
	Region               string               `json:"region"`
	TimeZone             string               `json:"timeZone"`
	Scenario             string               `json:"scenario"`
	NtpEnable            *bool                `json:"ntpEnable,omitempty"`
	NtpServers           *[]string            `json:"ntpServers,omitempty"`
	Dst                  *DST                 `json:"dst,omitempty"`
	Longitude            *int                 `json:"longitude,omitempty"`
	Latitude             *int                 `json:"latitude,omitempty"`
	Address              *string              `json:"address,omitempty"`
	DeviceAccountSetting DeviceAccountSetting `json:"deviceAccountSetting"`
	SupportES            *bool                `json:"supportES,omitempty"`
	SupportL2            *bool                `json:"supportL2,omitempty"`
	TagIds               *[]string            `json:"tagIds,omitempty"`
}

type DeviceAccountSetting struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
