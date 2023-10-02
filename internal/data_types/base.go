package data_types

import (
	"github.com/lowl11/boostef/internal/helpers/stringc"
	"github.com/lowl11/boostef/pkg/enums/sqls"
	"io"
	"strings"
)

type base struct {
	name          string
	autoIncrement bool
	isPrimary     bool
	isForeign     bool
	foreignTable  string
	notNull       bool
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
		_, _ = writer.Write([]byte(" NOT NULL"))
	}

	if dt.isPrimary {
		_, _ = writer.Write([]byte(" PRIMARY KEY"))
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

		// product_id uuid,
		// foreign key (product_id) references stock.products(id)

		//_, _ = writer.Write([]byte(" REFERENCES "))
		//_, _ = writer.Write([]byte(dt.foreignTable))
	}

	if dt.autoIncrement {
		if sql == sqls.Postgres {
			return
		}

		if (sql == sqls.MySQL || sql == sqls.MSSQL) && dt.name != "SERIAL" {
			_, _ = writer.Write([]byte(" AUTO_INCREMENT"))
		}
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
