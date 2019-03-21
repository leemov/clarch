package candidate

import (
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock //Composition of mock struct
}

//GetAllCandidate mock function
func (mc MockRepository) GetAllCandidate(electionID int64) ([]Model, error) {
	args := mc.Called(electionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Model), args.Error(1)
}

//GetByID mock function
func (mc MockRepository) GetByID(candidateID int64) (Model, error) {
	args := mc.Called(candidateID)
	return args.Get(0).(Model), args.Error(1)
}

//Save mock function
func (mc MockRepository) Save(candidate *Model) error {
	args := mc.Called(candidate)
	return args.Error(0)
}

//Update mock function
func (mc MockRepository) Update(candidate *Model) error {
	args := mc.Called(candidate)
	return args.Error(0)
}

//UpdateProperties mock function
func (mc MockRepository) UpdateProperties(ID int64, mapProperties map[string]string) error {
	args := mc.Called(mapProperties)
	return args.Error(0)
}
