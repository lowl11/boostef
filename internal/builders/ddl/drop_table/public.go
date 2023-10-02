package drop_table

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func (builder *Builder) Get(_ ...string) string {
	query := strings.Builder{}
	query.WriteString("DROP TABLE ")
	query.WriteString(builder.tableName)
	return query.String()
}

func (builder *Builder) Table(tableName string) iquery.DropTable {
	builder.tableName = tableName
	return builder
}
