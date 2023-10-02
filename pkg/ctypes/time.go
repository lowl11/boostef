package ctypes

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/data_types"
)

func Timestamp() iquery.DataType {
	return data_types.Timestamp()
}

func TimestampZ() iquery.DataType {
	return data_types.TimestampZ()
}

func SmallDateTime() iquery.DataType {
	return data_types.SmallDateTime()
}

func DateTime() iquery.DataType {
	return data_types.DateTime()
}

func DateTime2() iquery.DataType {
	return data_types.DateTime2()
}

func DateTimeOffset() iquery.DataType {
	return data_types.DateTimeOffset()
}

func Date() iquery.DataType {
	return data_types.Date()
}

func Time() iquery.DataType {
	return data_types.Time()
}

func Year() iquery.DataType {
	return data_types.Year()
}
