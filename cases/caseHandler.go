package cases

import (
	"bytes"
	"covidAss2/notifications"
	consts "covidAss2/variables"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

//CountriesCalls := make(map[string]int)

// Initialize signature (via init())
var SignatureKey = "X-SIGNATURE"

//var Mac hash.Hash
var Secret []byte

/*
 *	Entry point handler for Location information
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

/**
 *	Handles the get request for cases
 */
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

	if len(consts.CountriesCalls) == 0 {
		consts.CountriesCalls = make(map[string]int)
		fmt.Println("Mamma", consts.CountriesCalls[urlLastVal])
	}

	for i := range consts.CountriesCalls {
		if i == urlLastVal {
			consts.CountriesCalls[urlLastVal] += 1
			fmt.Println("Hei", consts.CountriesCalls[urlLastVal])

		} else {
			consts.CountriesCalls = make(map[string]int)
			fmt.Println("Pappa", consts.CountriesCalls[urlLastVal])

		}
	}

	// Iterate through registered webhooks and invoke based on registered URL, method, and with received content
	for i, v := range notifications.Webhooks {
		if notifications.Webhooks[i].Country == urlLastVal { // If country searched is in the webhooks
			if notifications.Webhooks[i].Calls == int64(consts.CountriesCalls[urlLastVal]) { // If number calls for that webhook occurs
				fmt.Println("Trigger event: Call to service endpoint with method " + http.MethodGet +
					" and content '" + string(writeCountry) + "'.")
				go CallUrl(v.Url, http.MethodGet, string(writeCountry))
			}
		}
	}
	fmt.Println(consts.CountriesCalls[urlLastVal])

}

/*
	Calls given URL with given content and awaits response (status and body).
*/
func CallUrl(url string, method string, content string) {
	fmt.Println("Attempting invocation of url " + url + " with content '" + content + "'.")
	//res, err := http.Post(url, "text/plain", bytes.NewReader([]byte(content)))
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(content)))
	if err != nil {
		log.Printf("%v Error during request creation. Error:", err)
		return
	}
	// Hash content (for content-based validation; not relevant for URL-based validation)
	mac := hmac.New(sha256.New, Secret)
	_, err = mac.Write([]byte(content))
	if err != nil {
		log.Printf("%v Error during content hashing. Error:", err)
		return
	}
	// Convert hash to string & add to header to transport to client for validation
	req.Header.Add(SignatureKey, hex.EncodeToString(mac.Sum(nil)))

	// Perform invocation
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error in HTTP request. Error:", err)
		return
	}

	// Read the response
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Something is wrong with invocation response. Error:", err)
		return
	}

	fmt.Println("Webhook invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(response))
}
