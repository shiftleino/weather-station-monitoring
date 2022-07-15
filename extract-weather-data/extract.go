package extractor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	fmt.Println("Requesting Access Token...")
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
		fmt.Printf("Error when creating request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error when sending request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	fmt.Printf("Got Response: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error when reading response body: %s\n", err)
		os.Exit(1)
	}
	var resBodyContent string = string(resBody)
	var resData map[string]interface{}
	json.Unmarshal([]byte(resBodyContent), &resData)
	accessToken := fmt.Sprint(resData["access_token"])
	return accessToken
}

func GetStationData() string {
	fmt.Println("Requesting Weather Station Data...")
	var searchUrl string = baseUrl + getStationDataReq
	req, err := http.NewRequest(http.MethodGet, searchUrl, nil)
	if err != nil {
		fmt.Printf("Error when creating request: %s\n", err)
		os.Exit(1)
	}

	var accessToken string = getAccessToken()
	var auth string = fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", auth)
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error when sending request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	fmt.Printf("Got response: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error when reading response body: %s\n", err)
		os.Exit(1)
	}

	return string(resBody)
}

func ExtractNetatmoData() string {
	var result string = "Hello World"
	return result
}
