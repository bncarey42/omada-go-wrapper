package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

const (
	host     = "https://noosphere-ctrl:8043/api"
	user     = "omada-tf-test"
	password = "V8qMh%k%J*xq&87aco"
)

func main() {
	//	client := &http.Client{}

	// SET UP

	//	infoURL := fmt.Sprintf("%s/%s", host, "/info")
	//	fmt.Printf("ATTEMPTING TO GET INFO \n")
	//	r, err := http.Get(infoURL)
	//	defer func() { _ = r.Body.Close() }()
	//	body, _ := io.ReadAll(r.Body)
	//	fmt.Printf("%s", body)

	// CAL API

	//	requestURL := fmt.Sprintf("%s/v2/%s", host, "login")
	//	fmt.Printf("ATTEMPTING TO LOG IN\n")
	//	jsonBody := []byte(fmt.Sprintf(`{"username": %s, "password": %s}`, user, password))
	//	bodyReader := bytes.NewReader(jsonBody)
	//
	//	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	//	if err != nil {
	//		fmt.Printf(fmt.Sprintf("ERROR 1::: %s", err.Error()))
	//		os.Exit(1)
	//	}

	// req.Header.Add("Content-Type", "application/json")

	//	resp, err := client.Do(req)
	//	if err != nil {
	//		fmt.Printf(fmt.Sprintf("ERROR 2::: %s", err.Error()))
	//		os.Exit(1)
	//	}
	//	defer func() { _ = resp.Body.Close() }()
	//	ret_body, _ := io.ReadAll(resp.Body)
	//	fmt.Printf("%s", ret_body)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	r, err := client.Get(fmt.Sprintf("%s/info", host))
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	body, _ := io.ReadAll(r.Body)
	fmt.Printf("%s", body)
}
