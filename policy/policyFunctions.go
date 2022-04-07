package policy

import (
	consts "covidAss2/variables"
)

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
