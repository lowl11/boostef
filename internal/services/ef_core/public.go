package ef_core

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/internal/helpers/dialect"
	"regexp"
	"time"
)

func (ef *EFCore) SetSQL(sql string) *EFCore {
	ef.sql = sql
	return ef
}

func (ef *EFCore) Dialect() string {
	return ef.sql
}

func (ef *EFCore) Schema() string {
	return ef.schema
}

func (ef *EFCore) SetSchema(schema string) *EFCore {
	ef.schema = schema
	return ef
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

	// cut schema
	var schema string
	r, _ := regexp.Compile("SearchPath=(.*?)")
	match := r.FindAllString(connectionString, -1)
	if len(match) > 1 {
		schema = match[0]
	}

	if len(schema) > 0 {
		ef.schema = schema
	}
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
