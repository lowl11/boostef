package iquery

type Where interface {
	SetAlias(alias string) Where

	Not(func(Where) Where) Where
	Or(func(Where) Where) Where

	Bool(field string, result bool) Where
	Equal(field string, value any) Where
	NotEqual(field string, value any) Where
	Is(field string, value any) Where
	IsNull(field string) Where
	IsNotNull(field string) Where
	In(field string, values []any) Where
	Like(field, value string) Where
	ILike(field, value string) Where
	Between(field string, left, right any) Where
	Gte(field string, value any) Where
	Gt(field string, value any) Where
	Lte(field string, value any) Where
	Lt(field string, value any) Where
}
