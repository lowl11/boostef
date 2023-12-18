package ef

import (
	"context"
	"github.com/lowl11/boostef/internal/transaction"
	"strings"
)

func BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := Connection().Beginx()
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, "boostef_transaction", tx), nil
}

func CloseTransaction(ctx context.Context, errors ...error) error {
	if ctx == nil {
		return nil
	}

	tx := transaction.Get(ctx)
	if tx == nil {
		return nil
	}

	var uplevelError error
	if len(errors) > 0 {
		uplevelError = errors[0]
	}

	if uplevelError != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	if err := tx.Rollback(); err != nil {
		if strings.Contains(err.Error(), "transaction has already been committed") {
			return nil
		}

		return err
	}

	return nil
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

func CommitTransaction(ctx context.Context) error {
	tx := transaction.Get(ctx)
	if tx == nil {
		return nil
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
