package ef

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/internal/services/ef_core"
	"time"
)

func Connection() *sqlx.DB {
	return ef_core.Get().Connection()
}

func SetDialect(sql string) {
	ef_core.Get().SetSQL(sql)
}

func SetConnectionString(connectionString string) {
	ef_core.Get().SetConnectionString(connectionString)
}

func SetMaxConnections(maxConnections int) {
	ef_core.Get().SetMaxConnections(maxConnections)
}

func SetMaxIdleConnections(maxIdleConnections int) {
	ef_core.Get().SetMaxIdleConnections(maxIdleConnections)
}

func SetMaxIdleLifetime(lifetime time.Duration) {
	ef_core.Get().SetMaxIdleLifeTime(lifetime)
}

func SetSettings(maxConnections, maxIdleConnections int, idleLifetime time.Duration) {
	ef_core.Get().
		SetMaxConnections(maxConnections).
		SetMaxIdleConnections(maxIdleConnections).
		SetMaxIdleLifeTime(idleLifetime)
}
