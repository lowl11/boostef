package updateq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/builders/dml/common/where"
	"github.com/lowl11/boostef/pkg/query"
)

type Builder struct {
	tableName string
	where     iquery.Where
	setPairs  []query.Pair
}

func New(tableName ...string) *Builder {
	builder := &Builder{
		where:    where.New(),
		setPairs: make([]query.Pair, 0),
	}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}
