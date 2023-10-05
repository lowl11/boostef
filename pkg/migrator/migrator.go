package migrator

import "github.com/jmoiron/sqlx"

type Migrator struct {
	connection *sqlx.DB
	dialect    string

	givenEntities []entity
	realEntities  []entity
}

func New(connection *sqlx.DB, dialect string) *Migrator {
	return &Migrator{
		connection: connection,
		dialect:    dialect,
	}
}
