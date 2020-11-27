package taps

// APIThreadrowsGetRes defines the response for a GET to /threadrows/?stk={stakeholder}&iter={iteration},
// /threadrows/?parent={parentid}&iter={iteration}, or /threadrows/?child={childid}
type APIThreadrowsGetRes struct {
	Ths []Threadrow `json:"ths"`
}
