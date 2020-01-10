package main

import (
	"flag"
	"log"
	"os"
)

var exp bool
var location string
var args []string
var file *os.File

func init() {
	// Config the flags for the program
	flag.BoolVar(&exp, "e", false, "Export the data retrieve from the web api into a json file.")
	flag.Parse()

	// Get the arguments
	args = os.Args[0:]

	// Set location data
	if len(args) < 2 {
		location = ""
	} else if exp {
		location = args[2]
	} else {
		location = args[1]
	}

	// Config the logger
	var err error
	file, err = os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer file.Close()
	log.SetOutput(file)
}
