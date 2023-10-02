package drop_index

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/enums/sqls"
	"strings"
)

func (builder *Builder) Get(_ ...string) string {
	query := strings.Builder{}

	if builder.sql == sqls.MySQL {
		query.WriteString("ALTER TABLE ")
		query.WriteString(builder.table)
		query.WriteString("\n")
	}

	query.WriteString("DROP INDEX ")

	switch builder.sql {
	case sqls.Postgres:
		query.WriteString(builder.table)
		query.WriteString(".")
		query.WriteString("\"")
		query.WriteString(builder.name)
		query.WriteString("\"")
	case sqls.MSSQL:
		query.WriteString(builder.name)
		query.WriteString(" ON ")
		query.WriteString(builder.table)
	case sqls.MySQL:
		query.WriteString(builder.name)
	}

	return query.String()
}

func (builder *Builder) SQL(sql string) iquery.DropIndex {
	builder.sql = sql
	return builder
}

func (builder *Builder) Name(name string) iquery.DropIndex {
	builder.name = name
	return builder
}

func (builder *Builder) Table(tableName string) iquery.DropIndex {
	builder.table = tableName
	return builder
}
