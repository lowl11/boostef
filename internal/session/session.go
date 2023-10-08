package session

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/data/interfaces/iquery"
)

const (
	defaultPageSize = 10
)

type Session[T any] struct {
	connection *sqlx.DB
	q          iquery.Select
	pageSize   int
}

func New[T any](connection *sqlx.DB, q iquery.Select) *Session[T] {
	return &Session[T]{
		connection: connection,
		q:          q,
		pageSize:   defaultPageSize,
	}
}
