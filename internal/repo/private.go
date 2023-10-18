package repo

import (
	"github.com/lowl11/boostef/pkg/query"
	"github.com/lowl11/flex"
	"strings"
)

func (r *repo[T]) eatEntity(entity T) {
	fs := flex.Struct(entity).Fields()
	for _, f := range fs {
		tags := flex.Field(f).Tag("ef")
		if len(tags) == 0 {
			continue
		}

		for _, tag := range tags {
			before, after, found := strings.Cut(tag, ":")
			if !found {
				continue
			}

			switch before {
			case "table":
				r.tableName = after
			case "alias":
				r.aliasName = after
			}
		}

		if len(r.tableName) > 0 && len(r.aliasName) > 0 {
			break
		}
	}

	r.fields = flex.Struct(entity).FieldsRow()
	columns := make([]string, 0, len(r.fields))
	for _, field := range r.fields {
		columns = append(columns, flex.Field(field).Tag("db")[0])
	}
	r.columns = columns
	r.entity = entity
}

func (r *repo[T]) getPairs(entity T) []query.Pair {
	pairs := make([]query.Pair, 0, len(r.columns))
	fStr := flex.Struct(entity)
	for _, column := range r.columns {
		pairs = append(pairs, query.Pair{
			Column: column,
			Value:  fStr.FieldValueByTag("db", column),
		})
	}
	return pairs
}

func (r *repo[T]) getColumns() []string {
	return r.columns
}

func (r *repo[T]) getTable() string {
	return r.tableName
}
