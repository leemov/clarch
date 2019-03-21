package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/clarch/infrastructure/hardcoded"
	"github.com/clarch/state/joinelection"

	"github.com/clarch/usecase"
)

// any dependency for presenter injected here
type JoinElectionPresenter struct {
	ViewModel *JoinElectionResponse
}

func (pst JoinElectionPresenter) Present(usecaseResp usecase.JoinElectionResponse) {
	// you can render what you want here. you can even get aditional data from external resource here (with external dependency injected)
	// you can format your view model need here
	onAllErrorMessage := "Some Internal Error happened."
	switch usecaseResp.ResultState {
	case
		joinelection.ErrorGetAllCandidates,
		joinelection.ErrorGetElection,
		joinelection.ErrorGetVoter:
		pst.ViewModel.Message = onAllErrorMessage
		pst.ViewModel.StatusCode = http.StatusInternalServerError
	case
		joinelection.NotFoundCandidates:
		pst.ViewModel.Message = "Candidates for election not found."
		pst.ViewModel.StatusCode = http.StatusNotFound
	case
		joinelection.NotFoundElection:
		pst.ViewModel.Message = "Election data not found."
		pst.ViewModel.StatusCode = http.StatusNotFound
	case
		joinelection.NotFoundVoter:
		pst.ViewModel.Message = "Voter data not found."
		pst.ViewModel.StatusCode = http.StatusNotFound
	case
		joinelection.NotValidAge:
		pst.ViewModel.Message = "Voter is not old enough data not found."
		pst.ViewModel.StatusCode = http.StatusBadRequest
	case
		joinelection.NotValidNationality:
		pst.ViewModel.Message = "Voter is not eligible to join specified Election."
		pst.ViewModel.StatusCode = http.StatusBadRequest
	case
		joinelection.Success:
		pst.ViewModel.Message = "Voter eligible to join Election"
		pst.ViewModel.StatusCode = http.StatusOK
		tmpCandidates := make([]JoinElectionCandidate, 0, len(usecaseResp.Candidates))
		for _, val := range usecaseResp.Candidates {
			jec := JoinElectionCandidate{}
			if val.Age >= 100 {
				jec.WordingAge = "Possibly Deadman ?"
			} else if val.Age < 100 && val.Age > 40 {
				jec.WordingAge = "Cukup Umur"
			} else {
				jec.WordingAge = "Noobie"
			}
			jec.Name = val.Name
			jec.Age = val.Age
			tmpCandidates = append(tmpCandidates, jec)
		}

		pst.ViewModel.Data.Candidates = tmpCandidates
		pst.ViewModel.Data.ElectionDate = usecaseResp.Election.Date.Format("Mon, 02 Jan 2006")
		pst.ViewModel.Data.ElectionName = usecaseResp.Election.Name
		pst.ViewModel.Data.ElectionNationality = usecaseResp.Election.Nationality
	}
}

func (cc RestController) RESTJoinElection(w http.ResponseWriter, r *http.Request) {
	//conversion from controller request model into interactor request model
	//all user input validation goes here too.

	//getting data from GET request params
	sUserID := r.URL.Query().Get("user_id")
	sElectionID := r.URL.Query().Get("election_id")

	var HTTPresponse JoinElectionResponse

	defer func() {
		jsonResponse, _ := json.Marshal(HTTPresponse)
		w.WriteHeader(HTTPresponse.StatusCode)
		w.Write(jsonResponse)
	}()
	//validation should happens here
	eID, err := strconv.ParseInt(sElectionID, 10, 64)
	if err != nil {
		HTTPresponse.StatusCode = http.StatusBadRequest
		HTTPresponse.Message = "Invalid electionID"
		return
	}

	vID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		HTTPresponse.StatusCode = http.StatusBadRequest
		HTTPresponse.Message = "Invalid voterID"
		return
	}

	hcVoter := hardcoded.HCVoter{}
	hcCandidate := hardcoded.HCCandidate{}
	hcElection := hardcoded.HCElection{}
	// pqElection := postgres.PQElection{
	// 	DB: cc.DB,
	// }
	ucase := usecase.NewJoinElectionInteractor(hcElection, hcCandidate, hcVoter)
	output := ucase.JoinElection(usecase.JoinElectionRequest{
		ElectionID: eID,
		VoterID:    vID,
	})

	viewModel := &JoinElectionResponse{}
	jePres := JoinElectionPresenter{
		ViewModel: viewModel,
	}

	jePres.Present(output)
	HTTPresponse = *jePres.ViewModel
	return
}
