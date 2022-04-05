package status

import (
	consts "covidAss2"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
Entry point handler for Location information
*/
func StatusHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		statusHandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}

func statusHandleGetRequest(w http.ResponseWriter, r *http.Request) {

	// Instantiate the client
	client := &http.Client{}

	// Issue request
	res, err := client.Get(consts.COVIDTRACKER)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}

	res2, err := client.Get(consts.COVIDGRAPHQL)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}

	t := time.Now()
	elapsed := t.Sub(consts.Start)

	statusInfo := map[string]interface{}{
		"cases_api":  res.StatusCode,
		"policy_api": res2.StatusCode, // Getting 400 Error, fix this before delivery
		//"webhooks":		,
		"version": res.TLS.Version,
		"uptime":  elapsed.Seconds(),
	}

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err = encoder.Encode(statusInfo)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
	}
}
