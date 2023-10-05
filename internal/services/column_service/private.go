package column_service

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/ctypes"
	"github.com/lowl11/flex"
	"reflect"
	"strings"
)

func defineType(t reflect.Type) iquery.DataType {
	ft := flex.Type(t)
	if ft.IsPtr() || ft.IsStruct() {
		if t.String() == "time.Time" {
			return ctypes.Timestamp()
		}

		ft.Reset(ft.Unwrap())
	} else if strings.Contains(strings.ToLower(t.String()), "uuid") {
		return ctypes.Uuid()
	}

	switch ft.Type().Kind() {
	case reflect.String:
		return ctypes.Text()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return ctypes.Integer()
	case reflect.Bool:
		return ctypes.Boolean()
	case reflect.Float32, reflect.Float64:
		return ctypes.Real()
	default:
		panic("Given column is not primitive variable: " + t.String())
	}

	return nil
}
