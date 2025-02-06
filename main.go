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

type Date struct {
	Month  int `json:"month"`
	Serial int `json:"serial"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

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

	siteService := NewOmadaService[SiteEntity](client, "sites")
	// var sites []SiteEntity
	// var err error
	sites, err := siteService.GetList(1, 200)
	if err != nil {
		log.Fatalf("Error getting Site List ::: %s", err)
	}

	for _, site := range sites {
		fmt.Println("\t", site.Name, site.SiteID)
	}

	fmt.Println()

	testSite := sites[0]

	site, err := siteService.GetOne(*testSite.SiteID)
	if err != nil {
		log.Fatalf("Error getting Site %s ::: %s", *testSite.SiteID, err.Error())
	}
	fmt.Println("\t", site.Name, site.SiteID)

	newSiteEntity := SiteEntity{Name: "test", Region: "United States", TimeZone: "America/Chicago", Scenario: "Home", DeviceAccountSetting: DeviceAccountSetting{Username: os.Getenv("DEVICE_UID"), Password: os.Getenv("DEVICE_PWD")}}
	err = siteService.Create(newSiteEntity)
	if err != nil {
		log.Fatalf("Error creating new Site :: %s", err.Error())
	}
	sites, err = siteService.GetList(1, 200)
	if err != nil {
		log.Fatalf("Error getting site list :: %s", err.Error())
	}

	newSite := sites[len(sites)-1]
	fmt.Printf("\t New Site :: %s : %s", newSite.Name, *newSite.SiteID)

	time.Sleep(10 * time.Second)

	err = siteService.Delete(*newSite.SiteID)
	if err != nil {
		log.Fatalf("Error deleteing site :: %s", err.Error())
	}
	fmt.Println(fmt.Printf("Deleting site :: %s", newSite.Name))

	// lanService := EntityService[]
}

/*sites, err = siteService.GetSiteList(1, 200)
	if err != nil {
		log.Fatalf("Error getting Site List ::: %s", err)
	}

	for _, site := range sites {
		fmt.Println("\t", site.Name, site.SiteID)
	}

	fmt.Println()

	err = siteService.CreateNewSite("test", ", "America/Chicago", "Home", os.Getenv("DEVICE_UID"), os.Getenv("DEVICE_PWD"))
	for err != nil {
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
}*/
