package taps

// APIStksPostReq defines the request for a POST to /stakeholders
type APIStksPostReq struct {
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Abbrev  string   `json:"abbrev"`
	ColorF  string   `json:"colorF"`
	ColorB  string   `json:"colorB"`
	Cadence Cadence  `json:"cadence"`
	Pas     []string `json:"parents"`
}

// APIStksGetRes defines the response for a GET to /stakeholders/{email}
type APIStksGetRes struct {
	Stk Stakeholder `json:"stakeholder"`
}

// APIStksDomGetRes defines the response for a GET to /stakeholders/?domain={domain}
type APIStksDomGetRes struct {
	Stks []StkInHier `json:"stakeholders"`
}
