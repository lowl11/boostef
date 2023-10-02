package iquery

type CreateTable interface {
	Query

	Table(tableName string) CreateTable
	IfNotExist() CreateTable
	Column(columns ...Column) CreateTable

	Sql(sql string) CreateTable
}

type DropTable interface {
	Query

	Table(tableName string) DropTable
}

type TruncateTable interface {
	Query

	Table(tableName string) TruncateTable
}

type AlterTable interface {
	Query

	Table(tableName string) AlterTable
	SQL(sql string) AlterTable
	Add(column string) AlterTable
	Drop(column string) AlterTable
	Rename(column, newName string) AlterTable
	Alter(column string) AlterTable
	DataType(dataType DataType) AlterTable
}
