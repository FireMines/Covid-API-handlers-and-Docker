package covidAss2

type Results struct {
	Name       string  `json:"name"`
	Date       string  `json:"date"`
	Confirmed  float64 `json:"confirmed"`
	Recovered  float64 `json:"recovered"`
	Deaths     float64 `json:"deaths"`
	GrowthRate float64 `json:"growthRate"`
}

type PolicyResults struct {
	Date             string `json:"date"`
	CountryCode      string `json:"country_code"`
	Confirmed        int    `json:"confirmed"`
	Deaths           int    `json:"deaths"`
	StringencyActual int    `json:"stringency_actual"`
	Stringency       int    `json:"stringency"`

	/* 	TypeCode           string `json:"policy_type_code"`
	   	TypeDisplay        string `json:"policy_type_display"`
	   	FlagValue          string `json:"flag_value_display_field"`
	   	PolicyValueDisplay string `json:"policy_value_display_field"`
	   	PolicyValue        int    `json:"policyvalue"`
	   	Flagged            string `json:"flagged"`
	   	Notes              string `json:"notes" */
}
