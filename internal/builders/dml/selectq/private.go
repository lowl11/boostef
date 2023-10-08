package selectq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/builders/dml/common/join"
	"strings"
)

func (builder *Builder) refreshColumns() *Builder {
	if len(builder.columns) == 0 {
		return builder
	}

	for i := 0; i < len(builder.columns); i++ {
		// aggregate
		if strings.Contains(builder.columns[i], "COUNT(") {
			continue
		}

		// already field name with dot. Example: product.title -> "product"."title"
		if before, after, found := strings.Cut(builder.columns[i], "."); found {
			if isNamed(builder.columns[i]) {
				continue
			}

			builder.columns[i] = makeName(before) + "." + makeName(after)
			continue
		}

		if len(builder.aliasName) != 0 {
			if isNamed(builder.columns[i]) && isAliased(builder.columns[i]) {
				continue
			}

			if isNamed(builder.columns[i]) {
				builder.columns[i] = builder.aliasName + "." + builder.columns[i]
			} else {
				builder.columns[i] = builder.aliasName + "." + makeName(builder.columns[i])
			}
		} else {
			if !isNamed(builder.columns[i]) {
				builder.columns[i] = makeName(builder.columns[i])
			}
		}
	}

	return builder
}

func (builder *Builder) orderByDescending() *Builder {
	builder.isDescending = true
	return builder
}

func (builder *Builder) setOrderByColumns(columns ...string) *Builder {
	builder.orderByColumns = columns
	return builder
}

func (builder *Builder) setColumns(columns ...string) *Builder {
	if len(columns) == 0 {
		return builder
	}

	builder.columns = columns
	return builder.refreshColumns()
}

func (builder *Builder) setTable(tableName string) *Builder {
	if len(tableName) == 0 {
		return builder
	}

	builder.tableName = makeName(tableName)
	return builder
}

func (builder *Builder) setAlias(aliasName string) *Builder {
	if len(aliasName) == 0 {
		return builder
	}

	builder.aliasName = makeName(aliasName)
	return builder.refreshColumns()
}

func (builder *Builder) applyWhere(whereFunc func(builder iquery.Where)) *Builder {
	whereFunc(builder.where)
	return builder
}

func (builder *Builder) setHaving(aggregateFunc func(aggregate iquery.Aggregate)) *Builder {
	aggregateFunc(builder.havingAggregate)
	return builder
}

func (builder *Builder) setGroupBy(columns ...string) *Builder {
	builder.groupByColumns = columns
	return builder
}

func (builder *Builder) setGroupByAggregate(aggregateFunc func(aggregate iquery.Aggregate)) *Builder {
	aggregateFunc(builder.groupAggregate)
	return builder
}

func (builder *Builder) setOffset(offset int) *Builder {
	if offset < 0 {
		return builder
	}

	builder.offset = offset
	return builder
}

func (builder *Builder) setLimit(limit int) *Builder {
	if limit < 0 {
		return builder
	}

	builder.limit = limit
	return builder
}

func (builder *Builder) addJoin(joinType, tableName, aliasName, joinColumn, mainColumn string) *Builder {
	builder.joins = append(builder.joins, join.
		New(joinType).
		Table(tableName).
		Alias(aliasName).
		JoinColumn(joinColumn).
		MainColumn(mainColumn))
	return builder
}
