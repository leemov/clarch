package election

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

//GetByID mock function
func (me MockRepository) GetByID(ID int64) (Model, error) {
	args := me.Called(ID)
	return args.Get(0).(Model), args.Error(1)
}
