package updateq

import (
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/query"
	"strings"
)

func (builder *Builder) Get(_ ...string) string {
	queryBuilder := strings.Builder{}
	appendUpdate(&queryBuilder, builder.tableName)
	appendSet(&queryBuilder, builder.setPairs)
	appendWhere(&queryBuilder, builder.where)
	return queryBuilder.String()
}

func (builder *Builder) GetParam() (string, bool) {
	var isParam bool
	if len(builder.setPairs) > 0 {
		isParam = types.ToString(builder.setPairs[0].Value) == ""
	}
	return builder.Get(), isParam
}

func (builder *Builder) From(tableName string) iquery.Update {
	builder.tableName = tableName
	return builder
}

func (builder *Builder) Set(pairs ...query.Pair) iquery.Update {
	builder.setPairs = pairs
	return builder
}

func (builder *Builder) Where(whereFunc func(builder iquery.Where)) iquery.Update {
	return builder.applyWhere(whereFunc)
}
