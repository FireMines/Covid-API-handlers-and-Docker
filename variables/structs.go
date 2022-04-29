package consts

/**
 *	Struct for the results we got in the cases
 */
type Results struct {
	Name       string  `json:"name"`
	Date       string  `json:"date"`
	Confirmed  float64 `json:"confirmed"`
	Recovered  float64 `json:"recovered"`
	Deaths     float64 `json:"deaths"`
	GrowthRate float64 `json:"growthRate"`
}

/**
 *	Struct for the polict results we got in the policies
 */
type PolicyResults struct {
	Date             string  `json:"date_value"`
	CountryCode      string  `json:"country_code"`
	Confirmed        float64 `json:"confirmed"`
	Deaths           float64 `json:"deaths"`
	StringencyActual float64 `json:"stringency_actual"`
	Stringency       float64 `json:"stringency"`
	Policy           int     `json:"policy"`
}

/**
 *	Handles the webhook data in a struct
 */
type WebhookRegistration struct {
	Weebhook_ID string `json:"webhook_id"`
	Url         string `json:"url"`
	Country     string `json:"country"`
	Calls       int64  `json:"calls"`
}
