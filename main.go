package main

import (
	"fmt"

	extractor "github.com/shiftleino/weather-data-pipeline/extract-weather-data"
)

func main() {
	var result string = string(extractor.GetNetatmoData())
	fmt.Println(result)
	fmt.Println("Success.")
}
