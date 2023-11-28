package irepo

import (
	"context"
	"github.com/lowl11/boostef/data/interfaces/iquery"
)

type Repo[T any] interface {
	Count(context.Context, func(iquery.Where)) (int, error)
	All(args ...any) Session[T]
	Create(context.Context, T) error
	Change(context.Context, T) error
	Remove(context.Context, T) error
}
