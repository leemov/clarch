package election

import "time"

//Model for election entity
type Model struct {
	ID          int64
	Name        string
	Nationality string
	Date        time.Time
}
