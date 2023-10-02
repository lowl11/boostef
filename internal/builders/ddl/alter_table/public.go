package alter_table

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/enums/sqls"
	"strings"
)

func (builder *Builder) Get(_ ...string) string {
	query := strings.Builder{}
	query.WriteString("ALTER TABLE ")
	query.WriteString(builder.tableName)
	query.WriteString("\n")
	query.WriteString(builder.mode)
	query.WriteString(" ")
	query.WriteString(builder.column)

	switch builder.mode {
	case "ADD", "ALTER COLUMN":
		query.WriteString(" ")
		builder.dataType.Write(builder.sql, &query)
	case "RENAME COLUMN":
		query.WriteString(" TO ")
		query.WriteString(builder.newName)
	}

	return query.String()
}

func (builder *Builder) Table(tableName string) iquery.AlterTable {
	builder.tableName = tableName
	return builder
}

func (builder *Builder) Add(column string) iquery.AlterTable {
	builder.mode = "ADD"
	builder.column = column
	return builder
}

func (builder *Builder) Drop(column string) iquery.AlterTable {
	builder.mode = "DROP COLUMN"
	builder.column = column
	return builder
}

func (builder *Builder) Rename(column, newName string) iquery.AlterTable {
	builder.mode = "RENAME COLUMN"
	builder.column = column
	builder.newName = newName
	return builder
}

func (builder *Builder) Alter(column string) iquery.AlterTable {
	if builder.sql == sqls.MySQL {
		builder.mode = "MODIFY COLUMN"
	} else {
		builder.mode = "ALTER COLUMN"
	}
	builder.column = column
	return builder
}

func (builder *Builder) SQL(sql string) iquery.AlterTable {
	builder.sql = sql
	return builder
}

func (builder *Builder) DataType(dataType iquery.DataType) iquery.AlterTable {
	builder.dataType = dataType
	return builder
}
