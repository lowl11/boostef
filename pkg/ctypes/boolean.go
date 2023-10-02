package ctypes

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/data_types"
)

func Bool() iquery.DataType {
	return data_types.Bool()
}

func Boolean() iquery.DataType {
	return data_types.Boolean()
}
