package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Send a GET request to the weather website
	resp, err := http.Get("https://weather.com/weather/today/")
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Convert the response body to a string
	htmlContent := string(body)

	// Extract the temperature <span ,,,>
	tempStartSubstring := `<span data-testid="TemperatureValue" class="CurrentConditions--tempValue--MHmYY">`
	tempStartIndex := strings.Index(htmlContent, tempStartSubstring)
	if tempStartIndex == -1 {
		fmt.Println("Error: Temperature not found")
		return
	}

	// </span> ain't unique
	tempStartIndex += len(tempStartSubstring)
	tempEndIndex := strings.Index(htmlContent[tempStartIndex:], "</span>")
	if tempEndIndex == -1 {
		fmt.Println("Error: Temperature not found")
		return
	}

	temp := htmlContent[tempStartIndex : tempStartIndex+tempEndIndex]

	// Extract the Condition <div ,,,>
	conditionsStartSubstring := `<div data-testid="wxPhrase" class="CurrentConditions--phraseValue--mZC_p">`
	conditionsStartIndex := strings.Index(htmlContent, conditionsStartSubstring)
	if conditionsStartIndex == -1 {
		fmt.Println("Error: Weather conditions not found")
		return
	}

	// </div> ain't unique
	conditionsStartIndex += len(conditionsStartSubstring)
	conditionsEndIndex := strings.Index(htmlContent[conditionsStartIndex:], "</div>")
	if conditionsEndIndex == -1 {
		fmt.Println("Error: Weather conditions not found")
		return
	}
	conditions := htmlContent[conditionsStartIndex : conditionsStartIndex+conditionsEndIndex]

	// Print extracted data
	fmt.Println(temp)
	fmt.Println(conditions)
}
