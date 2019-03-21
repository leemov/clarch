package hardcoded

import (
	"github.com/clarch/entities/candidate"
)

var mapCandidates map[int64]candidate.Model = map[int64]candidate.Model{
	1: candidate.Model{
		Age:        24,
		ID:         1,
		Name:       "Sansa Stark",
		ElectionID: 1,
	},
	2: candidate.Model{
		Age:        100,
		ID:         2,
		Name:       "Jon Snow",
		ElectionID: 1,
	},
}

type HCCandidate struct{}

func (hhc HCCandidate) GetAllCandidate(electionID int64) (result []candidate.Model, err error) {
	if len(mapCandidates) == 0 {
		return nil, nil
	}

	result = []candidate.Model{}
	for _, val := range mapCandidates {

		if electionID == val.ElectionID {
			result = append(result, val)
		}
	}

	return result, err
}

func (hhc HCCandidate) GetByID(candidateID int64) (candidate.Model, error) {
	return mapCandidates[candidateID], nil
}

func (hhc HCCandidate) Save(candidate *candidate.Model) error {
	return nil
}

func (hhc HCCandidate) Update(candidate *candidate.Model) error {
	return nil
}

func (hhc HCCandidate) UpdateProperties(ID int64, mapProperties map[string]string) error {
	return nil
}
