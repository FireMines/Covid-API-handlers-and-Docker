package cases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

/*
Entry point handler for Location information
*/
func CovidInfoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		covidCasesInfoGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}

func covidCasesInfoGetRequest(w http.ResponseWriter, r *http.Request) {

	urlLastVal := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")
	r.Header.Add("content-type", "application/json")
	if urlLastVal == "cases" {
		http.Error(w, "Looks like you forgot to add a country! Place do so next time ;)", http.StatusBadRequest)
		return
	}

	// Get country matching the name in the URL
	countryInfo := getCovidCasesPerCountry(urlLastVal)

	// Marshal them and write to Writer
	writeCountry, err := json.Marshal(countryInfo)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
		http.Error(w, "Error in response:", http.StatusInternalServerError)
	}

	w.Header().Add("content-type", "application/json")
	w.Write(writeCountry)

}
