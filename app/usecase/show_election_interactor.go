package usecase

import (
	"github.com/clarch/app/entities/candidate"
	"github.com/clarch/app/entities/election"
	"github.com/clarch/app/entities/voter"
)

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

//ShowElectionInteractor ShowElection Interactor
type ShowElectionInteractor struct {
	candidate candidate.Repository
	election  election.Repository
	voter     voter.External
}

//NewShowElectionInteractor Factory for ShowElection
func NewShowElectionInteractor() *ShowElectionInteractor {
	return &ShowElectionInteractor{}
}

//Init to avoid dependency confusion we pass exactly details on here
//Its the hacky way.. just to help another developer to understand well
func (it *ShowElectionInteractor) Init() {
	//why we need pointer ? this is the reason !!
}

//ShowElection implementation
func (it *ShowElectionInteractor) ShowElection(input ShowElectionRequest) (output ShowElectionResponse) {
	election, err := it.election.GetByID(input.ElectionID)
	if err != nil {
		output.IsError = true
		output.IsValid = false
		output.Message = "There is an error when getting election"
		output.Error = err
		return
	}

	if election.ID == 0 {
		output.IsError = false
		output.IsValid = false
		output.Message = "Election not found"
		output.Error = nil
		return
	}

	voter, err := it.voter.GetByID(input.VoterID)
	if err != nil {
		output.IsError = true
		output.IsValid = false
		output.Message = "There is an error when getting voter"
		output.Error = err
		return
	}

	if voter.ID == 0 {
		output.IsError = false
		output.IsValid = false
		output.Message = "Voter not found"
		output.Error = nil
		return
	}

	//validate nationality
	if election.Nationality != voter.Nationality {
		output.IsError = false
		output.IsValid = false
		output.Message = "Voter not eligible for election"
		output.Error = nil
		return
	}

	candidates, err := it.candidate.GetAllCandidate(election.ID)
	if err != nil {
		output.IsError = true
		output.IsValid = false
		output.Message = "There is an error when getting election candidates"
		output.Error = err
		return
	}

	if len(candidates) == 0 {
		output.IsError = false
		output.IsValid = false
		output.Message = "Candidates for election not found"
		output.Error = nil
		return
	}

	return ShowElectionResponse{
		Candidates: candidates,
		Election:   election,
		IsValid:    true,
		Message:    "Election found!",
	}
}
