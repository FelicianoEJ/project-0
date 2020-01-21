# project-0 by Emilio Feliciano aka FelicianoEJ

Project 0 is command line tool used to retrieve temperatures from cities around the globe through the use of web api.

To make it work simply type the command for the app and provide an argument.
If the city name has spaces between the type the argument unde double quotes.

*For example:*

`command-name Ponce,PR`

or

`command-name "San Diego"`

## How to build and use

Clone the project from github or use the `go get` tool to get this project.

`go get -u github.com/FelicianoEJ/project-0 `


**Note: Make sure to run with flag -u to get all the dependencies.**

The project is separated into 2 diferent executables unde the `/cmd` folder.

You can either build or run the projet.

If you decide to run the project provide an argument so the tool can retrieve the data from a web api.

**Note: Make sure to do this from the root directory.**

*For example:*

Under the `project-0` directory run either of this commands.

```bash
go run cmd/gweather/gweather.go Miami

go build cmd/gweather/gweather.go
```

*Expected output for go run:*

```s
Fetching...
Temperature for Miami is 73.8 degrees Fahrenheit.
```

## Flags

`-e` is used to export the weather data into a json file.

*For Example:*

```bash
gweather -e Miami
```

`-db` is used to store the weather data returned by the api into the specified sqlite database under the config.json file.

```bash
gweather -db "New York"
```

## Features
- [x] Documentation
- [x] Unit testing
- [x] Benchmarking
- [x] Logging
- [x] API Library
- [x] CLI flags
- [ ] Environment variables
- [x] Concurrency
- [x] Data Persistance
- [x] HTTP/HTTPS API

## Presentation
- [ ] 5 minute live demonstration
- [ ] Slides & visual aides
