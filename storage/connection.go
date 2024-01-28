package storage

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPool(builder *Builder) (*sqlx.DB, error) {
	pgxConfig, _ := pgx.ParseConfig(builder.connectionString)

	connection, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(pgxConfig))
	if err != nil {
		return nil, err
	}

	connection.SetMaxOpenConns(builder.maxConnections)
	connection.SetMaxIdleConns(builder.maxIdleConnections)

	connection.SetConnMaxLifetime(builder.maxLifetime)
	connection.SetConnMaxIdleTime(builder.maxIdleLifetime)

	return connection, nil
}
