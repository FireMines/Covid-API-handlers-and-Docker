package policy

import (
	consts "covidAss2/variables"
)

/**
 *	Stores data into its respective struct
 */
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
		// Check if the 'stringency_actual' field is nil, if it is: use the 'stringency' field instead.
		stringency := data["stringency_actual"].(float64)
		if data["stringency_actual"] == nil {
			stringency = data["stringency"].(float64)
		}

		return consts.PolicyResults{
			data["date_value"].(string),
			data["country_code"].(string),
			data["confirmed"].(float64),
			data["deaths"].(float64),
			stringency,
			0,
		}
	}
}
