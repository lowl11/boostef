package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"reflect"
	"sync"
)

const (
	defaultPageSize = 10
)

type Repository struct {
	connection *sqlx.DB
	entityType reflect.Type
	pageSize   int

	selectQuery iquery.Select
	insertQuery iquery.Insert
	updateQuery iquery.Update
	deleteQuery iquery.Delete

	predicate iquery.Where

	threadSafe bool
	mutex      sync.Mutex
}

func New(connection *sqlx.DB) *Repository {
	return &Repository{
		connection: connection,
		pageSize:   defaultPageSize,

		threadSafe: true,
	}
}
