package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type OmadaClient struct {
	baseUrl    string
	apiVersion string
	omadacID   string
	token      *authResponse
	httpClient *http.Client
}

type authResponse struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

type apiResponse[T any] struct {
	ErrorCode int32  `json:"errorCode"`
	Message   string `json:"msg"`
	Result    T      `json:"result"`
}

type PaginatedApiData[T any] struct {
	TotalRows   int64 `json:"totalRows"`
	CurrentPage int32 `json:"currentPage"`
	CurrentSize int32 `json:"currentSize"`
	Data        []T   `json:"data"`
}

func NewOmadaClient(baseUrl string, apiVersion string, omadacId string, httpClient *http.Client) *OmadaClient {
	return &OmadaClient{
		baseUrl:    baseUrl,
		apiVersion: apiVersion,
		omadacID:   omadacId,
		token:      nil,
		httpClient: httpClient,
	}
}

func (oc *OmadaClient) Login(apiClientID string, apiSecret string) error {
	url := oc.BuildApiURL("authorize/token")
	params := map[string]string{"grant_type": "client_credentials"}

	body := []byte(fmt.Sprintf(`{"omadacId": "%s", "client_id": "%s", "client_secret": "%s"}`, oc.omadacID, apiClientID, apiSecret))

	var err error

	oc.token, err = HttpRequest[authResponse]("POST", url, params, body, oc)
	if err != nil {
		return fmt.Errorf("Error Logging In :: %s", err.Error())
	}

	log.Println("LOGGED IN")
	return nil
}

func (oc *OmadaClient) isAuthenticated() bool {
	return oc.token != nil
}

func (oc *OmadaClient) BuildApiURL(slug string) string {
	appendToUrl := slug
	if oc.isAuthenticated() {
		appendToUrl = fmt.Sprintf("%s/%s/%s", oc.apiVersion, oc.omadacID, appendToUrl)
	}
	return fmt.Sprintf("https://%s/%s", oc.baseUrl, appendToUrl)
}

func HttpRequest[T any](method string, url string, params map[string]string, body []byte, client *OmadaClient) (*T, error) {
	var request *http.Request
	var err error

	switch method {
	case "POST":
		var bodyReader *bytes.Reader

		if body != nil {
			bodyReader = bytes.NewReader(body)
		} else {
			bodyReader = nil
		}

		request, err = http.NewRequest(method, url, bodyReader)

		request.Header.Add("content-type", "application/json")
	case "DELETE":
		request, err = http.NewRequest("DELETE", url, nil)
	case "GET":
		request, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to build request : %s : %s : %s", method, url, err.Error())
	}

	if client.isAuthenticated() {
		request.Header.Add("Authorization", fmt.Sprintf("AccessToken=%s", client.token.AccessToken))
	}

	if params != nil {
		// SET ENCODED QUERY PARAMS ON URL
		queryParams := request.URL.Query()

		for key, value := range params {
			queryParams.Set(key, value)
		}
		request.URL.RawQuery = queryParams.Encode()
	}

	log.Printf("Request URL: %s :: BODY: %s", request.URL.String(), string(body))

	// DO REQUEST
	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to submit request : %s : %s : %s", method, request.RequestURI, err.Error())
	}

	defer func() { _ = resp.Body.Close() }()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body : %s", err.Error())
	}

	var response *apiResponse[T]

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response : %s", err.Error())
	}

	if response.ErrorCode != 0 {
		return nil, fmt.Errorf("omada controller encountered an error : %d : %s : %s", response.ErrorCode, response.Message, request.URL.String())
	}

	return &response.Result, nil
}
