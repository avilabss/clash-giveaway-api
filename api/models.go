package api

const (
	SUCCESS = "success"
	ERROR   = "error"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type EligibleMembers struct {
	Members    int      `json:"members"`
	MemberList []Member `json:"memberList"`
}

type Member struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
