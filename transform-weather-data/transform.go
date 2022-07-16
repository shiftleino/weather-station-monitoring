package transformer

import (
	"encoding/json"
	"fmt"
	"time"
)

type stationDataTemplate struct {
	DateTime         time.Time `json:"dateTime"`
	HomeName         string    `json:"homeName"`
	Temperature      *float32  `json:"temperature"`
	CO2              *float32  `json:"co2"`
	Humidity         *float32  `json:"humidity"`
	Noise            *int      `json:"noise"`
	Pressure         *float32  `json:"pressure"`
	AbsolutePressure *float32  `json:"absolutePressure"`
	TempTrend        *string   `json:"temperatureTrend"`
	PressureTrend    *string   `json:"pressureTrend"`
}

func TransformStationData(stationData string) {
	fmt.Println("Transforming Weather Station Data...")
	var tmpStationData map[string]interface{}
	json.Unmarshal([]byte(stationData), &tmpStationData)

	transformedStationData := &stationDataTemplate{
		DateTime: time.Now(),
		HomeName: tmpStationData["body"]["devices"][0]["place"]["home_name"],
	}

	transformedStationDataJSON, err := json.Marshal()
}
