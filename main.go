package main

import (
	"log"

	extractor "github.com/shiftleino/weather-data-pipeline/extract-weather-data"
	loader "github.com/shiftleino/weather-data-pipeline/load-weather-data"
	transformer "github.com/shiftleino/weather-data-pipeline/transform-weather-data"
)

func main() {
	var stationData string = extractor.ExtractStationData()
	var transformedStationData string = transformer.TransformStationData(stationData)
	loader.LoadStationDataSheets(transformedStationData)
	log.Println("Weather Station Data Pipeline Executed Successfully.")
}
