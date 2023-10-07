package executor

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

func Exist(connection *sqlx.DB, query string, args ...any) (bool, error) {
	if connection == nil {
		return false, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	rows, err := connection.QueryxContext(ctx, query, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}

func GetStruct[T any](connection *sqlx.DB, query string, args ...any) ([]T, error) {
	if connection == nil {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	rows, err := connection.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]T, 0, 10)
	for rows.Next() {
		itemLink := new(T)
		item := *itemLink
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}
