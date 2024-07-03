package main

import (
	"net/http"
	"time"
)

type OmadaClient struct {
	Host       string
	username   string
	password   string
	HTTPClient *http.Client
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

func NewOmadaClient(host string, username string, password string) *OmadaClient {
	client := &OmadaClient{
		Host:     host,
		username: username,
		password: password,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	return client
}

func (c OmadaClient) Login() *ClientToken {

	return &ClientToken{}
}
