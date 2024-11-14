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
	Site       string
	BaseApiUrl string
	OmadacID   string
	Token      *authResult
	HTTPClient *http.Client
}

type HTTPMethod struct{ method string }

func (h HTTPMethod) String() string {
	return h.method
}

var (
	GET    = HTTPMethod{"GET"}
	POST   = HTTPMethod{"POST"}
	PUT    = HTTPMethod{"PUT"}
	PATCH  = HTTPMethod{"PATCH"}
	DELETE = HTTPMethod{"DELETE"}
)

type apiInfo struct {
	ErrorCode int           `json:"errorCode"`
	Message   string        `json:"msg"`
	Result    apiInfoResult `json:"result"`
}

type apiInfoResult struct {
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
	ErrorCode int        `json:"errorCode"`
	Message   string     `json:"msg"`
	Result    authResult `json:"result"`
}
type authResult struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}

func NewOmadaClient(host string, site string, apiUrl string, omadacId string, httpClient *http.Client) *OmadaClient {
	return &OmadaClient{
		Host:       host,
		Site:       site,
		BaseApiUrl: apiUrl,
		OmadacID:   omadacId,
		Token:      nil,
		HTTPClient: httpClient,
	}
}

func (c OmadaClient) buildURL(slug string, params map[string]string) string {
	var paramarr []string
	for key, value := range params {
		paramarr = append(paramarr, fmt.Sprintf("%s=%s", key, value))
	}
	return fmt.Sprintf("%s/%s/%s?%s", c.Host, c.BaseApiUrl, slug, strings.Join(paramarr, "&"))
}

func (c OmadaClient) Login(apiClientId string, apiToken string) error {
	var token authResponse

	url := c.buildURL("authorize/token", map[string]string{"grant_type": "client_credentials"})

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

func (c OmadaClient) Request(method string, urlSlug string, params map[string]string) (*any, error) {
	url := c.buildURL(urlSlug, params)
	log.Printf("%s :: %s", method, url)
	var bodyReader io.Reader

	request, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		log.Fatalf("ERROR :: Making request :: %s", err.Error())
	}

	request.Header.Add("access", fmt.Sprintf("AccessToken=%s", c.Token.AccessToken))
	switch method {
	case POST.method:
	case PATCH.method:
	case DELETE.method:
	case GET.method:
	default:
	}

	request, err := http.NewRequest(httpMethod, url, body)
}
