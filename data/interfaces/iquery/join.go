package iquery

type Join interface {
	Query

	Table(tableName string) Join
	Alias(aliasName string) Join
	JoinColumn(column string) Join
	MainColumn(column string) Join
}
