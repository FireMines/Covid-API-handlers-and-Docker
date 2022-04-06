package policy

import (
	consts "covidAss2"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"
)

/*
Entry point handler for policy information
*/
func PolicyHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		policyGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}

}

func policyGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Country you are searching for
	WantCountry := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20") // Gets the first output from path

	scope := r.URL.Query().Get("scope") // Gets the optional limit put on how many to output
	if scope == "" {
		currentTime := time.Now()
		scope = currentTime.Format("2006-01-02")
	}

	fmt.Println(consts.COVIDTRACKER + WantCountry + "/" + scope)
	resp, err := http.Get(consts.COVIDTRACKER + WantCountry + "/" + scope)
	if err != nil {
		http.Error(w, "1God damn it, you goofed up. Now back that badonkadonk up and fix it", http.StatusBadRequest)
		fmt.Println("Decoding1: " + err.Error())
		return
	}

	tester := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&tester)
	if err != nil {
		http.Error(w, "2God damn it, you goofed up. Now back that badonkadonk up and fix it", http.StatusBadRequest)
		fmt.Println("Decoding2: " + err.Error())
		return
	}

	cntry := storePolicyData(tester["stringencyData"].(map[string]interface{}))
	/* 	if cntry.Date == "" || cntry.CountryCode == "" || cntry.Confirmed == 0 || cntry.Deaths == 0 || cntry.StringencyActual == 0 || cntry.Stringency == 0 {
		http.Error(w, "Input did not contain complete policy specification. Recheck posted policy information and resubmit.", http.StatusBadRequest)
		fmt.Println("Empty ID on country:", cntry)

	} */

	defer resp.Body.Close()

	encoder := json.NewEncoder(w)
	err = encoder.Encode(cntry)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

	// Explicit specification of return status code --> will default to 200 if not provided.
	http.Error(w, "", http.StatusOK)

}

func storePolicyData(data map[string]interface{}) consts.PolicyResults {
	if _, ok := data["msg"]; ok {
		return consts.PolicyResults{
			"",
			"",
			-1,
			-1,
			-1,
			-1,
		}
	} else {
		return consts.PolicyResults{
			data["date_value"].(string),
			data["country_code"].(string),
			data["confirmed"].(float64),
			data["deaths"].(float64),
			data["stringency_actual"].(float64),
			data["stringency"].(float64),
		}
	}
}
