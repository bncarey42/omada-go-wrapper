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

	host_name := os.Getenv("OMADA_HOST")
	port := os.Getenv("OMADA_PORT")
	apiEndpoint := os.Getenv("OMADA_BASE_ENDPOINT")
	omadacid := os.Getenv("OMADA_CLIENT_OMADACID")
	sitename := os.Getenv("OMADA_SITENAME")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr, Timeout: time.Minute}

	omadaClient := NewOmadaClient(fmt.Sprintf("https://%s:%s", host_name, port), sitename, apiEndpoint, omadacid, httpClient)

	err := omadaClient.Login(os.Getenv("OMADA_CLIENT_ID"), os.Getenv("OMADA_CLIENT_TOKEN"))
	if err != nil {
		log.Fatalf("Error getting Topken for client", err.Error())
	}
}
