// Package config ...
package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

// APIKey is the structure that will contain the apikey to the openweather api.
// This structure makes possible the use of the json.Unmarshal() function.
type Config struct {
	APPID              string `json:"APPID"`
	DatabaseConnection string `json:"DatabaseConnection"`
}

// Exp is the boolean value for the -e flag.
var Exp bool

// Db is boolean value for the -db flag
var Db bool

// Location is the argument the user provides.
var Location string

// Args holds all the arguments.
var Args []string

// Configuration is the actual instance of the Config struct.
var Configuration Config

// Logfile is an instance of the struc File that is used to dump all the logs details.
var Logfile *os.File // defer Logfile.Close()

func init() {
	// Config the flags for the program
	flag.BoolVar(&Exp, "e", false, "Export the data retrieve from the web api into a json file.")
	flag.BoolVar(&Db, "db", false, "Export to persistent database specified at the default connection.")
	flag.Parse()

	// Get the all the arguments including the flags
	Args = os.Args[0:]
	argsSize := len(Args)

	// Set Location data depending on the argument
	if argsSize == 2 {
		Location = Args[1]
	} else if argsSize == 3 {
		Location = Args[2] // If flag Exp is set then Location is Args[2]
	} else {
		Location = "" // Defaults Location to an empty string
	}

	// Config the logger
	var err error
	Logfile, err = os.OpenFile("logs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(Logfile)

	// Set the api key
	var configData []byte
	configData, err = ioutil.ReadFile("config/appconfig.json")
	if err != nil {
		log.Println(err)
		log.Fatalln("error: could not read file appconfig.json.")
	}
	err = json.Unmarshal(configData, &Configuration)
	if err != nil || Configuration.APPID == "" {
		log.Println(err)
		log.Fatalln("error: appconfig.json is corrupted or you are missing apikey to access weather data.")
	}
}
