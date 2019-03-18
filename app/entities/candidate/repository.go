package candidate

//Repository methods for candidate entity
type Repository interface {
	GetAllCandidate(electionID int64) ([]Model, error)
	GetByID(candidateID int64) (Model, error)
	Save(candidate *Model) error
	Update(candidate *Model) error
	UpdateProperties(mapProperties map[string]string) error
}
