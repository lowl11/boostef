package data_types

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/data_types/all"
	"github.com/lowl11/boostef/internal/data_types/mssql"
	"github.com/lowl11/boostef/internal/data_types/mysql"
	"github.com/lowl11/boostef/internal/data_types/postgres"
	"github.com/lowl11/boostef/pkg/enums/sqls"
)

func Boolean() iquery.DataType {
	return createCustom(all.Boolean).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.Boolean
		case sqls.MySQL:
			return mysql.Boolean
		case sqls.MSSQL:
			return mssql.Bit
		default:
			return ""
		}
	})
}

func Bool() iquery.DataType {
	return createCustom(all.Boolean).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.Bool
		case sqls.Postgres:
			return all.Boolean
		case sqls.MSSQL:
			return mssql.Bit
		default:
			return ""
		}
	})
}
