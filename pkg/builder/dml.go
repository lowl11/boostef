package builder

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/builders/dml/deleteq"
	"github.com/lowl11/boostef/internal/builders/dml/insertq"
	"github.com/lowl11/boostef/internal/builders/dml/selectq"
	"github.com/lowl11/boostef/internal/builders/dml/updateq"
	"github.com/lowl11/boostef/pkg/query"
)

func Select(columns ...string) iquery.Select {
	return selectq.New(columns...)
}

func Delete(tableName ...string) iquery.Delete {
	return deleteq.New(tableName...)
}

func Update(tableName ...string) iquery.Update {
	return updateq.New(tableName...)
}

func Insert(pairs ...query.Pair) iquery.Insert {
	return insertq.New(pairs...)
}
