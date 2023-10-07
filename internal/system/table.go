package system

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/system/vars"
	"github.com/lowl11/boostef/pkg/builder"
	"github.com/lowl11/boostef/pkg/enums/sqls"
)

func GetTable(sql, tableValue string) string {
	tableName := vars.TableName(sql)

	return builder.
		Select(tableName).
		From(vars.Tables(sql)).
		Where(func(where iquery.Where) {
			where.Equal(tableName, tableValue)
		}).
		Get()
}

func GetColumns(sql, tableValue string) string {
	if sql == sqls.MySQL {
		return "SHOW COLUMNS FROM " + tableValue
	}

	return builder.
		Select(vars.Columns(sql)...).
		From(vars.ColumnTable(sql)).
		Where(func(where iquery.Where) {
			where.Equal(vars.TableName(sql), tableValue)
		}).
		Get()
}
