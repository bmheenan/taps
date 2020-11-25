package taps

// APIThreadrowsGetRes defines the response for a GET to /threadrows/?stk={stakeholder}&iter={iteration}
type APIThreadrowsGetRes struct {
	Ths []Threadrow `json:"ths"`
}
