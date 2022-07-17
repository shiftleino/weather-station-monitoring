package main

import (
	"fmt"

	extractor "github.com/shiftleino/weather-data-pipeline/extract-weather-data"
	transformer "github.com/shiftleino/weather-data-pipeline/transform-weather-data"
)

func main() {
	var stationData string = extractor.ExtractStationData()
	var transformedStationData string = transformer.TransformStationData(stationData)

	fmt.Println(transformedStationData)
	fmt.Println("Success.")
}
