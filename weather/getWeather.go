package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetWeather(location string) {
	const apiKey = "d22834dcbb6d91684237885b1830141b"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", location, apiKey)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making request %v:", err)
		return
	}

	defer response.Body.Close()

	var weatherData map[string]interface{}

	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		log.Fatalf("Error decoding response %v:", err)
		return
	}

	if weather, ok := weatherData["weather"].([]interface{}); ok {
		fmt.Printf("Weather: %s\n", weather[0].(map[string]interface{})["description"])
	}

	if main, ok := weatherData["main"].(map[string]interface{}); ok {
		fmt.Printf("Temperature: %.2f°C\n", main["temp"].(float64))
		fmt.Printf("Humidity: %.0f%%\n", main["humidity"].(float64))
	}

}
