package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/FelicianoEJ/project-0/config"
	"github.com/FelicianoEJ/project-0/internal/weathermodel"
	"github.com/FelicianoEJ/project-0/internal/weathertool"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// WeatherList if we want to modify data
type WeatherList struct {
	inventory []weathertool.Weather
	lock      sync.Mutex
}

func (wl *WeatherList) add(w weathertool.Weather) {
	wl.lock.Lock()
	wl.inventory = append(wl.inventory, w)
	wl.lock.Unlock()
}

func main() {
	db, err := gorm.Open("sqlite3", config.Configuration.DatabaseConnection)
	if err != nil {
		log.Fatalln("Failed to connect database: ", err)
	}
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("web")))

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("city")
		output, err := json.Marshal(db.Where("name = ?", name).First(&weathermodel.WeatherModel{}))
		if err != nil {
			w.WriteHeader(http.StatusTeapot)
			fmt.Fprintln(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(output))
	})

	http.HandleFunc("/weatherbyid", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.FormValue("id"))
		output, err := json.Marshal(db.Where("ID = ?", id).First(&weathermodel.WeatherModel{}))
		if err != nil {
			w.WriteHeader(http.StatusTeapot)
			fmt.Fprintln(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(output))
	})

	fmt.Println("Listening on ports 8080 (http) and ...") // 8081 (https)

	errorChan := make(chan error, 5)
	go func() {
		errorChan <- http.ListenAndServe(":8080", nil)
	}()

	// Need to get certificates to run https
	// go func() {
	// 	errorChan <- http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
	// }()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				log.Fatalln(err)
			}

		case sig := <-signalChan:
			fmt.Println("\nShutting down due to", sig)
			os.Exit(0)
		}
	}
}
