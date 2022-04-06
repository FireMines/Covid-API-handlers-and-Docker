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
	Date             string  `json:"date_value"`
	CountryCode      string  `json:"country_code"`
	Confirmed        float64 `json:"confirmed"`
	Deaths           float64 `json:"deaths"`
	StringencyActual float64 `json:"stringency_actual"`
	Stringency       float64 `json:"stringency"`

	/* 	TypeCode           string `json:"policy_type_code"`
	   	TypeDisplay        string `json:"policy_type_display"`
	   	FlagValue          string `json:"flag_value_display_field"`
	   	PolicyValueDisplay string `json:"policy_value_display_field"`
	   	PolicyValue        int    `json:"policyvalue"`
	   	Flagged            string `json:"flagged"`
	   	Notes              string `json:"notes" */
}

type WebhookRegistration struct {
	Url         string `json:"url"`
	Country     string `json:"country"`
	Calls       int    `json:"calls"`
	Weebhook_ID string `json:"webhook_id"`
}
