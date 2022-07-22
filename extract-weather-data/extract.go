package extractor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const baseUrl = "https://api.netatmo.com/"
const getAccessTokenReq = "oauth2/token"
const getStationDataReq = "api/getstationsdata?get_favorites=false"

func getAccessToken() string {
	log.Println("Requesting Access Token...")
	var searchUrl string = baseUrl + getAccessTokenReq
	godotenv.Load(".env")
	var clientSecret string = os.Getenv("CLIENT_SECRET")
	var clientId string = os.Getenv("CLIENT_ID")
	var username string = os.Getenv("USERNAME")
	var password string = os.Getenv("PASSWORD")

	urlParameters := url.Values{}
	urlParameters.Set("grant_type", "password")
	urlParameters.Set("client_id", clientId)
	urlParameters.Set("client_secret", clientSecret)
	urlParameters.Set("username", username)
	urlParameters.Set("password", password)
	urlParameters.Set("scope", "read_station")
	postBody := strings.NewReader(urlParameters.Encode())
	req, err := http.NewRequest(http.MethodPost, searchUrl, postBody)
	if err != nil {
		log.Fatalf("Error when creating request: %s\n", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error when sending request: %s\n", err)
	}
	defer res.Body.Close()
	log.Printf("Got Response: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error when reading response body: %s\n", err)
	}
	var resData map[string]interface{}
	json.Unmarshal(resBody, &resData)
	rawAccessToken, ok := resData["access_token"]
	if !ok {
		log.Fatalf("Value access_token does not exist.")
	}
	accessToken, ok := rawAccessToken.(string)
	if !ok {
		log.Fatalf("Value access_token is not a string.")
	}
	return accessToken
}

func ExtractStationData() string {
	var searchUrl string = baseUrl + getStationDataReq
	req, err := http.NewRequest(http.MethodGet, searchUrl, nil)
	if err != nil {
		log.Fatalf("Error when creating request: %s\n", err)
	}

	var accessToken string = getAccessToken()
	log.Println("Requesting Weather Station Data...")
	var auth string = fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", auth)
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error when sending request: %s\n", err)
	}
	defer res.Body.Close()
	log.Printf("Got response: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error when reading response body: %s\n", err)
	}

	return string(resBody)
}
