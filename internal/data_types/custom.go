package data_types

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/compares"
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

func (dt *custom) Equals(compare iquery.DataType) []string {
	if compare == nil {
		return []string{}
	}

	compareDataType := compare.(*custom)

	different := make([]string, 0)
	if dt.isPrimary != compareDataType.isPrimary {
		if dt.isPrimary {
			different = append(different, compares.IsPrimaryAdd)
		} else {
			different = append(different, compares.IsPrimaryRemove)
		}
	}

	if dt.notNull != compareDataType.notNull {
		var diff string
		if dt.notNull {
			diff = compares.NotNullAdd
		} else {
			diff = compares.NotNullRemove
		}
		different = append(different, diff)
	}

	if dt.name != compareDataType.name {
		different = append(different, compares.Type)
	}

	if dt.isUnique != compareDataType.isUnique && !dt.isPrimary && !dt.isForeign {
		var diff string
		if dt.isUnique {
			diff = compares.UniqueAdd
		} else {
			diff = compares.UniqueRemove
		}
		different = append(different, diff)
	}

	return different
}

func (dt *custom) Name() string {
	return dt.name
}

func (dt *custom) Size(size int) iquery.DataType {
	dt.size = size
	return dt
}

func (dt *custom) Default(defaultValue string) iquery.DataType {
	dt.defaultValue = defaultValue
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

func (dt *custom) Unique() iquery.DataType {
	dt.isUnique = true
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
