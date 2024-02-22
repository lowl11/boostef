package create_index

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func (builder *Builder) String(_ ...string) string {
	query := strings.Builder{}

	if builder.unique {
		query.WriteString("CREATE UNIQUE INDEX ")
	} else {
		query.WriteString("CREATE INDEX ")
	}

	if builder.ifNotExist {
		query.WriteString("IF NOT EXISTS ")
	}

	query.WriteString(builder.name)
	query.WriteString("\n")
	query.WriteString("ON ")
	query.WriteString(builder.table)
	query.WriteString(" (")
	for index, column := range builder.columns {
		query.WriteString(column)

		if index < len(builder.columns)-1 {
			query.WriteString(", ")
		}
	}
	query.WriteString(")")

	return query.String()
}

func (builder *Builder) IfNotExist() iquery.CreateIndex {
	builder.ifNotExist = true
	return builder
}

func (builder *Builder) Unique() iquery.CreateIndex {
	builder.unique = true
	return builder
}

func (builder *Builder) Name(name string) iquery.CreateIndex {
	builder.name = name
	return builder
}

func (builder *Builder) TableColumns(tableName string, columns ...string) iquery.CreateIndex {
	builder.table = tableName
	builder.columns = columns
	return builder
}
