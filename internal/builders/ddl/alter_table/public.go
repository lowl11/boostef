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
	case "ADD":
		query.WriteString(" ")
		builder.dataType.Write(builder.sql, &query)
	case "ALTER COLUMN":
		query.WriteString(" ")

		if builder.isSet {
			query.WriteString("SET ")
			query.WriteString(builder.setAttributes)
		} else if builder.isType {
			query.WriteString("TYPE ")
			query.WriteString(builder.dataType.Name())
		}
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

func (builder *Builder) AddColumn(column string) iquery.AlterTable {
	builder.mode = "ADD"
	builder.column = column
	return builder
}

func (builder *Builder) DropColumn(column string) iquery.AlterTable {
	builder.mode = "DROP COLUMN"
	builder.column = column
	return builder
}

func (builder *Builder) RenameColumn(column, newName string) iquery.AlterTable {
	builder.mode = "RENAME COLUMN"
	builder.column = column
	builder.newName = newName
	return builder
}

func (builder *Builder) AlterColumn(column string) iquery.AlterTable {
	if builder.sql == sqls.MySQL {
		builder.mode = "MODIFY COLUMN"
	} else {
		builder.mode = "ALTER COLUMN"
	}
	builder.column = column
	return builder
}

func (builder *Builder) Set(attributes string) iquery.AlterTable {
	builder.isSet = true
	builder.setAttributes = attributes
	return builder
}

func (builder *Builder) Type(dt iquery.DataType) iquery.AlterTable {
	builder.isType = true
	builder.dataType = dt
	return builder
}

func (builder *Builder) Add() iquery.AlterTable {
	builder.isAdd = true
	return builder
}

func (builder *Builder) Drop() iquery.AlterTable {
	builder.isDrop = true
	return builder
}

func (builder *Builder) Reset() iquery.AlterTable {
	builder.isReset = true
	return builder
}

func (builder *Builder) Restart() iquery.AlterTable {
	builder.isRestart = true
	return builder
}

func (builder *Builder) SQL(sql string) iquery.AlterTable {
	builder.sql = sql
	return builder
}
