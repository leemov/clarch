package voter

//External to get voter from the outside
type External interface {
	GetByID(voterID int64) (Model, error)
}
