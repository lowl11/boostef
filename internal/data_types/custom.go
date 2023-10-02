package data_types

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"io"
	"strings"
)

type custom struct {
	base
	size   int
	custom func(sql string) string
}

func createCustom(name string, size ...int) *custom {
	return Custom(name, size...).(*custom)
}

func Custom(name string, size ...int) iquery.DataType {
	dt := &custom{}
	dt.name = strings.ToUpper(name)

	if len(size) > 0 {
		dt.size = size[0]
	}

	return dt
}

func (dt *custom) setCustom(custom func(sql string) string) *custom {
	dt.custom = custom
	return dt
}

func (dt *custom) AutoIncrement() iquery.DataType {
	dt.setAutoIncrement()
	return dt
}

func (dt *custom) NotNull() iquery.DataType {
	dt.setNotNull()
	return dt
}

func (dt *custom) Primary() iquery.DataType {
	dt.setPrimary()
	return dt
}

func (dt *custom) Foreign(tableName string) iquery.DataType {
	dt.setForeign(tableName)
	return dt
}

func (dt *custom) Write(sql string, writer io.Writer) {
	if dt.custom != nil {
		_, _ = writer.Write([]byte(dt.custom(sql)))
	} else {
		_, _ = writer.Write([]byte(dt.name))
	}

	dt.writeSize(writer, dt.size)
	dt.append(sql, writer)
}
