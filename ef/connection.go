package ef

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/internal/ef_core"
	"time"
)

func Init(connectionString string) {
	ef_core.Get().SetConnectionString(connectionString)
}

func Close() error {
	connection := TryGet()
	if connection == nil {
		return nil
	}

	return connection.Close()
}

func SetSQL(sql string) {
	ef_core.Get().SetSQL(sql)
}

func SetupLimits(maxConnections, maxIdleConnections int, idleLifetime, lifetime time.Duration) {
	ef_core.Get().SetupConnection(maxConnections, maxIdleConnections, idleLifetime, lifetime)
}

func SetConnection(connection *sqlx.DB) {
	ef_core.Get().SetConnection(connection)
}

func Connection() *sqlx.DB {
	conn := TryGet()
	if conn == nil {
		panic("need to set connection before usage")
	}

	return conn
}

func TryGet() *sqlx.DB {
	return ef_core.Get().Connection()
}
