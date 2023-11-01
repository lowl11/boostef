package sql

import "github.com/lowl11/boostef/internal/sql"

type Repository = sql.Repository

func Inherit() sql.Repository {
	return sql.Inherit()
}
