package covidAss2

type Results struct {
	Name       string  `json:"name"`
	Date       string  `json:"date,"`
	Confirmed  float64 `json:"confirmed"`
	Recovered  float64 `json:"recovered"`
	Deaths     float64 `json:"deaths"`
	GrowthRate float64 `json:"growthRate"`
}
