package iquery

type Delete interface {
	Query

	From(tableName string) Delete
	Where(func(Where)) Delete
}
