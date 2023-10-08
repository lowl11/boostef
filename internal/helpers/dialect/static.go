package dialect

import "github.com/lowl11/boostef/pkg/enums/sqls"

func Get(sql string) string {
	switch sql {
	case sqls.Postgres:
		return "postgres"
	case sqls.MySQL:
		return "mysql"
	case sqls.MSSQL:
		return "odbc"
	}

	return "sql"
}
