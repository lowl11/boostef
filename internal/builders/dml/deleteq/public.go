package deleteq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func (builder *Builder) String(_ ...string) string {
	// builder
	query := strings.Builder{}
	query.Grow(300)

	// delete
	appendDelete(&query, builder.tableName)

	// where
	appendWhere(&query, builder.where)

	return query.String()
}

func (builder *Builder) From(tableName string) iquery.Delete {
	builder.tableName = tableName
	return builder
}

func (builder *Builder) Where(whereFunc func(builder iquery.Where)) iquery.Delete {
	return builder.applyWhere(whereFunc)
}
