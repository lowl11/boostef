package crud

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/data/interfaces/irepo"
)

func applyPredicate[T any](session irepo.Session[T], predicate func(iquery.Where)) {
	if predicate == nil {
		return
	}

	session.Where(predicate)
}
