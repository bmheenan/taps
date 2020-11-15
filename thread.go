package taps

// There are two ways to hold a thread's data

// Threadrow holds contextual info for a thread when it needs to be displayed in a hierarchical table
type Threadrow struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	State    State       `json:"state"`
	Cost     int         `json:"cost"` // Different for each stakeholder + total; depends on the context
	Owner    Stakeholder `json:"owner"`
	Iter     string      `json:"iter"` // The base iteration of the thread, not where its appearing
	Ord      int         `json:"ord"`
	Children []Threadrow `json:"children"`
}

// Thread holds info needed for a thread's details page, plus data used for calculating cost, order, and percentile
type Thread struct {
	ID         int64
	Name       string
	Desc       string
	State      State
	CostDir    int
	CostTot    int
	Owner      Stakeholder
	Iter       string
	Percentile float64
	Stks       map[string](struct {
		Iter string
		Ord  int
		Cost int
	})
	Parents map[int64](struct {
		Iter string
		Ord  int
	})
}

// State describes the possible states of a thread
type State string

const (
	// NotStarted is the default state; it hasn't been worked on (yet)
	NotStarted State = "not started"
	// InProgress is for when a thread is being worked on, but is not complete
	InProgress State = "in progress"
	// Done is for when the value has been delivered to the stakeholders
	Done State = "done"
	// Closed is appropriate for threads working as intended, unactionable, or duplicates
	Closed State = "closed"
	// Archived is for threads that were valid, but low enough priority that they were never addressed
	Archived State = "archived"
)
