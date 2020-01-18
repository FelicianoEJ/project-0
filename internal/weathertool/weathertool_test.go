package weathertool

import (
	"testing"
)

func TestGetWeather(t *testing.T) {
	resp := GetWeather("Ponce", "6c4aeca349de0e9daa392dc3b6c684ce")
	if resp.ID == 0 {
		t.Error("GetWeather failed")
	}
}
