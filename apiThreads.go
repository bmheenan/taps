package taps

// APIThreadsPostReq defines the request for a POST to /threads
type APIThreadsPostReq struct {
	Name  string  `json:"name"`
	Owner string  `json:"owner"`
	Iter  string  `json:"iter"`
	Cost  int     `json:"cost"`
	Pas   []int64 `json:"parents"`
	Chs   []int64 `json:"children"`
}

// APIThreadsPostRes defines the response for a POST to /threads
type APIThreadsPostRes struct {
	ID int64 `json:"id"`
}

// APIThreadsGetRes defines the response for a GET to /threads/{id}
type APIThreadsGetRes struct {
	Th Thread `json:"thread"`
}

// APIThreadsPutReq defines the request for a PUT to /threads
type APIThreadsPutReq struct {
	Ths        []int64 `json:"threads"`
	AddParents []int64 `json:"addParents"`
	RmParents  []int64 `json:"removeParents"`
	Iter       string  `json:"iter"`
	OrdBefore  int64   `json:"ordBefore"`
	OrdEnd     bool    `json:"ordEnd"`
}
