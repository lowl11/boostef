package truncate_table

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func (builder *Builder) String(_ ...string) string {
	query := strings.Builder{}
	query.WriteString("TRUNCATE TABLE ")
	query.WriteString(builder.tableName)
	return query.String()
}

func (builder *Builder) Table(tableName string) iquery.TruncateTable {
	builder.tableName = tableName
	return builder
}
