package transaction

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func Get(ctx context.Context) *sqlx.Tx {
	if ctx == nil {
		return nil
	}

	txValue := ctx.Value("boostef_transaction")
	if txValue == nil {
		return nil
	}

	return txValue.(*sqlx.Tx)
}
