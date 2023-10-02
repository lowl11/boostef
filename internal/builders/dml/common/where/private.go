package where

import "github.com/lowl11/boostef/data/interfaces/iquery"

func (where *Where) add(condition string) iquery.Where {
	where.conditions = append(where.conditions, condition)
	return where
}
