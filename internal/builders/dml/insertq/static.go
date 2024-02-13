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

func appendValues(query *strings.Builder, pairs []query.Pair, multiplePairs [][]query.Pair) {
	if len(pairs) == 0 {
		return
	}

	isMultiple := len(multiplePairs) > 0

	var isNamedValues bool
	if pairs[0].Value == nil {
		isNamedValues = true
	}

	if !isMultiple {
		query.WriteString("VALUES (")
	} else {
		query.WriteString("VALUES\n")
	}

	if !isMultiple {
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
	} else {
		for mIndex, mPairs := range multiplePairs {
			query.WriteString("\t(")
			for index, pair := range mPairs {
				query.WriteString(stringc.ToString(pair.Value))

				if index < len(mPairs)-1 {
					query.WriteString(", ")
				}
			}

			query.WriteString(")")
			if mIndex < len(multiplePairs)-1 {
				query.WriteString(",")
			}
			query.WriteString("\n")
		}
	}

	if !isMultiple {
		query.WriteString(")\n")
	}
}

func appendOnConflict(query *strings.Builder, conflict string) {
	if len(conflict) == 0 {
		return
	}

	query.WriteString("ON CONFLICT ")
	query.WriteString(conflict)
	query.WriteString("\n")
}
