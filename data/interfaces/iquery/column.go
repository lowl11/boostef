package iquery

import "io"

type Column interface {
	Query

	Name(name string) Column
	DataType(dataType DataType) Column
}

type DataType interface {
	Write(sql string, writer io.Writer)
	Size(size int) DataType
	Default(defaultValue string) DataType
	AutoIncrement() DataType
	Primary() DataType
	Foreign(string) DataType
	NotNull() DataType
}
