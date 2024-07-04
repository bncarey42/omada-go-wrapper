package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type OmadaClient struct {
	Host       string
	Site       string
	OmadacId   string
	HTTPClient *http.Client
}

type ApiInfo struct {
	ErrorCode int    `json:"errorCode"`
	Msg       string `json:"msg"`
	Result    struct {
		ControllerVer  string `json:"controllerVer"`
		APIVer         string `json:"apiVer"`
		Configured     bool   `json:"configured"`
		Type           int    `json:"type"`
		SupportApp     bool   `json:"supportApp"`
		OmadacID       string `json:"omadacId"`
		RegisteredRoot bool   `json:"registeredRoot"`
		OmadacCategory string `json:"omadacCategory"`
		MspMode        bool   `json:"mspMode"`
	} `json:"result"`
}

type ClientToken struct {
	OmadacId string `json:"omadac_id"`
	Token    string `json:"token"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func NewOmadaClient(host string, site string) *OmadaClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &OmadaClient{
		Host: host,
		Site: site,
		HTTPClient: &http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}
	return client
}

func (c OmadaClient) ApiInfo() {
	ret, err := c.HTTPClient.Get(fmt.Sprintf("%s/info", c.Host))
	if err != nil {
		panic(err)
	}
	defer func() { _ = ret.Body.Close() }()
	body, _ := io.ReadAll(ret.Body)
	var apiInfo ApiInfo
	err = json.Unmarshal(body, &apiInfo)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	c.OmadacId = apiInfo.Result.OmadacID
}

func (c OmadaClient) Login(user string, password string) *ClientToken {
	return &ClientToken{}
}
