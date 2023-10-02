package builder

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/builders/ddl/common/column"
	"github.com/lowl11/boostef/internal/builders/ddl/create_table"
)

func CreateTable(tableName ...string) iquery.CreateTable {
	return create_table.New(tableName...)
}

func Column(name ...string) iquery.Column {
	return column.New(name...)
}
