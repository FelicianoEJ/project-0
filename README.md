# project-0 by Emilio Feliciano aka FelicianoEJ

Project 0 is command line tool used to retrieve temperatures from cities around the globe through the use of web api.

To make it work simple type the command for the app and provide an argument.

*For example:*

`command-name Ponce,PR`

## How to build and use

Clone the project from github or use the `go get` tool to get this project.

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
