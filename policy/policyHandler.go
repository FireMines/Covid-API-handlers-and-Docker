package policy

import (
	outsideapi "covidAss2/outsideApi"
	consts "covidAss2/variables"
	"encoding/json"
	"net/http"
	"path"
	"strings"
	"time"
)

/*
 *	Entry point handler for policy information
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

/**
 *	Handles the get request for policy
 */
func policyGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Country you are searching for
	WantCountry := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20") // Gets the first output from path

	if WantCountry == "policy" {
		http.Error(w, "No country added", http.StatusUnprocessableEntity)
		return
	}
	scope := r.URL.Query().Get("scope") // Gets the optional limit put on how many to output
	if scope == "" {
		currentTime := time.Now()
		scope = currentTime.Format("2006-01-02")
	}

	dataRaw, err := outsideapi.GetHttpResponse(consts.COVIDTRACKER + WantCountry + "/" + scope)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Stores data into struct
	if dataRaw["stringencyData"] == nil {
		http.Error(w, "Could not find any data for the given country and date", http.StatusInternalServerError)
		return
	}
	cntry := storePolicyData(dataRaw["stringencyData"].(map[string]interface{}))
	policies := dataRaw["policyActions"].([]interface{})

	validPolicies := 0

	// Checks if policies does not exists, if it exists, make policies counter go up
	for _, j0 := range policies {
		j := j0.(map[string]interface{})
		if j["policy_type_code"].(string) != "NONE" {
			validPolicies++
		}
	}
	cntry.Policy = validPolicies

	// Encode
	encoder := json.NewEncoder(w)
	err = encoder.Encode(cntry)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

	// Explicit specification of return status code --> will default to 200 if not provided.
	http.Error(w, "", http.StatusOK)
}
