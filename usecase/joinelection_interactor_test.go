package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/clarch/entities/voter"

	"github.com/clarch/entities/election"
	state "github.com/clarch/state/joinelection"
)

func TestJoinElectionInteractor_JoinElection(t *testing.T) {
	type args struct {
		input JoinElectionRequest
	}
	tests := []struct {
		name       string
		args       args
		wantOutput JoinElectionResponse
		initMock   func() JoinElectionInteractor
	}{
		{
			name: "error election.GetByID",
			args: args{
				input: JoinElectionRequest{
					ElectionID: 1,
					VoterID:    2,
				},
			},
			initMock: func() JoinElectionInteractor {
				me := election.MockRepository{}
				me.On("GetByID", int64(1)).Return(election.Model{}, errors.New("Some error happened"))
				return JoinElectionInteractor{
					election: me,
				}
			},
			wantOutput: JoinElectionResponse{
				ResultState: state.ErrorGetElection,
			},
		},
		{
			name: "not Found election.GetByID",
			args: args{
				input: JoinElectionRequest{
					ElectionID: 1,
					VoterID:    2,
				},
			},
			initMock: func() JoinElectionInteractor {
				me := election.MockRepository{}
				me.On("GetByID", int64(1)).Return(election.Model{}, nil)
				return JoinElectionInteractor{
					election: me,
				}
			},
			wantOutput: JoinElectionResponse{
				ResultState: state.NotFoundElection,
				Election:    election.Model{},
			},
		},
		{
			name: "error voter.GetByID",
			args: args{
				input: JoinElectionRequest{
					ElectionID: 1,
					VoterID:    999,
				},
			},
			initMock: func() JoinElectionInteractor {
				me := election.MockRepository{}
				etime, _ := time.Parse("2006-01-02", "2019-04-29")
				me.On("GetByID", int64(1)).Return(election.Model{
					ID:          1,
					Name:        "Bob",
					Nationality: "ID",
					Date:        etime,
				}, nil)

				mv := voter.MockRepository{}
				mv.On("GetByID", int64(999)).Return(voter.Model{}, errors.New("Some error happened"))
				return JoinElectionInteractor{
					election: me,
					voter:    mv,
				}
			},
			wantOutput: JoinElectionResponse{
				ResultState: state.ErrorGetVoter,
			},
		},
		{
			name: "not Found voter.GetByID",
			args: args{
				input: JoinElectionRequest{
					ElectionID: 1,
					VoterID:    999,
				},
			},
			initMock: func() JoinElectionInteractor {
				me := election.MockRepository{}
				etime, _ := time.Parse("2006-01-02", "2019-04-29")
				me.On("GetByID", int64(1)).Return(election.Model{
					ID:          1,
					Name:        "Bob",
					Nationality: "ID",
					Date:        etime,
				}, nil)

				mv := voter.MockRepository{}
				mv.On("GetByID", int64(999)).Return(voter.Model{}, nil)
				return JoinElectionInteractor{
					election: me,
					voter:    mv,
				}
			},
			wantOutput: JoinElectionResponse{
				ResultState: state.ErrorGetVoter,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		it := tt.initMock()
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := it.JoinElection(tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("JoinElectionInteractor.JoinElection() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
