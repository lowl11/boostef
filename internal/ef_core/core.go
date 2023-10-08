package ef_core

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/pkg/enums/sqls"
	"time"
)

type Core struct {
	connection *sqlx.DB
	sql        string
	dialect    string

	maxConnections     int
	maxIdleConnections int
	maxIdleLifetime    time.Duration
}

var instance *Core

func Get() *Core {
	if instance != nil {
		return instance
	}

	instance = &Core{
		sql: sqls.Postgres,

		maxConnections:     10,
		maxIdleConnections: 10,
		maxIdleLifetime:    time.Minute * 5,
	}
	return instance
}
