package election

type Repository interface {
	GetByID(ID int64) (Model, error)
}
