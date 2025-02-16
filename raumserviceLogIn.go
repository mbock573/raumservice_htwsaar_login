package raumserviceLogin

import (
	"github.com/mbock573/httpClientHelper"
	raumserviceURLs "github.com/mbock573/raumservice_htwsaar_login/internal"
	"io"
	"log"
	"net/http"
	"net/url"
)

const nutzername = "mbock"
const hi = "hi"
const passwort = "X2hahxnr7!"

func Run(client *http.Client) error {
	//initial request
	httpResult, err := httpClientHelper.HttpGetRequest(client, raumserviceURLs.BaseURL)
	if err != nil {
		log.Fatalf("Error while request on %s: %v", raumserviceURLs.BaseURL, err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error while closing Body on %s request: %v", raumserviceURLs.BaseURL, err)
		}
	}(httpResult.Body)

	//login request
	formBody := url.Values{
		"username": {nutzername},
		"password": {passwort},
		"action":   {"login"},
	}
	httpResult, err = httpClientHelper.HttpPostRequest(client, raumserviceURLs.LoginURL, formBody)
	if err != nil {
		log.Fatalf("Error while HTTP POST request on %s: %v", raumserviceURLs.LoginURL, err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error while closing Body on %s request: %v", raumserviceURLs.LoginURL, err)
		}
	}(httpResult.Body)
	return nil
}
