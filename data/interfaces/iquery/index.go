package iquery

type CreateIndex interface {
	Query

	Unique() CreateIndex
	Name(name string) CreateIndex
	TableColumns(tableName string, columns ...string) CreateIndex
}

type DropIndex interface {
	Query

	SQL(sql string) DropIndex
	Name(name string) DropIndex
	Table(tableName string) DropIndex
}
