package transformer

import (
	"encoding/json"
	"log"
	"time"
)

type stationBody struct {
	Body *stationDevices
}

type stationDevices struct {
	Devices []*stationGeneralInfo
}

type stationGeneralInfo struct {
	Home_name      string
	Dashboard_data *stationMonitoringData
}

type stationMonitoringData struct {
	Temperature      float32
	CO2              float32
	Humidity         float32
	Noise            int
	Pressure         float32
	AbsolutePressure float32
	Temp_trend       string
	Pressure_trend   string
	Time_utc         float64
}

type resultTemplate struct {
	HomeName         string   `json:"homeName"`
	DataDatetime     string   `json:"dataDatetime"`
	Temperature      *float32 `json:"temperature"`
	CO2              *float32 `json:"co2"`
	Humidity         *float32 `json:"humidity"`
	Noise            *int     `json:"noise"`
	Pressure         *float32 `json:"pressure"`
	AbsolutePressure *float32 `json:"absolutePressure"`
	TempTrend        *string  `json:"tempTrend"`
	PressureTrend    *string  `json:"pressureTrend"`
	InputDatetime    string   `json:"inputDatetime"`
}

func TransformStationData(stationData string) string {
	log.Println("Transforming Weather Station Data...")
	var stationContent *stationBody
	json.Unmarshal([]byte(stationData), &stationContent)

	stationBody := *stationContent.Body
	stationDevice := *stationBody.Devices[0]
	stationDeviceData := *stationDevice.Dashboard_data
	var homeName string = stationDevice.Home_name
	timeUtc := time.Unix(int64(stationDeviceData.Time_utc), 0)
	var dataDatetime string = timeUtc.Local().Format(time.RFC3339)
	currentTime := time.Now()
	var timestampNow string = currentTime.Format(time.RFC3339)

	transformedData := &resultTemplate{
		HomeName:         homeName,
		DataDatetime:     dataDatetime,
		Temperature:      &stationDeviceData.Temperature,
		CO2:              &stationDeviceData.CO2,
		Humidity:         &stationDeviceData.Humidity,
		Noise:            &stationDeviceData.Noise,
		Pressure:         &stationDeviceData.Pressure,
		AbsolutePressure: &stationDeviceData.AbsolutePressure,
		TempTrend:        &stationDeviceData.Temp_trend,
		PressureTrend:    &stationDeviceData.Pressure_trend,
		InputDatetime:    timestampNow,
	}

	transformedDataJSON, err := json.Marshal(transformedData)
	if err != nil {
		log.Fatalf("Error when creating the transformed data JSON.")
	}

	log.Println("Weather Station Data Transformed Successfully.")
	return string(transformedDataJSON)
}
