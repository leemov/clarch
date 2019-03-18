package usecase

import (
	"github.com/clarch/app/entities/candidate"
	"github.com/clarch/app/entities/election"
)

//ShowElectionResponse response model for this usecase
type ShowElectionResponse struct {
	IsElectionExist bool
	IsUserValid     bool
	Election        election.Model
	Candidates      []candidate.Model
	Error           error
}

//ShowElectionRequest request model for this usecase
type ShowElectionRequest struct {
	ElectionID int64
	VoterID    int64
}

//InputBoundary gateway from delivery mech. to usecase
type InputBoundary interface {
	ShowElection(electionID int64) ShowElectionResponse
}

//OutputBoundary gateway from usecase to delivery mech.
type OutputBoundary interface {
	Present(output ShowElectionResponse)
}
