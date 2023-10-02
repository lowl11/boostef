package iquery

type Select interface {
	Query

	Select(columns ...string) Select
	From(tableName string) Select
	SetAlias(aliasName string) Select
	Where(func(Where)) Select
	OrderBy(columns ...string) Select
	OrderByDescending(columns ...string) Select
	Having(func(Aggregate)) Select
	GroupBy(columns ...string) Select
	GroupByAggregate(func(Aggregate)) Select
	Offset(offset int) Select
	Limit(limit int) Select
	Page(pageSize, pageNumber int) Select
}
