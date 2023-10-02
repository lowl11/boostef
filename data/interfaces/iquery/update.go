package iquery

import "github.com/lowl11/boostef/pkg/query"

type Update interface {
	Query

	From(tableName string) Update
	Set(pairs ...query.Pair) Update
	Where(func(Where)) Update
}
