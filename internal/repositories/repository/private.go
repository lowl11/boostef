package repository

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/builder"
)

func (repo *Repository) lock() {
	if !repo.threadSafe {
		return
	}

	repo.mutex.Lock()
}

func (repo *Repository) unlock() {
	if !repo.threadSafe {
		return
	}

	repo.mutex.Unlock()
}

func (repo *Repository) getSelectQuery() iquery.Select {
	repo.lock()
	defer repo.unlock()

	if repo.selectQuery != nil {
		return repo.selectQuery
	}

	repo.selectQuery = builder.Select()
	return repo.selectQuery
}
