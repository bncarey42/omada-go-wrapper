package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host_name := os.Getenv("OMADA_HOST")
	port := os.Getenv("OMADA_PORT")
	apiEndpoint := os.Getenv("OMADA_BASE_ENDPOINT")

	sitename := os.Getenv("OMADA_SITENAME")

	omadaClient := NewOmadaClient(fmt.Sprintf("https://%s:%s/%s", host_name, port, apiEndpoint), sitename)

	omadaClient.Login(os.Getenv("OMADA_USER"), os.Getenv("OMADA_PASSWORD"))
}
