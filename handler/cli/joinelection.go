package cli

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"

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
	onAllErrorMessage := "XX Sorry some error happened. please try again..."
	switch usecaseResp.ResultState {
	case
		joinelection.ErrorGetAllCandidates,
		joinelection.ErrorGetElection,
		joinelection.ErrorGetVoter:
		pst.ViewModel.Message = onAllErrorMessage
		pst.ViewModel.IsValid = false
		pst.ViewModel.IsError = true
	case
		joinelection.NotFoundCandidates:
		pst.ViewModel.Message = "XX We cannot find candidates with election specified."
		pst.ViewModel.IsValid = false
		pst.ViewModel.IsError = false
	case
		joinelection.NotFoundElection:
		pst.ViewModel.Message = "XX We cannot find election with id specified."
		pst.ViewModel.IsValid = false
		pst.ViewModel.IsError = false
	case
		joinelection.NotFoundVoter:
		pst.ViewModel.Message = "XX We cannot find You as voter."
		pst.ViewModel.IsValid = false
		pst.ViewModel.IsError = false
	case
		joinelection.NotValidAge:
		pst.ViewModel.Message = "XX Voter should be more than 17."
		pst.ViewModel.IsValid = false
		pst.ViewModel.IsError = false
	case
		joinelection.NotValidNationality:
		pst.ViewModel.Message = "XX Voter nationality not the same as Election Nationality."
		pst.ViewModel.IsValid = false
		pst.ViewModel.IsError = false
	case
		joinelection.Success:
		pst.ViewModel.Message = ">> Yes you can join the election"
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

		pst.ViewModel.Candidates = tmpCandidates
		pst.ViewModel.ElectionDate = usecaseResp.Election.Date.Format("Mon, 02 Jan 2006")
		pst.ViewModel.ElectionName = usecaseResp.Election.Name
		pst.ViewModel.IsValid = true
		pst.ViewModel.IsError = false
		pst.ViewModel.ElectionNationality = usecaseResp.Election.Nationality
	}
}

type CliController struct {
	DB *sqlx.DB //actually this should be interface // so its independent of any framework
}

func (cc CliController) CJoinElection(reqModel JoinElectionRequest) {
	//conversion from controller request model into interactor request model
	//all user input validation goes here too.
	eID, err := strconv.ParseInt(reqModel.ElectionID, 10, 64)
	if err != nil {
		fmt.Println(">> Your election ID is not valid number")
		return
	}

	vID, err := strconv.ParseInt(reqModel.VoterID, 10, 64)
	if err != nil {
		fmt.Println(">> Your voter ID is not valid number")
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
	//render here
	JoinElectionCLIRender(viewModel)
}

func JoinElectionCLIRender(viewModel *JoinElectionResponse) {
	if viewModel.IsError {
		fmt.Println("[ERROR]", viewModel.Message)
	} else if !viewModel.IsValid {
		fmt.Println("[INVALID]", viewModel.Message)
	} else {
		fmt.Println(viewModel.Message)
		fmt.Println("YOU JOINED THE ELECTION : ", viewModel.ElectionName, " at ", viewModel.ElectionDate)
		fmt.Println("Nationality : ", viewModel.ElectionNationality)
		fmt.Println("Select Candidate : ")

		for idx, candidate := range viewModel.Candidates {
			fmt.Println("[ ", idx+1, " ]")
			fmt.Println("Name : ", candidate.Name)
			fmt.Println("Age : ", candidate.Age, "( ", candidate.WordingAge, " )")
		}
	}
}
