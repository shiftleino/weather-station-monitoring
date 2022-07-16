package main

import (
	"fmt"

	extractor "github.com/shiftleino/weather-data-pipeline/extract-weather-data"
)

func main() {
	var stationData string = extractor.ExtractStationData()
	fmt.Println(stationData)
	fmt.Println("Success.")
}
