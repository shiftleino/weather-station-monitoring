package loader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type jsonContent struct {
	HomeName         string
	DataDatetime     string
	Temperature      float32
	CO2              float32
	Humidity         float32
	Noise            int
	Pressure         float32
	AbsolutePressure float32
	TempTrend        string
	PressureTrend    string
	InputDatetime    string
}

func createSheetsService(confFilePath string) *sheets.SpreadsheetsService {
	ctx := context.Background()
	service, err := sheets.NewService(ctx, option.WithCredentialsFile(confFilePath))

	if err != nil {
		log.Fatalf("Error when creating the Google Sheets Service: %v\n", err)
	}

	return service.Spreadsheets
}

func LoadStationDataSheets(stationData string, env string) {
	log.Println("Loading Weather Station Data into Google Sheets...")

	godotenv.Load(".env")
	var confFilePath string = os.Getenv("CONF_FILEPATH")
	var spredsheetId string = os.Getenv("SPREADSHEET_ID")
	service := createSheetsService(confFilePath)

	var jsonData *jsonContent
	json.Unmarshal([]byte(stationData), &jsonData)
	jsonDataValues := []interface{}{
		jsonData.HomeName,
		jsonData.DataDatetime,
		jsonData.Temperature,
		jsonData.CO2,
		jsonData.Humidity,
		jsonData.Noise,
		jsonData.Pressure,
		jsonData.AbsolutePressure,
		jsonData.TempTrend,
		jsonData.PressureTrend,
		jsonData.InputDatetime,
	}

	var data sheets.ValueRange
	appendRange := fmt.Sprintf("%s!A:K", env)
	data.Values = append(data.Values, jsonDataValues)

	_, err := service.Values.Append(spredsheetId, appendRange, &data).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Error when appending data to Google Sheets: %v\n", err)
	}
}
