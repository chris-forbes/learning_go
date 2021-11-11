package main


import (
	"fmt"
	"log"
	services "webtrisapiapp/services"
)



func main() {
	baseUrl := "https://webtris.highwaysengland.co.uk/api/v1/areas"
	areas := services.DownloadWebtrisData(baseUrl)
	log.Println("Response: ")
	log.Println(fmt.Sprintf("Response from server: %+v",areas))
}