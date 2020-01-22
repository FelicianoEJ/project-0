package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/FelicianoEJ/project-0/config"
	"github.com/FelicianoEJ/project-0/internal/weathermodel"
	"github.com/FelicianoEJ/project-0/internal/weathertool"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	if config.Location != "" {
		fmt.Println("Fetching...")

		// Gets the weather from the web api
		var currWeather weathertool.Weather = weathertool.GetWeather(config.Location, config.Configuration.APPID)

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
		} else if config.Db {
			db, err := gorm.Open("sqlite3", config.Configuration.DatabaseConnection)
			if err != nil {
				panic("failed to connect database")
			}
			defer db.Close()

			// Migrate the schema
			db.AutoMigrate(&weathermodel.WeatherModel{})

			// Create
			var model weathermodel.WeatherModel = weathermodel.WeatherModel{
				Name:        currWeather.Name,
				CoordLon:    currWeather.Coord.Lon,
				CoordLat:    currWeather.Coord.Lat,
				WeatherMain: currWeather.Weather[0].Main,
				WeatherDesc: currWeather.Weather[0].Description,
				Base:        currWeather.Base,
				Temp:        currWeather.Main.Temp,
				Pressure:    currWeather.Main.Pressure,
				Humidity:    currWeather.Main.Humidity,
				TempMin:     currWeather.Main.TempMin,
				TempMax:     currWeather.Main.TempMax,
				Visibility:  currWeather.Visibility,
				WindSpeed:   currWeather.Wind.Speed,
				WindDeg:     currWeather.Wind.Deg,
				Country:     currWeather.Sys.Country,
				Sunrise:     currWeather.Sys.Sunrise,
				Sunset:      currWeather.Sys.Sunset,
			}
			db.Create(&model)
		}

		config.Logfile.Close() // Properly close the logger
	} else {
		fmt.Println("Please provide the location you want the weather off after the program call as an argument.")
		fmt.Println("If the location name has spaces in between the use double quotes.")
		fmt.Println("For example: $ <command-name> \"San Diego\"")
	}
}
