package policy

import "net/http"

/* 	"encoding/json"
"fmt"
"io"
"net/http"
"path"
"strconv"
"strings"
"time"
*/ //consts "covidAss2"

/*
Entry point handler for policy information
*/
func PolicyHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		//handleMapPostRequest(w, r)
	case http.MethodGet:
		//policyGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only POST and GET are supported.", http.StatusNotImplemented)
		return
	}

}

/*
func policyGetRequest(w http.ResponseWriter, r *http.Request) {
	// Country you are searching for
	WantCountry := strings.ReplaceAll(path.Base(path.Dir(r.URL.Path)), " ", "%20") // Gets the first output from path
	WantDate := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")              // Gets the second output from path

	scope, err := strconv.Atoi(r.URL.Query().Get("scope")) // Gets the optional limit put on how many to output
	noScopeSet := false
	var dateNow string
	if err != nil {
		noScopeSet = true
		currentTime := time.Now()
		dateNow = currentTime.Format("2022-03-01")
	}

	// Step 1:
	// One country (ex. Norway) => Bordering country alphas (ex. [FIN, RUS, SWE])
	var cntry []consts.Results
	allInfo, _ := http.Get(consts.COVIDTRACKER + WantCountry)
	countrybody, _ := io.ReadAll(allInfo.Body)
	err = json.Unmarshal(countrybody, &cntry)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
		http.Error(w, "Error in response:", http.StatusInternalServerError)
	}

	bord := cntry[0].Borders
	var alphaBody []byte

	// Step 2:
	// Bordering country alphas => Full names (ex. ["Finland", "Russian Federation", "Sweden"])
	tempurl := "https://restcountries.com/v3.1/alpha?codes="
	for i, temp := range bord {
		if i < len(bord)-1 {
			temp += ","
		}
		tempurl += temp
	}

	fullUrl, _ := http.Get(tempurl)
	fullUrlBody, _ := io.ReadAll(fullUrl.Body)
	alphaBody = append(alphaBody, fullUrlBody...)

	var borderingCountries []consts.Results
	err = json.Unmarshal(alphaBody, &borderingCountries)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}

	// Step 3:
	// Full names => Universities in each country, matching (partially) with `WantUni`
	var doneAppending bool = false
	var writeBorderUnis []consts.Results
	for i, _ := range borderingCountries {
		if doneAppending {
			break
		}
		unis := UnisGetByNameAndCountry(WantDate, borderingCountries[i].Name.Common)

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
	unis := UnisGetByNameAndCountry(WantDate, WantCountry)
	writeBorderUnis = append(writeBorderUnis, unis...)

	// Step 5: (last step)
	// Write and return
	writeBody, _ := json.Marshal(writeBorderUnis)
	w.Header().Add("content-type", "application/json")
	w.Write(writeBody)

}
*/