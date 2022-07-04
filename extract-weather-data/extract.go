package extractor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const baseUrl = "https://api.netatmo.com/"
const getStationsData = "api/getstationsdata?get_favorites=false"

func GetNetatmoData() []byte {
	var searchUrl string = baseUrl + getStationsData
	req, err := http.NewRequest(http.MethodGet, searchUrl, nil)
	if err != nil {
		fmt.Printf("Error when creating request: %s\n", err)
		os.Exit(1)
	}

	godotenv.Load(".env")
	var token string = os.Getenv("ACCESS_TOKEN")
	var auth string = fmt.Sprintf("Bearer %s", token)
	fmt.Println(auth)
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
	fmt.Printf("Got response: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error when reading response body: %s\n", err)
		os.Exit(1)
	}
	return resBody
}

func ExtractNetatmoData() string {
	var result string = "Hello World"
	return result
}
