package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	hostName := os.Getenv("OMADA_HOST")
	port := os.Getenv("OMADA_PORT")
	apiEndpoint := os.Getenv("OMADA_BASE_ENDPOINT")
	apiVersion := os.Getenv("OMADA_API_VERSION")
	omadacid := os.Getenv("OMADA_CLIENT_OMADACID")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr, Timeout: time.Minute}

	omadaClient := NewOmadaClient(
		fmt.Sprintf("https://%s:%s", hostName, port),
		apiEndpoint,
		omadacid,
		apiVersion,
		httpClient)

	//	err := omadaClient.Login(os.Getenv("OMADA_CLIENT_ID"), os.Getenv("OMADA_CLIENT_TOKEN"))
	err := omadaClient.NewLogin(os.Getenv("OMADA_CLIENT_ID"), os.Getenv("OMADA_CLIENT_TOKEN"))
	if err != nil {
		log.Fatalf("Error getting Topken for client", err.Error())
	}
}
