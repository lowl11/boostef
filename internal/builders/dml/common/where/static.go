package where

import (
	"fmt"
	"github.com/lowl11/boostef/internal/helpers/stringc"
	"strings"
)

const (
	and = " AND "
	or  = " OR "
)

func build(alias, field, sign string, value any) string {
	valueString := stringc.ToString(value)

	builder := strings.Builder{}
	if len(alias) > 0 {
		if !strings.Contains(field, ".") {
			_, _ = fmt.Fprintf(&builder, "%s.%s %s %s", alias, field, sign, valueString)
		} else {
			_, after, _ := strings.Cut(field, ".")
			_, _ = fmt.Fprintf(&builder, "%s.%s %s %s", alias, after, sign, valueString)
		}
	} else {
		_, _ = fmt.Fprintf(&builder, "%s %s %s", field, sign, valueString)
	}

	return builder.String()
}

func buildBetween(field string, left, right any) string {
	leftString := stringc.ToString(left)
	rightString := stringc.ToString(right)

	builder := strings.Builder{}
	_, _ = fmt.Fprintf(&builder, "%s BETWEEN %s AND %s", field, leftString, rightString)
	return builder.String()
}

func buildArray(field, sign, value string) string {
	builder := strings.Builder{}
	_, _ = fmt.Fprintf(&builder, "%s %s %s", field, sign, value)
	return builder.String()
}

func not(condition string) string {
	builder := strings.Builder{}
	builder.Grow(len(condition) + 5)
	_, _ = fmt.Fprintf(&builder, "NOT(%s)", condition)
	return builder.String()
}

func brackets(condition string) string {
	builder := strings.Builder{}
	builder.Grow(len(condition) + 5)
	_, _ = fmt.Fprintf(&builder, "(%s)", condition)
	return builder.String()
}
