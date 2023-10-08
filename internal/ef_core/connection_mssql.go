package ef_core

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/microsoft/go-mssqldb"
)

func connectionMSSQL(connectionString string) (*sqlx.DB, error) {
	connection, err := sqlx.Open("odbc", connectionString)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
