package repository

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/data/interfaces/irepository"
	"github.com/lowl11/boostef/internal/repositories/common/session"
)

func (repo *Repository) getSession(query iquery.Query) irepository.Session {
	return session.
		New(repo.connection, query).
		SetPageSize(repo.pageSize)
}
