package taps

// Stakeholder defines a person or a team of people- an entity who can own or be a stakeholder of threads
type Stakeholder struct {
	Email   string      `json:"email"`
	Domain  string      `json:"domain"`
	Name    string      `json:"name"`
	Abbrev  string      `json:"abbrev"`
	ColorF  string      `json:"colorF"`
	ColorB  string      `json:"colorB"`
	Cadence IterCadence `json:"iterCadence"`
}

// IterCadence describes the granularity that a stakeholder uses for iterations
type IterCadence string

const (
	// Yearly means each year is an iteration. Appropriate for long term objectives of big orgs
	Yearly IterCadence = "yearly"
	// Quarterly means each quarter (3 months) is an iteration
	Quarterly IterCadence = "quarterly"
	// Monthly means each month is an iteration
	Monthly IterCadence = "monthly"
	// Biweekly means every two weeks is an iteration, starting on the first day of the year
	Biweekly IterCadence = "biweekly"
)