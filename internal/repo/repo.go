package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/data/interfaces/irepo"
	"github.com/lowl11/boostef/internal/ef_core"
	"reflect"
)

type repo[T any] struct {
	connection *sqlx.DB
	entity     any
	tableName  string
	fields     []reflect.StructField
	columns    []string
}

func Inherit[T any]() irepo.Repo[T] {
	r := &repo[T]{
		connection: ef_core.Get().Connection(),
	}
	r.eatEntity(*new(T))
	return r
}
