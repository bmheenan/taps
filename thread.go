package taps

// There are three ways to hold a thread's data

// Threadrow holds contextual info for a thread when it needs to be displayed in a hierarchical table
type Threadrow struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	State    ThreadState `json:"state"`
	Cost     int         `json:"costCtx"` // Different for each stakeholder + total; depends on the context
	Owner    Stakeholder `json:"owner"`
	Iter     string      `json:"iter"` // The base iteration of the thread, not where its appearing
	Ord      int         `json:"ord"`
	Children []Threadrow `json:"children"`
}

// Thread holds info needed for a thread's details page, plus data used for calculating cost, order, and percentile
type Thread struct {
	ID           int64
	Name         string
	Desc         string
	State        string
	CostDir      int
	CostTot      int
	Owner        string
	Iter         string
	Percentile   float64
	Stakeholders map[string](struct {
		Iter string
		Ord  int
	})
	Parents map[int64](struct {
		Iter string
		Ord  int
	})
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
