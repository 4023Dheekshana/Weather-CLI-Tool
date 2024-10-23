package main

import (
	"fmt"
	"os"
	"weathertool/weather"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter the location. ")
		return
	}
	location := os.Args[1]
	weather.GetWeather(location)
}
