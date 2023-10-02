package deleteq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/builders/dml/common/where"
)

type Builder struct {
	tableName string
	where     iquery.Where
}

func New(tableName ...string) *Builder {
	builder := &Builder{
		where: where.New(),
	}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}
