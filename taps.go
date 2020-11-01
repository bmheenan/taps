package taps

// Personteam defines a person or a team, with the info needed to display them to the user
type Personteam struct {
	Email      string     `json:"email"` // primary key
	Domain     string     `json:"domain"`
	Name       string     `json:"name"`
	Abbrev     string     `json:"abbrev"`
	ColorF     string     `json:"colorF"`
	ColorB     string     `json:"colorB"`
	IterTiming IterTiming `json:"iterTiming"`
}

// IterTiming describes the granularity of iterations, allowing different people and teams to specify their cadence
type IterTiming string

const (
	// Yearly means each year is an iteration. Appropriate for long term objectives of big orgs
	Yearly IterTiming = "yearly"
	// Quarterly means each quarter (3 months) is an iteration
	Quarterly IterTiming = "quarterly"
	// Monthly means each month is an iteration
	Monthly IterTiming = "monthly"
	// Biweekly means every two weeks is an iteration, starting on the first day of the year
	Biweekly IterTiming = "biweekly"
)

// Threadrow holds the summary info for a thread, needed to display it in a table with other threads
type Threadrow struct {
	ID         int64       `json:"id"`
	Domain     string      `json:"domain"`
	Name       string      `json:"name"`
	State      ThreadState `json:"state"`
	CostCtx    int         `json:"costCtx"` // The work owned by the personteam being displayed to complete this thread
	Owner      Personteam  `json:"owner"`
	Iteration  string      `json:"iteration"`
	Order      int         `json:"order"`
	Percentile float32     `json:"percentile"`
	Children   []Threadrow `json:"children"`
}

// Threaddetail holds detailed information on a thread, needed for the drilldown view
type Threaddetail struct {
	ID           int64        `json:"id"`
	Domain       string       `json:"domain"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	State        ThreadState  `json:"state"`
	CostDirect   int          `json:"costDirect"`
	CostTotal    int          `json:"costTotal"`
	Owner        Personteam   `json:"owner"`
	Stakeholders []Personteam `json:"stakeholders"`
	Iteration    string       `json:"iteration"`
	Order        int          `json:"order"`
	Percentile   float32      `json:"percentile"`
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

// Threadrel tracks just the thread data the backend uses for calculating cost, order, and percentile
type Threadrel struct {
	ID               int64
	State            string
	CostDirect       int
	Owner            string
	Iteration        string
	Order            int
	Percentile       float64
	StakeholderMatch bool // States if the thread has a given stakeholder, specified in a query
}
