package usecase

import (
	"github.com/clarch/entities/candidate"
	"github.com/clarch/entities/election"
	"github.com/clarch/entities/voter"
	"github.com/clarch/infrastructure/hardcoded"
	"github.com/clarch/infrastructure/postgres"
	state "github.com/clarch/state/joinelection"
)

// In this layer, we will implement usecase

//Business logic :
// 1. Check if election ID exist
// 2. Check if voter ID exist
// 3. Check if election x voter nationality match
// 4. Get All election candidates
// 5. Output

//Entities involved :
// Candidate, Election, Voter

//Exception :
// 1. invalid request election not exist
// 2. invalid request voter not exist
// 3. invalid request voter not allowed
// 4. system error

//JoinElectionInteractor JoinElection Interactor
type JoinElectionInteractor struct {
	candidate candidate.Repository
	election  election.Repository
	voter     voter.Repository
}

//NewJoinElectionInteractor Factory for JoinElection
// func NewJoinElectionInteractor(election postgres.PQElection, candidate hardcoded.HCCandidate, voter hardcoded.HCVoter) JoinElectionInteractor {
// 	return JoinElectionInteractor{
// 		election:  election,
// 		candidate: candidate,
// 		voter:     voter,
// 	}
// }

func NewJoinElectionInteractor(election postgres.PQElection, candidate hardcoded.HCCandidate, voter hardcoded.HCVoter) JoinElectionInteractor {
	return JoinElectionInteractor{
		election:  election,
		candidate: candidate,
		voter:     voter,
	}
}

//JoinElection implementation
func (it JoinElectionInteractor) JoinElection(input JoinElectionRequest) (output JoinElectionResponse) {
	election, err := it.election.GetByID(input.ElectionID)
	if err != nil {
		output.ResultState = state.ErrorGetElection
		return
	}

	if election.ID == 0 {
		output.ResultState = state.NotFoundElection
		return
	}

	voter, err := it.voter.GetByID(input.VoterID)
	if err != nil {
		output.ResultState = state.ErrorGetVoter
		return
	}

	if voter.ID == 0 {
		output.ResultState = state.NotFoundVoter
		return
	}

	//validate Age
	if voter.Age < 17 {
		output.ResultState = state.NotValidAge
		return
	}

	//validate nationality
	if election.Nationality != voter.Nationality {
		output.ResultState = state.NotValidNationality
		return
	}

	candidates, err := it.candidate.GetAllCandidate(election.ID)
	if err != nil {
		output.ResultState = state.ErrorGetAllCandidates
		return
	}

	if len(candidates) == 0 {
		output.ResultState = state.NotFoundCandidates
		return
	}

	return JoinElectionResponse{
		Candidates:  candidates,
		Election:    election,
		ResultState: state.Success,
	}
}
