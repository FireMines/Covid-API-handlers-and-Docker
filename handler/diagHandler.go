package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
Entry point handler for Location information
*/
func DiagHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		diagHandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}

func diagHandleGetRequest(w http.ResponseWriter, r *http.Request) {

	// Instantiate the client
	client := &http.Client{}

	// Issue request
	res, err := client.Get(urlCountries)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}

	res2, err := client.Get(urlUnis)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}

	t := time.Now()
	elapsed := t.Sub(Start)

	diagInfo := map[string]interface{}{
		"countriesapi":    res.StatusCode,
		"universitiesapi": res2.StatusCode,
		"version":         "v1",
		"uptime":          elapsed.Seconds(),
	}

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err = encoder.Encode(diagInfo)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}
}
