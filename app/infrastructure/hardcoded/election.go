package hardcoded

import (
	"time"

	"github.com/clarch/entities/election"
)

type HCElection struct{}

func (hce HCElection) GetByID(ID int64) (election.Model, error) {
	electionDate, _ := time.Parse("2006-01-02", "2019-04-14")

	return election.Model{
		ID:          1,
		Nationality: "Winterfell",
		Name:        "Pemilihan Lurah",
		Date:        electionDate,
	}, nil
}
