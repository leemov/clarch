package voter

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

//GetByID mock function
func (mr MockRepository) GetByID(voterID int64) (Model, error) {
	args := mr.Called(voterID)
	return args.Get(0).(Model), args.Error(1)
}
