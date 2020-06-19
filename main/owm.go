package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	zipCode, tempExtreme := initFlags()
	*zipCode = correctZip(*zipCode)
	url := fmt.Sprintf("%s%s%s%s%s", urlPrefix, *zipCode, apiKeyPrefix, apiKey, unitsSuffix)
	resp, respErr := http.Get(url)
	if respErr != nil {
		fmt.Printf("An error occurred while contacting OpenWeatherMap.\n%s", respErr.Error())
		return
	}
	dispCurrentForecast(resp, *tempExtreme)
}

func initFlags() (*string, *string) {
	zipCode := flag.String("zip", "10001", "US ZIP code")
	tempExtreme := flag.String("extreme", "", "High or low temp for location")
	flag.Parse()
	return zipCode, tempExtreme
}

func correctZip(zipCode string) string {
	zipInt, convErr := strconv.Atoi(zipCode)
	reader := bufio.NewReader(os.Stdin)
	for (zipInt < 501) || (zipInt > 99950) || (len(zipCode) != 5) || (convErr != nil) {
		fmt.Println("You entered an incorrect zip code. Please try again.")
		zipCode, _ := reader.ReadString('\n')
		zipInt, convErr = strconv.Atoi(zipCode)
	}
	return zipCode
}

func dispCurrentForecast(resp *http.Response, tempExtreme string) {
	weather := new(CurrentWeatherInfo)
	decodeErr := json.NewDecoder(resp.Body).Decode(weather)
	if decodeErr != nil {
		fmt.Printf("An error occured while processing forecast data.\n%s\n", decodeErr.Error())
		return
	}
	fmt.Printf("-- Location --\n%s\n", weather.City)
	fmt.Printf("-- Today's Forecast --\n%s\n", weather.Description[0].Detail)
	fmt.Printf("Current temperature: %s\n", formatTemp(weather.CurrentTempInfo["temp"].(float64)))
	dispCurrentExtreme(weather, tempExtreme)
}

func dispCurrentExtreme(weather *CurrentWeatherInfo, tempExtreme string) {
	if tempExtreme == "" {
		return
	} else if tempExtreme == "high" {
		fmt.Printf("High temperature: %s\n", formatTemp(weather.CurrentTempInfo["temp_max"].(float64)))
	} else if tempExtreme == "low" {
		fmt.Printf("Low temperature: %s\n", formatTemp(weather.CurrentTempInfo["temp_min"].(float64)))
	}
}

func formatTemp(temp float64) string {
	return fmt.Sprintf("%d°F", int(temp))
}
