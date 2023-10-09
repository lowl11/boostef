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
	AddColumn(column string) AlterTable
	DropColumn(column string) AlterTable
	RenameColumn(column, newName string) AlterTable
	AlterColumn(column string) AlterTable
	Set(string) AlterTable
	Type(DataType) AlterTable
	Reset() AlterTable
	Restart() AlterTable
	Add() AlterTable
	Drop() AlterTable
}
