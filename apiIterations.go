package taps

// APIItersGetRes defines the response for a GET to /iterations/?stk={stakeholder} or /iterations/?parent={parentid}
type APIItersGetRes struct {
	Iters []string `json:"iters"`
}
