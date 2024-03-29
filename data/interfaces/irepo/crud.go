package irepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/lowl11/boostef/data/interfaces/iquery"
)

type Crud[T any] interface {
	SetPredicate(func(where iquery.Where)) Crud[T]

	Exist(context.Context, func(iquery.Where)) (bool, error)
	Count(context.Context, func(iquery.Where)) (int, error)

	Single(context.Context, func(iquery.Where)) (*T, error)
	List(context.Context, func(iquery.Where)) ([]T, error)
	ListPage(context.Context, func(iquery.Where), int) ([]T, error)

	GetAll(context.Context) ([]T, error)
	GetPage(context.Context, int) ([]T, error)
	GetById(context.Context, uuid.UUID) (*T, error)
	GetByIDs(ctx context.Context, ids []uuid.UUID) ([]T, error)

	Add(context.Context, T) error
	AddList(ctx context.Context, entities []T) error
	Update(context.Context, T) error
	Delete(context.Context, T) error
	DeleteBy(ctx context.Context, filter func(iquery.Where)) error
}
