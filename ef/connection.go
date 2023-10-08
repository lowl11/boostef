package ef

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/internal/ef_core"
)

func Init(connectionString string) {
	ef_core.Get().SetConnectionString(connectionString)
}

func SetConnection(connection *sqlx.DB) {
	ef_core.Get().SetConnection(connection)
}

func Connection() *sqlx.DB {
	return ef_core.Get().Connection()
}
