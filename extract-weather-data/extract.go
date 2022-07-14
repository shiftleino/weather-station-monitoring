package extractor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const baseUrl = "https://api.netatmo.com/"
const getAccessTokenReq = "oauth2/token"
const getStationDataReq = "api/getstationsdata?get_favorites=false"

func getAccessToken() string {
	var seachUrl string = baseUrl + getAccessTokenReq
	godotenv.Load(".env")
	var clientSecret string = os.Getenv("CLIENT_SECRET")
	var clientId string = os.Getenv("CLIENT_ID")
	var username string = os.Getenv("USERNAME")
	var password string = os.Getenv("PASSWORD")

	postBodyBytes, _ := json.Marshal(map[string]string{
		"grant_type":    "password",
		"client_id":     clientId,
		"client_secret": clientSecret,
		"username":      username,
		"password":      password,
		"scope":         "read_station",
	})
	postBody := bytes.NewBuffer(postBodyBytes)
	res, err := http.Post(seachUrl, "application/x-www-form-urlencoded", postBody)
	if err != nil {
		fmt.Printf("Error when sending request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error when reading response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(string(resBody))

	return "Hello World"
}

func GetStationData() string {
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
