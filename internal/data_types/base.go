package data_types

import (
	"github.com/lowl11/boostef/internal/helpers/stringc"
	"github.com/lowl11/boostef/pkg/enums/sqls"
	"io"
	"strings"
)

const (
	notNull       = " NOT NULL"
	primaryKey    = " PRIMARY KEY"
	autoIncrement = " AUTO_INCREMENT"
	defaultValue  = " DEFAULT "
)

type base struct {
	name          string
	autoIncrement bool
	isPrimary     bool
	isForeign     bool
	foreignTable  string
	notNull       bool
	defaultValue  string
	isUnique      bool
}

func (dt *base) setAutoIncrement() {
	dt.autoIncrement = true
}

func (dt *base) setNotNull() {
	dt.notNull = true
}

func (dt *base) setPrimary() {
	dt.isPrimary = true
}

func (dt *base) setForeign(tableName string) {
	dt.isForeign = true
	dt.foreignTable = tableName
}

func (dt *base) append(sql string, writer io.Writer) {
	var last *string
	defer func() {
		if last != nil {
			_, _ = writer.Write([]byte(*last))
		}
	}()

	if dt.notNull {
		_, _ = writer.Write([]byte(notNull))
	}

	if dt.isPrimary {
		_, _ = writer.Write([]byte(primaryKey))
	} else if dt.isForeign {
		query := strings.Builder{}
		query.WriteString(",\n\tFOREIGN KEY (% FIELD_NAME %) REFERENCES ")

		var tableName string
		var fieldName string
		beforeDot, afterDot, found := strings.Cut(dt.foreignTable, ".")
		if !found {
			tableName = dt.foreignTable
			fieldName = "% FIELD_NAME %"
		} else {
			tableName = beforeDot
			fieldName = afterDot
		}

		query.WriteString(tableName)
		query.WriteString("(")
		query.WriteString(fieldName)
		query.WriteString(")")

		q := query.String()
		last = &q
	}

	if dt.autoIncrement {
		if sql == sqls.Postgres {
			return
		}

		if (sql == sqls.MySQL || sql == sqls.MSSQL) && dt.name != "SERIAL" {
			_, _ = writer.Write([]byte(autoIncrement))
		}
	}

	if dt.defaultValue != "" {
		_, _ = writer.Write([]byte(defaultValue))
		_, _ = writer.Write([]byte(dt.defaultValue))
	}

	if dt.isUnique {
		_, _ = writer.Write([]byte(" UNIQUE"))
	}
}

func (dt *base) writeSize(writer io.Writer, size int) {
	if size == 0 {
		return
	}

	_, _ = writer.Write([]byte("("))
	_, _ = writer.Write([]byte(stringc.ToString(size)))
	_, _ = writer.Write([]byte(")"))
}
