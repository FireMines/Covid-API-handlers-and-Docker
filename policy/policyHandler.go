package policy

import (
	consts "covidAss2/variables"
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
		http.Error(w, "Fault while getting covidtracker and country", http.StatusBadRequest)
		fmt.Println("Decoding1: " + err.Error())
		return
	}

	dataRaw := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&dataRaw)
	if err != nil {
		http.Error(w, "Fault while decoding", http.StatusBadRequest)
		fmt.Println("Decoding2: " + err.Error())
		return
	}

	cntry := storePolicyData(dataRaw["stringencyData"].(map[string]interface{}))
	policies := dataRaw["policyActions"].([]interface{})

	validPolicies := 0

	for _, j0 := range policies {
		j := j0.(map[string]interface{})
		if j["policy_type_code"].(string) != "NONE" {
			validPolicies++
		}
	}
	cntry.Policy = validPolicies
	fmt.Println(len(policies))

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
