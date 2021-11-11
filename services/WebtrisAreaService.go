package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"webtrisapiapp/model"
)

func getHttpClient() http.Client {
	return http.Client{
		Timeout: time.Second * 60,
	}
}

func buildGetRequest(uri string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		log.Fatal("Failed to init http request")
	} else {
		return req
	}
	return nil
}

func performRequest(httpClient *http.Client, request *http.Request) *http.Response {
	res, err := httpClient.Do(request)
	if err != nil {
		log.Fatal("something when wrong performing the request", err)
	} else {
		return res
	}
	return nil
}

func DownloadWebtrisData(baseurl string) *model.WebtrisAreas {
	log.Println(fmt.Sprintf("Prepairing to call %s", baseurl))
	httpClient := getHttpClient()
	httpRequest := buildGetRequest(baseurl) //pointer to http request
	log.Println("Performing GET request")
	response := performRequest(&httpClient, httpRequest)
	body, err := parseHttpResponseBody(response)
	areas, done := responseBodyToJson(err, body)
	if done {
		return areas
	}
	return nil
}

func parseHttpResponseBody(response *http.Response) ([]byte, error) {
	log.Println("Reading body from response")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read response body")
	}
	return body, err
}

func responseBodyToJson(err error, body []byte) (*model.WebtrisAreas, bool) {
	log.Println("Unmarshalling body to objects")
	availableSites := model.WebtrisAreas{}
	err = json.Unmarshal(body, &availableSites)
	if err != nil {
		log.Fatal("Failed during unmarshal")
	} else {
		return &availableSites, true
	}
	return nil, false
}