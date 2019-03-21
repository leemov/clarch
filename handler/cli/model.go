package cli

type JoinElectionCandidate struct {
	Name       string
	Age        int
	WordingAge string
}

type JoinElectionResponse struct {
	ElectionDate        string
	Candidates          []JoinElectionCandidate
	ElectionName        string
	ElectionNationality string
	IsError             bool
	IsValid             bool
	Message             string
}

type JoinElectionRequest struct {
	VoterID    string
	ElectionID string
}
