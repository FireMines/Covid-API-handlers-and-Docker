package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
)

/*
Entry point handler for Location information
*/
func UniInfoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		uniInfoHandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}

/**
 *	Gets all unis (partially) matching the input name, and their corresponding data.
 */
func UnisGetByNameAndCountry(uniName string, countryName string) []University {
	var unis []University
	var allUnis []University

	url := "http://universities.hipolabs.com/search?name=" + uniName
	if countryName != "" {
		url += "&country=" + countryName
	}

	write, _ := http.Get(url)
	body, _ := io.ReadAll(write.Body)

	err := json.Unmarshal(body, &unis)
	if err != nil {
		fmt.Println("Error in response:", err.Error()) //log.Fatal("Not able to unmarshal")
	}

	// Get languages and map data
	var output []CountryAndNeighbours
	for _, uni := range unis {

		country := uni.Country
		res, _ := http.Get("https://restcountries.com/v3.1/name/" + country)

		data, _ := io.ReadAll(res.Body)

		_ = json.Unmarshal(data, &output)

		allUnis = append(allUnis, University{
			Name:      uni.Name,
			Country:   uni.Country,
			A2C:       uni.A2C,
			Webpages:  uni.Webpages,
			Languages: output[0].Languages,
			Map:       output[0].Map,
		})
	}

	return allUnis
}

func uniInfoHandleGetRequest(w http.ResponseWriter, r *http.Request) {
	urlLastVal := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")

	// Get all unis (partially) matching the name in the URL
	allUnis := UnisGetByNameAndCountry(urlLastVal, "")

	// Marshal them and write to Writer
	writeUnis, _ := json.Marshal(allUnis)
	w.Header().Add("content-type", "application/json")
	w.Write(writeUnis)
}
