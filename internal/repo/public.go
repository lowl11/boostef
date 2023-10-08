package repo

import (
	"context"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/data/interfaces/irepo"
	"github.com/lowl11/boostef/internal/session"
	"github.com/lowl11/boostef/pkg/builder"
	"github.com/lowl11/boostef/pkg/query"
	"github.com/lowl11/flex"
	"reflect"
)

func (r *repo[T]) Count(ctx context.Context, filter func(iquery.Where)) (int, error) {
	rows, err := r.connection.QueryxContext(ctx, builder.
		Select("COUNT(*)").
		From(r.getTable()).
		Where(filter).
		Get())
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

func (r *repo[T]) All() irepo.Session[T] {
	return session.New[T](r.connection, builder.
		Select(r.getColumns()...).
		From(r.getTable()),
	)
}

func (r *repo[T]) Create(ctx context.Context, entity T) error {
	q := builder.
		Insert(r.getPairs(entity)...).
		To(r.getTable())

	_, err := r.connection.ExecContext(ctx, q.Get())
	if err != nil {
		return err
	}

	return nil
}

func (r *repo[T]) Change(ctx context.Context, entity T) error {
	baseEntity := flex.
		Struct(entity).
		FieldByType(reflect.TypeOf(query.Entity{})).(query.Entity)

	q := builder.
		Update(r.getTable()).
		Set(r.getPairs(entity)...).
		Where(func(where iquery.Where) {
			where.Equal("id", baseEntity.ID)
		})

	_, err := r.connection.ExecContext(ctx, q.Get())
	if err != nil {
		return err
	}

	return nil
}

func (r *repo[T]) Remove(ctx context.Context, entity T) error {
	baseEntity := flex.
		Struct(entity).
		FieldByType(reflect.TypeOf(query.Entity{})).(query.Entity)

	q := builder.
		Delete(r.getTable()).
		Where(func(where iquery.Where) {
			where.Equal("id", baseEntity.ID)
		})

	_, err := r.connection.ExecContext(ctx, q.Get())
	if err != nil {
		return err
	}

	return nil
}
