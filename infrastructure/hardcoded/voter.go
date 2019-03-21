package hardcoded

import "github.com/clarch/entities/voter"

type HCVoter struct{}

var mapVoter map[int64]voter.Model = map[int64]voter.Model{
	1: voter.Model{
		ID:          1,
		Name:        "Brandon Stark",
		Nationality: "Winterfell",
		Age:         15,
	},
	2: voter.Model{
		ID:          2,
		Name:        "Arya Stark",
		Nationality: "Winterfell",
		Age:         20,
	},
	3: voter.Model{
		ID:          3,
		Name:        "Theon Greyjoy",
		Nationality: "IronIsland",
		Age:         34,
	},
}
var gotResult voter.Model = voter.Model{}

//GetByID implement for harcoded data
func (vhc HCVoter) GetByID(voterID int64) (result voter.Model, err error) {
	// this is harcoded datasource for voter entity
	return mapVoter[voterID], nil
}
