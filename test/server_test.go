package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

// GetAlot ...
func GetAlot() bool {
	resp, _ := http.Get("http://localhost:8080/weather?id=1")
	body, _ := ioutil.ReadAll(resp.Body)
	if len(string(body)) < 1 {
		return false
	}
	// resp, _ = http.Get("http://localhost:8080/weather?id=2")
	// body, _ = ioutil.ReadAll(resp.Body)
	// if len(string(body)) < 1 {
	// 	return false
	// }
	// resp, _ = http.Get("http://localhost:8080/weather?id=3")
	// body, _ = ioutil.ReadAll(resp.Body)
	// if len(string(body)) < 1 {
	// 	return false
	// }
	return true
}

// TestGetAlot ...
func TestGetAlot(t *testing.T) {
	// Logger
	logfile, _ := os.OpenFile("logs/testlog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logfile)

	for i := 0; i < 1000; i++ {
		var ret bool = GetAlot()
		t.Log("test", i+1, "returned", ret)
		if !ret {
			t.Errorf("Something unexpected happened. The return body is: %t", ret)
		}
	}
}
