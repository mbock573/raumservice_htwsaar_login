package raumserviceLogin

import (
	"github.com/mbock573/httpClientHelper"
	"io"
	"log"
	"net/http"
	"net/url"
)

const nutzername = "mbock"
const passwort = "X2hahxnr7!"

const BaseURL = "https://raumservice.htwsaar.de/"
const LoginURL = "https://raumservice.htwsaar.de/index.php?mid=998&muid=0&muuid=0"
const TimetableOptionsURL = "https://raumservice.htwsaar.de/index.php?mid=100&muid=0&muuid=0"
const TimetableParsingURL = "https://raumservice.htwsaar.de/index.php?mid=100&muid=0&muuid=0"

func Run(client *http.Client) error {
	//initial request
	httpResult, err := httpClientHelper.HttpGetRequest(client, BaseURL)
	if err != nil {
		log.Fatalf("Error while request on %s: %v", BaseURL, err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error while closing Body on %s request: %v", BaseURL, err)
		}
	}(httpResult.Body)

	//login request
	formBody := url.Values{
		"username": {nutzername},
		"password": {passwort},
		"action":   {"login"},
	}
	httpResult, err = httpClientHelper.HttpPostRequest(client, LoginURL, formBody)
	if err != nil {
		log.Fatalf("Error while HTTP POST request on %s: %v", LoginURL, err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error while closing Body on %s request: %v", LoginURL, err)
		}
	}(httpResult.Body)
	return nil
}
