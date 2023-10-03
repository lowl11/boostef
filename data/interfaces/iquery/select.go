package iquery

type Select interface {
	Query

	Select(columns ...string) Select
	From(tableName string) Select
	SetAlias(aliasName string) Select
	Join(tableName, aliasName, joinColumn, mainColumn string) Select
	LeftJoin(tableName, aliasName, joinColumn, mainColumn string) Select
	RightJoin(tableName, aliasName, joinColumn, mainColumn string) Select
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
