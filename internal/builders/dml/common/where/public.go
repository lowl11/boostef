package where

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/helpers/stringc"
	"github.com/lowl11/boostef/pkg/enums/signs"
	"strings"
)

func (where *Where) String(_ ...string) string {
	separator := and + "\n\t"
	if where.or {
		separator = or
	}

	result := strings.Join(where.conditions, separator)
	if where.or {
		return brackets(result)
	}

	return result
}

func (where *Where) SetAlias(alias string) iquery.Where {
	where.alias = alias
	return where
}

func (where *Where) Not(condition func(iquery.Where) iquery.Where) iquery.Where {
	where.add(not(condition(New(where.alias)).(iquery.Query).String()))
	return where
}

func (where *Where) Or(condition func(iquery.Where) iquery.Where) iquery.Where {
	where.add(condition(NewOr(where.alias)).(iquery.Query).String())
	return where
}

func (where *Where) Bool(field string, result bool) iquery.Where {
	if !result {
		return where.add(not(field))
	}

	return where.add(field)
}

func (where *Where) Equal(field string, value any) iquery.Where {
	return where.add(build(where.alias, field, signs.Equal, value))
}

func (where *Where) NotEqual(field string, value any) iquery.Where {
	return where.add(build(where.alias, field, signs.NotEqual, value))
}

func (where *Where) In(field string, values []any) iquery.Where {
	queryArray := strings.Builder{}

	queryArray.WriteString("(")
	for index, value := range values {
		queryArray.WriteString(stringc.ToString(value))

		if index < len(values)-1 {
			queryArray.WriteString(", ")
		}
	}
	queryArray.WriteString(")")

	return where.add(buildArray(field, signs.In, queryArray.String()))
}

func (where *Where) Is(field string, value any) iquery.Where {
	return where.add(build(where.alias, field, signs.Is, value))
}

func (where *Where) IsNull(field string) iquery.Where {
	return where.add(build(where.alias, field, signs.Is, "$NULL"))
}

func (where *Where) IsNotNull(field string) iquery.Where {
	return where.add(build(where.alias, field, signs.IsNot, "$NULL"))
}

func (where *Where) Like(field, value string) iquery.Where {
	return where.add(build(where.alias, field, signs.Like, value))
}

func (where *Where) ILike(field, value string) iquery.Where {
	return where.add(build(where.alias, field, signs.ILike, value))
}

func (where *Where) Between(field string, left, right any) iquery.Where {
	return where.add(buildBetween(field, left, right))
}

func (where *Where) Gte(field string, value any) iquery.Where {
	return where.add(build(where.alias, field, signs.Gte, value))
}

func (where *Where) Gt(field string, value any) iquery.Where {
	return where.add(build(where.alias, field, signs.Gt, value))
}

func (where *Where) Lte(field string, value any) iquery.Where {
	return where.add(build(where.alias, field, signs.Lte, value))
}

func (where *Where) Lt(field string, value any) iquery.Where {
	return where.add(build(where.alias, field, signs.Lt, value))
}
