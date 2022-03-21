package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
)

/*
Entry point handler for Location information
*/
func NeighbourUniHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		neighbourHandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}

func neighbourHandleGetRequest(w http.ResponseWriter, r *http.Request) {
	// Country you are searching for
	WantUni := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")               // Gets the second output from path
	WantCountry := strings.ReplaceAll(path.Base(path.Dir(r.URL.Path)), " ", "%20") // Gets the first output from path

	limit, err := strconv.Atoi(r.URL.Query().Get("limit")) // Gets the optional limit put on how many to output
	noLimitSet := false
	if err != nil {
		noLimitSet = true
	}

	// Step 1:
	// One country (ex. Norway) => Bordering country alphas (ex. [FIN, RUS, SWE])
	var cntry []CountryAndNeighbours
	allInfo, _ := http.Get(COVIDTRACKER + WantCountry)
	countrybody, _ := io.ReadAll(allInfo.Body)
	err = json.Unmarshal(countrybody, &cntry)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}

	bord := cntry[0].Borders
	var alphaBody []byte

	// Step 2:
	// Bordering country alphas => Full names (ex. ["Finland", "Russian Federation", "Sweden"])
	tempurl := "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/date-range/"
	for i, temp := range bord {
		if i < len(bord)-1 {
			temp += ","
		}
		tempurl += temp
	}

	fullUrl, _ := http.Get(tempurl)
	fullUrlBody, _ := io.ReadAll(fullUrl.Body)
	alphaBody = append(alphaBody, fullUrlBody...)

	var borderingCountries []CountryAndNeighbours
	err = json.Unmarshal(alphaBody, &borderingCountries)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}

	// Step 3:
	// Full names => Universities in each country, matching (partially) with `WantUni`
	var doneAppending bool = false
	var writeBorderUnis []University
	for i, _ := range borderingCountries {
		if doneAppending {
			break
		}
		unis := UnisGetByNameAndCountry(WantUni, borderingCountries[i].Name.Common)

		for _, uni := range unis {
			writeBorderUnis = append(writeBorderUnis, uni)
			if len(writeBorderUnis) >= limit && !noLimitSet {
				doneAppending = true
				break
			}
		}
	}

	// Step 4.
	// Add local universities (universities of `WantCountry` that matches `WantUni`)
	unis := UnisGetByNameAndCountry(WantUni, WantCountry)
	writeBorderUnis = append(writeBorderUnis, unis...)

	// Step 5: (last step)
	// Write and return
	writeBody, _ := json.Marshal(writeBorderUnis)
	w.Header().Add("content-type", "application/json")
	w.Write(writeBody)
}
