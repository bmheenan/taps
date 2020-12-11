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
	Name string `json:"name"`
	Desc struct {
		New   bool   `json:"new"`
		Value string `json:"value"`
	} `json:"desc"`
	Cost struct {
		New   bool `json:"new"`
		Value int  `json:"value"`
	} `json:"cost"`
	Owner      string   `json:"owner"`
	State      State    `json:"state"`
	AddStks    []string `json:"addStks"`
	RmStks     []string `json:"removeStks"`
	AddParents []int64  `json:"addParents"`
	RmParents  []int64  `json:"removeParents"`
	Iter       string   `json:"iter"`
	Ord        struct {
		Pa  int64  `json:"parent"`
		Stk string `json:"stk"`
		Val int64  `json:"value"`
	} `json:"ord"`
}
