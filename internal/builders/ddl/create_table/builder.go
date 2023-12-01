package create_table

import "github.com/lowl11/boostef/data/interfaces/iquery"

type Builder struct {
	tableName        string
	columns          []iquery.Column
	ifNotExist       bool
	partitionColumns []string

	sql string
}

func New(tableName ...string) *Builder {
	builder := &Builder{
		columns:          make([]iquery.Column, 0, 10),
		partitionColumns: make([]string, 0),
	}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}
