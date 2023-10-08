package selectq

import (
	"fmt"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/helpers/stringc"
	"strings"
)

func appendTable(query *strings.Builder, tableName, aliasName string, columns []string) {
	query.WriteString("SELECT\n\t")
	if len(columns) == 0 {
		query.WriteString("*")
	} else {
		query.WriteString(strings.Join(columns, ", \n\t"))
	}
	query.WriteString("\nFROM ")
	query.WriteString(tableName)
	if len(aliasName) > 0 {
		query.WriteString(" AS ")
		query.WriteString(aliasName)
	}
}

func appendJoins(query *strings.Builder, joins []iquery.Join) {
	for _, join := range joins {
		query.WriteString("\n\t")
		query.WriteString(join.Get())
	}
}

func appendWhere(query *strings.Builder, whereQuery iquery.Where) {
	whereClause := whereQuery.(iquery.Query).Get()
	if len(whereClause) == 0 {
		return
	}

	query.WriteString("\nWHERE \n\t")
	query.WriteString(whereClause)
	query.WriteString("\n")
}

func appendOrderBy(query *strings.Builder, orderByColumns []string, isDescending bool) {
	if len(orderByColumns) > 0 {
		query.WriteString("ORDER BY ")
		query.WriteString(strings.Join(orderByColumns, ", "))
		if !isDescending {
			query.WriteString(" ASC")
		} else {
			query.WriteString(" DESC")
		}
		query.WriteString("\n")
	}
}

func appendHaving(query *strings.Builder, aggregate iquery.Aggregate) {
	if aggregate == nil {
		return
	}

	query.WriteString("HAVING ")
	query.WriteString(aggregate.(iquery.Query).Get())
	query.WriteString("\n")
}

func appendGroupBy(query *strings.Builder, aggregate iquery.Aggregate, columns ...string) {
	if aggregate == nil {
		return
	}

	query.WriteString("GROUP BY ")
	if len(columns) > 0 {
		for index, column := range columns {
			query.WriteString(column)

			if index < len(columns)-1 {
				query.WriteString(", ")
			}
		}
	} else {
		query.WriteString(aggregate.(iquery.Query).Get())
	}
	query.WriteString("\n")
}

func appendOffset(query *strings.Builder, offset int) {
	if offset > -1 {
		query.WriteString("\nOFFSET " + stringc.ToString(offset) + "\n")
	}
}

func appendLimit(query *strings.Builder, limit int) {
	if limit > -1 {
		query.WriteString("LIMIT " + stringc.ToString(limit) + "\n")
	}
}

func makeName(name string) string {
	query := strings.Builder{}
	query.Grow(len(name) + 2)
	if strings.Contains(name, ".") {
		before, after, _ := strings.Cut(name, ".")
		_, _ = fmt.Fprintf(&query, "\"%s\".\"%s\"", before, after)
	} else {
		_, _ = fmt.Fprintf(&query, "\"%s\"", name)
	}
	return query.String()
}

func isNamed(name string) bool {
	return strings.Contains(name, "\"")
}

func isAliased(name string) bool {
	return strings.Contains(name, ".")
}
