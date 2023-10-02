package updateq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/helpers/stringc"
	"github.com/lowl11/boostef/pkg/query"
	"strings"
)

func appendUpdate(query *strings.Builder, tableName string) {
	if len(tableName) == 0 {
		return
	}

	query.WriteString("UPDATE ")
	query.WriteString(tableName)
	query.WriteString("\n")
}

func appendSet(query *strings.Builder, pairs []query.Pair) {
	if len(pairs) == 0 {
		return
	}

	query.WriteString("SET\n")
	for index, pair := range pairs {
		query.WriteString("\t")
		query.WriteString(pair.Column)
		query.WriteString(" = ")
		query.WriteString(stringc.ToString(pair.Value))
		if index < len(pairs)-1 {
			query.WriteString(",\n")
		}
	}
	query.WriteString("\n")
}

func appendWhere(query *strings.Builder, where iquery.Where) {
	whereClause := where.(iquery.Query).Get()
	if len(whereClause) == 0 {
		return
	}

	query.WriteString("WHERE \n\t")
	query.WriteString(whereClause)
	query.WriteString("\n")
}
