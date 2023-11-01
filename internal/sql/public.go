package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/log"
)

func (repo Repository) CloseRows(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		log.Error(err, "Closing rows error")
	}
}

func (repo Repository) Transaction(transactionActions func(tx *sqlx.Tx) error) error {
	transaction, err := repo.Connection.Beginx()
	if err != nil {
		return err
	}
	defer rollback(transaction)

	if err = transactionActions(transaction); err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
