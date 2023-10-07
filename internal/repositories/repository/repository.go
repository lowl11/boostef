package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	defaultPageSize = 10
)

type Repository struct {
	connection *sqlx.DB
	pageSize   int
}

func New(connection *sqlx.DB) *Repository {
	return &Repository{
		connection: connection,
		pageSize:   defaultPageSize,
	}
}
