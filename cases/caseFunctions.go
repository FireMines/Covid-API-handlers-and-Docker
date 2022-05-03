package cases

import (
	outsideapi "covidAss2/outsideApi"
	consts "covidAss2/variables"
)

/**
 *	Gets all the info about covid cases and stores them into a struct
 *
 * @return country - Struct with all info about covid cases
 */
func getCovidCasesPerCountry(countryName string) consts.Results {
	var country consts.Results

	graphqlBody := `
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
	`

	graphqlResponse, err := outsideapi.GetGraphqlResponse(graphqlBody, consts.COVIDGRAPHQL)
	if err != nil {
		return consts.Results{} // Return empty struct if something went wrong
	}

	country = storeData(graphqlResponse)
	return country
}

/**
 *	Stores all data into a struct
 */
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
