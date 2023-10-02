package ctypes

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/data_types"
)

func Bit(size ...int) iquery.DataType {
	return data_types.Bit(size...)
}

func TinyInt(size ...int) iquery.DataType {
	return data_types.TinyInt(size...)
}

func Int(size ...int) iquery.DataType {
	return data_types.Int(size...)
}

func MediumInt(size ...int) iquery.DataType {
	return data_types.MediumInt(size...)
}

func Decimal(size ...int) iquery.DataType {
	return data_types.Decimal(size...)
}

func Dec(size ...int) iquery.DataType {
	return data_types.Dec(size...)
}

func Double(size ...int) iquery.DataType {
	return data_types.Double(size...)
}

func Float(size ...int) iquery.DataType {
	return data_types.Float(size...)
}

func SmallMoney(size ...int) iquery.DataType {
	return data_types.SmallMoney(size...)
}

func Money(size ...int) iquery.DataType {
	return data_types.Money(size...)
}

func SmallSerial() iquery.DataType {
	return data_types.SmallSerial()
}

func Serial() iquery.DataType {
	return data_types.Serial()
}

func BigSerial() iquery.DataType {
	return data_types.BigSerial()
}

func SmallInt() iquery.DataType {
	return data_types.SmallInt()
}

func Integer(size ...int) iquery.DataType {
	return data_types.Integer(size...)
}

func Bigint() iquery.DataType {
	return data_types.Bigint()
}

func Real() iquery.DataType {
	return data_types.Real()
}
