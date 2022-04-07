package cases

import (
	"context"
	consts "covidAss2/variables"
	"fmt"

	"github.com/machinebox/graphql"
)

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
