package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/log"
	"strings"
)

func rollback(transaction *sqlx.Tx) {
	if err := transaction.Rollback(); err != nil {
		if !strings.Contains(
			err.Error(),
			"sql: transaction has already been committed or rolled back",
		) {
			log.Error(err, "Rollback transaction error")
		}
	}
}
