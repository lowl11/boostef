package ef_core

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/internal/helpers/dialect"
	"time"
)

func (ef *EFCore) SetSQL(sql string) *EFCore {
	ef.sql = sql
	return ef
}

func (ef *EFCore) Dialect() string {
	return ef.sql
}

func (ef *EFCore) Connection() *sqlx.DB {
	ef.mutex.Lock()
	defer ef.mutex.Unlock()

	return ef.connection
}

func (ef *EFCore) SetConnection(connection *sqlx.DB) *EFCore {
	ef.mutex.Lock()
	defer ef.mutex.Unlock()

	ef.connection = connection
	return ef
}

func (ef *EFCore) SetConnectionString(connectionString string) *EFCore {
	// connection pool for Postgres
	connectionPool, err := sqlx.Open(dialect.Get(ef.sql), connectionString)
	if err != nil {
		panic(err)
	}

	// setting connection pool configurations
	connectionPool.SetMaxOpenConns(ef.maxConnections)
	connectionPool.SetMaxIdleConns(ef.maxIdleConnections)
	connectionPool.SetConnMaxIdleTime(ef.maxIdleLifetime)

	ef.connection = connectionPool
	return ef
}

func (ef *EFCore) SetMaxConnections(maxConnections int) *EFCore {
	ef.maxConnections = maxConnections
	return ef
}

func (ef *EFCore) SetMaxIdleConnections(maxIdleConnections int) *EFCore {
	ef.maxIdleConnections = maxIdleConnections
	return ef
}

func (ef *EFCore) SetMaxIdleLifeTime(lifetime time.Duration) *EFCore {
	ef.maxIdleLifetime = lifetime
	return ef
}
