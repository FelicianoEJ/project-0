package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/FelicianoEJ/project-0/config"
	"github.com/FelicianoEJ/project-0/internal/weathertool"
)

func main() {
	if config.Location != "" {
		fmt.Println("Fetching...")
		var currWeather weathertool.Weather = weathertool.GetWeather(config.Location, config.Apikey.APPID)
		tempF := weathertool.KelvinToFahrenheit(currWeather.Main.Temp)
		fmt.Println("Temperature for", config.Location, "is", tempF, "degrees Fahrenheit.")
		if config.Exp {
			expJSON := weathertool.Weather2Json(currWeather)
			err := ioutil.WriteFile("WeatherData.json", expJSON, 0644)
			if err != nil {
				log.Println("error:", err)
			}
		}
	} else {
		fmt.Println("Please provide the location of you want the weather of after the program call as an argument.")
	}
}
