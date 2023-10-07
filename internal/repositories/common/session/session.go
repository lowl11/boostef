package session

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/data/interfaces/iquery"
)

const (
	defaultPageSize = 10
)

type Session struct {
	connection *sqlx.DB

	pageSize int

	query     iquery.Query
	predicate iquery.Where
}

func New(connection *sqlx.DB, query iquery.Query) *Session {
	return &Session{
		connection: connection,
		query:      query,
		pageSize:   defaultPageSize,
	}
}
