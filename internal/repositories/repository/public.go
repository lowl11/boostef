package repository

import (
	"github.com/lowl11/boostef/data/interfaces/irepository"
	"github.com/lowl11/boostef/pkg/builder"
)

func (repo *Repository) Select(columns ...string) irepository.Session {
	return repo.getSession(builder.Select(columns...))
}
