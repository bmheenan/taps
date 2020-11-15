package taps

// Stakeholder defines a person or a team of people- an entity who can own or be a stakeholder of threads
type Stakeholder struct {
	Email   string  `json:"email"`
	Domain  string  `json:"domain"`
	Name    string  `json:"name"`
	Abbrev  string  `json:"abbrev"`
	ColorF  string  `json:"colorF"`
	ColorB  string  `json:"colorB"`
	Cadence Cadence `json:"cadence"`
}

// Team holds the Stakeholder info for a team, and an array of all children of that team
type Team struct {
	Stk     Stakeholder `json:"stk"`
	Members []Team      `json:"members"`
}

// Cadence describes the granularity that a stakeholder uses for iterations
type Cadence string

const (
	// Yearly means each year is an iteration. Appropriate for long term objectives of big orgs
	Yearly Cadence = "yearly"
	// Quarterly means each quarter (3 months) is an iteration
	Quarterly Cadence = "quarterly"
	// Monthly means each month is an iteration
	Monthly Cadence = "monthly"
	// Biweekly means every two weeks is an iteration, starting on the first day of the year
	Biweekly Cadence = "biweekly"
)
