package repo

import (
	"github.com/lowl11/boostef/data/interfaces/irepo"
	"github.com/lowl11/boostef/internal/repo"
)

func Inherit[T any]() irepo.Repo[T] {
	return repo.Inherit[T]()
}
