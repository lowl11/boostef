package executor

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

func Exec(connection *sqlx.DB, query string) error {
	if connection == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := connection.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
