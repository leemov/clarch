package usecase

import (
	"github.com/clarch/entities/candidate"
	"github.com/clarch/entities/election"
)

// In this layer, should be only a contract of what function to be implemented and what request/response model will be

//JoinElectionResponse response model for this usecase
type JoinElectionResponse struct {
	Election    election.Model
	Candidates  []candidate.Model
	ResultState int
}

//JoinElectionRequest request model for this usecase
type JoinElectionRequest struct {
	ElectionID int64
	VoterID    int64
}

//JoinElectionIBoundary gateway from delivery mech. to usecase
type JoinElectionIBoundary interface {
	JoinElection(input JoinElectionRequest) JoinElectionResponse
}

//JoinElectionOBoundary gateway from usecase to delivery mech.
type JoinElectionOBoundary interface {
	Present(output JoinElectionResponse)
}
