package migrator

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/services/executor"
	"github.com/lowl11/boostef/pkg/builder"
)

type Migrator struct {
	connection *sqlx.DB
	dialect    string
	keys       []tableKey

	givenEntities []entity
	realEntities  []entity
}

func New(connection *sqlx.DB, dialect, schema string) *Migrator {
	keys, err := executor.GetStruct[tableKey](connection, builder.
		Select("constraint_name", "table_name", "column_name").
		From("information_schema.key_column_usage").
		Where(func(where iquery.Where) {
			where.Equal("table_schema", schema)
		}).
		Get())
	if err != nil {
		panic(err)
	}

	return &Migrator{
		connection: connection,
		dialect:    dialect,
		keys:       keys,
	}
}
