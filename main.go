package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

const errorMessage = "Invalid Time Zone!"

type CurrentTime struct {
	CurrentTime time.Time `json:"current_time,omitempty"`
}

func main() {

	//mux := http.NewServeMux()

	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:3000", router))

}

func getTime(w http.ResponseWriter, r *http.Request) {

	timeZone := "UTC"
	response := make(map[string]string)

	if r.URL.Query().Has("tz") {
		timeZone = r.URL.Query().Get("tz")

	}

	res1 := strings.Split(timeZone, ",")

	for _, tzdb := range res1 {
		location, err := time.LoadLocation(string(tzdb))

		if err != nil {

			response[string(tzdb)] = "Invalid Time Zone!"

		} else {

			response[string(tzdb)] = time.Now().In(location).String()

		}
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

}
