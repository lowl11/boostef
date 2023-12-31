package ef

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boostef/internal/transaction"
	"strings"
)

func BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := Connection().BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, "boostef_transaction", tx), nil
}

func MustBeginTransaction(ctx context.Context) context.Context {
	newCtx, err := BeginTransaction(ctx)
	if err != nil {
		return ctx
	}

	return newCtx
}

func RollbackTransaction(ctx context.Context) error {
	tx := transaction.Get(ctx)
	if tx == nil {
		return nil
	}

	if err := tx.Rollback(); err != nil {
		if strings.Contains(err.Error(), "transaction has already been committed") {
			return nil
		}

		return err
	}

	return nil
}

func MustRollbackTransaction(ctx context.Context) {
	if err := RollbackTransaction(ctx); err != nil {
		log.Error(err, "Rollback transaction error")
	}
}

func CommitTransaction(ctx context.Context) error {
	tx := transaction.Get(ctx)
	if tx == nil {
		return nil
	}

	if err := tx.Commit(); err != nil {
		if strings.Contains(err.Error(), "transaction has already been committed") {
			return nil
		}

		return err
	}

	return nil
}

func MustCommitTransaction(ctx context.Context) {
	if err := CommitTransaction(ctx); err != nil {
		log.Error(err, "Commit transaction error")
	}
}

func GetTransaction(ctx context.Context) *sqlx.Tx {
	return transaction.Get(ctx)
}
