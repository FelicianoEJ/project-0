package weathermodel

import "github.com/jinzhu/gorm"

// WeatherModel ...
type WeatherModel struct {
	gorm.Model
	Name        string
	CoordLon    float64
	CoordLat    float64
	WeatherMain string
	WeatherDesc string
	Base        string
	Temp        float64
	Pressure    int
	Humidity    int
	TempMin     float64
	TempMax     float64
	Visibility  int
	WindSpeed   float64
	WindDeg     int
	Country     string
	Sunrise     int
	Sunset      int
}
