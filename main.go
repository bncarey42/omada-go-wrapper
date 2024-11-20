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

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr, Timeout: time.Minute}

	baseUrl := fmt.Sprintf("%s:%s/%s",
		os.Getenv("OMADA_HOST"),
		os.Getenv("OMADA_PORT"),
		os.Getenv("OMADA_BASE_ENDPOINT"))

	fmt.Println(baseUrl)

	client := NewOmadaClient(
		baseUrl,
		os.Getenv("OMADA_API_VERSION"),
		os.Getenv("OMADA_CLIENT_OMADACID"),
		httpClient)
	if err := client.Login(
		os.Getenv("OMADA_CLIENT_ID"),
		os.Getenv("OMADA_CLIENT_TOKEN")); err != nil {
		log.Fatalf("Error getting Token for client ::: %s", err.Error())
	}

	siteService := NewSiteService(client)
	sites, err := siteService.GetSiteList(1, 200)
	if err != nil {
		log.Fatalf("Error getting Site List ::: %s", err)
	}
	log.Printf("%d", len(sites))
	for _, site := range sites {
		fmt.Println(site.Name, site.Type, site.SiteID)
	}
	deviceService := NewDeviceService(client)
	devices, err := deviceService.GetSiteDeviceList("65b1fc94e6e82c26c5155f37", 1, 200)
	if err != nil {
		log.Fatalf("Error getting Device List ::: %s", err.Error())
	}

	log.Println("GOT DEVICES")

	log.Printf("%d", len(devices))

	for _, device := range devices {
		fmt.Println(device.Name, device.Mac, device.IP, device.Type)
	}
}
