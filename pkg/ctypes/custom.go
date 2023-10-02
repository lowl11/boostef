package ctypes

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/data_types"
)

func Custom(name string) iquery.DataType {
	return data_types.Custom(name)
}
