package crud

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/data/interfaces/irepo"
	"github.com/lowl11/boostef/internal/repo"
)

type crud[T any] struct {
	repo      irepo.Repo[T]
	predicate func(iquery.Where)
}

func Inherit[T any]() irepo.Crud[T] {
	return &crud[T]{
		repo: repo.Inherit[T](),
	}
}
