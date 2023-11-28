package insertq

import (
	"github.com/lowl11/boostef/internal/helpers/stringc"
	"github.com/lowl11/boostef/pkg/query"
	"strings"
)

func appendInsert(query *strings.Builder, tableName string, pairs []query.Pair) {
	if len(pairs) == 0 {
		return
	}

	query.WriteString("INSERT INTO ")
	query.WriteString(tableName)
	query.WriteString(" (")
	for index, pair := range pairs {
		query.WriteString(pair.Column)

		if index < len(pairs)-1 {
			query.WriteString(", ")
		}
	}
	query.WriteString(")\n")
}

func appendValues(query *strings.Builder, pairs []query.Pair) {
	if len(pairs) == 0 {
		return
	}

	var isNamedValues bool
	if pairs[0].Value == nil {
		isNamedValues = true
	}

	query.WriteString("VALUES (")
	for index, pair := range pairs {
		if isNamedValues {
			query.WriteString(":")
			query.WriteString(pair.Column)
		} else {
			query.WriteString(stringc.ToString(pair.Value))
		}

		if index < len(pairs)-1 {
			query.WriteString(", ")
		}
	}

	query.WriteString(")\n")
}

func appendOnConflict(query *strings.Builder, conflict string) {
	if len(conflict) == 0 {
		return
	}

	query.WriteString("ON CONFLICT ")
	query.WriteString(conflict)
	query.WriteString("\n")
}
