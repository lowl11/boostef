package ef

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := Connection().Beginx()
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, "boostef_transaction", tx), nil
}

func CloseTransaction(ctx context.Context) error {
	if ctx == nil {
		return nil
	}

	txValue := ctx.Value("boostef_transaction")
	if txValue == nil {
		return nil
	}

	tx := txValue.(*sqlx.Tx)
	if err := tx.Commit(); err != nil {
		return err
	}

	if err := tx.Rollback(); err != nil {
		return err
	}
	
	return nil
}
