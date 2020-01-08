package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FelicianoEJ/project-0/weathertool"
)

var exp bool
var location string

func init() {
	// Config the flags for the program
	flag.BoolVar(&exp, "e", false, "Export the data retrieve from the web api into a json file.")
	flag.Parse()
}

func main() {
	args := os.Args[1:]
	if exp {
		location = args[1]
	} else {
		location = args[0]
	}
	if location != "" {
		fmt.Println("Fetching...")
		var currWeather weathertool.Weather = weathertool.GetWeatherTemp(location)
		tempF := weathertool.KelvinToFahrenheit(currWeather.Main.Temp)
		fmt.Println("Temperature for", location, "is", tempF, "degrees Fahrenheit.")
	} else {
		fmt.Println("Please provide an argumet after the program call.")
	}
}
