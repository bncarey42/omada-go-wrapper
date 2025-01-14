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

	fmt.Println()

	siteService := NewSiteService(client)
	sites, err := siteService.GetSiteList(1, 200)
	if err != nil {
		log.Fatalf("Error getting Site List ::: %s", err)
	}

	for _, site := range sites {
		fmt.Println("\t", site.Name, site.SiteID)
	}

	fmt.Println()

	sites, err = siteService.GetSiteList(1, 200)
	if err != nil {
		log.Fatalf("Error getting Site List ::: %s", err)
	}

	for _, site := range sites {
		fmt.Println("\t", site.Name, site.SiteID)
	}

	fmt.Println()

	err = siteService.CreateNewSite("test", "United States", "America/Chicago", "Home", "", "")
	if err != nil {
		log.Fatalf("err created new site :: %s", err.Error())
	}

	fmt.Println()

	sites, err = siteService.GetSiteList(1, 200)
	if err != nil {
		log.Fatalf("Error getting Site List ::: %s", err)
	}

	for _, site := range sites {
		fmt.Println("\t", site.Name, site.SiteID)
	}

	var site *SiteEntity
	site, err = siteService.GetSiteInfo(sites[1].SiteID)
	if err != nil {
		log.Fatalf("Error getting Site Info :: %s", err.Error())
	}

	fmt.Println("GOT SITE INFO")

	fmt.Println("\t", site.Name, site.SiteID)
	/*
	   siteId := site.SiteID

	   	fmt.Println()

	   	   err = siteService.DeleteSite(siteId)

	   	   	if err != nil {
	   	   		log.Fatalf("Error deleteing Site ::: %s", err)
	   	   	}

	   	   fmt.Println("\t", "DELETED SITE", siteId)

	   	   fmt.Println()

	   	   sites, err = siteService.GetSiteList(1, 200)

	   	   	if err != nil {
	   	   		log.Fatalf("Error getting Site List ::: %s", err)
	   	   	}

	   	   	for _, site := range sites {
	   	   		fmt.Println("\t", site.Name, site.SiteID)
	   	   	}

	   	   deviceService := NewDeviceService(client)
	   	   devices, err := deviceService.GetSiteDeviceList(siteId, 1, 200)

	   	   	if err != nil {
	   	   		log.Fatalf("Error getting Device List ::: %s", err.Error())
	   	   	}

	   	   	log.Println("GOT DEVICES")

	   	   		for _, device := range devices {
	   	   			fmt.Println("\t", device.Name, device.Mac, device.IP, device.Type)
	   	   		}
	*/
}
