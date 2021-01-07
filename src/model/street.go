package model

type Street struct {
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	Suburb            string       `json:"suburb"`
	SuburbDescription string       `json:"suburbDescription"`
	Type              string       `json:"type"`
	TypeShort         string       `json:"typeShort"`
	Districts         []string     `json:"districts"`
	AllNames          []StreetName `json:"allNames"`
}

type StreetName struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Obsolete    bool   `json:"obsolete"`
}
