package irepo

import (
	"context"
	"github.com/lowl11/boostef/data/interfaces/iquery"
)

type Session[T any] interface {
	Get(context.Context) ([]T, error)
	Where(func(iquery.Where)) Session[T]
	SetPage(int) Session[T]
}
