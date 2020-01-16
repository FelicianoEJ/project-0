// Package config ...
package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

// APIKey ...
type APIKey struct {
	APPID string `json:"APPID"`
}

// Exp ...
var Exp bool

// Location ...
var Location string

// Args ...
var Args []string

// Apikey ...
var Apikey APIKey

// Logfile ...
var Logfile *os.File // defer Logfile.Close()

func init() {
	// Config the flags for the program
	flag.BoolVar(&Exp, "e", false, "Export the data retrieve from the web api into a json file.")
	flag.Parse()

	// Get the arguments and flags
	Args = os.Args[0:]

	// Set Location data depending on the argument
	if len(Args) < 2 {
		Location = ""
	} else if Exp {
		Location = Args[2] // If flag Exp is set then Location is Args[2]
	} else {
		Location = Args[1]
	}

	// Config the logger
	var err error
	Logfile, err = os.OpenFile("logs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(Logfile)

	// Set the api key
	var config []byte
	config, err = ioutil.ReadFile("config/appconfig.json")
	if err != nil {
		log.Println(err)
		log.Fatalln("Fatal error: could not read file appconfig.json.")
	}
	err = json.Unmarshal(config, &Apikey)
	if err != nil || Apikey.APPID == "" {
		log.Println(err)
		log.Fatalln("Fatal error: appconfig.json is corrupted or you are missing apikey to access weather data.")
	}
}
