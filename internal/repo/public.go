package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/data/interfaces/irepo"
	"github.com/lowl11/boostef/ef"
	"github.com/lowl11/boostef/internal/session"
	"github.com/lowl11/boostef/internal/transaction"
	"github.com/lowl11/boostef/pkg/builder"
	"github.com/lowl11/flex"
	"reflect"
)

func (r *repo[T]) Count(ctx context.Context, filter func(iquery.Where)) (int, error) {
	selectBuilder := builder.
		Select("COUNT(*)").
		SetAlias(r.aliasName).
		From(r.getTable()).
		Where(filter)

	if len(r.aliasName) > 0 {
		selectBuilder.SetAlias(r.aliasName)
	}

	q := selectBuilder.Get()

	ef.DebugPrint(q)

	tx := transaction.Get(ctx)

	var rows *sqlx.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryxContext(ctx, q)
	} else {
		rows, err = r.connection.QueryxContext(ctx, q)
	}
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, nil
	}

	var count int
	if err = rows.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repo[T]) All(args ...any) irepo.Session[T] {
	selectBuilder := builder.
		Select(r.getColumns()...).
		From(r.getTable())

	if len(r.aliasName) > 0 {
		selectBuilder.SetAlias(r.aliasName)
	}

	return session.New[T](r.connection, selectBuilder, args...)
}

func (r *repo[T]) Create(ctx context.Context, entity T) error {
	q, isParam := builder.
		Insert(r.getPairs(entity)...).
		To(r.getTable()).
		GetParamStatus()

	ef.DebugPrint(q)

	tx := transaction.Get(ctx)

	var err error
	if isParam {
		var statement *sqlx.NamedStmt

		if tx != nil {
			statement, err = tx.PrepareNamedContext(ctx, q)
		} else {
			statement, err = r.connection.PrepareNamedContext(ctx, q)
		}
		if err != nil {
			return err
		}

		_, err = statement.ExecContext(ctx, entity)
	} else {
		if tx != nil {
			_, err = tx.ExecContext(ctx, q)
		} else {
			_, err = r.connection.ExecContext(ctx, q)
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *repo[T]) Change(ctx context.Context, entity T) error {
	baseEntity := flex.
		Struct(entity).
		FieldByType(reflect.TypeOf(ef.Entity{})).(ef.Entity)

	q, isParam := builder.
		Update(r.getTable()).
		Set(r.getPairs(entity)...).
		Where(func(where iquery.Where) {
			where.Equal("id", baseEntity.ID)
		}).GetParam()

	ef.DebugPrint(q)

	tx := transaction.Get(ctx)

	if isParam {
		var statement *sqlx.NamedStmt
		var err error
		if tx != nil {
			statement, err = r.connection.PrepareNamedContext(ctx, q)
		} else {
			statement, err = r.connection.PrepareNamedContext(ctx, q)
		}
		if err != nil {
			return err
		}

		_, err = statement.ExecContext(ctx, entity)
		if err != nil {
			return err
		}
	} else {
		var err error
		if tx != nil {
			_, err = tx.ExecContext(ctx, q)
		} else {
			_, err = r.connection.ExecContext(ctx, q)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repo[T]) Remove(ctx context.Context, entity T) error {
	baseEntity := flex.
		Struct(entity).
		FieldByType(reflect.TypeOf(ef.Entity{})).(ef.Entity)

	q := builder.
		Delete(r.getTable()).
		Where(func(where iquery.Where) {
			where.Equal("id", "$1")
		}).Get()

	ef.DebugPrint(q)

	tx := transaction.Get(ctx)

	var statement *sqlx.Stmt
	var err error
	if tx != nil {
		statement, err = tx.PreparexContext(ctx, q)
	} else {
		statement, err = r.connection.PreparexContext(ctx, q)
	}
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, baseEntity.ID.String())
	if err != nil {
		return err
	}

	return nil
}

func (r *repo[T]) RemoveBy(ctx context.Context, where func(iquery.Where)) error {
	q := builder.
		Delete(r.getTable()).
		Where(where).
		Get()

	ef.DebugPrint(q)

	tx := transaction.Get(ctx)

	var statement *sqlx.Stmt
	var err error
	if tx != nil {
		statement, err = tx.PreparexContext(ctx, q)
	} else {
		statement, err = r.connection.PreparexContext(ctx, q)
	}
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
