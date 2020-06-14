package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

const (
	urlPrefix    = "http://api.openweathermap.org/data/2.5/weather?zip="
	apiKeyPrefix = "&appid="
	apiKey       = "ea32cfa63f93679a01cb6a8ebcbf8e48"
	unitsSuffix  = "&units=imperial"
)

type CurrentWeatherInfo struct {
	Description     []CurrentDescription   `json:"weather"`
	CurrentTempInfo map[string]interface{} `json:"main"`
	City            string                 `json:"name"`
}

type CurrentDescription struct {
	Detail string `json:"description"`
}

func main() {
	zipCode := *flag.String("zip", "10001", "US ZIP code")
	flag.Parse()
	url := fmt.Sprintf("%s%s%s%s%s", urlPrefix, zipCode, apiKeyPrefix, apiKey, unitsSuffix)
	resp, respErr := http.Get(url)
	if respErr != nil {
		fmt.Printf("An error occurred while contacting OpenWeatherMap.\n%s", respErr.Error())
		return
	}
	dispCurrentForecast(resp)
}

func dispCurrentForecast(resp *http.Response) {
	weather := new(CurrentWeatherInfo)
	decodeErr := json.NewDecoder(resp.Body).Decode(weather)
	if decodeErr != nil {
		fmt.Printf("An error occured while processing forecast data.\n%s\n", decodeErr.Error())
		return
	}
	fmt.Printf("-- Location --\n%s\n", weather.City)
	fmt.Printf("-- Forecast --\n%s\n", weather.Description[0].Detail)
	fmt.Printf("Current temperature: %s\n", formatTemp(weather.CurrentTempInfo["temp"].(float64)))
}

func formatTemp(temp float64) string {
	return fmt.Sprintf("%d°F", int(temp))
}
