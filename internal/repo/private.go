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

		before, after, found := strings.Cut(tags[0], ":")
		if !found {
			continue
		}

		if before == "table" {
			r.tableName = after
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
