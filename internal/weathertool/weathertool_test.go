package weathertool

import (
	"testing"
)

func TestGetWeather(t *testing.T) {
	resp := GetWeather("Ponce", "")
	if resp.ID == 0 {
		t.Error("GetWeather failed")
	}
}
