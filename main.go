// Command line tool to get current weather for particular location
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
)

// OpenWeather weather data struct
type weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func main() {
	location := os.Args[1]
	if location != "" {
		fmt.Println("Fetching...")
		getWeatherTemp(location)
	} else {
		fmt.Println("Please provide an argumet after the program call.")
	}
}

// The function getWeather utilizes HTTP GET protocols to
// retrieve web api from OpenWeather in order to use it in the cli app
func getWeatherTemp(location string) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + location + "&APPID=f60c52fb199f51b3d8c425bc6f2c2752")
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			var currentWeather weather
			json.Unmarshal(body, &currentWeather)
			tempF := KelvinToFahrenheit(currentWeather.Main.Temp)
			fmt.Println("Temperature for", location, "is", tempF, "degrees Fahrenheit.")
		}
	}
}

// KelvinToFahrenheit converts kelvin to fahrenheit.
func KelvinToFahrenheit(k float64) float64 {
	value := round(float64((k * 9 / 5) - 459.67))
	return value
}

// Round a value to 1 decimal place.
func round(val float64) float64 {
	var round float64
	digit := 10 * val
	if _, div := math.Modf(digit); div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / 10
}
