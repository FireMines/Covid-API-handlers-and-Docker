package cases

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	consts "covidAss2"

	"github.com/machinebox/graphql"
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

	// Get country matching the name in the URL
	countryInfo := getCovidCasesPerCountry(urlLastVal, "3/10/2020")

	// Marshal them and write to Writer
	writeUnis, _ := json.Marshal(countryInfo)
	w.Header().Add("content-type", "application/json")
	w.Write(writeUnis)

}

func getCovidCasesPerCountry(countryName string, date string) []consts.Results {
	var country []consts.Results

	//url := COVIDGRAPHQL

	graphqlClient := graphql.NewClient(consts.RESOURCE_ROOT_PATH + consts.COVIDCASES)
	graphqlRequest := graphql.NewRequest(`
	query {
		# time series data
		results (countries: [` + countryName + ` ], date: { lt: ` + date + ` }) {
		  country {
			name
		  }
		  date
		  confirmed
		  deaths
		  recovered
		  growthRate
		}
	  
		# by country
		country(name: "Sweden") {
		  name
		  mostRecent {
			date(format: "yyyy-MM-dd")
			confirmed
		  }
		}
	  } 
	`)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	fmt.Println(graphqlResponse)

	return country
}
