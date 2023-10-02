package builder

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/builders/ddl/alter_table"
	"github.com/lowl11/boostef/internal/builders/ddl/common/column"
	"github.com/lowl11/boostef/internal/builders/ddl/create_index"
	"github.com/lowl11/boostef/internal/builders/ddl/create_table"
	"github.com/lowl11/boostef/internal/builders/ddl/drop_index"
	"github.com/lowl11/boostef/internal/builders/ddl/drop_table"
	"github.com/lowl11/boostef/internal/builders/ddl/truncate_table"
)

func CreateTable(tableName ...string) iquery.CreateTable {
	return create_table.New(tableName...)
}

func Column(name ...string) iquery.Column {
	return column.New(name...)
}

func DropTable(tableName ...string) iquery.DropTable {
	return drop_table.New(tableName...)
}

func TruncateTable(tableName ...string) iquery.TruncateTable {
	return truncate_table.New(tableName...)
}

func AlterTable(tableName ...string) iquery.AlterTable {
	return alter_table.New(tableName...)
}

func CreateIndex(name ...string) iquery.CreateIndex {
	return create_index.New(name...)
}

func DropIndex(name ...string) iquery.DropIndex {
	return drop_index.New(name...)
}
