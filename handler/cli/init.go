package cli

import "github.com/jmoiron/sqlx"

type CliController struct {
	DB *sqlx.DB //actually this should be interface // so its independent of any framework
}
