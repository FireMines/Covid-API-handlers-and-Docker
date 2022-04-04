package covidAss2

/* type NameCommon struct {
	Common string `json:"common"`
} */

type Country struct {
	Name string `json:"name"`
}

type Results struct {
	Country    Country `json:"country"` // Not loaded in yet
	Date       string  `json:"date,"`
	Confirmed  int     `json:"confirmed"`
	Deaths     int     `json:"deaths"`
	Recovered  int     `json:"recovered"`
	GrowthRate float64 `json:"growthRate"`
	//Recovered map[string]string `json:"recovered"` // Not loaded in yet
}

/* type CountryAndNeighbours struct {
	Name      NameCommon        `json:"name"`
	Country   string            `json:"country,omitempty"` // Suppress field in JSON output if it is empty
	A2C       string            `json:"cca2"`
	A3C       string            `json:"cca3"`
	Webpages  []string          `json:"web_pages"`
	Languages map[string]string `json:"languages"` // Not loaded in yet
	Map       MapsStreet        `json:"maps"`      // Not loaded in yet
	Borders   []string          `json:"borders"`
}
*/
