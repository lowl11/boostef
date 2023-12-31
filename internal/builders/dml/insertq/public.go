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

func (builder *Builder) GetParamStatus() (string, bool) {
	var isParam bool
	if len(builder.pairs) > 0 {
		isParam = builder.pairs[0].Value == nil
	}
	return builder.Get(), isParam
}

func (builder *Builder) Pairs(pairs ...query.Pair) iquery.Insert {
	if len(pairs) == 0 {
		return builder
	}

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
