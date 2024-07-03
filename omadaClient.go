package main

import (
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
	ErrorCode int           `json:errorCode`
	Msg       string        `json:msg`
	Result    ApiInfoResult `json:result`
}

type ApiInfoResult struct {
	// TODO
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
	client := &OmadaClient{
		Host: host,
		Site: site,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	return client
}

func (c OmadaClient) Login(user string, password string) *ClientToken {
	return &ClientToken{}
}
