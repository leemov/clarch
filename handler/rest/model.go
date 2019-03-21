package rest

type JoinElectionCandidate struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	WordingAge string `json:"wording_age"`
}

type JoinElectionData struct {
	ElectionDate        string                  `json:"election_date"`
	Candidates          []JoinElectionCandidate `json:"candidates"`
	ElectionName        string                  `json:"election_name"`
	ElectionNationality string                  `json:"nationality"`
}

type JoinElectionResponse struct {
	Data       JoinElectionData `json:"data"`
	Message    string           `json:"message"`
	StatusCode int              `json:"status_code"`
}
