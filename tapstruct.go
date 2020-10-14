package tapstruct

// Personteam defines a person or a team, with the info needed to display them to the user
type Personteam struct {
	Email  string `json:"email"` // primary key
	Name   string `json:"name"`
	Abbrev string `json:"abbrev"`
	ColorF string `json:"colorF"`
	ColorB string `json:"colorB"`
}
