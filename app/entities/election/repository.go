package election

type Repo interface {
	GetByID(ID int64) (Model, error)
}
