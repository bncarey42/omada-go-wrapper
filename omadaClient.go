package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type OmadaClient struct {
	Host       string
	BaseApiUrl string
	ApiVersion string
	OmadacID   string
	Token      *authResponse
	HTTPClient *http.Client
}

type ApiResponse[T any] struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"msg"`
	Result    T      `json:"result"`
}

type apiInfoResponse struct {
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

type authResponse struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}

func NewOmadaClient(host string, apiUrl string, omadacId string, apiVersion string, httpClient *http.Client) *OmadaClient {
	return &OmadaClient{
		Host:       host,
		BaseApiUrl: apiUrl,
		ApiVersion: apiVersion,
		OmadacID:   omadacId,
		Token:      nil,
		HTTPClient: httpClient,
	}
}

func (c OmadaClient) buildURL(slug string, params map[string]string, useVersion bool) string {
	var paramarr []string
	for key, value := range params {
		paramarr = append(paramarr, fmt.Sprintf("%s=%s", key, value))
	}
	if useVersion {
		return fmt.Sprintf("%s/%s/%s/%s?%s", c.Host, c.BaseApiUrl, c.ApiVersion, slug, strings.Join(paramarr, "&"))
	}
	return fmt.Sprintf("%s/%s/%s?%s", c.Host, c.BaseApiUrl, slug, strings.Join(paramarr, "&"))
}

func (c OmadaClient) Login(apiClientId string, apiToken string) error {
	var token ApiResponse[authResponse]

	url := c.buildURL("authorize/token", map[string]string{"grant_type": "client_credentials"}, false)

	log.Printf("login url::: %s", url)

	jsonBody := []byte(fmt.Sprintf(`{"omadacId": "%s", "client_id": "%s", "client_secret": "%s"}`, c.OmadacID, apiClientId, apiToken))
	log.Println(string(jsonBody))
	bodyReader := bytes.NewReader(jsonBody)

	ret, err := c.HTTPClient.Post(url, "application/json", bodyReader)
	if err != nil {
		log.Fatalf("Unable to get access token due to %s", err.Error())
		return err
	}
	defer func() { _ = ret.Body.Close() }()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		log.Fatalf("Error reading Login body::: %s", err.Error())
		return err
	}
	err = json.Unmarshal(body, &token)
	if err != nil {
		log.Fatalf("error Unmarshalling::: %s", err.Error())
		return err
	}
	log.Printf("ERROR_CODE: %d, MESSAGE: %s", token.ErrorCode, token.Message)

	c.Token = &token.Result

	return nil
}

func (c OmadaClient) NewLogin(apiClientId string, apiToken string) error {
	url := c.buildURL("authorize/token", map[string]string{"grant_type": "client_credentials"}, false)

	log.Printf("login url::: %s", url)

	jsonBody := []byte(fmt.Sprintf(`{"omadacId": "%s", "client_id": "%s", "client_secret": "%s"}`, c.OmadacID, apiClientId, apiToken))
	log.Println(string(jsonBody))

	token, err := HttpRequest[authResponse]("POST", url, jsonBody, false, c)
	if err != nil {
		log.Fatalf("ERROR ::: Getting Auth Token :: %s", err.Error())
	}

	c.Token = token

	return nil
}

func HttpRequest[T any](method string, url string, jsonBody []byte, doAuth bool, client OmadaClient) (*T, error) {
	var response *ApiResponse[T]
	var request *http.Request

	switch method {
	case "POST":
	case "PUT":
	case "PATCH":
		bodyReader := bytes.NewReader(jsonBody)

		request, err := http.NewRequest(method, url, bodyReader)
		if err != nil {
			log.Fatalf("ERROR::Unable to build request::%s:%s:%s", method, url, err.Error())
			return nil, err
		}
		request.Header.Add("Content-Type", "application/json")
		break

	case "GET":
	default:
		break
	}

	if doAuth {
		request.Header.Add("access", fmt.Sprintf("Bearer:%s", client.Token))
	}

	ret, err := client.HTTPClient.Do(request)
	if err != nil {
		log.Fatalf("ERROR::Unable to do request::%s:%s:%s", method, url, err.Error())
		return nil, err
	}

	defer func() { _ = ret.Body.Close() }()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		log.Fatalf("Error reading response body::: %s", err.Error())
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("error Unmarshalling::: %s", err.Error())
		return nil, err
	}

	return &response.Result, nil
}
