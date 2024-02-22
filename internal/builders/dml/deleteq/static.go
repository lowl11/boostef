package deleteq

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func appendDelete(query *strings.Builder, tableName string) {
	query.WriteString("DELETE FROM ")
	query.WriteString(tableName)
	query.WriteString("\n")
}

func appendWhere(query *strings.Builder, whereQuery iquery.Where) {
	whereClause := whereQuery.(iquery.Query).String()
	if len(whereClause) == 0 {
		return
	}

	query.WriteString("WHERE \n\t")
	query.WriteString(whereClause)
	query.WriteString("\n")
}
