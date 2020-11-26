package taps

// APIThreadsPostReq defines the request for a POST to /threads
type APIThreadsPostReq struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Iter  string `json:"iter"`
	Cost  int    `json:"cost"`
	Pas   []int  `json:"parents"`
	Chs   []int  `json:"children"`
}

// APIThreadsPostRes defines the response for a POST to /threads
type APIThreadsPostRes struct {
	ID int `json:"id"`
}
