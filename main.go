package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/FelicianoEJ/project-0/weathertool"
)

func main() {
	if location != "" {
		fmt.Println("Fetching...")
		var currWeather weathertool.Weather = weathertool.GetWeather(location)
		tempF := weathertool.KelvinToFahrenheit(currWeather.Main.Temp)
		fmt.Println("Temperature for", location, "is", tempF, "degrees Fahrenheit.")
		if exp {
			expJSON := weathertool.Weather2Json(currWeather)
			err := ioutil.WriteFile("WeatherData", expJSON, 0644)
			if err != nil {
				log.Println("error:", err)
			}
		}
	} else {
		fmt.Println("Please provide the location of you want the weather of after the program call as an argument.")
	}
}
