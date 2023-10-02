package alter_table

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/enums/sqls"
)

type Builder struct {
	tableName string
	mode      string
	column    string
	newName   string
	dataType  iquery.DataType
	sql       string
}

func New(tableName ...string) *Builder {
	builder := &Builder{
		sql: sqls.Postgres,
	}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}
