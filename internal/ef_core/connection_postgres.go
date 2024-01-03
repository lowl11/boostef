package ef_core

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func connectPostgres(connectionString string) (*sqlx.DB, error) {
	pgxConfig, _ := pgx.ParseConfig(connectionString)

	connection, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(pgxConfig))
	if err != nil {
		return nil, err
	}

	return connection, nil
}
