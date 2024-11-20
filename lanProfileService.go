package main

import (
	"encoding/json"
	"fmt"
)

type StormCtrl struct {
	UnknownUnicastEnable bool `json:"unknownUnicastEnable,omitempty"`
	UnknownUnicast       *int `json:"unknownUnicast,omitempty"`
	MulticastEnable      bool `json:"multicastEnable"`
	Multicast            *int `json:"multicast,omitempty"`
	BroadcastEnable      bool `json:"broadcastEnable"`
	Broadcast            *int `json:"broadcast,omitempty"`
	Action               int  `json:"action"`
	RecoverTime          *int `json:"recoverTime,omitempty"`
}

type BandCtrl struct {
	EgressEnable  bool `json:"egressEnable"`
	EgressLimit   *int `json:"egressLimit,omitempty"`
	EgressUnit    *int `json:"egressUnit,omitempty"`
	IngressEnable bool `json:"ingressEnable"`
	IngressLimit  *int `json:"ingressLimit,omitempty"`
	IngressUnit   *int `json:"ingressUnit,omitempty"`
}

type DhcpL2RelaySettings struct {
	Enable *bool `json:"enable,omitempty"`
	Format *int  `json:"format,omitempty"`
}

type LanProfileConfig struct {
	Name                          string               `json:"name"`
	Poe                           int32                `json:"poe"`
	NativeNetworkID               string               `json:"nativeNetworkId"`
	TagNetworkIds                 *[]string            `json:"tagNetworkIds,omitempty"`
	UntagNetworkIds               *[]string            `json:"untagNetworkIds,omitempty"`
	VoiceNetworkID                *string              `json:"voiceNetworkId,omitempty"`
	Dot1X                         int32                `json:"dot1x"`
	PortIsolationEnable           bool                 `json:"portIsolationEnable"`
	LldpMedEnable                 bool                 `json:"lldpMedEnable"`
	BandWidthCtrlType             int32                `json:"bandWidthCtrlType"`
	StormCtrl                     *StormCtrl           `json:"stormCtrl,omitempty"`
	BandCtrl                      *BandCtrl            `json:"bandCtrl,omitempty"`
	SpanningTreeEnable            bool                 `json:"spanningTreeEnable"`
	LoopbackDetectEnable          bool                 `json:"loopbackDetectEnable"`
	EeeEnable                     *bool                `json:"eeeEnable,omitempty"`
	FlowControlEnable             *bool                `json:"flowControlEnable,omitempty"`
	LoopbackDetectVlanBasedEnable *bool                `json:"loopbackDetectVlanBasedEnable,omitempty"`
	DhcpL2RelaySettings           *DhcpL2RelaySettings `json:"dhcpL2RelaySettings,omitempty"`
}

type LANProfileService struct {
	omadaClient *OmadaClient
	baseUrl     string
}

func NewLanProfileService(client *OmadaClient, siteID string) *LANProfileService {
	return &LANProfileService{omadaClient: client, baseUrl: fmt.Sprintf("sites/%s/lan-profiles", siteID)}
}

func (lps LANProfileService) CreateNewLANProfile(name string, poe int32, nativeNetworkId string, dot1x int32, portIsolationEnable bool, lldpMedEnable bool, bandwidthCtrlType int32, spanningTreeEnable bool, loopbackDetectEnable bool) (*string, error) {
	lanProfileConfig := LanProfileConfig{
		Name:                 name,
		Poe:                  poe,
		NativeNetworkID:      nativeNetworkId,
		Dot1X:                dot1x,
		PortIsolationEnable:  portIsolationEnable,
		LldpMedEnable:        lldpMedEnable,
		BandWidthCtrlType:    bandwidthCtrlType,
		SpanningTreeEnable:   spanningTreeEnable,
		LoopbackDetectEnable: loopbackDetectEnable,
	}

	jsonBodyStr, err := json.Marshal(lanProfileConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal new lan-profile entity :: %s", err.Error())
	}
	body := []byte(jsonBodyStr)

	url := lps.omadaClient.BuildApiURL(lps.baseUrl)

	response, err := HttpRequest[struct {
		Id string `json:"id"`
	}]("GET", url, nil, body, lps.omadaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new LAN Profile :: %s", err.Error())
	}
	return &response.Id, nil
}
