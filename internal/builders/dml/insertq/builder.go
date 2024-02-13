package insertq

import "github.com/lowl11/boostef/pkg/query"

type Builder struct {
	tableName     string
	pairs         []query.Pair
	conflict      string
	multiplePairs [][]query.Pair
}

func New(pairs ...query.Pair) *Builder {
	return &Builder{
		pairs:         pairs,
		multiplePairs: make([][]query.Pair, 0),
	}
}
