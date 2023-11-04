package ctypes

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/data_types"
)

func Binary(size ...int) iquery.DataType {
	return data_types.Binary(size...)
}

func VarBinary(size ...int) iquery.DataType {
	return data_types.VarBinary(size...)
}

func Char(size int) iquery.DataType {
	return data_types.Char(size)
}

func NChar(size int) iquery.DataType {
	return data_types.NChar(size)
}

func Character(size int) iquery.DataType {
	return data_types.Character(size)
}

func Varchar(size int) iquery.DataType {
	return data_types.Varchar(size)
}

func NVarchar(size int) iquery.DataType {
	return data_types.NVarchar(size)
}

func CharacterVarying(size int) iquery.DataType {
	return data_types.CharacterVarying(size)
}

func TinyText(size ...int) iquery.DataType {
	return data_types.TinyText(size...)
}

func Text(size ...int) iquery.DataType {
	return data_types.Text(size...)
}

func NText(size ...int) iquery.DataType {
	return data_types.NText(size...)
}

func MediumText(size ...int) iquery.DataType {
	return data_types.MediumText(size...)
}

func LongText(size ...int) iquery.DataType {
	return data_types.LongText(size...)
}

func Uuid() iquery.DataType {
	return data_types.UUID()
}

func JsonB() iquery.DataType {
	return data_types.JsonB()
}
