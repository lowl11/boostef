package ef_core

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func connectionMy(connectionString string) (*sqlx.DB, error) {
	connection, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
