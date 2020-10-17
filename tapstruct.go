package tapstruct

// Personteam defines a person or a team, with the info needed to display them to the user
type Personteam struct {
	Email       string       `json:"email"` // primary key
	Domain      string       `json:"domain"`
	Name        string       `json:"name"`
	Abbrev      string       `json:"abbrev"`
	ColorF      string       `json:"colorF"`
	ColorB      string       `json:"colorB"`
	HasChildren bool         `json:"hasChildren"`
	Children    []Personteam `json:"childrenDetails"` // Can be empty even if HasChildren is true
}

// Threadrow holds the summary info for a thread, needed to display it in a table with other threads
type Threadrow struct {
	ID          int         `json:"id"`
	Domain      string      `json:"domain"`
	Name        string      `json:"name"`
	State       ThreadState `json:"state"`
	Cost        int         `json:"cost"`
	Owner       Personteam  `json:"owner"`
	Iteration   string      `json:"iteration"`
	HasChildren bool        `json:"hasChildren"`
	Children    []Threadrow `json:"children"` // Can be empty even if HasChildren is true
}

// Threaddetail holds detailed information on a thread, needed for the drilldown view
type Threaddetail struct {
	ID           int          `json:"id"`
	Domain       string       `json:"domain"`
	Name         string       `json:"name"`
	State        ThreadState  `json:"state"`
	CostDirect   int          `json:"costDirect"`
	CostTotal    int          `json:"costTotal"`
	Owner        Personteam   `json:"owner"`
	Stakeholders []Personteam `json:"stakeholders"`
	Iteration    string       `json:"iteration"`
	// Children and parents aren't required; we can use functions that query using ID as the parent/child to get them
}

// ThreadState describes the possible states for a thread
type ThreadState string

const (
	// NotStarted is the default state; it hasn't been worked on (yet)
	NotStarted ThreadState = "not started"
	// InProgress is for when a thread is being worked on, but is not complete
	InProgress ThreadState = "in progress"
	// Done is for when the value has been delivered to the stakeholders
	Done ThreadState = "done"
	// Closed is appropriate for threads working as intended, unactionable, or duplicates
	Closed ThreadState = "closed"
	// Archived is for threads that were valid, but low enough priority that they were never addressed
	Archived ThreadState = "archived"
)
