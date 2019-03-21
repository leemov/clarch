package rest

import "github.com/jmoiron/sqlx"

type RestController struct {
	DB *sqlx.DB //actually this should be interface // so its independent of any framework
}
