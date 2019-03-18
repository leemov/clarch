package vote

//Model for vote / save choice of a voter
type Model struct {
	ID          int64
	VoterID     int64
	CandidateID int64
	ElectionID  int64
}
