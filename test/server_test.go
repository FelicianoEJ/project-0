package test

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

// GetAlot ...
func GetAlot() bool {
	for i := 1; i < 4; i++ {
		resp, _ := http.Get("http://localhost:8080/weather?id=" + strconv.Itoa(i))
		body, _ := ioutil.ReadAll(resp.Body)
		if string(body) == "" && resp.StatusCode != 200 {
			return false
		}
	}
	return true
}

// TestGetAlot ...
func TestGetAlot(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var ret bool = GetAlot()
		t.Log("test", i+1, "returned", ret)
		if !ret {
			t.Errorf("Something unexpected happened. The return body is: %t", ret)
		}
	}
}
