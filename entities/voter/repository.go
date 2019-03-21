package voter

//Repository to get voter
type Repository interface {
	GetByID(voterID int64) (Model, error)
}
