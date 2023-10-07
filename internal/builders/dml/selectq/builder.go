package selectq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/builders/dml/common/where"
)

type Builder struct {
	columns         []string
	tableName       string
	aliasName       string
	joins           []iquery.Join
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
		columns:        columns,
		joins:          make([]iquery.Join, 0, 2),
		where:          where.New(),
		groupByColumns: make([]string, 0),

		offset: -1,
		limit:  -1,
	}
	builder.refreshColumns()
	return builder
}
