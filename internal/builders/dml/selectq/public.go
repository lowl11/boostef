package selectq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func (builder *Builder) Get(_ ...string) string {
	// builder
	query := strings.Builder{}
	query.Grow(300)

	// select
	appendTable(&query, builder.tableName, builder.aliasName, builder.columns)

	// where
	appendWhere(&query, builder.where)

	// order by
	appendOrderBy(&query, builder.orderByColumns, builder.isDescending)

	// having
	appendHaving(&query, builder.havingAggregate)

	// group by
	appendGroupBy(&query, builder.groupAggregate, builder.groupByColumns...)

	// offset
	appendOffset(&query, builder.offset)

	// limit
	appendLimit(&query, builder.limit)

	return query.String()
}

func (builder *Builder) Select(columns ...string) iquery.Select {
	return builder.setColumns(columns...)
}

func (builder *Builder) From(tableName string) iquery.Select {
	return builder.setTable(tableName)
}

func (builder *Builder) SetAlias(aliasName string) iquery.Select {
	return builder.setAlias(aliasName)
}

func (builder *Builder) Where(whereFunc func(builder iquery.Where)) iquery.Select {
	return builder.applyWhere(whereFunc)
}

func (builder *Builder) OrderBy(columns ...string) iquery.Select {
	return builder.setOrderByColumns(columns...)
}

func (builder *Builder) OrderByDescending(columns ...string) iquery.Select {
	return builder.
		setOrderByColumns(columns...).
		orderByDescending()
}

func (builder *Builder) Having(aggregateFunc func(aggregate iquery.Aggregate)) iquery.Select {
	return builder.setHaving(aggregateFunc)
}

func (builder *Builder) GroupBy(columns ...string) iquery.Select {
	return builder.setGroupBy(columns...)
}

func (builder *Builder) GroupByAggregate(aggregateFunc func(aggregate iquery.Aggregate)) iquery.Select {
	return builder.setGroupByAggregate(aggregateFunc)
}

func (builder *Builder) Offset(offset int) iquery.Select {
	return builder.setOffset(offset)
}

func (builder *Builder) Limit(limit int) iquery.Select {
	return builder.setLimit(limit)
}

func (builder *Builder) Page(pageSize, pageNumber int) iquery.Select {
	if pageSize < 0 {
		return builder
	}

	from := pageSize * (pageNumber - 1)
	to := from + pageSize
	return builder.setOffset(from).setLimit(to)
}
