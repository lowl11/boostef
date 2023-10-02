package selectq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/builders/dml/common/aggregate"
	"github.com/lowl11/boostef/internal/builders/dml/common/where"
)

type Builder struct {
	columns         []string
	tableName       string
	aliasName       string
	where           iquery.Where
	orderByColumns  []string
	isDescending    bool
	havingAggregate iquery.Aggregate
	groupAggregate  iquery.Aggregate
	groupByColumns  []string
	offset          int
	limit           int
}

func New(columns ...string) *Builder {
	builder := &Builder{
		columns:         columns,
		where:           where.New(),
		havingAggregate: aggregate.New(),
		groupAggregate:  aggregate.New(),
		groupByColumns:  make([]string, 0),
	}
	builder.refreshColumns()
	return builder
}
