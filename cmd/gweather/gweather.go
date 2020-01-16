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

		// Gets the weather from the web api
		var currWeather weathertool.Weather = weathertool.GetWeather(config.Location, config.Apikey.APPID)

		// Converts the temperature provided by the api fahrenheit
		tempF := weathertool.KelvinToFahrenheit(currWeather.Main.Temp)

		fmt.Println("Temperature for", config.Location, "is", tempF, "degrees Fahrenheit.")

		// If -e flag was specified then the code below will execute
		if config.Exp {
			expJSON := weathertool.Weather2Json(currWeather)
			err := ioutil.WriteFile("WeatherData.json", expJSON, 0644)
			if err != nil {
				log.Println("error:", err)
			}
		}

		config.Logfile.Close() // Properly close the logger
	} else {
		fmt.Println("Please provide the location of you want the weather of after the program call as an argument.")
	}
}
