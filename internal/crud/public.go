package crud

import (
	"context"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/data/interfaces/irepo"
)

func (c *crud[T]) SetPredicate(predicate func(where iquery.Where)) irepo.Crud[T] {
	c.predicate = predicate
	return c
}

func (c *crud[T]) Exist(ctx context.Context, condition func(where iquery.Where)) (bool, error) {
	results, err := c.repo.All().Where(condition).Get(ctx)
	if err != nil {
		return false, err
	}
	return len(results) > 0, nil
}

func (c *crud[T]) Count(ctx context.Context, condition func(where iquery.Where)) (int, error) {
	return c.repo.Count(ctx, condition)
}

func (c *crud[T]) GetAll(ctx context.Context) ([]T, error) {
	return c.List(ctx, func(where iquery.Where) {})
}

func (c *crud[T]) GetById(ctx context.Context, id string) (*T, error) {
	return c.Single(ctx, func(where iquery.Where) {
		where.Equal("id", id)
	})
}

func (c *crud[T]) GetByIDs(ctx context.Context, ids []string) ([]T, error) {
	anyIDs := make([]any, 0, len(ids))
	for _, id := range ids {
		anyIDs = append(anyIDs, id)
	}

	return c.List(ctx, func(where iquery.Where) {
		where.In("id", anyIDs)
	})
}

func (c *crud[T]) GetPage(ctx context.Context, page int) ([]T, error) {
	current := c.repo.All()
	applyPredicate(current, c.predicate)
	current.SetPage(page)
	return current.Get(ctx)
}

func (c *crud[T]) Single(ctx context.Context, filter func(iquery.Where)) (*T, error) {
	current := c.repo.All()
	applyPredicate(current, c.predicate)

	result, err := current.Where(filter).Get(ctx)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func (c *crud[T]) List(ctx context.Context, filter func(iquery.Where)) ([]T, error) {
	current := c.repo.All()
	applyPredicate(current, c.predicate)
	current.Where(filter)
	return current.Get(ctx)
}

func (c *crud[T]) ListPage(ctx context.Context, filter func(iquery.Where), page int) ([]T, error) {
	current := c.repo.All()
	applyPredicate(current, c.predicate)
	current.Where(filter)
	current.SetPage(page)
	return current.Get(ctx)
}

func (c *crud[T]) Add(ctx context.Context, entity T) error {
	return c.repo.Create(ctx, entity)
}

func (c *crud[T]) Update(ctx context.Context, entity T) error {
	return c.repo.Change(ctx, entity)
}

func (c *crud[T]) Delete(ctx context.Context, entity T) error {
	return c.repo.Remove(ctx, entity)
}

func (c *crud[T]) DeleteBy(ctx context.Context, filter func(iquery.Where)) error {
	return c.repo.RemoveBy(ctx, filter)
}
