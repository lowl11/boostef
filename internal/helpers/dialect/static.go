package dialect

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/microsoft/go-mssqldb"

	"github.com/lowl11/boostef/pkg/enums/sqls"
)

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
