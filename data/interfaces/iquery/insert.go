package iquery

import "github.com/lowl11/boostef/pkg/query"

type Insert interface {
	Query

	GetParamStatus() (string, bool)
	To(tableName string) Insert
	OnConflict(query string) Insert
	Values(pairs ...query.Pair) Insert
}
