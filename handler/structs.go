package handler

type NameCommon struct {
	Common string `json:"common"`
}

type MapsStreet struct {
	OpenStreetMap string `json:"openStreetMaps"`
}

type University struct {
	Name      string            `json:"name"`
	Country   string            `json:"country,omitempty"` // Suppress field in JSON output if it is empty
	A2C       string            `json:"alpha_two_code"`
	Webpages  []string          `json:"web_pages"`
	Languages map[string]string `json:"languages"` // Not loaded in yet
	Map       MapsStreet        `json:"maps"`      // Not loaded in yet
}

type CountryAndNeighbours struct {
	Name      NameCommon        `json:"name"`
	Country   string            `json:"country,omitempty"` // Suppress field in JSON output if it is empty
	A2C       string            `json:"cca2"`
	A3C       string            `json:"cca3"`
	Webpages  []string          `json:"web_pages"`
	Languages map[string]string `json:"languages"` // Not loaded in yet
	Map       MapsStreet        `json:"maps"`      // Not loaded in yet
	Borders   []string          `json:"borders"`
}
