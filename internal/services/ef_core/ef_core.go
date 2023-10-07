package ef_core

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/pkg/enums/sqls"
	"sync"
	"time"
)

type EFCore struct {
	sql                string
	connection         *sqlx.DB
	maxConnections     int
	maxIdleConnections int
	maxIdleLifetime    time.Duration
	schema             string

	mutex sync.Mutex
}

var instance *EFCore

func Get() *EFCore {
	if instance != nil {
		return instance
	}

	instance = &EFCore{
		sql:    sqls.Postgres,
		schema: "public",
	}
	return instance
}
