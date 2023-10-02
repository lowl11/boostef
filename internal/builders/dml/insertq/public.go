package insertq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/query"
	"strings"
)

func (builder *Builder) Get(_ ...string) string {
	q := strings.Builder{}
	appendInsert(&q, builder.tableName, builder.pairs)
	appendValues(&q, builder.pairs)
	appendOnConflict(&q, builder.conflict)
	return q.String()
}

func (builder *Builder) Pairs(pairs ...query.Pair) iquery.Insert {
	builder.pairs = pairs
	return builder
}

func (builder *Builder) To(tableName string) iquery.Insert {
	builder.tableName = tableName
	return builder
}

func (builder *Builder) OnConflict(query string) iquery.Insert {
	builder.conflict = query
	return builder
}
