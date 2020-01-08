package weathertool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

// Weather is a struct that contains the data from the json OpenWeather weather api
type Weather struct {
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

// GetWeatherTemp utilizes HTTP GET protocols to retrieve
// web api from OpenWeather in order to use it in the cli app
func GetWeatherTemp(location string) Weather {
	var currentWeather Weather

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + location + "&APPID=f60c52fb199f51b3d8c425bc6f2c2752")
	defer resp.Body.Close()

	if err != nil {
		log.Println("error:", err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("error:", err)
		} else {
			json.Unmarshal(body, &currentWeather)
		}
	}
	return currentWeather
}

// Weather2Json uses json marshaller to transform the weather struct into a []byte of json
func Weather2Json(weatherData Weather) []byte {
	b, err := json.Marshal(weatherData)
	if err != nil {
		log.Println("error:", err)
	}
	return b
}

// KelvinToFahrenheit converts kelvin to fahrenheit.
func KelvinToFahrenheit(k float64) float64 {
	value := Round(float64((k * 9 / 5) - 459.67))
	return value
}

// Round a value to 1 decimal place.
func Round(val float64) float64 {
	var round float64
	digit := 10 * val
	if _, div := math.Modf(digit); div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / 10
}
