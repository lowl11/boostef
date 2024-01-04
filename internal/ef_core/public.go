package ef_core

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/internal/helpers/dialect"
	"github.com/lowl11/boostef/pkg/enums/sqls"
	"regexp"
	"strings"
	"time"
)

/*
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
*/

func (core *Core) SetConnection(connection *sqlx.DB) *Core {
	core.connection = connection
	return core
}

func (core *Core) SetConnectionString(connectionString string) *Core {
	var connection *sqlx.DB
	var err error

	switch core.sql {
	case sqls.Postgres:
		connection, err = connectPostgres(connectionString)
	case sqls.MySQL:
		connection, err = connectionMy(connectionString)
	case sqls.MSSQL:
		connection, err = connectionMSSQL(connectionString)
	default:
		panic("Unknown type of database: " + core.sql)
	}
	if err != nil {
		panic(err)
	}

	dbSchema := "public"
	searchPath := regexp.MustCompile("search_path=(.*)").FindAllString(connectionString, -1)
	if len(searchPath) > 0 {
		_, after, found := strings.Cut(searchPath[0], "=")
		if found {
			dbSchema = after
		} else {
			dbSchema = searchPath[0]
		}
	}

	core.SetSchema(dbSchema)

	// setting connection pool configurations
	connection.SetMaxOpenConns(core.maxConnections)
	connection.SetMaxIdleConns(core.maxIdleConnections)
	connection.SetConnMaxIdleTime(core.maxIdleLifetime)

	core.connection = connection
	return core
}

func (core *Core) Connection() *sqlx.DB {
	return core.connection
}

func (core *Core) SetSQL(sql string) *Core {
	core.sql = sql
	return core.SetDialect(dialect.Get(sql))
}

func (core *Core) SetSchema(schema string) *Core {
	core.schema = schema
	return core
}

func (core *Core) Schema() string {
	return core.schema
}

func (core *Core) SQL() string {
	return core.sql
}

func (core *Core) SetDialect(dialect string) *Core {
	core.dialect = dialect
	return core
}

func (core *Core) Dialect() string {
	return core.dialect
}

func (core *Core) SetupConnection(maxConnections, maxIdleConnections int, idleLifetime, lifetime time.Duration) *Core {
	return core.
		SetMaxConnections(maxConnections).
		SetMaxIdleConnections(maxIdleConnections).
		SetMaxIdleLifetime(idleLifetime).
		SeMaxLifetime(lifetime)
}

func (core *Core) SetMaxConnections(maxConnections int) *Core {
	core.maxConnections = maxConnections
	return core
}

func (core *Core) SetMaxIdleConnections(maxIdleConnections int) *Core {
	core.maxIdleConnections = maxIdleConnections
	return core
}

func (core *Core) SetMaxIdleLifetime(lifetime time.Duration) *Core {
	core.maxIdleLifetime = lifetime
	return core
}

func (core *Core) SeMaxLifetime(lifetime time.Duration) *Core {
	core.maxLifetime = lifetime
	return core
}
