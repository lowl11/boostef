package iquery

type CreateTable interface {
	Query

	Table(tableName string) CreateTable
	IfNotExist() CreateTable
	Column(columns ...Column) CreateTable

	Sql(sql string) CreateTable
}

type DropTable interface {
	//
}
