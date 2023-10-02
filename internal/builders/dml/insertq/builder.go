package insertq

import "github.com/lowl11/boostef/pkg/query"

type Builder struct {
	tableName string
	pairs     []query.Pair
	conflict  string
}

func New(pairs ...query.Pair) *Builder {
	return &Builder{
		pairs: pairs,
	}
}
