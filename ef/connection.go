package ef

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/internal/ef_core"
)

func Init(connectionString string) {
	ef_core.Get().SetConnectionString(connectionString)
}

func SetSQL(sql string) {
	ef_core.Get().SetSQL(sql)
}

func SetConnection(connection *sqlx.DB) {
	ef_core.Get().SetConnection(connection)
}

func Connection() *sqlx.DB {
	conn := ef_core.Get().Connection()
	if conn == nil {
		panic("need to set connection before usage")
	}

	return conn
}

func TryGet() *sqlx.DB {
	return ef_core.Get().Connection()
}
