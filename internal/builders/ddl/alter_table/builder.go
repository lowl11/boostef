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
	sql       string

	dataType       iquery.DataType
	setAttributes  string
	dropAttributes string

	isSet     bool
	isType    bool
	isReset   bool
	isRestart bool
	isAdd     bool
	isDrop    bool
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
