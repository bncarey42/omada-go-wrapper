package main

import (
	"fmt"
)

type DeviceInfo struct {
	Mac              string   `json:"mac"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Subtype          string   `json:"subtype"`
	DeviceSeriesType int      `json:"deviceSeriesType"`
	Model            string   `json:"model"`
	IP               string   `json:"ip"`
	Ipv6             []string `json:"ipv6"`
	Uptime           string   `json:"uptime"`
	Status           int      `json:"status"`
	LastSeen         int      `json:"lastSeen"`
	CPUUtil          int      `json:"cpuUtil"`
	MemUtil          int      `json:"memUtil"`
	SerialNumber     string   `json:"sn"`
	LicenseStatus    int      `json:"licenseStatus"`
	TagName          string   `json:"tagName"`
	UplinkDeviceMac  string   `json:"uplinkDeviceMac"`
	UplinkDeviceName string   `json:"uplinkDeviceName"`
	UplinkDevicePort string   `json:"uplinkDevicePort"`
	LinkSpeed        int      `json:"linkSpeed"`
	Duplex           int      `json:"duplex"`
	SwitchConsistent bool     `json:"switchConsistent"`
	PublicIP         string   `json:"publicIp"`
	FirmwareVersion  string   `json:"firmwareVersion"`
}

type DeviceService struct {
	omadaClient *OmadaClient
	baseUrl     string
}

func NewDeviceService(client *OmadaClient) *DeviceService {
	return &DeviceService{omadaClient: client, baseUrl: "sites"}
}

func (ds *DeviceService) GetSiteDeviceList(siteId string, page int32, pageSize int32) ([]DeviceInfo, error) {
	url := ds.omadaClient.BuildApiURL(fmt.Sprintf("%s/%s/devices", ds.baseUrl, siteId))
	devices, err := HttpRequest[PaginatedApiData[DeviceInfo]](
		"GET",
		url,
		map[string]string{"page": fmt.Sprint(page), "pageSize": fmt.Sprint(pageSize)},
		nil,
		ds.omadaClient)
	if err != nil {
		return nil, fmt.Errorf("error getting SiteDeviceList :: %s", err.Error())
	}

	return devices.Data, nil
}
