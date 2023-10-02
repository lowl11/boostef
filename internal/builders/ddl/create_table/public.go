package create_table

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func (builder *Builder) Get(_ ...string) string {
	query := strings.Builder{}
	query.Grow(500)
	query.WriteString("CREATE TABLE ")
	if builder.ifNotExist {
		query.WriteString("IF NOT EXISTS ")
	}
	query.WriteString(builder.tableName)
	query.WriteString(" (\n")

	for index, column := range builder.columns {
		query.WriteString("\t")
		query.WriteString(column.Get(builder.sql))
		if index < len(builder.columns)-1 {
			query.WriteString(",\n")
		}
	}

	query.WriteString("\n)")
	return query.String()
}

func (builder *Builder) Table(tableName string) iquery.CreateTable {
	builder.tableName = tableName
	return builder
}

func (builder *Builder) IfNotExist() iquery.CreateTable {
	builder.ifNotExist = true
	return builder
}

func (builder *Builder) Column(columns ...iquery.Column) iquery.CreateTable {
	builder.columns = append(builder.columns, columns...)
	return builder
}

func (builder *Builder) Sql(sql string) iquery.CreateTable {
	builder.sql = sql
	return builder
}
