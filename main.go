package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	city := os.Args[1]
	if city != "" {
		fmt.Println("Fetching...")
		checkWeather()
	} else {
		fmt.Println("Please provide an argumet after the program call.")
	}
}

func checkWeather() {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=London,uk&APPID=f60c52fb199f51b3d8c425bc6f2c2752")
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			//Forecast struct
			//json.Unmarshal
			fmt.Println(string(body))
		}
	}
}
