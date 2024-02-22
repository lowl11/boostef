package insertq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/query"
	"strings"
)

func (builder *Builder) String(_ ...string) string {
	q := strings.Builder{}
	appendInsert(&q, builder.tableName, builder.pairs)
	appendValues(&q, builder.pairs, builder.multiplePairs)
	appendOnConflict(&q, builder.conflict)
	return q.String()
}

func (builder *Builder) GetParamStatus() (string, bool) {
	var isParam bool
	if len(builder.pairs) > 0 {
		isParam = builder.pairs[0].Value == nil
	}
	return builder.String(), isParam
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

func (builder *Builder) Values(pairs ...query.Pair) iquery.Insert {
	builder.multiplePairs = append(builder.multiplePairs, pairs)
	return builder
}
