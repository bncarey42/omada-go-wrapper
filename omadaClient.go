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
	ApiInfo    *ApinInfoResult
	HTTPClient *http.Client
}

type ApiInfo struct {
	ErrorCode int            `json:"errorCode"`
	Msg       string         `json:"msg"`
	Result    ApinInfoResult `json:"result"`
}

type ApinInfoResult struct {
	ControllerVer  string `json:"controllerVer"`
	APIVer         string `json:"apiVer"`
	Configured     bool   `json:"configured"`
	Type           int    `json:"type"`
	SupportApp     bool   `json:"supportApp"`
	OmadacID       string `json:"omadacId"`
	RegisteredRoot bool   `json:"registeredRoot"`
	OmadacCategory string `json:"omadacCategory"`
	MspMode        bool   `json:"mspMode"`
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

func (c OmadaClient) BuildURL(slug string) (string, error) {
	if c.ApiInfo == nil {
		apiInfo, err := c.GetApiInfo()
		if err != nil {
			return "", err
		}
		c.ApiInfo = &apiInfo.Result
	}

	return fmt.Sprintf("%s/v%d/%s/%s", c.Host, c.ApiInfo.APIVer, c.ApiInfo.OmadacID, slug), nil
}

func (c OmadaClient) GetApiInfo() (*ApiInfo, error) {
	var apiInfo ApiInfo
	ret, err := c.HTTPClient.Get(fmt.Sprintf("%s/info", c.Host))
	if err != nil {
		log.Fatalf("Unable to get apiInfo due to %s", err.Error())
		return nil, err
	}
	defer func() { _ = ret.Body.Close() }()
	body, _ := io.ReadAll(ret.Body)
	err = json.Unmarshal(body, &apiInfo)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
		return nil, err
	}
	return &apiInfo, nil
}

func (c OmadaClient) Login(user string, password string) (*ClientToken, error) {
	url, err := c.BuildURL("login")
	if err != nil {
		return nil, err
	}

	log.Printf("login url::: %s", url)
	return nil, nil
}
