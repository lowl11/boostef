package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/ef"
)

type Repository struct {
	Connection *sqlx.DB
}

func Inherit() Repository {
	return Repository{
		Connection: ef.Connection(),
	}
}
