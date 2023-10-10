package iquery

import "io"

type Column interface {
	Query

	Name(name string) Column
	GetName() string
	DataType(dataType DataType) Column
	GetDataType() DataType
}

type DataType interface {
	Write(sql string, writer io.Writer)
	Name() string
	Size(size int) DataType
	Default(defaultValue string) DataType
	AutoIncrement() DataType
	Primary() DataType
	Foreign(string) DataType
	NotNull() DataType
	Unique() DataType
	Equals(DataType) []string
}
