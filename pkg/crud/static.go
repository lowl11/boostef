package crud

import (
	"github.com/lowl11/boostef/data/interfaces/irepo"
	"github.com/lowl11/boostef/internal/crud"
)

func Inherit[T any]() irepo.Crud[T] {
	return crud.Inherit[T]()
}
