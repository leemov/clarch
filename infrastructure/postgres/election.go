package postgres

import (
	"database/sql"
	"log"

	"github.com/lib/pq"

	"github.com/clarch/entities/election"
	"github.com/jmoiron/sqlx"
)

type PQElection struct {
	//actually should be interface so it could be loosely coupled
	DB *sqlx.DB
}

type ElectionDB struct {
	ID           int64          `db:"id"`
	Name         sql.NullString `db:"name"`
	ElectionTime pq.NullTime    `db:"election_time"`
	District     string         `db:"district"`
}

func (pqe PQElection) GetByID(ID int64) (election.Model, error) {
	query := `SELECT id, name, election_time, district FROM elections WHERE id=:id`
	args := map[string]interface{}{
		"id": ID,
	}

	query, namedArgs, err := sqlx.Named(query, args)
	if err != nil {
		log.Printf("There is an error while election.GetByID sqlx.Named err : %s", err.Error())
		return election.Model{}, err
	}

	query = pqe.DB.Rebind(query)
	var electionDB ElectionDB
	err = pqe.DB.Get(&electionDB, query, namedArgs...)
	if err != nil {
		log.Printf("There is an error while election.GetByID err : %s", err.Error())
		return election.Model{}, err
	}

	//mapping to entity

	//broken data, entity need valid time but data not
	if !electionDB.ElectionTime.Valid {
		log.Println("There is an error while election.GetByID data broken election_time not valid")
		return election.Model{}, err
	}

	//broken data, entity need valid string but data not
	if !electionDB.Name.Valid {
		log.Println("There is an error while election.GetByID data broken name not valid")
		return election.Model{}, err
	}

	return election.Model{
		Date:        electionDB.ElectionTime.Time,
		Name:        electionDB.Name.String,
		ID:          electionDB.ID,
		Nationality: electionDB.District,
	}, nil
}
