package ef_core

import (
	"github.com/jmoiron/sqlx"
	
	_ "github.com/lib/pq"
)

func connectPostgres(connectionString string) (*sqlx.DB, error) {
	connection, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
