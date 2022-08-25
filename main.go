package main

import (
	"flag"
	"log"

	extractor "github.com/shiftleino/weather-data-pipeline/extract-weather-data"
	loader "github.com/shiftleino/weather-data-pipeline/load-weather-data"
	transformer "github.com/shiftleino/weather-data-pipeline/transform-weather-data"
)

func main() {
	env := flag.String("env", "test", "Environment name, used in Sheets")
	flag.Parse()
	var stationData string = extractor.ExtractStationData()
	var transformedStationData string = transformer.TransformStationData(stationData)
	loader.LoadStationDataSheets(transformedStationData, *env)
	log.Println("Success.")
}
