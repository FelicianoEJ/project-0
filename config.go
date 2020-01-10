package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

type apiKey struct {
	APPID string `json:"APPID"`
}

var exp bool
var location string
var args []string
var apikey apiKey
var logfile *os.File

func init() {
	// Config the flags for the program
	flag.BoolVar(&exp, "e", false, "Export the data retrieve from the web api into a json file.")
	flag.Parse()

	// Get the arguments and flags
	args = os.Args[0:]

	// Set location data depending on the argument
	if len(args) < 2 {
		location = ""
	} else if exp {
		location = args[2] // If flag exp is set then location is args[2]
	} else {
		location = args[1]
	}

	// Config the logger
	var err error
	logfile, err = os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// defer file.Close()
	log.SetOutput(logfile)

	// Set the api key
	config, configerr := ioutil.ReadFile("appconfig.json")
	if configerr != nil {
		log.Fatalln("fatal error: could not read file appconfig.json.")
	}
	jsonerr := json.Unmarshal(config, &apikey)
	if jsonerr != nil || apikey.APPID == "" {
		log.Fatalln("error: appconfig.json is corrupted or you are missing apikey to access weather data.")
	}
}
