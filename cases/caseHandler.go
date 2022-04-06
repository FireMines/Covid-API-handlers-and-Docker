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
	r.Header.Add("content-type", "application/json")

	// Get country matching the name in the URL
	countryInfo := getCovidCasesPerCountry(urlLastVal)

	// Marshal them and write to Writer
	writeUnis, err := json.Marshal(countryInfo)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
		http.Error(w, "Error in response:", http.StatusInternalServerError)
	}

	w.Header().Add("content-type", "application/json")
	w.Write(writeUnis)

}

func getCovidCasesPerCountry(countryName string) consts.Results {
	var country consts.Results
	graphqlClient := graphql.NewClient(consts.COVIDGRAPHQL)
	graphqlRequest := graphql.NewRequest(`
	query {
		# by country
		country(name:"` + countryName + `") {
		  name
		  mostRecent {
			date(format: "yyyy-MM-dd")
			confirmed
			recovered
			deaths
			growthRate
		  }
		}
	  } 
	`)
	var graphqlResponse map[string]interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
		//fmt.Println("Error in response:", err.Error())
		//http.Error(w, "Error in response:", http.StatusInternalServerError)
	}

	fmt.Println(graphqlResponse)

	country = storeData(graphqlResponse)

	return country
}

func storeData(data map[string]interface{}) consts.Results {
	data = (data["country"].(map[string]interface{}))
	allData := data["mostRecent"].(map[string]interface{})

	return consts.Results{
		data["name"].(string),
		allData["date"].(string),
		allData["confirmed"].(float64),
		allData["deaths"].(float64),
		allData["recovered"].(float64),
		allData["growthRate"].(float64),
	}
}
